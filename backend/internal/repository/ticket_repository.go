package repository

import (
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
	"gorm.io/gorm"
)

type TicketRepository interface {
	Create(ticket *entity.Ticket) error
	Assign(ticketID, userID string) error
	FindByID(id string) (*entity.Ticket, error)
	Update(ticket *entity.Ticket) error
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db: db}
}

func (r *ticketRepository) Create(ticket *entity.Ticket) error {
	return r.db.Create(ticket).Error
}

func (r *ticketRepository) Assign(ticketID, userID string) error {
	return r.db.Model(&entity.Ticket{}).Where("id = ?", ticketID).Update("assigned_agent_id", userID).Error
}

func (r *ticketRepository) FindByID(id string) (*entity.Ticket, error) {
	var ticket entity.Ticket
	if err := r.db.First(&ticket, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (r *ticketRepository) Update(ticket *entity.Ticket) error {
	return r.db.Save(ticket).Error
}
