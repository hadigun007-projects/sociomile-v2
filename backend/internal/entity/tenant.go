package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tenant struct {
	ID    string `gorm:"type:varchar(36);primaryKey"`
	Name  string `gorm:"type:varchar(255);not null"`
	Users []User `gorm:"foreignKey:TenantID"`
	Base
}

func (b *Tenant) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
