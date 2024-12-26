package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waenSaran/vending-machine-api/app/models"
	"github.com/waenSaran/vending-machine-api/db"
)

// UpdateProductAmount func update a product amount.
// @Summary update only a product amount by id
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} models.Products
// @Router /product/{id}/amount [put]
func UpdateProductAmount(c *fiber.Ctx) error {
	id := c.Params("id")
	product := new(models.Products)
	db.DB.First(&product, id)

	if product.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Could not parse product data: " + err.Error(),
		})
	}

	if err := db.DB.Model(&product).Update("amount", product.Amount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update product amount: " + err.Error(),
		})
	}

	// Fetch the updated product
	if err := db.DB.First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not fetch updated product: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}
