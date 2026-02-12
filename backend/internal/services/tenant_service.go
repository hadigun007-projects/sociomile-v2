package services

import (
	"errors"
	"fmt"

	"github.com/cinnamorollofficials/sociomile-v2/backend/config"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/repository"
	"gorm.io/gorm"
)

type TenantService interface {
	GetTenants() ([]entity.Tenant, error)
	CreateTenant(tenant *entity.Tenant) error
	UpdateTenant(id string, input map[string]interface{}) (*entity.Tenant, error)
	DeleteTenant(id string) error
}

type tenantService struct {
	tenantHandler repository.TenantRepository
	cfg           *config.Config
}

func NewTenantService(tenantHandler repository.TenantRepository, cfg *config.Config) TenantService {
	return &tenantService{tenantHandler: tenantHandler, cfg: cfg}
}

func (s *tenantService) GetTenants() ([]entity.Tenant, error) {
	return s.tenantHandler.GetTenants()
}

func (s *tenantService) CreateTenant(tenant *entity.Tenant) error {
	return s.tenantHandler.Create(tenant)
}

func (s *tenantService) UpdateTenant(id string, input map[string]interface{}) (*entity.Tenant, error) {
	tenant, err := s.tenantHandler.FindByID(id)
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

	if err := s.tenantHandler.Update(tenant); err != nil {
		return nil, err
	}

	return tenant, nil
}

func (s *tenantService) DeleteTenant(id string) error {
	return s.tenantHandler.Delete(id)
}
