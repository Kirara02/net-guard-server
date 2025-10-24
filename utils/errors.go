package utils

import "errors"

// Common application errors
var (
	ErrValidation     = errors.New("validation error")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
	ErrNotFound       = errors.New("not found")
	ErrConflict       = errors.New("conflict")
	ErrInternalServer = errors.New("internal server error")
)

// AppError represents a custom application error
type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e AppError) Error() string {
	return e.Message
}

// NewAppError creates a new application error
func NewAppError(code, message, details string) AppError {
	return AppError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

// ValidationError creates a validation error
func ValidationError(message string) AppError {
	return NewAppError("VALIDATION_ERROR", message, "")
}

// NotFoundError creates a not found error
func NotFoundError(resource string) AppError {
	return NewAppError("NOT_FOUND", "Resource not found", resource)
}

// UnauthorizedError creates an unauthorized error
func UnauthorizedError() AppError {
	return NewAppError("UNAUTHORIZED", "Unauthorized access", "")
}

// ForbiddenError creates a forbidden error
func ForbiddenError() AppError {
	return NewAppError("FORBIDDEN", "Access forbidden", "")
}

// ConflictError creates a conflict error
func ConflictError(message string) AppError {
	return NewAppError("CONFLICT", message, "")
}