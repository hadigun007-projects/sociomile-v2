package database

import (
	"log"

	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.Tenant{},
		&entity.User{},
		&entity.RefreshToken{},
		&entity.Customer{},
		&entity.Conversation{},
		&entity.Message{},
		&entity.Ticket{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	db.Exec("ALTER TABLE tickets MODIFY COLUMN status ENUM('requested', 'open', 'in_progress', 'resolved', 'closed') DEFAULT 'open'")

	log.Println("Database Migration Completed!")
}
