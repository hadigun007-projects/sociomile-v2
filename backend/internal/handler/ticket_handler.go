package handler

import (
	"math"
	"net/http"
	"strconv"

	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/services"
	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	ticketService services.TicketService
}

func NewTicketHandler(ticketService services.TicketService) *TicketHandler {
	return &TicketHandler{ticketService: ticketService}
}

func (h *TicketHandler) GetTickets(c *gin.Context) {
	tenantID, exists := c.Get("user_tenant_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Tenant ID not found in context"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	tickets, total, err := h.ticketService.GetTicketsWithPagination(tenantID.(string), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	c.JSON(http.StatusOK, gin.H{
		"data": tickets,
		"meta": gin.H{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}

type UpdateTicketStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=open in_progress resolved closed"`
}

func (h *TicketHandler) UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	tenantID, exists := c.Get("user_tenant_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Tenant ID not found in context"})
		return
	}

	var req UpdateTicketStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticket, err := h.ticketService.UpdateTicketStatus(id, tenantID.(string), req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ticket status updated successfully",
		"data":    ticket,
	})
}
