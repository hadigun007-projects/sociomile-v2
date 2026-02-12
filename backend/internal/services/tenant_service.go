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
	GetTenantsWithPagination(page, limit int) ([]entity.Tenant, int64, error)
	CreateTenant(req dto.CreateTenantRequest) (*entity.Tenant, error)
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
	return s.tenantRepo.GetAll()
}

func (s *tenantService) GetTenantsWithPagination(page, limit int) ([]entity.Tenant, int64, error) {
	return s.tenantRepo.GetAllWithPagination(page, limit)
}

func (s *tenantService) CreateTenant(req dto.CreateTenantRequest) (*entity.Tenant, error) {
	// 1. Create Tenant
	tenant := &entity.Tenant{
		Name: req.Name,
		Plan: req.Plan,
	}

	if err := s.tenantRepo.Create(tenant); err != nil {
		return nil, fmt.Errorf("failed to create tenant: %w", err)
	}

	// 2. Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.AdminPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// 3. Create Default Admin User
	adminUser := &entity.User{
		TenantID: tenant.ID,
		Email:    req.AdminEmail,
		Password: string(hashedPassword),
		Role:     "admin",
	}

	if err := s.userRepo.Create(adminUser); err != nil {
		// Ideally rollback
		return nil, fmt.Errorf("failed to create admin user: %w", err)
	}

	return tenant, nil
}

func (s *tenantService) UpdateTenant(id string, input map[string]interface{}) (*entity.Tenant, error) {
	tenant, err := s.tenantRepo.GetByID(id)
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
