package repository

import (
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
	"gorm.io/gorm"
)

type TenantRepository interface {
	Create(tenant *entity.Tenant) error
	GetByID(id string) (*entity.Tenant, error)
	GetByName(name string) (*entity.Tenant, error)
	GetAll() ([]entity.Tenant, error)
	GetAllWithPagination(page, limit int) ([]entity.Tenant, int64, error)
	Update(tenant *entity.Tenant) error
	Delete(id string) error
}

type tenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) TenantRepository {
	return &tenantRepository{db: db}
}

func (r *tenantRepository) Create(tenant *entity.Tenant) error {
	return r.db.Create(tenant).Error
}

func (r *tenantRepository) GetByID(id string) (*entity.Tenant, error) {
	var tenant entity.Tenant
	if err := r.db.Where("id = ?", id).First(&tenant).Error; err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (r *tenantRepository) GetByName(name string) (*entity.Tenant, error) {
	var tenant entity.Tenant
	if err := r.db.Where("name = ?", name).First(&tenant).Error; err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (r *tenantRepository) GetAll() ([]entity.Tenant, error) {
	var tenants []entity.Tenant
	if err := r.db.Find(&tenants).Error; err != nil {
		return nil, err
	}
	return tenants, nil
}

func (r *tenantRepository) GetAllWithPagination(page, limit int) ([]entity.Tenant, int64, error) {
	var tenants []entity.Tenant
	var total int64

	offset := (page - 1) * limit

	query := r.db.Model(&entity.Tenant{})

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&tenants).Error; err != nil {
		return nil, 0, err
	}

	return tenants, total, nil
}

func (r *tenantRepository) Update(tenant *entity.Tenant) error {
	return r.db.Save(tenant).Error
}

func (r *tenantRepository) Delete(id string) error {
	return r.db.Delete(&entity.Tenant{}, "id = ?", id).Error
}
