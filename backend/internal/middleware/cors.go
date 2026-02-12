package middleware

import (
	"time"

	"github.com/cinnamorollofficials/sociomile-v2/backend/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS(cfg config.Config) gin.HandlerFunc {
	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{cfg.AllowedMethods},
		AllowHeaders:     []string{cfg.AllowedHeaders},
		ExposeHeaders:    []string{cfg.ExposeHeaders},
		AllowCredentials: cfg.AllowCredentials,
		MaxAge:           time.Duration(cfg.MaxAge),
	}

	return cors.New(config)
}
