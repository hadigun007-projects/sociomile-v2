package handler

import (
	"net/http"

	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/services"
	"github.com/gin-gonic/gin"
)

type ConversationHandler struct {
	conversationService services.ConversationService
}

func NewConversationHandler(conversationService services.ConversationService) *ConversationHandler {
	return &ConversationHandler{conversationService: conversationService}
}

func (h *ConversationHandler) GetConversations(c *gin.Context) {
	tenantID, exists := c.Get("user_tenant_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Tenant ID not found in context"})
		return
	}

	conversations, err := h.conversationService.GetConversations(tenantID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": conversations})
}

func (h *ConversationHandler) GetConversationByID(c *gin.Context) {
	id := c.Param("id")
	tenantID, exists := c.Get("user_tenant_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Tenant ID not found in context"})
		return
	}

	conversation, err := h.conversationService.GetConversationByID(id, tenantID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": conversation})
}

type AssignAgentRequest struct {
	AgentID string `json:"agent_id" binding:"required"`
}

func (h *ConversationHandler) AssignAgent(c *gin.Context) {
	id := c.Param("id")
	tenantID, exists := c.Get("user_tenant_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Tenant ID not found in context"})
		return
	}

	var req AssignAgentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conversation, err := h.conversationService.AssignAgent(id, tenantID.(string), req.AgentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": conversation})
}

type ReplyMessageRequest struct {
	Message string `json:"message" binding:"required"`
}

func (h *ConversationHandler) ReplyToConversation(c *gin.Context) {
	id := c.Param("id")
	tenantID, exists := c.Get("user_tenant_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Tenant ID not found in context"})
		return
	}

	var req ReplyMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.conversationService.ReplyToConversation(id, tenantID.(string), req.Message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reply sent successfully"})
}
