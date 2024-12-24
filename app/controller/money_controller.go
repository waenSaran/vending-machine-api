package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waenSaran/vending-machine-api/app/models"
	"github.com/waenSaran/vending-machine-api/app/services"
	"github.com/waenSaran/vending-machine-api/db"
)

// GetMoneyList func gets money list.
// @Summary get money list
// @Tags Money
// @Accept json
// @Produce json
// @Success 200 {array} models.Money
// @Router /money [get]
func GetMoneyList(c *fiber.Ctx) error {
	var money []models.Money
	db.DB.Model(&models.Money{}).Order("ID ASC").Find(&money)
	return c.JSON(money)
}

// UpdateMoneyInStock func update money in stock.
// @Summary update money in stock
// @Tags Money
// @Accept json
// @Produce json
// @Success 200 {array} models.Money
// @Router /money [put]
func UpdateMoneyInStock(c *fiber.Ctx) error {
	return services.UpdateMoneyInStock(c)
}
