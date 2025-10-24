package services

import (
	"NetGuardServer/models"
	"NetGuardServer/repository"
	"NetGuardServer/utils"
	"errors"
	"strings"

	"github.com/google/uuid"
)

// AuthService defines the interface for authentication business logic
type AuthService interface {
	Register(name, email, password, division, phone, role string) (*models.User, string, error)
	Login(email, password string) (*models.User, string, error)
	GetProfile(userID uuid.UUID) (*models.User, error)
	UpdateProfile(userID uuid.UUID, name, division, phone string) (*models.User, error)
}

// authService implements AuthService
type authService struct {
	userRepo repository.UserRepository
}

// NewAuthService creates a new auth service instance
func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

// Register handles user registration business logic
func (s *authService) Register(name, email, password, division, phone, role string) (*models.User, string, error) {
	// Validate input
	if strings.TrimSpace(name) == "" {
		return nil, "", errors.New("name is required")
	}
	if strings.TrimSpace(email) == "" {
		return nil, "", errors.New("email is required")
	}
	if len(password) < 6 {
		return nil, "", errors.New("password must be at least 6 characters long")
	}

	// Check if user already exists
	existingUser, err := s.userRepo.FindByEmail(strings.ToLower(email))
	if err == nil && existingUser != nil {
		return nil, "", errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, "", errors.New("failed to hash password")
	}

	// Create user
	user := &models.User{
		Name:         strings.TrimSpace(name),
		Email:        strings.ToLower(strings.TrimSpace(email)),
		PasswordHash: hashedPassword,
		Division:     division,
		Phone:        phone,
		Role:         role,
		IsActive:     true,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, "", errors.New("failed to create user")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(*user)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	// Remove password hash from response
	user.PasswordHash = ""

	return user, token, nil
}

// Login handles user authentication business logic
func (s *authService) Login(email, password string) (*models.User, string, error) {
	// Validate input
	if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" {
		return nil, "", errors.New("email and password are required")
	}

	// Find user
	user, err := s.userRepo.FindByEmail(strings.ToLower(email))
	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	// Check password
	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return nil, "", errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(*user)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	// Remove password hash from response
	user.PasswordHash = ""

	return user, token, nil
}

// GetProfile handles getting user profile business logic
func (s *authService) GetProfile(userID uuid.UUID) (*models.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Remove password hash from response
	user.PasswordHash = ""

	return user, nil
}

// UpdateProfile handles updating user profile business logic
func (s *authService) UpdateProfile(userID uuid.UUID, name, division, phone string) (*models.User, error) {
	// Get current user
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Update fields if provided
	if name != "" {
		user.Name = name
	}
	if division != "" {
		user.Division = division
	}
	if phone != "" {
		user.Phone = phone
	}

	// Save updated user
	if err := s.userRepo.Update(user); err != nil {
		return nil, errors.New("failed to update profile")
	}

	// Remove password hash from response
	user.PasswordHash = ""

	return user, nil
}