package repository

import (
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
	"gorm.io/gorm"
)

type ConversationRepository interface {
	FindByCustomerID(tenantID, customerID string) (*entity.Conversation, error)
	Create(conversation *entity.Conversation) error
	GetByTenantID(tenantID string) ([]entity.Conversation, error)
	GetByTenantIDWithPagination(tenantID string, page, limit int) ([]entity.Conversation, int64, error)
	FindByID(id, tenantID string) (*entity.Conversation, error)
	Update(conversation *entity.Conversation) error
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

func (r *conversationRepository) GetByTenantID(tenantID string) ([]entity.Conversation, error) {
	var conversations []entity.Conversation
	if err := r.db.Where("tenant_id = ?", tenantID).
		Preload("Messages").
		Preload("AssignedAgent").
		Preload("Ticket").
		Find(&conversations).Error; err != nil {
		return nil, err
	}
	return conversations, nil
}

func (r *conversationRepository) GetByTenantIDWithPagination(tenantID string, page, limit int) ([]entity.Conversation, int64, error) {
	var conversations []entity.Conversation
	var total int64

	offset := (page - 1) * limit

	query := r.db.Model(&entity.Conversation{}).Where("tenant_id = ?", tenantID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Preload("Messages").Preload("AssignedAgent").Preload("Ticket").Order("created_at DESC").Offset(offset).Limit(limit).Find(&conversations).Error; err != nil {
		return nil, 0, err
	}

	return conversations, total, nil
}

func (r *conversationRepository) FindByID(id, tenantID string) (*entity.Conversation, error) {
	var conversation entity.Conversation
	query := r.db.Where("id = ?", id)

	// Only filter by tenant_id if it's provided
	if tenantID != "" {
		query = query.Where("tenant_id = ?", tenantID)
	}

	if err := query.
		Preload("Messages").
		Preload("AssignedAgent").
		Preload("Ticket").
		First(&conversation).Error; err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (r *conversationRepository) Update(conversation *entity.Conversation) error {
	return r.db.Save(conversation).Error
}
