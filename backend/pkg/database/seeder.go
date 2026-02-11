package database

import (
	"log"

	"github.com/cinnamorollofficials/sociomile-v2/backend/config"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedData(db *gorm.DB, cfg *config.Config) {
	tenant := entity.Tenant{
		Name: "Sociomile Enterprise",
	}
	if err := db.Where("name = ?", tenant.Name).FirstOrCreate(&tenant).Error; err != nil {
		log.Printf("Failed to seed tenant: %v", err)
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	admin := entity.User{
		TenantID: tenant.ID,
		Email:    cfg.AdminEmail,
		Role:     "admin",
	}
	if err := db.Where("email = ?", admin.Email).Attrs(entity.User{Password: string(hashedPassword)}).FirstOrCreate(&admin).Error; err != nil {
		log.Printf("Failed to seed admin: %v", err)
	}

	hashedAgentPassword, _ := bcrypt.GenerateFromPassword([]byte(cfg.AgentPassword), bcrypt.DefaultCost)

	agent := entity.User{
		TenantID: tenant.ID,
		Email:    cfg.AgentEmail,
		Role:     "agent",
	}
	if err := db.Where("email = ?", agent.Email).Attrs(entity.User{Password: string(hashedAgentPassword)}).FirstOrCreate(&agent).Error; err != nil {
		log.Printf("Failed to seed agent: %v", err)
	}

	log.Println("Seeding completed!")
	log.Printf("Login Admin: %s | %s", cfg.AdminEmail, cfg.AdminPassword)
	log.Printf("Login Agent: %s | %s", cfg.AgentEmail, cfg.AgentPassword)
}
