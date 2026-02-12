package handler

import (
	"net/http"

	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/handler/dto"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/services"
	"github.com/gin-gonic/gin"
)

type TenantHandler struct {
	tenantService services.TenantService
}

func NewTenantHandler(tenantService services.TenantService) *TenantHandler {
	return &TenantHandler{tenantService: tenantService}
}

func (h *TenantHandler) GetTenants(c *gin.Context) {
	tenants, err := h.tenantService.GetTenants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tenants})
}

func (h *TenantHandler) CreateTenant(c *gin.Context) {
	var req dto.CreateTenantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenant, err := h.tenantService.CreateTenant(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": tenant})
}

func (h *TenantHandler) UpdateTenant(c *gin.Context) {
	id := c.Param("id")
	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenant, err := h.tenantService.UpdateTenant(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tenant})
}

func (h *TenantHandler) DeleteTenant(c *gin.Context) {
	id := c.Param("id")
	if err := h.tenantService.DeleteTenant(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "tenant deleted successfully"})
}

func (h *TenantHandler) GetPublicTenants(c *gin.Context) {
	tenants, err := h.tenantService.GetTenants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return only ID and Name for public consumption
	var publicTenants []gin.H
	for _, tenant := range tenants {
		publicTenants = append(publicTenants, gin.H{
			"id":   tenant.ID,
			"name": tenant.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": publicTenants})
}
