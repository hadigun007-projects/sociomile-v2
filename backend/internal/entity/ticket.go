package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ticket struct {
	ID              string  `gorm:"type:varchar(36);primaryKey"`
	TenantID        string  `gorm:"type:varchar(36);not null;index"`
	ConversationID  string  `gorm:"type:varchar(36);unique;not null"`
	Title           string  `gorm:"type:varchar(255);not null"`
	Description     string  `gorm:"type:text"`
	Status          string  `gorm:"type:enum('open', 'in_progress', 'resolved', 'closed');default:'open'"`
	Priority        string  `gorm:"type:enum('low', 'medium', 'high', 'urgent');default:'medium'"`
	AssignedAgentID *string `gorm:"type:varchar(36)"`
	Base
}

func (b *Ticket) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
