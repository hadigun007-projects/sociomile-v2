package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ticket struct {
	ID              string        `gorm:"type:varchar(36);primaryKey" json:"id"`
	TenantID        string        `gorm:"type:varchar(36);not null;index" json:"tenant_id"`
	ConversationID  string        `gorm:"type:varchar(36);unique;not null" json:"conversation_id"`
	Title           string        `gorm:"type:varchar(255);not null" json:"title"`
	Description     string        `gorm:"type:text" json:"description"`
	Status          string        `gorm:"type:enum('requested', 'open', 'in_progress', 'resolved', 'closed');default:'open'" json:"status"`
	Priority        string        `gorm:"type:enum('low', 'medium', 'high', 'urgent');default:'medium'" json:"priority"`
	AssignedAgentID *string       `gorm:"type:varchar(36)" json:"assigned_agent_id"`
	AssignedAgent   *User         `gorm:"foreignKey:AssignedAgentID" json:"assigned_agent"`
	Conversation    *Conversation `gorm:"foreignKey:ConversationID" json:"conversation"`
	Base            `json:",inline"`
}

func (b *Ticket) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
