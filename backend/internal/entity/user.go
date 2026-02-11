package entity

type User struct {
	Base
	TenantID string `gorm:"type:varchar(36);not null;index"`
	Email    string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Role     string `gorm:"type:enum('admin', 'agent');not null"`
}
