package middleware

import (
	"backend/utils/httpresponse"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"strings"
	"time"
)

func (config *MiddlewareConfig) Refresh(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")
		if header == "" {
			return httpresponse.Error(c, "wrong header", nil)
		}

		tokenParts := strings.Split(header, " ")
		if len(tokenParts) != 2 || strings.EqualFold(tokenParts[0], "Bearer") {
			return httpresponse.Error(c, "wrong format authorization", nil)
		}

		token := tokenParts[1]

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("wrong signing method")
			}

			return []byte(config.Viper.GetString("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			if int64(claims["exp"].(float64)) <= time.Now().Unix() {
				c.Set("user", jwt.MapClaims{
					"sub":  claims["sub"],
					"name": claims["name"],
				})
				return next(c)
			}
			return httpresponse.Error(c, err.Error(), nil, 400)
		}

		return httpresponse.Error(c, "token not expired", nil, 400)

	}
}
