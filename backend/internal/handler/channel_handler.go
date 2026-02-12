package handler

import (
	"net/http"

	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/services"
	"github.com/gin-gonic/gin"
)

type ChannelHandler struct {
	channelService services.ChannelService
}

func NewChannelHandler(channelService services.ChannelService) *ChannelHandler {
	return &ChannelHandler{channelService: channelService}
}

type WebhookPayload struct {
	TenantID           string `json:"tenant_id"`
	CustomerExternalID string `json:"customer_external_id"`
	ConversationID     string `json:"conversation_id"`
	Message            string `json:"message" binding:"required"`
}

func (h *ChannelHandler) HandleWebhook(c *gin.Context) {
	var payload WebhookPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate based on mode
	if payload.ConversationID != "" {
		// Mode 1: Sending to existing conversation (only conversation_id needed)
		if err := h.channelService.HandleWebhook(payload.TenantID, payload.CustomerExternalID, payload.ConversationID, payload.Message); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		// Mode 2: Creating new conversation (need tenant_id and customer_external_id)
		if payload.TenantID == "" || payload.CustomerExternalID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "tenant_id and customer_external_id are required when conversation_id is not provided"})
			return
		}

		if err := h.channelService.HandleWebhook(payload.TenantID, payload.CustomerExternalID, payload.ConversationID, payload.Message); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	response := gin.H{
		"message": "Message received successfully",
		"data":    gin.H{},
	}

	if payload.CustomerExternalID != "" {
		response["data"].(gin.H)["customer_external_id"] = payload.CustomerExternalID
	}
	if payload.ConversationID != "" {
		response["data"].(gin.H)["conversation_id"] = payload.ConversationID
	}

	c.JSON(http.StatusOK, response)
}
