package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/waenSaran/vending-machine-api/app/middleware"
	"github.com/waenSaran/vending-machine-api/db"
	_ "github.com/waenSaran/vending-machine-api/docs"
	"github.com/waenSaran/vending-machine-api/routes"
)

// @title Vending Machine System API
// @version 1.0
// @description This is a Vending Machine System API for BluePi assignment
// @termsOfService http://swagger.io/terms/
// @contact.name Saranya
// @contact.email saranya.hiransatakun@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app := fiber.New()

	// loading env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to database
	db.InitDB()

	// Swagger Route
	routes.SwaggerRoutes(app)

	// Auth Routes
	routes.AuthRoutes(app)

	// Middleware
	middleware.LoggingMiddleware(app)
	middleware.CorsMiddleware(app)

	routes.CategoryRoutes(app)
	routes.MoneyRoutes(app)
	routes.ProductRoutes(app)
	routes.RestrictedProductRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	app.Listen(":" + port)
}
