package models

import (
	"gorm.io/gorm"
)

type BaseProducts struct {
	Name          string `json:"name"`
	SubCategoryId int    `json:"sub_category_id"`
	Price         int    `json:"price"`
	Amount        int    `json:"amount"`
	ImageURL      string `json:"image_url"`
	UpdatedBy     string `json:"updated_by"`
}

type Products struct {
	gorm.Model
	BaseProducts
}
