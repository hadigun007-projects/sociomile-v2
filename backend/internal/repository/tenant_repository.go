package repository

import (
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
	"gorm.io/gorm"
)

type TenantRepository interface {
	GetTenants() ([]entity.Tenant, error)
	Create(tenant *entity.Tenant) error
	FindByID(id string) (*entity.Tenant, error)
	Update(tenant *entity.Tenant) error
	Delete(id string) error
}

type tenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) TenantRepository {
	return &tenantRepository{db: db}
}

func (r *tenantRepository) GetTenants() ([]entity.Tenant, error) {
	var tenants []entity.Tenant
	if err := r.db.Find(&tenants).Error; err != nil {
		return nil, err
	}
	return tenants, nil
}

func (r *tenantRepository) Create(tenant *entity.Tenant) error {
	return r.db.Create(tenant).Error
}

func (r *tenantRepository) FindByID(id string) (*entity.Tenant, error) {
	var tenant entity.Tenant
	if err := r.db.First(&tenant, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (r *tenantRepository) Update(tenant *entity.Tenant) error {
	return r.db.Save(tenant).Error
}

func (r *tenantRepository) Delete(id string) error {
	return r.db.Delete(&entity.Tenant{}, "id = ?", id).Error
}
