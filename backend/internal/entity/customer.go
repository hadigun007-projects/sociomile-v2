package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID         string `gorm:"type:varchar(36);primaryKey"`
	TenantID   string `gorm:"type:varchar(36);not null;uniqueIndex:idx_customer_tenant_ext"`
	ExternalID string `gorm:"type:varchar(255);not null;uniqueIndex:idx_customer_tenant_ext"`
	Name       string `gorm:"type:varchar(255)"`
	Base
}

func (b *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
