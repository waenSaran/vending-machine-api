package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waenSaran/vending-machine-api/app/models"
	"github.com/waenSaran/vending-machine-api/db"
)

func UpdateMoneyInStock(c *fiber.Ctx) error {
	moneyList := new([]models.Money)
	if err := c.BodyParser(&moneyList); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	for _, item := range *moneyList {
		if err := db.DB.Model(&models.Money{}).Where("id = ?", item.ID).Updates(item).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update item with ID " + string(item.ID),
			})
		}
	}
	return c.JSON(moneyList)
}
