package models

import "gorm.io/gorm"

type BaseCategories struct {
	Name string `json:"name"`
}

type BaseSubCategories struct {
	CategoryID int        `json:"category_id"`
	Category   Categories `json:"categories,omitempty" gorm:"foreignKey:CategoryID;references:ID"`
	Name       string     `json:"name"`
}

type Categories struct {
	gorm.Model
	BaseCategories
}

type SubCategories struct {
	gorm.Model
	BaseSubCategories
}

type CategoriesResponse struct {
	Category    string `json:"category"`
	SubCategory string `json:"sub_category"`
}
