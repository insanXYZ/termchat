package middleware

import "github.com/spf13/viper"

type MiddlewareConfig struct {
	Viper *viper.Viper
}

func NewMiddleware(viper *viper.Viper) *MiddlewareConfig {
	return &MiddlewareConfig{Viper: viper}
}
