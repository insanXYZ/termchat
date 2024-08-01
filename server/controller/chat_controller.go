package controller

import (
	"backend/model"
	"backend/model/converter"
	"backend/service"
	"backend/utils/httpresponse"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ChatController struct {
	ChatService *service.ChatService
}

func NewChatController(chatService *service.ChatService) *ChatController {
	return &ChatController{ChatService: chatService}
}

func (controller *ChatController) GetChats(c echo.Context) error {
	claims := c.Get("user").(jwt.MapClaims)
	chats, err := controller.ChatService.GetChats(claims)
	if err != nil {
		return httpresponse.Error(c, err.Error(), nil)
	}

	res := make([]*model.WriteMessage, 0)
	for _, chat := range *chats {
		res = append(res, converter.ChatToWriteMessage(&chat))
	}

	return httpresponse.Success(c, "success get all chats", res)

}

func (controller *ChatController) WsChat(c echo.Context) error {
	claims := c.Get("user").(jwt.MapClaims)
	return controller.ChatService.Chat(claims, c.Response(), c.Request())
}
