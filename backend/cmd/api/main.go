package main

import (
	"fmt"

	"github.com/cinnamorollofficials/sociomile-v2/backend/config"
)

func main() {
	cfg := config.LoadConfig()

	// address := fmt.Sprintf(":%s", cfg.ServerPort)
	fmt.Printf("Server starting on port %s...\n", cfg.ServerPort)
}
