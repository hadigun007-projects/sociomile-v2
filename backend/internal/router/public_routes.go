package router

import (
	"net/http"

	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/handler"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/repository"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func (r *Router) setupPublicRoutes(router *gin.Engine) {

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "UP",
			"app":     "Social Mile",
			"version": "2.0.0",
		})
	})

	// auth handler
	authRepository := repository.NewUserRepository(r.db)
	authService := services.NewAuthService(authRepository, r.cfg)
	authHandler := handler.NewAuthHandler(authService)

	// auth routes
	auth := router.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
	}
}
