package main

import (
	"log"

	"NetGuardServer/config"
	"NetGuardServer/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using environment vars")
	}

	// Load configuration
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("❌ Failed to load configuration: %v", err)
	}

	log.Println("✅ Database connected and migrated successfully")

	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Setup routes
	routes.SetupRoutes(app)

	// Route sederhana
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🚀 NetGuard Backend is running!")
	})

	// Jalankan server
	log.Printf("✅ Server running at http://localhost:%s", config.AppConfig.Server.Port)
	log.Fatal(app.Listen(":" + config.AppConfig.Server.Port))
}
