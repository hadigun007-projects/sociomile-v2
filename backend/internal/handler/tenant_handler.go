package handler

import (
	"net/http"

	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
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
	var tenant entity.Tenant
	if err := c.ShouldBindJSON(&tenant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.tenantService.CreateTenant(&tenant); err != nil {
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
