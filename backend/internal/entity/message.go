package entity

type Message struct {
	Base
	ConversationID string `gorm:"type:varchar(36);not null;index"`
	SenderType     string `gorm:"type:enum('customer', 'agent');not null"`
	Message        string `gorm:"type:text;not null"`
}
