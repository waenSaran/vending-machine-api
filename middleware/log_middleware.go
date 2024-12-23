package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware(app *fiber.App) {
	app.Use(durationLogging)
}

// loggingMiddleware logs the processing time for each request
func durationLogging(c *fiber.Ctx) error {
	// Start timer
	start := time.Now()

	// Process request
	err := c.Next()

	// Calculate processing time
	duration := time.Since(start)

	// Log the information
	fmt.Printf("Request URL: %s - Method: %s - Duration: %s\n", c.OriginalURL(), c.Method(), duration)

	return err
}
