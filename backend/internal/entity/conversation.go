package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Conversation struct {
	ID              string    `gorm:"type:varchar(36);primaryKey"`
	TenantID        string    `gorm:"type:varchar(36);not null;index"`
	CustomerID      string    `gorm:"type:varchar(36);not null"`
	Status          string    `gorm:"type:enum('open', 'assigned', 'closed');default:'open';index"`
	AssignedAgentID *string   `gorm:"type:varchar(36)"`
	Messages        []Message `gorm:"foreignKey:ConversationID"`
	Ticket          *Ticket   `gorm:"foreignKey:ConversationID"`
	Base
}

func (b *Conversation) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
