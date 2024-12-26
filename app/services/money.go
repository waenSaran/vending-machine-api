package services

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/waenSaran/vending-machine-api/app/models"
	"github.com/waenSaran/vending-machine-api/db"
)

func UpdateMoneyInStock(c *fiber.Ctx) error {
	moneyList := new([]models.Money)

	// Parse the request body
	if err := c.BodyParser(&moneyList); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	updatedList := []models.Money{}

	for _, item := range *moneyList {
		// Update the record
		if err := db.DB.Model(&item).Update("amount", item.Amount).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update item with ID " + fmt.Sprintf("%d", item.ID),
			})
		}

		// Fetch the updated record
		updatedItem := models.Money{}
		if err := db.DB.Where("id = ?", item.ID).First(&updatedItem).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch updated item with ID " + fmt.Sprintf("%d", item.ID),
			})
		}

		updatedList = append(updatedList, updatedItem)
	}

	return c.Status(fiber.StatusOK).JSON(updatedList)
}
