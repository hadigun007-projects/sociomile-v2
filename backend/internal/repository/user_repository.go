package repository

import (
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindByID(id, tenantID string) (*entity.User, error)
	FindByIDSystem(id string) (*entity.User, error) // For internal use (e.g. Refresh Token)
	FindByEmail(email string) (*entity.User, error) // Email is globally unique
	Update(user *entity.User) error
	Delete(id, tenantID string) error
	GetUsers(tenantID string) ([]entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByID(id, tenantID string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ? AND tenant_id = ?", id, tenantID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByIDSystem(id string) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id, tenantID string) error {
	return r.db.Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&entity.User{}).Error
}

func (r *userRepository) GetUsers(tenantID string) ([]entity.User, error) {
	var users []entity.User
	if err := r.db.Where("tenant_id = ?", tenantID).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
