package main

import (
	"log"
	database "med-manager/database"
	routes "med-manager/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create Fiber app
	app := fiber.New()

	// Add middleware
	app.Use(logger.New())

	// Setup routes
	routes.SetupRoutes(app, db)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
