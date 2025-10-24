package dto

// RegisterRequest represents user registration request
type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Division string `json:"division,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

// LoginRequest represents user login request
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// UpdateProfileRequest represents profile update request
type UpdateProfileRequest struct {
	Name     string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Division string `json:"division,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

// AuthResponse represents authentication response
type AuthResponse struct {
	Token string    `json:"token"`
	User  UserDTO   `json:"user"`
}

// UserDTO represents user data transfer object
type UserDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Division  string `json:"division,omitempty"`
	Phone     string `json:"phone,omitempty"`
	CreatedAt string `json:"created_at"`
}