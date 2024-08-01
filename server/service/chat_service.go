package service

import (
	"backend/entity"
	"backend/model"
	"backend/repository"
	"backend/service/websocket"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ChatService struct {
	DB             *gorm.DB
	Viper          *viper.Viper
	Validator      *validator.Validate
	Hub            *websocket.Hub
	UserRepository *repository.UserRepository
	ChatRepository *repository.ChatRepository
}

func NewChatService(DB *gorm.DB, viper *viper.Viper, validator *validator.Validate, userRepository *repository.UserRepository, chatRepository *repository.ChatRepository) *ChatService {
	hub := websocket.NewHub(DB, userRepository)
	go hub.Run()
	return &ChatService{DB: DB, Viper: viper, Validator: validator, Hub: hub, UserRepository: userRepository, ChatRepository: chatRepository}
}

func (service *ChatService) Chat(claims jwt.MapClaims, response http.ResponseWriter, request *http.Request) error {
	upgrade, err := websocket.Upgrader.Upgrade(response, request, nil)
	if err != nil {
		return err
	}

	user := &entity.User{
		ID: claims["sub"].(string),
	}

	err = service.UserRepository.Take(service.DB, user)
	if err != nil {
		return err
	}

	client := &websocket.Client{
		ChatRepository: service.ChatRepository,
		Hub:            service.Hub,
		User:           user,
		Conn:           upgrade,
		Send:           make(chan *model.SendMessage),
		DB:             service.DB,
	}

	client.Hub.Register <- client

	go client.ReadPump()
	go client.WritePump()

	return nil

}

func (service *ChatService) GetChats(claims jwt.MapClaims) (*[]entity.Chat, error) {
	chats := new([]entity.Chat)
	err := service.ChatRepository.GetChats(service.DB, claims["sub"].(string), chats)
	if err != nil {
		return nil, err
	}

	return chats, nil
}
