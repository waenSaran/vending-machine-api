package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/waenSaran/vending-machine-api/constant"
)

func AuthMiddleware(app *fiber.App) {
	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(constant.Config["SECRET_KEY"])},
	}))

	// Middleware to extract user data from the JWT token
	app.Use(ExtractUserFromJWT)
}

func ExtractUserFromJWT(c *fiber.Ctx) error {
	user := &constant.UserData{}

	// Extract the token from the Fiber context (inserted by the JWT middleware)
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	user.Email = claims["email"].(string)
	user.Name = claims["name"].(string)
	user.Role = claims["role"].(string)

	// Store the user data in the Fiber context
	c.Locals("user", user)

	return c.Next()
}
