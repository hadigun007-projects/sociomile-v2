package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	ID             string `gorm:"type:varchar(36);primaryKey" json:"id"`
	ConversationID string `gorm:"type:varchar(36);not null;index" json:"conversation_id"`
	SenderType     string `gorm:"type:enum('customer', 'agent');not null" json:"sender_type"`
	Message        string `gorm:"type:text;not null" json:"message"`
	Base           `json:",inline"`
}

func (b *Message) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
