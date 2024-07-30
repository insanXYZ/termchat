package repository

import "backend/entity"

type ChatRepository struct {
	Repository[entity.Chat]
}

func NewChatRepository() *ChatRepository {
	return &ChatRepository{}
}
