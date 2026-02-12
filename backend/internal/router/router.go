package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cinnamorollofficials/sociomile-v2/backend/config"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewRouter(cfg *config.Config, db *gorm.DB) *Router {
	return &Router{cfg: cfg, db: db}
}

func (r *Router) SetupRouter() *gin.Engine {

	router := gin.New()

	// middleware
	router.Use(middleware.CORS(*r.cfg))
	router.Use(middleware.LoggerMiddleware("storage/logs"))

	// Setup public routes
	r.setupPublicRoutes(router)

	// private routes
	r.setupPrivateRoutes(router)

	return router
}

func (r *Router) Run() {
	server := &http.Server{
		Addr:           ":" + r.cfg.ServerPort,
		Handler:        r.SetupRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	go func() {
		fmt.Printf("🚀 Server running on port %s\n", r.cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("❌ Server failed to start")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Server forced to shutdown")
	}

	fmt.Println("Server exited successfully")
}
