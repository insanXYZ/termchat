package repository

import (
	"backend/entity"
	"gorm.io/gorm"
)

type ChatRepository struct {
	Repository[entity.Chat]
}

func NewChatRepository() *ChatRepository {
	return &ChatRepository{}
}

func (r *ChatRepository) GetChats(db *gorm.DB, id string, chats *[]entity.Chat) error {
	return db.Joins("Sender").Joins("Receiver").Where("sender_id = ? or receiver_id = ?", id, id).Find(chats).Error

}
