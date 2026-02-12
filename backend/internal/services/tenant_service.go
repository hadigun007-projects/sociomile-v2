package services

import (
	"errors"
	"fmt"

	"github.com/cinnamorollofficials/sociomile-v2/backend/config"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/handler/dto"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type TenantService interface {
	GetTenants() ([]entity.Tenant, error)
	CreateTenant(input dto.CreateTenantRequest) (*entity.Tenant, error)
	UpdateTenant(id string, input map[string]interface{}) (*entity.Tenant, error)
	DeleteTenant(id string) error
}

type tenantService struct {
	tenantRepo repository.TenantRepository
	userRepo   repository.UserRepository
	cfg        *config.Config
}

func NewTenantService(tenantRepo repository.TenantRepository, userRepo repository.UserRepository, cfg *config.Config) TenantService {
	return &tenantService{tenantRepo: tenantRepo, userRepo: userRepo, cfg: cfg}
}

func (s *tenantService) GetTenants() ([]entity.Tenant, error) {
	return s.tenantRepo.GetTenants()
}

func (s *tenantService) CreateTenant(input dto.CreateTenantRequest) (*entity.Tenant, error) {
	// 1. Create Tenant
	tenant := &entity.Tenant{
		Name: input.Name,
		Plan: input.Plan,
	}

	if err := s.tenantRepo.Create(tenant); err != nil {
		return nil, fmt.Errorf("failed to create tenant: %w", err)
	}

	// 2. Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.AdminPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// 3. Create Default Admin User
	adminUser := &entity.User{
		TenantID: tenant.ID,
		Email:    input.AdminEmail,
		Password: string(hashedPassword),
		Role:     "admin", // Fixed as admin for the first user
	}

	if err := s.userRepo.Create(adminUser); err != nil {
		// Ideally we should rollback tenant creation here if not using transaction
		// For simplicity, we just return error. In production, use transaction.
		return nil, fmt.Errorf("failed to create admin user: %w", err)
	}

	return tenant, nil
}

func (s *tenantService) UpdateTenant(id string, input map[string]interface{}) (*entity.Tenant, error) {
	tenant, err := s.tenantRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("tenant not found")
		}
		return nil, err
	}

	if val, ok := input["name"]; ok {
		tenant.Name = val.(string)
	}

	if val, ok := input["plan"]; ok {
		tenant.Plan = val.(string)
	}

	if err := s.tenantRepo.Update(tenant); err != nil {
		return nil, err
	}

	return tenant, nil
}

func (s *tenantService) DeleteTenant(id string) error {
	return s.tenantRepo.Delete(id)
}
