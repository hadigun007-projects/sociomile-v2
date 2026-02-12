package repository

import (
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindByExternalID(tenantID, externalID string) (*entity.Customer, error)
	Create(customer *entity.Customer) error
	GetByTenantID(tenantID string) ([]entity.Customer, error)
	GetByTenantIDWithPagination(tenantID string, page, limit int) ([]entity.Customer, int64, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) FindByExternalID(tenantID, externalID string) (*entity.Customer, error) {
	var customer entity.Customer
	if err := r.db.Where("tenant_id = ? AND external_id = ?", tenantID, externalID).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) Create(customer *entity.Customer) error {
	return r.db.Create(customer).Error
}

func (r *customerRepository) GetByTenantID(tenantID string) ([]entity.Customer, error) {
	var customers []entity.Customer
	if err := r.db.Where("tenant_id = ?", tenantID).Order("created_at DESC").Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *customerRepository) GetByTenantIDWithPagination(tenantID string, page, limit int) ([]entity.Customer, int64, error) {
	var customers []entity.Customer
	var total int64

	offset := (page - 1) * limit

	query := r.db.Model(&entity.Customer{}).Where("tenant_id = ?", tenantID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&customers).Error; err != nil {
		return nil, 0, err
	}

	return customers, total, nil
}
