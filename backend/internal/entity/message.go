package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	ID             string `gorm:"type:varchar(36);primaryKey"`
	ConversationID string `gorm:"type:varchar(36);not null;index"`
	SenderType     string `gorm:"type:enum('customer', 'agent');not null"`
	Message        string `gorm:"type:text;not null"`
	Base
}

func (b *Message) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
