package middleware

import (
	"backend/utils/httpresponse"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (config *MiddlewareConfig) QueryParamToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.QueryParam("token")
		if token == "" {
			return httpresponse.Error(c, "authorization failed", nil)
		}

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("wrong signing method")
			}

			return []byte(config.Viper.GetString("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			return httpresponse.Error(c, "authorization failed", nil)
		}

		c.Set("user", claims)
		return next(c)
	}
}
