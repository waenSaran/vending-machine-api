package services

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/waenSaran/vending-machine-api/constant"
)

// Dummy user for example
var dummyUser = struct {
	Name     string
	Email    string
	Role     string
	Password string
}{
	Name:     "John Doe",
	Email:    "user@example.com",
	Role:     "admin",
	Password: "password123",
}

func Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	pass := c.FormValue("password")

	// Throws Unauthorized error
	if email != dummyUser.Email || pass != dummyUser.Password {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"email": dummyUser.Email,
		"role":  "admin",
		"name":  dummyUser.Name,
		"exp":   time.Now().Add(time.Minute * 5).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte(constant.Config["SECRET_KEY"])

	// Generate encoded token and send it as response.
	t, err := token.SignedString(secretKey)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
