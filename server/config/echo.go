package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func NewEcho(viper *viper.Viper) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use()
	return e
}
