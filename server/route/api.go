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
	g.GET("/chat", c.ChatController.Chat, c.Middlewares.QueryParamToken)

	user := g.Group("/user")
	user.Use(c.Middlewares.JwtBase())
	user.GET("", c.UserController.GetUser)

}
