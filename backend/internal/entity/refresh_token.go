package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshToken struct {
	ID        string    `gorm:"type:varchar(36);primaryKey"`
	UserID    string    `gorm:"not null;index" json:"user_id"`
	Token     string    `gorm:"uniqueIndex;not null;type:varchar(512)" json:"token"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	Revoked   bool      `gorm:"default:false" json:"revoked"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	Base
}

func (b *RefreshToken) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
