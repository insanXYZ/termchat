package valid

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func HandleValidatorStruct(err error) error {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var e string
		for _, validationError := range validationErrors {
			if validationError.Param() == "" {
				e += fmt.Sprintf("%s field must be %s \n", validationError.Field(), validationError.Tag())
			} else {
				e += fmt.Sprintf("%s must have a %s of %s characters \n", validationError.Field(), validationError.Tag(), validationError.Param())
			}
		}

		err = errors.New(e)
	}

	return err

}
