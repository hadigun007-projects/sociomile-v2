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
	privateRoute.Use(middleware.JWTAuth(r.cfg.JWTSecret))

	// handler
	userRepository := repository.NewUserRepository(r.db)
	userService := services.NewUserService(userRepository, r.cfg)
	userHandler := handler.NewUserHandler(userService)

	tenantRepository := repository.NewTenantRepository(r.db)
	tenantService := services.NewTenantService(tenantRepository, userRepository, r.cfg)
	tenantHandler := handler.NewTenantHandler(tenantService)

	// tenants routes (owner only)
	tenantRoutes := privateRoute.Group("/tenants")
	tenantRoutes.Use(middleware.RBACMiddleware("owner"))
	{
		tenantRoutes.GET("", tenantHandler.GetTenants)
		tenantRoutes.POST("", tenantHandler.CreateTenant)
		tenantRoutes.PUT("/:id", tenantHandler.UpdateTenant)
		tenantRoutes.DELETE("/:id", tenantHandler.DeleteTenant)
	}
	// get users
	privateRoute.GET("/users", userHandler.GetUsers)
	// update user
	privateRoute.PUT("/users/:id", userHandler.UpdateUser)
	// delete user
	privateRoute.DELETE("/users/:id", userHandler.DeleteUser)
	// create user
	privateRoute.POST("/users", userHandler.CreateUser)
}
