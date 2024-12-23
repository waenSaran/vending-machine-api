package models

type LoginResponse struct {
	Token string
}

type UserData struct {
	Email string
	Name  string
	Role  string
}
