package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Conversation struct {
	ID              string    `gorm:"type:varchar(36);primaryKey" json:"id"`
	TenantID        string    `gorm:"type:varchar(36);not null;index" json:"tenant_id"`
	CustomerID      string    `gorm:"type:varchar(36);not null" json:"customer_id"`
	Status          string    `gorm:"type:enum('open', 'assigned', 'closed');default:'open';index" json:"status"`
	AssignedAgentID *string   `gorm:"type:varchar(36)" json:"assigned_agent_id,omitempty"`
	Messages        []Message `gorm:"foreignKey:ConversationID" json:"messages"`
	Ticket          *Ticket   `gorm:"foreignKey:ConversationID" json:"ticket,omitempty"`
	Base            `json:",inline"`
}

func (b *Conversation) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
