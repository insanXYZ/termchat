package model

type UserResponse struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Bio   string `json:"bio,omitempty"`
	Token string `json:"token,omitempty"`
}

type RegisterUser struct {
	Name     string `json:"name" validate:"max=8,required"`
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"min=8,required"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"min=8,required"`
}

type UpdateUser struct {
	Name     string `json:"name" validate:"omitempty,max=8"`
	Email    string `json:"email" validate:"omitempty,email"`
	Bio      string `json:"bio" validate:"omitempty,max=20"`
	Password string `json:"password" validate:"omitempty,min=8"`
}

type GetUser struct {
	ID   string `query:"id" validate:"required_if=Name ''"`
	Name string `query:"name" validate:"required_if=ID ''"`
}
