package converter

import (
	"backend/entity"
	"backend/model"
)

func UserToLogin(user *entity.User, token *string) *model.UserResponse {
	return &model.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Bio:   user.Bio,
		Token: *token,
	}
}

func UserToToken(token *string) *model.UserResponse {
	return &model.UserResponse{
		Token: *token,
	}
}

func UserToGet(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Bio:   user.Bio,
		Email: user.Email,
	}
}

func UserToChatUser(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:   user.ID,
		Name: user.Name,
	}
}
