package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waenSaran/vending-machine-api/app/controller"
)

func MoneyRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	route.Get("/money", controller.GetMoneyList)
	route.Put("/money", controller.UpdateMoneyInStock)
}
