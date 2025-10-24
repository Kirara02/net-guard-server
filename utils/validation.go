package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validate is the global validator instance
var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

// ValidateStruct validates a struct using validator tags
func ValidateStruct(s interface{}) error {
	return Validate.Struct(s)
}

// Custom validation functions
func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	return len(password) >= 6
}

func ValidateUUID(fl validator.FieldLevel) bool {
	uuid := fl.Field().String()
	return len(uuid) == 36 && strings.Contains(uuid, "-")
}

// SanitizeString removes leading/trailing whitespace
func SanitizeString(s string) string {
	return strings.TrimSpace(s)
}

// ValidateEmail validates email format
func ValidateEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// ValidateRequired validates required fields
func ValidateRequired(value string) bool {
	return strings.TrimSpace(value) != ""
}