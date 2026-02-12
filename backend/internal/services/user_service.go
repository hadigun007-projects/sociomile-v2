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

type UserService interface {
	GetUsers() ([]entity.User, error)
	UpdateUser(id string, input map[string]interface{}) (*entity.User, error)
	DeleteUser(id string) error
	CreateUser(user *entity.User) error
}

type userService struct {
	userRepo repository.UserRepository
	cfg      *config.Config
}

func NewUserService(userRepo repository.UserRepository, cfg *config.Config) UserService {
	return &userService{userRepo: userRepo, cfg: cfg}
}

func (s *userService) GetUsers() ([]entity.User, error) {
	return s.userRepo.GetUsers()
}

func (s *userService) UpdateUser(id string, input map[string]interface{}) (*entity.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	if val, ok := input["role"]; ok {
		user.Role = val.(string)
	}

	if val, ok := input["password"]; ok {
		password := val.(string)
		if password != "" {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				return nil, err
			}
			user.Password = string(hashedPassword)
		}
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id string) error {
	return s.userRepo.Delete(id)
}

func (s *userService) CreateUser(user *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.userRepo.Create(user)
}
