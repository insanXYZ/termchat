package entity

import (
	"gorm.io/gorm"
	"time"
)

type Chat struct {
	ID         int       `gorm:"primaryKey;column:id;autoIncrement"`
	Message    string    `gorm:"column:message"`
	SenderID   string    `gorm:"column:sender_id"`
	ReceiverID string    `gorm:"column:receiver_id"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	Sender     *User     `gorm:"foreignKey:sender_id;references:id"`
	Receiver   *User     `gorm:"foreignKey:receiver_id;references:id"`
}

func (c Chat) TableName(db *gorm.DB) string {
	return "chats"
}
