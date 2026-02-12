package main

import (
	"log"

	"github.com/cinnamorollofficials/sociomile-v2/backend/config"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/router"
	"github.com/cinnamorollofficials/sociomile-v2/backend/pkg/database"
)

func main() {
	cfg := config.LoadConfig()

	// 1. database connection
	db, err := database.NewMySQLConnection(&cfg)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	log.Println("Database connected successfully")

	// 2. database migration
	database.Migrate(db)

	// 3. database seeder
	database.SeedData(db, &cfg)

	// 4. run server
	r := router.NewRouter(&cfg, db)
	r.Run()
}
