package controllers

import (
	"NetGuardServer/services"
	"NetGuardServer/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// HistoryController handles history-related HTTP requests
type HistoryController struct {
	historyService services.HistoryService
}

// NewHistoryController creates a new history controller
func NewHistoryController(historyService services.HistoryService) *HistoryController {
	return &HistoryController{
		historyService: historyService,
	}
}

// CreateHistoryRequest represents the create history request payload
type CreateHistoryRequest struct {
	ServerID   string `json:"server_id"`
	ServerName string `json:"server_name"`
	URL        string `json:"url"`
	Status     string `json:"status"`
}

// CreateHistory handles history record creation
func (ctrl *HistoryController) CreateHistory(c *fiber.Ctx) error {
	// Get user ID from JWT
	userIDStr, ok := c.Locals("user_id").(string)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid user ID")
	}

	var req CreateHistoryRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Validate required fields
	if req.ServerID == "" || req.ServerName == "" || req.URL == "" || req.Status == "" {
		return utils.SendError(c, fiber.StatusBadRequest, "ServerID, ServerName, URL, and Status are required")
	}

	// Parse server ID
	serverID, err := uuid.Parse(req.ServerID)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid server ID")
	}

	// Validate status
	if req.Status != "DOWN" {
		return utils.SendError(c, fiber.StatusBadRequest, "Status must be DOWN for history records")
	}

	history, err := ctrl.historyService.CreateHistory(serverID, req.ServerName, req.URL, req.Status, userID)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SendSuccess(c, "History record created successfully", history)
}

// ResolveHistoryRequest represents the resolve history request payload
type ResolveHistoryRequest struct {
	ResolveNote string `json:"resolve_note"`
}

// ResolveHistory handles history record resolution
func (ctrl *HistoryController) ResolveHistory(c *fiber.Ctx) error {
	historyIDStr := c.Params("id")
	historyID, err := uuid.Parse(historyIDStr)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid history ID")
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

	var req ResolveHistoryRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if req.ResolveNote == "" {
		return utils.SendError(c, fiber.StatusBadRequest, "Resolve note is required")
	}

	history, err := ctrl.historyService.ResolveHistory(historyID, userID, req.ResolveNote)
	if err != nil {
		if err.Error() == "history record not found" {
			return utils.SendError(c, fiber.StatusNotFound, err.Error())
		}
		if err.Error() == "history record already resolved" {
			return utils.SendError(c, fiber.StatusConflict, err.Error())
		}
		return utils.SendError(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SendSuccess(c, "History record resolved successfully", history)
}

// GetMonthlyReport handles monthly report generation
func (ctrl *HistoryController) GetMonthlyReport(c *fiber.Ctx) error {
	yearStr := c.Query("year")
	monthStr := c.Query("month")

	if yearStr == "" || monthStr == "" {
		return utils.SendError(c, fiber.StatusBadRequest, "Year and month parameters are required")
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil || year < 2020 || year > 2030 {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid year")
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil || month < 1 || month > 12 {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid month")
	}

	report, err := ctrl.historyService.GetMonthlyReport(year, month)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SendSuccess(c, "Monthly report generated successfully", fiber.Map{
		"year":   year,
		"month":  month,
		"report": report,
	})
}

// GetHistory handles getting all history records
func (ctrl *HistoryController) GetHistory(c *fiber.Ctx) error {
	// Check if server_id is provided
	serverIDStr := c.Query("server_id")
	limitStr := c.Query("limit", "50")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 50
	}

	// Limit maximum records
	if limit > 1000 {
		limit = 1000
	}

	var histories interface{}
	if serverIDStr != "" {
		// Get history for specific server
		serverID, err := uuid.Parse(serverIDStr)
		if err != nil {
			return utils.SendError(c, fiber.StatusBadRequest, "Invalid server ID")
		}

		histories, err = ctrl.historyService.GetHistoryByServerID(serverID)
		if err != nil {
			return utils.SendError(c, fiber.StatusInternalServerError, err.Error())
		}
	} else {
		// Get all history records from all users
		histories, err = ctrl.historyService.GetAllHistory(limit)
		if err != nil {
			return utils.SendError(c, fiber.StatusInternalServerError, err.Error())
		}
	}

	return utils.SendData(c, histories)
}