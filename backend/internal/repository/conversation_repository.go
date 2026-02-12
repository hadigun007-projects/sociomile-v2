package repository

import (
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
	"gorm.io/gorm"
)

type ConversationRepository interface {
	FindByCustomerID(tenantID, customerID string) (*entity.Conversation, error)
	Create(conversation *entity.Conversation) error
}

type conversationRepository struct {
	db *gorm.DB
}

func NewConversationRepository(db *gorm.DB) ConversationRepository {
	return &conversationRepository{db: db}
}

func (r *conversationRepository) FindByCustomerID(tenantID, customerID string) (*entity.Conversation, error) {
	var conversation entity.Conversation
	if err := r.db.Where("tenant_id = ? AND customer_id = ? AND status != ?", tenantID, customerID, "closed").
		Order("created_at DESC").
		First(&conversation).Error; err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (r *conversationRepository) Create(conversation *entity.Conversation) error {
	return r.db.Create(conversation).Error
}
