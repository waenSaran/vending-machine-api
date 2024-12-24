package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waenSaran/vending-machine-api/app/controller"
	"github.com/waenSaran/vending-machine-api/app/middleware"
)

func ProductRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	route.Get("/products", controller.GetProductList)
	route.Get("/product/:id", controller.GetProductById)
	route.Put("/product/:id", controller.UpdateProduct)
}

func RestrictedProductRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	middleware.AuthMiddleware(app)

	route.Post("/product", controller.AddProduct)
	route.Delete("/product/:id", controller.DeleteProduct)
}
