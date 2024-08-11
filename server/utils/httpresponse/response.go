package httpresponse

import (
	"backend/model"
	"backend/utils/valid"

	"github.com/labstack/echo/v4"
)

func response(c echo.Context, statuscode int, message string, data any) error {
	return c.JSON(statuscode, model.Response{
		Message: message,
		Data:    data,
	})
}

func Success(c echo.Context, message string, data any, statuscode ...int) error {
	if len(statuscode) == 0 {
		return response(c, 200, message, data)
	}
	return response(c, statuscode[0], message, data)
}

func Error(c echo.Context, message error, data any, statuscode ...int) error {
	if len(statuscode) == 0 {
		statuscode = append(statuscode, 400)
	}
	return response(c, statuscode[0], valid.HandleValidatorStruct(message).Error(), data)

}
