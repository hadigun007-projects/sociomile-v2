package services

import (
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/repository"
)

type TicketService interface {
	GetTicketsWithPagination(tenantID string, page, limit int) ([]entity.Ticket, int64, error)
	UpdateTicketStatus(id, tenantID, status string) (*entity.Ticket, error)
}

type ticketService struct {
	ticketRepo repository.TicketRepository
}

func NewTicketService(ticketRepo repository.TicketRepository) TicketService {
	return &ticketService{ticketRepo: ticketRepo}
}

func (s *ticketService) GetTicketsWithPagination(tenantID string, page, limit int) ([]entity.Ticket, int64, error) {
	return s.ticketRepo.GetByTenantIDWithPagination(tenantID, page, limit)
}

func (s *ticketService) UpdateTicketStatus(id, tenantID, status string) (*entity.Ticket, error) {
	ticket, err := s.ticketRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if ticket.TenantID != tenantID {
		return nil, entity.ErrTicketNotFound
	}

	ticket.Status = status
	if err := s.ticketRepo.Update(ticket); err != nil {
		return nil, err
	}

	return ticket, nil
}
