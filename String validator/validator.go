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
		return errors.New("Invalid email format")
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
