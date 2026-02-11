package entity

type Ticket struct {
	Base
	TenantID        string  `gorm:"type:varchar(36);not null;index"`
	ConversationID  string  `gorm:"type:varchar(36);unique;not null"`
	Title           string  `gorm:"type:varchar(255);not null"`
	Description     string  `gorm:"type:text"`
	Status          string  `gorm:"type:enum('open', 'in_progress', 'resolved', 'closed');default:'open'"`
	Priority        string  `gorm:"type:enum('low', 'medium', 'high', 'urgent');default:'medium'"`
	AssignedAgentID *string `gorm:"type:varchar(36)"`
}
