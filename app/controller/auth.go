package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waenSaran/vending-machine-api/app/services"
	_ "github.com/waenSaran/vending-machine-api/docs"
)

// Login func for back-office
// @Summary login
// @Tags Authentication
// @Accept json
// @Produce json
// @Param email query string true "Email"
// @Param password query string true "Password"
// @Success 200 {object} models.LoginResponse
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	return services.Login(c)
}
