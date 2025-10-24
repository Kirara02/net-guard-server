package routes

import (
	"NetGuardServer/di"
	"NetGuardServer/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App) {
	// Initialize all dependencies using Wire
	appContainer, err := di.InitializeApp()
	if err != nil {
		panic("Failed to initialize application: " + err.Error())
	}

	// API routes
	api := app.Group("/api")

	// Public routes
	auth := api.Group("/auth")
	auth.Post("/register", appContainer.AuthController.Register)
	auth.Post("/login", appContainer.AuthController.Login)

	// Protected routes
	protected := api.Group("", middleware.JWTMiddleware)
	protected.Get("/auth/me", appContainer.AuthController.GetProfile)
	protected.Put("/auth/profile", appContainer.AuthController.UpdateProfile)

	// Server routes
	servers := protected.Group("/servers")
	servers.Post("", appContainer.ServerController.CreateServer)
	servers.Get("", appContainer.ServerController.GetServers)
	servers.Get("/:id", appContainer.ServerController.GetServer)
	servers.Put("/:id", appContainer.ServerController.UpdateServer)
	servers.Delete("/:id", appContainer.ServerController.DeleteServer)
	servers.Patch("/:id/status", appContainer.ServerController.UpdateServerStatus)

	// History routes
	history := protected.Group("/history")
	history.Get("", appContainer.HistoryController.GetHistory)
	history.Patch("/:id/resolve", appContainer.HistoryController.ResolveHistory)
	history.Get("/report/monthly", appContainer.HistoryController.GetMonthlyReport)
}
