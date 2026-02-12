package services

import (
	"errors"
	"fmt"

	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/repository"
	"gorm.io/gorm"
)

type ConversationService interface {
	GetConversations(tenantID string) ([]entity.Conversation, error)
	GetConversationByID(id, tenantID string) (*entity.Conversation, error)
	AssignAgent(conversationID, tenantID, agentID string) (*entity.Conversation, error)
	ReplyToConversation(conversationID, tenantID, message string) error
}

type conversationService struct {
	conversationRepo repository.ConversationRepository
	messageRepo      repository.MessageRepository
}

func NewConversationService(
	conversationRepo repository.ConversationRepository,
	messageRepo repository.MessageRepository,
) ConversationService {
	return &conversationService{
		conversationRepo: conversationRepo,
		messageRepo:      messageRepo,
	}
}

func (s *conversationService) GetConversations(tenantID string) ([]entity.Conversation, error) {
	return s.conversationRepo.GetByTenantID(tenantID)
}

func (s *conversationService) GetConversationByID(id, tenantID string) (*entity.Conversation, error) {
	return s.conversationRepo.FindByID(id, tenantID)
}

func (s *conversationService) AssignAgent(conversationID, tenantID, agentID string) (*entity.Conversation, error) {
	conversation, err := s.conversationRepo.FindByID(conversationID, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("conversation not found")
		}
		return nil, err
	}

	conversation.AssignedAgentID = &agentID
	conversation.Status = "assigned"

	if err := s.conversationRepo.Update(conversation); err != nil {
		return nil, err
	}

	return conversation, nil
}

func (s *conversationService) ReplyToConversation(conversationID, tenantID, message string) error {
	// Verify conversation exists and belongs to tenant
	_, err := s.conversationRepo.FindByID(conversationID, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("conversation not found")
		}
		return err
	}

	// Create agent message
	newMessage := &entity.Message{
		ConversationID: conversationID,
		SenderType:     "agent",
		Message:        message,
	}

	return s.messageRepo.Create(newMessage)
}
