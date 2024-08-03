package model

type User struct {
	Name, Email, ID, Bio string
}

type UpdateUser struct {
	Name, Email, Password, Bio string
}
