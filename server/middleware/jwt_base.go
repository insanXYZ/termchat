package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (config *MiddlewareConfig) JwtBase() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SuccessHandler: func(c echo.Context) {
			c.Set("user", c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims))
		},
		SigningKey:    []byte(config.Viper.GetString("JWT_SECRET_KEY")),
		SigningMethod: echojwt.AlgorithmHS256,
	})
}
