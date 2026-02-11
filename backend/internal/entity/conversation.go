package entity

type Conversation struct {
	Base
	TenantID        string    `gorm:"type:varchar(36);not null;index"`
	CustomerID      string    `gorm:"type:varchar(36);not null"`
	Status          string    `gorm:"type:enum('open', 'assigned', 'closed');default:'open';index"`
	AssignedAgentID *string   `gorm:"type:varchar(36)"`
	Messages        []Message `gorm:"foreignKey:ConversationID"`
	Ticket          *Ticket   `gorm:"foreignKey:ConversationID"`
}
