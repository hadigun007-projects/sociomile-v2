package services

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/middleware"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/repository"
	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(userID, email, role, tenantID string) (string, error)
	GenerateRefreshToken(userID string) (string, error)
}

type jwtService struct {
	secret           string
	expiryHours      int
	refreshTokenRepo repository.RefreshTokenRepository
	userRepo         repository.UserRepository
}

func NewJWTService(secret string, expiryHours int, refreshTokenRepo repository.RefreshTokenRepository, userRepo repository.UserRepository) JWTService {
	return &jwtService{
		secret:           secret,
		expiryHours:      expiryHours,
		refreshTokenRepo: refreshTokenRepo,
		userRepo:         userRepo,
	}
}

func (s *jwtService) GenerateToken(userID, email, role, tenantID string) (string, error) {
	claims := &middleware.JWTClaims{
		UserID:   userID,
		Email:    email,
		Role:     role,
		TenantID: tenantID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}

func (s *jwtService) GenerateRefreshToken(userID string) (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("failed to generate random token: %w", err)
	}
	tokenString := base64.URLEncoding.EncodeToString(b)

	refreshToken := &entity.RefreshToken{
		UserID:    userID,
		Token:     tokenString,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour), // 7 days
		Revoked:   false,
	}

	if err := s.refreshTokenRepo.Create(refreshToken); err != nil {
		return "", fmt.Errorf("failed to store refresh token: %w", err)
	}

	return tokenString, nil
}

func (s *jwtService) ValidateToken(tokenString string) (*middleware.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &middleware.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*middleware.JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func (s *jwtService) RefreshAccessToken(refreshToken string) (string, string, error) {
	storedToken, err := s.refreshTokenRepo.FindByToken(refreshToken)
	if err != nil {
		return "", "", fmt.Errorf("invalid refresh token")
	}

	user, err := s.userRepo.FindByIDSystem(storedToken.UserID)
	if err != nil {
		return "", "", fmt.Errorf("failed to find user: %w", err)
	}

	accessToken, err := s.GenerateToken(storedToken.UserID, user.Email, user.Role, user.TenantID)
	if err != nil {
		return "", "", err
	}
	newRefreshToken, err := s.GenerateRefreshToken(storedToken.UserID)
	if err != nil {
		return "", "", err
	}

	_ = s.refreshTokenRepo.RevokeByToken(refreshToken)

	return accessToken, newRefreshToken, nil
}
