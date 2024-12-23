package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waenSaran/vending-machine-api/services"
)

func Login(c *fiber.Ctx) error {
	return services.Login(c)
}
