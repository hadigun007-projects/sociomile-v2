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

	userRepository := repository.NewUserRepository(r.db)
	userService := services.NewUserService(userRepository, r.cfg)
	userHandler := handler.NewUserHandler(userService)

	// get tenants
	privateRoute.GET("/tenants", tenantHandler.GetTenants)
	// create tenant
	privateRoute.POST("/tenants", tenantHandler.CreateTenant)
	// update tenant
	privateRoute.PUT("/tenants/:id", tenantHandler.UpdateTenant)
	// delete tenant
	privateRoute.DELETE("/tenants/:id", tenantHandler.DeleteTenant)
	// get users
	privateRoute.GET("/users", userHandler.GetUsers)
	// update user
	privateRoute.PUT("/users/:id", userHandler.UpdateUser)
	// delete user
	privateRoute.DELETE("/users/:id", userHandler.DeleteUser)
	// create user
	privateRoute.POST("/users", userHandler.CreateUser)
}
