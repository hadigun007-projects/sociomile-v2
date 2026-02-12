package services

import (
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/repository"
)

type CustomerService interface {
	GetCustomers(tenantID string) ([]entity.Customer, error)
}

type customerService struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerService(customerRepo repository.CustomerRepository) CustomerService {
	return &customerService{customerRepo: customerRepo}
}

func (s *customerService) GetCustomers(tenantID string) ([]entity.Customer, error) {
	return s.customerRepo.GetByTenantID(tenantID)
}
