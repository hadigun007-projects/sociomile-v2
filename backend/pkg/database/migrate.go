package database

import (
	"log"

	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.RefreshToken{},
		&entity.Tenant{},
		&entity.Conversation{},
		&entity.Message{},
		&entity.Ticket{},
		&entity.RefreshToken{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database Migration Completed!")
}
