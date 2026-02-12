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
	EscalateToTicket(conversationID, tenantID, title, description string) (*entity.Ticket, error)
	GetConversationsWithPagination(tenantID string, page, limit int) ([]entity.Conversation, int64, error)
}

type conversationService struct {
	conversationRepo repository.ConversationRepository
	messageRepo      repository.MessageRepository
	ticketRepo       repository.TicketRepository
}

func NewConversationService(
	conversationRepo repository.ConversationRepository,
	messageRepo repository.MessageRepository,
	ticketRepo repository.TicketRepository,
) ConversationService {
	return &conversationService{
		conversationRepo: conversationRepo,
		messageRepo:      messageRepo,
		ticketRepo:       ticketRepo,
	}
}

func (s *conversationService) GetConversations(tenantID string) ([]entity.Conversation, error) {
	return s.conversationRepo.GetByTenantID(tenantID)
}

func (s *conversationService) GetConversationsWithPagination(tenantID string, page, limit int) ([]entity.Conversation, int64, error) {
	return s.conversationRepo.GetByTenantIDWithPagination(tenantID, page, limit)
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

func (s *conversationService) EscalateToTicket(conversationID, tenantID, title, description string) (*entity.Ticket, error) {
	// 1. Verify conversation exists and belongs to tenant
	conversation, err := s.conversationRepo.FindByID(conversationID, tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("conversation not found")
		}
		return nil, err
	}

	// 2. Check if ticket already exists for this conversation
	// (Note: Ticket has unique constraint on ConversationID)

	// 3. Create Ticket with status 'requested'
	ticket := &entity.Ticket{
		TenantID:       tenantID,
		ConversationID: conversationID,
		Title:          title,
		Description:    description,
		Status:         "requested",
		Priority:       "medium",
	}

	if conversation.AssignedAgentID != nil {
		ticket.AssignedAgentID = conversation.AssignedAgentID
	}

	if err := s.ticketRepo.Create(ticket); err != nil {
		return nil, err
	}

	return ticket, nil
}
