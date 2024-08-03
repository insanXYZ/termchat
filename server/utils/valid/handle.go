package valid

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func HandleValidatorStruct(err error) error {
	if vErrors, ok := err.(validator.ValidationErrors); ok {
		var s string
		for _, err := range vErrors {
			s += fmt.Sprintf("the %s field requirement is %s \n", err.Field(), err.Tag())
		}
		err = errors.New(s)
	}

	return err

}
