package services

import (
	"errors"
	"fmt"

	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/repository"
	"gorm.io/gorm"
)

type ChannelService interface {
	HandleWebhook(tenantID, customerExternalID, conversationID, message string) error
}

type channelService struct {
	customerRepo     repository.CustomerRepository
	conversationRepo repository.ConversationRepository
	messageRepo      repository.MessageRepository
}

func NewChannelService(
	customerRepo repository.CustomerRepository,
	conversationRepo repository.ConversationRepository,
	messageRepo repository.MessageRepository,
) ChannelService {
	return &channelService{
		customerRepo:     customerRepo,
		conversationRepo: conversationRepo,
		messageRepo:      messageRepo,
	}
}

func (s *channelService) HandleWebhook(tenantID, customerExternalID, conversationID, message string) error {
	var targetConversation *entity.Conversation
	var err error

	// Mode 1: Direct to existing conversation
	if conversationID != "" {
		// When conversation_id is provided, we can get tenant_id from the conversation itself
		targetConversation, err = s.conversationRepo.FindByID(conversationID, tenantID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("conversation not found")
			}
			return err
		}
	} else {
		// Mode 2: Find or create customer, then find or create conversation
		// 1. Find or create customer
		customer, err := s.customerRepo.FindByExternalID(tenantID, customerExternalID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Create new customer
				customer = &entity.Customer{
					TenantID:   tenantID,
					ExternalID: customerExternalID,
					Name:       customerExternalID, // Use external ID as name initially
				}
				if err := s.customerRepo.Create(customer); err != nil {
					return err
				}
			} else {
				return err
			}
		}

		// 2. Find or create conversation
		targetConversation, err = s.conversationRepo.FindByCustomerID(tenantID, customer.ID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Create new conversation
				targetConversation = &entity.Conversation{
					TenantID:   tenantID,
					CustomerID: customer.ID,
					Status:     "open",
				}
				if err := s.conversationRepo.Create(targetConversation); err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	// 3. Create message
	newMessage := &entity.Message{
		ConversationID: targetConversation.ID,
		SenderType:     "customer",
		Message:        message,
	}
	if err := s.messageRepo.Create(newMessage); err != nil {
		return err
	}

	return nil
}
