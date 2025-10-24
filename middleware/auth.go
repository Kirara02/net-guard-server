package middleware

import (
	"NetGuardServer/repository"
	"NetGuardServer/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// JWTMiddleware validates JWT tokens for protected routes
func JWTMiddleware(c *fiber.Ctx) error {
	// Get token from Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header is required",
		})
	}

	// Check if it starts with "Bearer "
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid authorization format. Use 'Bearer <token>'",
		})
	}

	// Extract token
	token := strings.TrimPrefix(authHeader, "Bearer ")

	// Validate token
	claims, err := utils.ValidateJWT(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	// Set user information in context
	if userID, ok := claims["user_id"].(string); ok {
		c.Locals("user_id", userID)
	}
	if email, ok := claims["email"].(string); ok {
		c.Locals("email", email)
	}
	if name, ok := claims["name"].(string); ok {
		c.Locals("name", name)
	}

	return c.Next()
}

// AdminMiddleware checks if user has ADMIN role
func AdminMiddleware(c *fiber.Ctx) error {
	// First check if user is authenticated
	userID, ok := c.Locals("user_id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// Parse UUID and check user role from database
	uid, err := uuid.Parse(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	// Get user from database to check role
	userRepo := repository.NewUserRepository()
	user, err := userRepo.FindByID(uid)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Check if user is active
	if !user.IsActive {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Account is deactivated",
		})
	}

	// Check admin role
	if user.Role != "ADMIN" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Admin access required",
		})
	}

	// Set role in context for further use
	c.Locals("user_role", user.Role)

	return c.Next()
}