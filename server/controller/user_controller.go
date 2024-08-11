package controller

import (
	"backend/model"
	"backend/model/converter"
	"backend/service"
	"backend/utils/httpresponse"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (controller *UserController) Register(c echo.Context) error {
	req := new(model.RegisterUser)
	err := c.Bind(req)
	if err != nil {
		return httpresponse.Error(c, err, nil)
	}
	err = controller.UserService.Register(req)
	if err != nil {
		return httpresponse.Error(c, err, nil)
	}
	return httpresponse.Success(c, "success register user", req)
}

func (controller *UserController) Login(c echo.Context) error {
	req := new(model.LoginUser)
	err := c.Bind(req)
	if err != nil {
		return httpresponse.Error(c, err, nil)
	}

	user, token, err := controller.UserService.Login(req)
	if err != nil {
		return httpresponse.Error(c, err, nil)
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = *token
	cookie.Path = "/"

	c.SetCookie(cookie)

	return httpresponse.Success(c, "success login", converter.UserToLogin(user, token))

}

func (controller *UserController) Refresh(c echo.Context) error {
	claims := c.Get("user").(jwt.MapClaims)

	token, err := controller.UserService.Refresh(claims)
	if err != nil {
		return httpresponse.Error(c, err, nil)
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = *token
	cookie.Path = "/"
	c.SetCookie(cookie)

	return httpresponse.Success(c, "success refresh token", converter.UserToToken(token))
}

func (controller *UserController) GetUser(c echo.Context) error {
	req := new(model.GetUser)
	err := c.Bind(req)
	if err != nil {
		return httpresponse.Error(c, err, nil)
	}
	users, err := controller.UserService.GetUser(req)
	if err != nil {
		return httpresponse.Error(c, err, nil)
	}

	res := make([]*model.UserResponse, len(*users))
	for i, user := range *users {
		res[i] = converter.UserToResponse(&user)
	}

	return httpresponse.Success(c, "success get user", res)
}

func (controller *UserController) UpdateUser(c echo.Context) error {
	claims := c.Get("user").(jwt.MapClaims)
	req := new(model.UpdateUser)
	err := c.Bind(req)
	if err != nil {
		return httpresponse.Error(c, err, nil)
	}
	user, err := controller.UserService.UpdateUser(claims, req)
	if err != nil {
		return httpresponse.Error(c, err, nil)
	}

	return httpresponse.Success(c, "success update user", converter.UserToResponse(user))
}
