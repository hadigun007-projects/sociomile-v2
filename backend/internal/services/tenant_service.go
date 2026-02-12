package services

import (
	"github.com/cinnamorollofficials/sociomile-v2/backend/config"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/repository"
)

type TenantService interface {
	GetTenants() ([]entity.Tenant, error)
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
