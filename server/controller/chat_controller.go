package controller

import (
	"backend/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ChatController struct {
	ChatService *service.ChatService
}

func NewChatController(chatService *service.ChatService) *ChatController {
	return &ChatController{ChatService: chatService}
}

//
//func (controller *ChatController) GetChat(c echo.Context) error {
//
//}

func (controller *ChatController) WsChat(c echo.Context) error {
	claims := c.Get("user").(jwt.MapClaims)
	return controller.ChatService.Chat(claims, c.Response(), c.Request())
}
