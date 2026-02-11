package entity

type Tenant struct {
	Base
	Name  string `gorm:"type:varchar(255);not null"`
	Users []User `gorm:"foreignKey:TenantID"`
}
