package converter

import (
	"backend/entity"
	"backend/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Bio:   user.Bio,
	}
}

func UserToLogin(user *entity.User, token *string) *model.UserResponse {
	response := UserToResponse(user)
	response.Token = *token
	return response
}

func UserToToken(token *string) *model.UserResponse {
	return &model.UserResponse{
		Token: *token,
	}
}
