package controllers

import (
	"NetGuardServer/dto"
	"NetGuardServer/services"
	"NetGuardServer/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// AuthController handles authentication HTTP requests
type AuthController struct {
	authService services.AuthService
}

// NewAuthController creates a new auth controller
func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Register handles user registration
func (ctrl *AuthController) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Validation failed: "+err.Error())
	}

	user, token, err := ctrl.authService.Register(req.Name, req.Email, req.Password, req.Division, req.Phone, "USER")
	if err != nil {
		if appErr, ok := err.(utils.AppError); ok {
			switch appErr.Code {
			case "VALIDATION_ERROR":
				return utils.SendError(c, fiber.StatusBadRequest, appErr.Message)
			case "CONFLICT":
				return utils.SendError(c, fiber.StatusConflict, appErr.Message)
			default:
				return utils.SendError(c, fiber.StatusInternalServerError, appErr.Message)
			}
		}
		return utils.SendError(c, fiber.StatusInternalServerError, err.Error())
	}

	// Convert to DTO
	userDTO := dto.UserDTO{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		Division:  user.Division,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	response := dto.AuthResponse{
		Token: token,
		User:  userDTO,
	}

	return utils.SendSuccess(c, "User registered successfully", response)
}

// Login handles user authentication
func (ctrl *AuthController) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Validation failed: "+err.Error())
	}

	user, token, err := ctrl.authService.Login(req.Email, req.Password)
	if err != nil {
		if appErr, ok := err.(utils.AppError); ok {
			switch appErr.Code {
			case "VALIDATION_ERROR", "UNAUTHORIZED":
				return utils.SendError(c, fiber.StatusUnauthorized, appErr.Message)
			default:
				return utils.SendError(c, fiber.StatusInternalServerError, appErr.Message)
			}
		}
		return utils.SendError(c, fiber.StatusUnauthorized, err.Error())
	}

	// Convert to DTO
	userDTO := dto.UserDTO{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		Division:  user.Division,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	response := dto.AuthResponse{
		Token: token,
		User:  userDTO,
	}

	return utils.SendSuccess(c, "Login successful", response)
}

// GetProfile returns the current user's profile
func (ctrl *AuthController) GetProfile(c *fiber.Ctx) error {
	// Get user ID from JWT claims (set by middleware)
	userID, ok := c.Locals("user_id").(string)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	// Parse UUID
	uid, err := uuid.Parse(userID)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid user ID")
	}

	user, err := ctrl.authService.GetProfile(uid)
	if err != nil {
		if appErr, ok := err.(utils.AppError); ok && appErr.Code == "NOT_FOUND" {
			return utils.SendError(c, fiber.StatusNotFound, appErr.Message)
		}
		return utils.SendError(c, fiber.StatusInternalServerError, err.Error())
	}

	// Convert to DTO
	userDTO := dto.UserDTO{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		Division:  user.Division,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	return utils.SendData(c, userDTO)
}

// UpdateProfile handles profile updates
func (ctrl *AuthController) UpdateProfile(c *fiber.Ctx) error {
	// Get user ID from JWT claims (set by middleware)
	userID, ok := c.Locals("user_id").(string)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	// Parse UUID
	uid, err := uuid.Parse(userID)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid user ID")
	}

	var req dto.UpdateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Validation failed: "+err.Error())
	}

	user, err := ctrl.authService.UpdateProfile(uid, req.Name, req.Division, req.Phone)
	if err != nil {
		if appErr, ok := err.(utils.AppError); ok && appErr.Code == "NOT_FOUND" {
			return utils.SendError(c, fiber.StatusNotFound, appErr.Message)
		}
		return utils.SendError(c, fiber.StatusInternalServerError, err.Error())
	}

	// Convert to DTO
	userDTO := dto.UserDTO{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		Division:  user.Division,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	return utils.SendSuccess(c, "Profile updated successfully", userDTO)
}
