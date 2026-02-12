package services

import (
	"errors"
	"fmt"

	"github.com/cinnamorollofficials/sociomile-v2/backend/config"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(email, password string) (*entity.User, error)
}

type authService struct {
	userRepo   repository.UserRepository
	bcryptCost int
}

func NewAuthService(
	userRepo repository.UserRepository,
	cfg *config.Config,
) AuthService {
	return &authService{
		userRepo:   userRepo,
		bcryptCost: bcrypt.DefaultCost,
	}
}

func (s *authService) Login(email, password string) (*entity.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("invalid email or password")
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	return user, nil
}
