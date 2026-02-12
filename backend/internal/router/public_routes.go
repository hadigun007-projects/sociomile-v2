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
	userRepository := repository.NewUserRepository(r.db)
	refreshTokenRepository := repository.NewRefreshTokenRepository(r.db)
	authService := services.NewAuthService(userRepository, r.cfg)
	jwtService := services.NewJWTService(r.cfg.JWTSecret, 24, refreshTokenRepository, userRepository)
	authHandler := handler.NewAuthHandler(authService, jwtService)

	// auth routes
	auth := router.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
	}

	// channel webhook (simulated external source)
	customerRepository := repository.NewCustomerRepository(r.db)
	conversationRepository := repository.NewConversationRepository(r.db)
	messageRepository := repository.NewMessageRepository(r.db)
	channelService := services.NewChannelService(customerRepository, conversationRepository, messageRepository)
	channelHandler := handler.NewChannelHandler(channelService)

	channel := router.Group("/channel")
	{
		channel.POST("/webhook", channelHandler.HandleWebhook)
	}

	// public tenant list (for simulator dropdown)
	tenantRepository := repository.NewTenantRepository(r.db)
	tenantService := services.NewTenantService(tenantRepository, userRepository, r.cfg)
	tenantHandler := handler.NewTenantHandler(tenantService)

	router.GET("/tenants/public", tenantHandler.GetPublicTenants)

	// public conversation endpoint (for simulator)
	ticketRepository := repository.NewTicketRepository(r.db)
	conversationService := services.NewConversationService(conversationRepository, messageRepository, ticketRepository)
	conversationHandler := handler.NewConversationHandler(conversationService)

	router.GET("/conversations/:id/public", conversationHandler.GetConversationByIDPublic)
}
