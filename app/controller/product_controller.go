package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waenSaran/vending-machine-api/app/models"
	"github.com/waenSaran/vending-machine-api/db"
	_ "github.com/waenSaran/vending-machine-api/docs"
)

// GetProductList func gets all products.
// @Summary get all products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func GetProductList(c *fiber.Ctx) error {
	var products []models.Products
	db.DB.Model(&models.Products{}).Find(&products)
	return c.JSON(products)
}

// GetProductById func gets a product.
// @Summary get a product by id
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} models.Product
// @Router /product/{id} [get]
func GetProductById(c *fiber.Ctx) error {
	var product models.Products
	id := c.Params("id")
	db.DB.Model(&product).First(&product, id)

	if product.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	return c.JSON(product)
}

// AddProduct func adds a product.
// @Summary add a product
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} models.Products
// @Router /product [post]
func AddProduct(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.UserData)
	product := new(models.Products)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Could not parse product data: " + err.Error(),
		})
	}

	// Set the createdBy field using the user data
	product.UpdatedBy = user.Email

	if err := db.DB.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create product: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(product)
}

// UpdateProduct func update a product.
// @Summary update a product by id
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} models.Products
// @Router /product [post]
func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product := new(models.Products)
	user := c.Locals("user").(*models.UserData)
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

	product.UpdatedBy = user.Email

	if err := db.DB.Save(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create product: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product := new(models.Products)
	db.DB.First(&product, id)

	if product.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	if err := db.DB.Delete(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not delete product: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(product)
}
