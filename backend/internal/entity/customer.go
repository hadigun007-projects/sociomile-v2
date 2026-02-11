package entity

type Customer struct {
	Base
	TenantID   string `gorm:"type:varchar(36);not null;uniqueIndex:idx_customer_tenant_ext"`
	ExternalID string `gorm:"type:varchar(255);not null;uniqueIndex:idx_customer_tenant_ext"`
	Name       string `gorm:"type:varchar(255)"`
}
