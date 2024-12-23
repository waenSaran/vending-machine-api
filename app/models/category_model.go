package models

import "gorm.io/gorm"

type BaseCategories struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type BaseSubCategories struct {
	ID         int    `json:"id"`
	CategoryID int    `json:"category_id"`
	Name       string `json:"name"`
}

type Categories struct {
	gorm.Model
	BaseCategories
}

type SubCategories struct {
	gorm.Model
	BaseSubCategories
}
