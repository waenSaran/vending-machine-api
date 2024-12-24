package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waenSaran/vending-machine-api/app/controller"
)

func CategoryRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	route.Get("/category", controller.GetCategoryList)
}
