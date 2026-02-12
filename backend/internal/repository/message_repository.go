package repository

import (
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
	"gorm.io/gorm"
)

type MessageRepository interface {
	Create(message *entity.Message) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) Create(message *entity.Message) error {
	return r.db.Create(message).Error
}
