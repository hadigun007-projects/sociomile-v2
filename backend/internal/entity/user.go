package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"type:varchar(36);primaryKey" json:"id"`
	TenantID string `gorm:"type:varchar(36);not null;index" json:"tenant_id"`
	Email    string `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
	Role     string `gorm:"type:enum('owner', 'admin', 'agent');not null" json:"role"`
	Base
}

func (b *User) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
