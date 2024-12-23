package models

import "gorm.io/gorm"

type LoginResponse struct {
	Token string
}

type UserData struct {
	gorm.Model
	Email string
	Name  string
	Role  string
}
