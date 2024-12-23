package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waenSaran/vending-machine-api/constant"
)

func GetProductList(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   "product list",
	})
}

func GetProductById(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   "product by id",
	})
}

func AddProduct(c *fiber.Ctx) error {
	user := c.Locals("user").(*constant.UserData)
	return c.JSON(fiber.Map{
		"status":    "success",
		"data":      "product added",
		"createdBy": user.Email,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   "product updated",
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   "product deleted",
	})
}
