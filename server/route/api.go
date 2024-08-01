package route

import (
	"backend/controller"
	"backend/middleware"

	"github.com/labstack/echo/v4"
)

type RouteConfig struct {
	Echo           *echo.Echo
	Middlewares    *middleware.MiddlewareConfig
	UserController *controller.UserController
	ChatController *controller.ChatController
}

func (c *RouteConfig) Setup() {
	g := c.Echo.Group("/api")
	g.POST("/register", c.UserController.Register)
	g.POST("/login", c.UserController.Login)

	jwtGroup := g.Group("/")
	jwtGroup.Use(c.Middlewares.JwtBase())
	jwtGroup.GET("ws/chat", c.ChatController.WsChat)
	jwtGroup.GET("user", c.UserController.GetUser)
	jwtGroup.GET("chat", c.ChatController.GetChats)
}
