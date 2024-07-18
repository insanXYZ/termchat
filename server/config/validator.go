package config

import "github.com/go-playground/validator/v10"

func NewValidator() *validator.Validate {
	validate := validator.New()
	return validate
}
