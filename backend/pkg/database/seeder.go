package database

import (
	"log"

	"github.com/cinnamorollofficials/sociomile-v2/backend/config"
	"github.com/cinnamorollofficials/sociomile-v2/backend/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedData(db *gorm.DB, cfg *config.Config) {
	// Common password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	// 1. Create Tenant Alpha
	tenantAlpha := entity.Tenant{
		Name: "Tenant Alpha",
		Plan: "enterprise",
	}
	if err := db.Where("name = ?", tenantAlpha.Name).FirstOrCreate(&tenantAlpha).Error; err != nil {
		log.Printf("Failed to seed Tenant Alpha: %v", err)
	}

	// 2. Create Tenant Beta
	tenantBeta := entity.Tenant{
		Name: "Tenant Beta",
		Plan: "basic",
	}
	if err := db.Where("name = ?", tenantBeta.Name).FirstOrCreate(&tenantBeta).Error; err != nil {
		log.Printf("Failed to seed Tenant Beta: %v", err)
	}

	// 3. Create Owner (Assign to Tenant Alpha for now, or could be a system tenant)
	owner := entity.User{
		TenantID: tenantAlpha.ID,
		Email:    cfg.OwnerEmail,
		Role:     "owner",
	}
	if err := db.Where("email = ?", owner.Email).Attrs(entity.User{Password: string(hashedPassword)}).FirstOrCreate(&owner).Error; err != nil {
		log.Printf("Failed to seed owner: %v", err)
	}

	// 4. Create Users for Tenant Alpha
	adminAlpha := entity.User{
		TenantID: tenantAlpha.ID,
		Email:    cfg.AdminAlphaEmail,
		Role:     "admin",
	}
	if err := db.Where("email = ?", adminAlpha.Email).Attrs(entity.User{Password: string(hashedPassword)}).FirstOrCreate(&adminAlpha).Error; err != nil {
		log.Printf("Failed to seed admin alpha: %v", err)
	}

	agentAlpha := entity.User{
		TenantID: tenantAlpha.ID,
		Email:    cfg.AgentAlphaEmail,
		Role:     "agent",
	}
	if err := db.Where("email = ?", agentAlpha.Email).Attrs(entity.User{Password: string(hashedPassword)}).FirstOrCreate(&agentAlpha).Error; err != nil {
		log.Printf("Failed to seed agent alpha: %v", err)
	}

	// 5. Create Users for Tenant Beta
	adminBeta := entity.User{
		TenantID: tenantBeta.ID,
		Email:    cfg.AdminBetaEmail,
		Role:     "admin",
	}
	if err := db.Where("email = ?", adminBeta.Email).Attrs(entity.User{Password: string(hashedPassword)}).FirstOrCreate(&adminBeta).Error; err != nil {
		log.Printf("Failed to seed admin beta: %v", err)
	}

	agentBeta := entity.User{
		TenantID: tenantBeta.ID,
		Email:    cfg.AgentBetaEmail,
		Role:     "agent",
	}
	if err := db.Where("email = ?", agentBeta.Email).Attrs(entity.User{Password: string(hashedPassword)}).FirstOrCreate(&agentBeta).Error; err != nil {
		log.Printf("Failed to seed agent beta: %v", err)
	}

	log.Println("Seeding completed!")
	log.Println("---------------------------------------------------")
	log.Printf("Login Owner:      %s       | %s", cfg.OwnerEmail, cfg.OwnerPassword)
	log.Println("---------------------------------------------------")
	log.Println("Tenant Alpha (Enterprise):")
	log.Printf("  Admin:          %s | %s", cfg.AdminAlphaEmail, cfg.AdminAlphaPassword)
	log.Printf("  Agent:          %s | %s", cfg.AgentAlphaEmail, cfg.AgentAlphaPassword)
	log.Println("---------------------------------------------------")
	log.Println("Tenant Beta (Basic):")
	log.Printf("  Admin:          %s  | %s", cfg.AdminBetaEmail, cfg.AdminBetaPassword)
	log.Printf("  Agent:          %s  | %s", cfg.AgentBetaEmail, cfg.AgentBetaPassword)
	log.Println("---------------------------------------------------")
}
