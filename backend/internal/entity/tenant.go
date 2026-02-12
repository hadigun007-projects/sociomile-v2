package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tenant struct {
	ID    string `gorm:"type:varchar(36);primaryKey" json:"id"`
	Name  string `gorm:"type:varchar(255);not null" json:"name"`
	Plan  string `gorm:"type:varchar(50);default:'basic'" json:"plan"`
	Users []User `gorm:"foreignKey:TenantID" json:"users,omitempty"`
	Base
}

func (b *Tenant) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
