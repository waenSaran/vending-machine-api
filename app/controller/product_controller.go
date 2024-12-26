package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waenSaran/vending-machine-api/app/models"
	"github.com/waenSaran/vending-machine-api/app/services"
	"github.com/waenSaran/vending-machine-api/db"
	_ "github.com/waenSaran/vending-machine-api/docs"
)

// GetProductList func gets all products.
// @Summary get all products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {array} models.ProductResponse
// @Router /products [get]
func GetProductList(c *fiber.Ctx) error {
	var products []models.ProductsResponse
	db.DB.Table("products").Select("products.*", "categories.id as category_id").Joins("JOIN sub_categories ON sub_categories.id = products.sub_category_id").Joins("JOIN categories ON categories.id = sub_categories.category_id").Order("products.ID ASC").Find(&products)
	return c.JSON(products)
}

// GetProductById func gets a product.
// @Summary get a product by id
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} models.ProductResponse
// @Router /product/{id} [get]
func GetProductById(c *fiber.Ctx) error {
	var product models.ProductsResponse
	id := c.Params("id")
	db.DB.Table("products").Select("products.*", "categories.id as category_id").Joins("JOIN sub_categories ON sub_categories.id = products.sub_category_id").Joins("JOIN categories ON categories.id = sub_categories.category_id").First(&product, id)

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
// @Router /product [put]
func UpdateProduct(c *fiber.Ctx) error {
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

	if err := db.DB.Save(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update product: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(product)
}

func UpdateProductAmount(c *fiber.Ctx) error {
	return services.UpdateProductAmount(c)
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

	if err := db.DB.Delete(&product, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not delete product: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(product)
}
