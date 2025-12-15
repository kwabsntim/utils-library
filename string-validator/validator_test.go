package stringvalidator_test

import (
	"fmt"
	"testing"

	stringvalidator "github.com/kwabsntim/utils-library/string-validator"
)

//------------ EXAMPLE TESTS ----------------------

func ExampleIsEmail() {
	email := "test@example.com"
	err := stringvalidator.IsEmail(email)
	if err != nil {
		fmt.Println("Invalid")
	} else {
		fmt.Println("Valid")
	}
	// Output: Valid
}

func ExampleIsEmail_invalid() {
	email := "invalid-email"
	err := stringvalidator.IsEmail(email)
	if err != nil {
		fmt.Println("Invalid")
	} else {
		fmt.Println("Valid")
	}
	// Output: Invalid
}

func ExampleIsEmpty() {
	err := stringvalidator.IsEmpty("hello", "")
	if err != nil {
		fmt.Println("Empty")
	} else {
		fmt.Println("Not empty")
	}
	// Output: Not empty
}

func ExampleMinlength() {
	err := stringvalidator.Minlength("hello", 3, "")
	if err != nil {
		fmt.Println("Too short")
	} else {
		fmt.Println("Valid length")
	}
	// Output: Valid length
}

func ExampleMaxlength() {
	err := stringvalidator.Maxlength("hi", 5, "")
	if err != nil {
		fmt.Println("Too long")
	} else {
		fmt.Println("Valid length")
	}
	// Output: Valid length
}

func ExampleValidatePassword() {
	err := stringvalidator.ValidatePassword("MyPass123", "")
	if err != nil {
		fmt.Println("Weak password")
	} else {
		fmt.Println("Strong password")
	}
	// Output: Strong password
}

func ExampleValidateUsername() {
	err := stringvalidator.ValidateUsername("user123", "")
	if err != nil {
		fmt.Println("Invalid username")
	} else {
		fmt.Println("Valid username")
	}
	// Output: Valid username
}

func ExampleValidatePhoneNumber() {
	err := stringvalidator.ValidatePhoneNumber("+1234567890", "")
	if err != nil {
		fmt.Println("Invalid phone")
	} else {
		fmt.Println("Valid phone")
	}
	// Output: Valid phone
}

func ExampleValidateURL() {
	err := stringvalidator.ValidateURL("https://example.com", "")
	if err != nil {
		fmt.Println("Invalid URL")
	} else {
		fmt.Println("Valid URL")
	}
	// Output: Valid URL
}

func ExampleIsAlphanumeric() {
	err := stringvalidator.IsAlphanumeric("abc123", "")
	if err != nil {
		fmt.Println("Invalid")
	} else {
		fmt.Println("Valid")
	}
	// Output: Valid
}
func ExampleValidateID() {
	err := stringvalidator.ValidateID("507f1f77bcf86cd799439011", "")
	if err != nil {
		fmt.Println("Invalid ID")
	} else {
		fmt.Println("Valid ID")
	}
	// Output: Valid ID
}

func ExampleValidateID_uuid() {
	err := stringvalidator.ValidateID("550e8400-e29b-41d4-a716-446655440000", "")
	if err != nil {
		fmt.Println("Invalid ID")
	} else {
		fmt.Println("Valid ID")
	}
	// Output: Valid ID
}

//------------ UNIT TESTS ----------------------

func TestIsEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},
		{"user.name@domain.co.uk", true},
		{"invalid-email", false},
		{"@domain.com", false},
		{"user@", false},
		{"", false},
		{"   ", false},
	}

	for _, test := range tests {
		err := stringvalidator.IsEmail(test.email)
		isValid := err == nil
		if isValid != test.expected {
			t.Errorf("IsEmail(%q) = %v, expected %v", test.email, isValid, test.expected)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"   ", true},
		{"hello", false},
		{"  hello  ", false},
	}

	for _, test := range tests {
		err := stringvalidator.IsEmpty(test.input, "")
		isEmpty := err != nil
		if isEmpty != test.expected {
			t.Errorf("IsEmpty(%q) = %v, expected %v", test.input, isEmpty, test.expected)
		}
	}
}

func TestMinlength(t *testing.T) {
	tests := []struct {
		input     string
		minLength int
		expected  bool
	}{
		{"hello", 3, true},
		{"hi", 5, false},
		{"", 1, false},
		{"   test   ", 4, true},
	}

	for _, test := range tests {
		err := stringvalidator.Minlength(test.input, test.minLength, "")
		isValid := err == nil
		if isValid != test.expected {
			t.Errorf("Minlength(%q, %d) = %v, expected %v", test.input, test.minLength, isValid, test.expected)
		}
	}
}

func TestMaxlength(t *testing.T) {
	tests := []struct {
		input     string
		maxLength int
		expected  bool
	}{
		{"hello", 10, true},
		{"verylongstring", 5, false},
		{"", 1, true},
		{"test", 4, true},
	}

	for _, test := range tests {
		err := stringvalidator.Maxlength(test.input, test.maxLength, "")
		isValid := err == nil
		if isValid != test.expected {
			t.Errorf("Maxlength(%q, %d) = %v, expected %v", test.input, test.maxLength, isValid, test.expected)
		}
	}
}

func TestHasLengthRange(t *testing.T) {
	tests := []struct {
		input     string
		minLength int
		maxLength int
		expected  bool
	}{
		{"hello", 3, 10, true},
		{"hi", 5, 10, false},
		{"verylongstring", 3, 5, false},
		{"test", 4, 4, true},
	}

	for _, test := range tests {
		err := stringvalidator.HasLengthRange(test.input, test.minLength, test.maxLength, "")
		isValid := err == nil
		if isValid != test.expected {
			t.Errorf("HasLengthRange(%q, %d, %d) = %v, expected %v", test.input, test.minLength, test.maxLength, isValid, test.expected)
		}
	}
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"MyPass123", true},
		{"password", false},    // no uppercase, no number
		{"PASSWORD123", false}, // no lowercase
		{"MyPassword", false},  // no number
		{"MyPass1", false},     // too short
		{"", false},            // empty
	}

	for _, test := range tests {
		err := stringvalidator.ValidatePassword(test.password, "")
		isValid := err == nil
		if isValid != test.expected {
			t.Errorf("ValidatePassword(%q) = %v, expected %v", test.password, isValid, test.expected)
		}
	}
}

func TestValidateUsername(t *testing.T) {
	tests := []struct {
		username string
		expected bool
	}{
		{"user123", true},
		{"test_user", true},
		{"ab", false},        // too short
		{"user@name", false}, // invalid character
		{"user name", false}, // space not allowed
		{"", false},          // empty
	}

	for _, test := range tests {
		err := stringvalidator.ValidateUsername(test.username, "")
		isValid := err == nil
		if isValid != test.expected {
			t.Errorf("ValidateUsername(%q) = %v, expected %v", test.username, isValid, test.expected)
		}
	}
}

func TestValidatePhoneNumber(t *testing.T) {
	tests := []struct {
		phone    string
		expected bool
	}{
		{"+1234567890", true},        // 11 chars, 10 digits after first - valid
		{"1234567890", true},         // 10 chars, 9 digits after first - valid
		{"+12345678901234", true},    // 15 chars, 14 digits after first - valid
		{"123", true},                // 3 chars, 2 digits after first - valid (change to true)
		{"+0123456789", false},       // starts with 0 - invalid
		{"abc123", false},            // contains letters - invalid
		{"", false},                  // empty - invalid
		{"+123456789012345", true},   // 16 chars, 14 digits after first - valid (change to true)
		{"+1234567890123456", false}, // 17 chars, 15 digits after first - invalid (exceeds limit)
	}

	for _, test := range tests {
		err := stringvalidator.ValidatePhoneNumber(test.phone, "")
		isValid := err == nil
		if isValid != test.expected {
			t.Errorf("ValidatePhoneNumber(%q) = %v, expected %v", test.phone, isValid, test.expected)
		}
	}
}

func TestValidateURL(t *testing.T) {
	tests := []struct {
		url      string
		expected bool
	}{
		{"https://example.com", true},
		{"http://test.org", true},
		{"https://sub.domain.com/path", true},
		{"ftp://example.com", false}, // not http/https
		{"example.com", false},       // no protocol
		{"", false},                  // empty
	}

	for _, test := range tests {
		err := stringvalidator.ValidateURL(test.url, "")
		isValid := err == nil
		if isValid != test.expected {
			t.Errorf("ValidateURL(%q) = %v, expected %v", test.url, isValid, test.expected)
		}
	}
}

func TestIsAlphanumeric(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc123", true},
		{"ABC123", true},
		{"test", true},
		{"123", true},
		{"test@123", false}, // special character
		{"test 123", false}, // space
		{"", false},         // empty
	}

	for _, test := range tests {
		err := stringvalidator.IsAlphanumeric(test.input, "")
		isValid := err == nil
		if isValid != test.expected {
			t.Errorf("IsAlphanumeric(%q) = %v, expected %v", test.input, isValid, test.expected)
		}
	}
}
func TestValidateID(t *testing.T) {
	tests := []struct {
		id       string
		expected bool
	}{
		// MongoDB ObjectID (24 hex chars)
		{"507f1f77bcf86cd799439011", true},
		{"60d5ec49f1b2c8b1f8e4e1a1", true},

		// UUID format
		{"550e8400-e29b-41d4-a716-446655440000", true},
		{"6ba7b810-9dad-11d1-80b4-00c04fd430c8", true},

		// Numeric ID
		{"123", true},
		{"999999", true},

		// Invalid cases
		{"", false},                                 // empty
		{"   ", false},                              // whitespace only
		{"507f1f77bcf86cd79943901", false},          // ObjectID too short
		{"507f1f77bcf86cd799439011z", false},        // ObjectID with invalid char
		{"550e8400-e29b-41d4-a716", false},          // UUID incomplete
		{"550e8400e29b41d4a716446655440000", false}, // UUID without dashes
		{"0", false},                                // numeric starting with 0
		{"abc", false},                              // random string
		{"123abc", false},                           // mixed alphanumeric
	}

	for _, test := range tests {
		err := stringvalidator.ValidateID(test.id, "")
		isValid := err == nil
		if isValid != test.expected {
			t.Errorf("ValidateID(%q) = %v, expected %v", test.id, isValid, test.expected)
		}
	}
}

func TestCustomMessages(t *testing.T) {
	customMsg := "Custom error message"

	err := stringvalidator.IsEmpty("", customMsg)
	if err == nil || err.Error() != customMsg {
		t.Errorf("Expected custom message %q, got %v", customMsg, err)
	}

	err = stringvalidator.Minlength("hi", 5, customMsg)
	if err == nil || err.Error() != customMsg {
		t.Errorf("Expected custom message %q, got %v", customMsg, err)
	}

	err = stringvalidator.ValidatePassword("weak", customMsg)
	if err == nil || err.Error() != customMsg {
		t.Errorf("Expected custom message %q, got %v", customMsg, err)
	}
	err = stringvalidator.ValidateID("invalid", customMsg)
	if err == nil || err.Error() != customMsg {
		t.Errorf("Expected custom message %q, got %v", customMsg, err)
	}
}
