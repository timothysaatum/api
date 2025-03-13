package main

import (
	"fmt"
	"log"
	"os"

	"api/database"
	"api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	// Initialize Fiber app
	app := fiber.New()

	// Connect to database
	database.ConnectDB()

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(app.Listen(":" + port))
}
