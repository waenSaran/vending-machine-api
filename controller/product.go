package controller

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/waenSaran/vending-machine-api/docs"
	"github.com/waenSaran/vending-machine-api/models"
)

// GetProductList func gets all products.
// @Summary get all products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func GetProductList(c *fiber.Ctx) error {
	products := []models.Product{
		{
			ID:       1,
			Name:     "Product 1",
			Category: "Category 1",
			Price:    100,
			Quantity: 10,
			ImageURL: "https://example.com/product1.jpg",
		},
		{
			ID:       2,
			Name:     "Product 2",
			Category: "Category 2",
			Price:    200,
			Quantity: 20,
		},
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   products,
	})
}

// GetProductById func gets a product.
// @Summary get a product by id
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} models.Product
// @Router /product/{id} [get]
func GetProductById(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   "product by id",
	})
}

// AddProduct func adds a product.
// @Summary add a product
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} models.AddProductResponse
// @Router /product [post]
func AddProduct(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.UserData)
	reponse := models.AddProductResponse{
		Status:    "success",
		Data:      "product added",
		CreatedBy: user.Email,
	}
	if err := c.BodyParser(&reponse); err != nil {
		return err
	}

	return c.JSON(reponse)
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
