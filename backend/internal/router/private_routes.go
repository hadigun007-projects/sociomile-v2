package router

import (
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/handler"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/middleware"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/repository"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func (r *Router) setupPrivateRoutes(router *gin.Engine) {

	// private route middleware
	privateRoute := router.Group("/")
	privateRoute.Use(middleware.APIKeyMiddleware(r.cfg.APIKey))

	// handler
	tenantRepository := repository.NewTenantRepository(r.db)
	tenantService := services.NewTenantService(tenantRepository, r.cfg)
	tenantHandler := handler.NewTenantHandler(tenantService)

	// get tenants
	privateRoute.GET("/tenants", tenantHandler.GetTenants)
}
