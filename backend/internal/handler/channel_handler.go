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
	TenantID           string `json:"tenant_id" binding:"required"`
	CustomerExternalID string `json:"customer_external_id" binding:"required"`
	Message            string `json:"message" binding:"required"`
}

func (h *ChannelHandler) HandleWebhook(c *gin.Context) {
	var payload WebhookPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.channelService.HandleWebhook(payload.TenantID, payload.CustomerExternalID, payload.Message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Message received successfully",
		"data": gin.H{
			"tenant_id":            payload.TenantID,
			"customer_external_id": payload.CustomerExternalID,
		},
	})
}
