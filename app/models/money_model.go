package models

import "gorm.io/gorm"

type BaseMoney struct {
	Value  int `json:"value"`
	Amount int `json:"amount"`
}

type Money struct {
	gorm.Model
	BaseMoney
}
