package controllers

import (
	"log"
	"NetGuardServer/repository"
	"NetGuardServer/services"
	"NetGuardServer/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// ServerController handles server-related HTTP requests
type ServerController struct {
	serverService       services.ServerService
	notificationService services.NotificationService
}

// NewServerController creates a new server controller
func NewServerController(serverService services.ServerService, notificationService services.NotificationService) *ServerController {
	return &ServerController{
		serverService:       serverService,
		notificationService: notificationService,
	}
}

// CreateServerRequest represents the create server request payload
type CreateServerRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// UpdateServerRequest represents the update server request payload
type UpdateServerRequest struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

// CreateServer handles server creation
func (ctrl *ServerController) CreateServer(c *fiber.Ctx) error {
	// Get user ID from JWT
	userIDStr, ok := c.Locals("user_id").(string)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid user ID")
	}

	var req CreateServerRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Validate required fields
	if req.Name == "" || req.URL == "" {
		return utils.SendError(c, fiber.StatusBadRequest, "Name and URL are required")
	}

	server, err := ctrl.serverService.CreateServer(userID, req.Name, req.URL)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SendSuccess(c, "Server created successfully", server)
}

// GetServers handles getting all servers (all users)
func (ctrl *ServerController) GetServers(c *fiber.Ctx) error {
	servers, err := ctrl.serverService.GetAllServers()
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SendData(c, servers)
}

// GetServer handles getting a specific server by ID
func (ctrl *ServerController) GetServer(c *fiber.Ctx) error {
	serverIDStr := c.Params("id")
	serverID, err := uuid.Parse(serverIDStr)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid server ID")
	}

	// Get user ID from JWT for authorization
	userIDStr, ok := c.Locals("user_id").(string)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid user ID")
	}

	server, err := ctrl.serverService.GetServerByID(serverID)
	if err != nil {
		return utils.SendError(c, fiber.StatusNotFound, "Server not found")
	}

	// Check if user owns this server
	if server.CreatedBy != userID {
		return utils.SendError(c, fiber.StatusForbidden, "Access denied")
	}

	return utils.SendData(c, server)
}

// UpdateServer handles server updates
func (ctrl *ServerController) UpdateServer(c *fiber.Ctx) error {
	serverIDStr := c.Params("id")
	serverID, err := uuid.Parse(serverIDStr)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid server ID")
	}

	// Get user ID from JWT for authorization
	userIDStr, ok := c.Locals("user_id").(string)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid user ID")
	}

	var req UpdateServerRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	server, err := ctrl.serverService.UpdateServer(serverID, userID, req.Name, req.URL)
	if err != nil {
		if err.Error() == "server not found" {
			return utils.SendError(c, fiber.StatusNotFound, err.Error())
		}
		if err.Error() == "access denied" {
			return utils.SendError(c, fiber.StatusForbidden, err.Error())
		}
		return utils.SendError(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SendSuccess(c, "Server updated successfully", server)
}

// DeleteServer handles server deletion
func (ctrl *ServerController) DeleteServer(c *fiber.Ctx) error {
	serverIDStr := c.Params("id")
	serverID, err := uuid.Parse(serverIDStr)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid server ID")
	}

	// Get user ID from JWT for authorization
	userIDStr, ok := c.Locals("user_id").(string)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid user ID")
	}

	err = ctrl.serverService.DeleteServer(serverID, userID)
	if err != nil {
		if err.Error() == "server not found" {
			return utils.SendError(c, fiber.StatusNotFound, err.Error())
		}
		if err.Error() == "access denied" {
			return utils.SendError(c, fiber.StatusForbidden, err.Error())
		}
		return utils.SendError(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SendSuccess(c, "Server deleted successfully", nil)
}

// UpdateServerStatus handles server status updates from mobile clients
func (ctrl *ServerController) UpdateServerStatus(c *fiber.Ctx) error {
	serverIDStr := c.Params("id")
	serverID, err := uuid.Parse(serverIDStr)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid server ID")
	}

	// Get user ID from JWT
	userIDStr, ok := c.Locals("user_id").(string)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid user ID")
	}

	var req struct {
		Status       string `json:"status"`
		ResponseTime int64  `json:"response_time"`
	}

	if err := c.BodyParser(&req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if req.Status == "" {
		return utils.SendError(c, fiber.StatusBadRequest, "Status is required")
	}

	// Validate status
	if req.Status != "UP" && req.Status != "DOWN" && req.Status != "UNKNOWN" {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid status. Must be UP, DOWN, or UNKNOWN")
	}

	// Get server info for history
	server, err := ctrl.serverService.GetServerByID(serverID)
	if err != nil {
		return utils.SendError(c, fiber.StatusNotFound, "Server not found")
	}

	// If server is DOWN, create history record and send notifications
	if req.Status == "DOWN" {
		// Create history record
		historyRepo := repository.NewHistoryRepository()
		historyService := services.NewHistoryService(historyRepo)

		_, err := historyService.CreateHistory(serverID, server.Name, server.URL, "DOWN", userID)
		if err != nil {
			// Log error but don't fail the request
			log.Printf("ERROR: Failed to create history record for server %s: %v", serverID, err)
		} else {
			log.Printf("INFO: History record created for server DOWN: %s (%s)", server.Name, server.URL)
		}

		// Send FCM notification
		err = ctrl.notificationService.SendServerDownNotification(serverID, server.Name, server.URL, userID)
		if err != nil {
			// Log error but don't fail the request
			log.Printf("ERROR: Failed to send FCM notification for server %s: %v", serverID, err)
		} else {
			log.Printf("INFO: FCM notification sent for server DOWN: %s (%s)", server.Name, server.URL)
		}
	}

	return utils.SendSuccess(c, "Server status updated successfully", server)
}