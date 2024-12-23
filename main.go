package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/waenSaran/vending-machine-api/middleware"
	"github.com/waenSaran/vending-machine-api/routes"
)

func main() {
	app := fiber.New()

	// loading env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Auth Routes
	routes.AuthRoutes(app)

	// Middleware
	middleware.LoggingMiddleware(app)
	middleware.CorsMiddleware(app)

	routes.ProductRoutes(app)
	routes.RestrictedProductRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	app.Listen(":" + port)
}
