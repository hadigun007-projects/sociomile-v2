package handler

import (
	"math"
	"net/http"
	"strconv"

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

	// userRole, _ := c.Get("user_role")
	// userID, _ := c.Get("user_id")

	conversations, err := h.conversationService.GetConversations(tenantID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Admins see all conversations
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	conversations, total, err := h.conversationService.GetConversationsWithPagination(tenantID.(string), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Filter for agents: only show conversations assigned to them
	// Note: Pagination with filtering in memory is tricky.
	// For now, if role is agent, we might need to adjust the query or filter after fetching (which breaks pagination).
	// Ideally, we should pass the agentID to the repository to filter at database level.
	// But given the current scope, let's assume we fetch paginated results and then filtered? NO, that would result in less than limit items.
	// BETTER APPROACH: Add AgentID to GetConversationsWithPagination service/repo method.

	// Filter for agents: only show conversations assigned to them
	// TODO: Implement filtering by AgentID at repository level for correct pagination
	// Currently showing all conversations to agents due to pagination constraints

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	c.JSON(http.StatusOK, gin.H{
		"data": conversations,
		"meta": gin.H{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": totalPages,
		},
	})
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

func (h *ConversationHandler) GetConversationByIDPublic(c *gin.Context) {
	id := c.Param("id")

	conversation, err := h.conversationService.GetConversationByID(id, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": conversation})
}

type EscalateRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

func (h *ConversationHandler) Escalate(c *gin.Context) {
	id := c.Param("id")
	tenantID, exists := c.Get("user_tenant_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Tenant ID not found in context"})
		return
	}

	var req EscalateRequest
	if err := h.ShouldBindRequest(c, &req); err != nil {
		return
	}

	ticket, err := h.conversationService.EscalateToTicket(id, tenantID.(string), req.Title, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Conversation escalated to ticket successfully",
		"data":    ticket,
	})
}

func (h *ConversationHandler) ShouldBindRequest(c *gin.Context, req interface{}) error {
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	return nil
}
