package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waenSaran/vending-machine-api/app/models"
	"github.com/waenSaran/vending-machine-api/db"
)

// GetCategoryList func gets category list.
// @Summary get category list
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {array} models.Category
// @Router /category [get]
func GetCategoryList(c *fiber.Ctx) error {
	var category []models.Categories
	db.DB.Model(&models.Categories{}).Order("ID ASC").Find(&category)
	return c.JSON(category)
}
