package stringvalidator

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"
	"strings"
)

var (
	upperRegex    = regexp.MustCompile(`[A-Z]`)
	lowerRegex    = regexp.MustCompile(`[a-z]`)
	numberRegex   = regexp.MustCompile(`[0-9]`)
	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
)

// This package  is used to validate strings for various formats
func IsEmail(email string) error {
	//trim white space
	email = strings.TrimSpace(email)
	if email == "" {
		return errors.New("email cannot be empty")
	}
	//parsing the email
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("invalid email format")
	}
	return nil

}
func IsEmpty(s string, customMessage string) error {
	//trim white space
	s = strings.TrimSpace(s)
	if s == "" {
		if customMessage != "" {
			return errors.New(customMessage)
		}
		return errors.New("string cannot be empty")
	}
	return nil
}

// checking for the minimum length of the string
func Minlength(s string, minLength int, customMessage string) error {
	//trim white space
	s = strings.TrimSpace(s)
	if len(s) < minLength {
		if customMessage != "" {
			return errors.New(customMessage)
		}
		return fmt.Errorf("string must be at least %d characters long, got %d", minLength, len(s))
	}
	return nil
}
func Maxlength(s string, maxLength int, customMessage string) error {
	//trim white space
	s = strings.TrimSpace(s)
	if len(s) > maxLength {
		if customMessage != "" {
			return errors.New(customMessage)
		}
		return fmt.Errorf("string must be at most %d characters long, got %d", maxLength, len(s))
	}
	return nil
}

// HasLengthRange checks if string is within min and max length
func HasLengthRange(s string, minLength, maxLength int, customMessage string) error {
	s = strings.TrimSpace(s)
	length := len(s)

	if length < minLength {
		if customMessage != "" {
			return errors.New(customMessage)
		}
		return fmt.Errorf("string must be at least %d characters long, got %d", minLength, length)
	}
	if length > maxLength {
		if customMessage != "" {
			return errors.New(customMessage)
		}
		return fmt.Errorf("string must be at most %d characters long, got %d", maxLength, length)
	}
	return nil
}

// ValidatePhoneNumber validates phone number format
func ValidatePhoneNumber(phone string, customMessage string) error {
	phone = strings.TrimSpace(phone)
	phoneRegex := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	if !phoneRegex.MatchString(phone) {
		if customMessage != "" {
			return errors.New(customMessage)
		}
		return errors.New("invalid phone number format")
	}
	return nil
}

// ValidateURL validates URL format
func ValidateURL(url string, customMessage string) error {
	url = strings.TrimSpace(url)
	urlRegex := regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`)
	if !urlRegex.MatchString(url) {
		if customMessage != "" {
			return errors.New(customMessage)
		}
		return errors.New("invalid URL format")
	}
	return nil
}

// ValidatePassword validates password strength
func ValidatePassword(password string, customMessage string) error {
	password = strings.TrimSpace(password)

	if len(password) < 8 {
		if customMessage != "" {
			return errors.New(customMessage)
		}
		return errors.New("password must be at least 8 characters")
	}

	if !upperRegex.MatchString(password) {
		if customMessage != "" {
			return errors.New(customMessage)
		}
		return errors.New("password must contain uppercase letter")
	}

	if !lowerRegex.MatchString(password) {
		if customMessage != "" {
			return errors.New(customMessage)
		}
		return errors.New("password must contain lowercase letter")
	}

	if !numberRegex.MatchString(password) {
		if customMessage != "" {
			return errors.New(customMessage)
		}
		return errors.New("password must contain number")
	}

	return nil
}

// ValidateUsername validates username format
func ValidateUsername(username string, customMessage string) error {
	username = strings.TrimSpace(username)
	if len(username) < 3 {
		if customMessage != "" {
			return errors.New(customMessage)
		}
		return errors.New("username must be at least 3 characters")
	}
	if !usernameRegex.MatchString(username) {
		if customMessage != "" {
			return errors.New(customMessage)
		}
		return errors.New("username can only contain letters, numbers, and underscores")
	}
	return nil
}

// IsAlphanumeric checks if string contains only letters and numbers
func IsAlphanumeric(s string, customMessage string) error {
	s = strings.TrimSpace(s)
	alphanumericRegex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !alphanumericRegex.MatchString(s) {
		if customMessage != "" {
			return errors.New(customMessage)
		}
		return errors.New("string must contain only letters and numbers")
	}
	return nil
}

// ValidateID validates various ID formats (UUID, ObjectID, numeric ID)
func ValidateID(id string, customMessage string) error {
	id = strings.TrimSpace(id)

	if id == "" {
		if customMessage != "" {
			return errors.New(customMessage)
		}
		return errors.New("ID cannot be empty")
	}

	// Check for UUID format (8-4-4-4-12 hex digits)
	uuidRegex := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
	if uuidRegex.MatchString(id) {
		return nil
	}

	// Check for MongoDB ObjectID format (24 hex characters)
	objectIdRegex := regexp.MustCompile(`^[0-9a-fA-F]{24}$`)
	if objectIdRegex.MatchString(id) {
		return nil
	}

	// Check for numeric ID (positive integers)
	numericIdRegex := regexp.MustCompile(`^[1-9]\d*$`)
	if numericIdRegex.MatchString(id) {
		return nil
	}

	if customMessage != "" {
		return errors.New(customMessage)
	}
	return errors.New("invalid ID format - must be UUID, ObjectID, or numeric")
}
