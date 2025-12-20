# String Validator Library

A comprehensive Go library for validating various string formats including emails, passwords, usernames, phone numbers, URLs, and IDs.

## Installation

```bash
go get github.com/kwabsntim/utils-library/string-validator
```

## Import

```go
import stringvalidator "github.com/kwabsntim/utils-library/string-validator"
```

## Quick Start

```go
package main

import (
    "fmt"
    stringvalidator "github.com/kwabsntim/utils-library/string-validator"
)

func main() {
    // Validate email
    if err := stringvalidator.IsEmail("user@example.com"); err != nil {
        fmt.Println("Invalid email:", err)
    } else {
        fmt.Println("Valid email!")
    }
    
    // Validate password
    if err := stringvalidator.ValidatePassword("MyPass123", ""); err != nil {
        fmt.Println("Weak password:", err)
    } else {
        fmt.Println("Strong password!")
    }
}
```

## API Reference

### Email Validation

#### `IsEmail(email string) error`
Validates email format using Go's built-in mail parser.

**Parameters:**
- `email` (string): Email address to validate

**Returns:** `error` - nil if valid, error message if invalid

**Example:**
```go
err := stringvalidator.IsEmail("test@example.com")
if err != nil {
    fmt.Println("Invalid email")
}
```

---

### String Length Validation

#### `IsEmpty(s string, customMessage string) error`
Checks if string is empty or contains only whitespace.

**Parameters:**
- `s` (string): String to check
- `customMessage` (string): Custom error message (optional, use "" for default)

**Returns:** `error` - error if empty, nil if not empty

**Example:**
```go
err := stringvalidator.IsEmpty("hello", "Field cannot be empty")
```

#### `Minlength(s string, minLength int, customMessage string) error`
Validates minimum string length.

**Parameters:**
- `s` (string): String to validate
- `minLength` (int): Minimum required length
- `customMessage` (string): Custom error message (optional)

**Example:**
```go
err := stringvalidator.Minlength("hello", 3, "")
```

#### `Maxlength(s string, maxLength int, customMessage string) error`
Validates maximum string length.

**Parameters:**
- `s` (string): String to validate
- `maxLength` (int): Maximum allowed length
- `customMessage` (string): Custom error message (optional)

**Example:**
```go
err := stringvalidator.Maxlength("hi", 10, "")
```

#### `HasLengthRange(s string, minLength, maxLength int, customMessage string) error`
Validates string length within a range.

**Parameters:**
- `s` (string): String to validate
- `minLength` (int): Minimum required length
- `maxLength` (int): Maximum allowed length
- `customMessage` (string): Custom error message (optional)

**Example:**
```go
err := stringvalidator.HasLengthRange("hello", 3, 10, "")
```

---

### Password Validation

#### `ValidatePassword(password string, customMessage string) error`
Validates password strength with the following requirements:
- Minimum 8 characters
- At least one uppercase letter
- At least one lowercase letter  
- At least one number

**Parameters:**
- `password` (string): Password to validate
- `customMessage` (string): Custom error message (optional)

**Example:**
```go
err := stringvalidator.ValidatePassword("MyPass123", "")
```

---

### Username Validation

#### `ValidateUsername(username string, customMessage string) error`
Validates username format:
- Minimum 3 characters
- Only letters, numbers, and underscores allowed

**Parameters:**
- `username` (string): Username to validate
- `customMessage` (string): Custom error message (optional)

**Example:**
```go
err := stringvalidator.ValidateUsername("user_123", "")
```

---

### Phone Number Validation

#### `ValidatePhoneNumber(phone string, customMessage string) error`
Validates international phone number format:
- Optional + prefix
- Must start with 1-9 (not 0)
- 2-15 digits total

**Parameters:**
- `phone` (string): Phone number to validate
- `customMessage` (string): Custom error message (optional)

**Valid formats:**
- `+1234567890`
- `1234567890`
- `+12345678901234`

**Example:**
```go
err := stringvalidator.ValidatePhoneNumber("+1234567890", "")
```

---

### URL Validation

#### `ValidateURL(url string, customMessage string) error`
Validates HTTP/HTTPS URL format.

**Parameters:**
- `url` (string): URL to validate
- `customMessage` (string): Custom error message (optional)

**Valid formats:**
- `https://example.com`
- `http://subdomain.example.com/path`

**Example:**
```go
err := stringvalidator.ValidateURL("https://example.com", "")
```

---

### Alphanumeric Validation

#### `IsAlphanumeric(s string, customMessage string) error`
Checks if string contains only letters and numbers (no spaces or special characters).

**Parameters:**
- `s` (string): String to validate
- `customMessage` (string): Custom error message (optional)

**Example:**
```go
err := stringvalidator.IsAlphanumeric("abc123", "")
```

---

### ID Validation

#### `ValidateIDString(id string, customMessage string) error`
Validates various string ID formats:
- **UUID**: `550e8400-e29b-41d4-a716-446655440000`
- **MongoDB ObjectID**: `507f1f77bcf86cd799439011` (24 hex characters)
- **Numeric ID**: `123`, `999999` (positive integers, no leading zeros)

**Parameters:**
- `id` (string): ID string to validate
- `customMessage` (string): Custom error message (optional)

**Example:**
```go
// UUID
err := stringvalidator.ValidateIDString("550e8400-e29b-41d4-a716-446655440000", "")

// MongoDB ObjectID
err := stringvalidator.ValidateIDString("507f1f77bcf86cd799439011", "")

// Numeric ID
err := stringvalidator.ValidateIDString("12345", "")
```

#### `ValidateObjectID(id primitive.ObjectID, customMessage string) error`
Validates MongoDB primitive.ObjectID (checks if not nil/zero).

**Parameters:**
- `id` (primitive.ObjectID): MongoDB ObjectID to validate
- `customMessage` (string): Custom error message (optional)

**Example:**
```go
import "go.mongodb.org/mongo-driver/bson/primitive"

objectID, _ := primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")
err := stringvalidator.ValidateObjectID(objectID, "")
```

---

## Error Handling

All functions return an `error` type:
- `nil` if validation passes
- `error` with descriptive message if validation fails

### Custom Error Messages

Most functions accept a `customMessage` parameter:
- Pass `""` for default error messages
- Pass your custom message to override defaults

```go
// Default message
err := stringvalidator.IsEmpty("", "")
// Returns: "string cannot be empty"

// Custom message  
err := stringvalidator.IsEmpty("", "Name is required")
// Returns: "Name is required"
```

## Complete Example

```go
package main

import (
    "fmt"
    stringvalidator "github.com/kwabsntim/utils-library/string-validator"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func validateUser(email, password, username, phone, idString string) {
    // Validate email
    if err := stringvalidator.IsEmail(email); err != nil {
        fmt.Printf("Email error: %v\n", err)
        return
    }
    
    // Validate password
    if err := stringvalidator.ValidatePassword(password, ""); err != nil {
        fmt.Printf("Password error: %v\n", err)
        return
    }
    
    // Validate username
    if err := stringvalidator.ValidateUsername(username, ""); err != nil {
        fmt.Printf("Username error: %v\n", err)
        return
    }
    
    // Validate phone
    if err := stringvalidator.ValidatePhoneNumber(phone, ""); err != nil {
        fmt.Printf("Phone error: %v\n", err)
        return
    }
    
    // Validate ID string
    if err := stringvalidator.ValidateIDString(idString, ""); err != nil {
        fmt.Printf("ID error: %v\n", err)
        return
    }
    
    // Convert and validate ObjectID
    objectID, err := primitive.ObjectIDFromHex(idString)
    if err != nil {
        fmt.Printf("ObjectID conversion error: %v\n", err)
        return
    }
    
    if err := stringvalidator.ValidateObjectID(objectID, ""); err != nil {
        fmt.Printf("ObjectID validation error: %v\n", err)
        return
    }
    
    fmt.Println("All validations passed!")
}

func main() {
    validateUser(
        "user@example.com", 
        "MyPass123", 
        "user_123", 
        "+1234567890",
        "507f1f77bcf86cd799439011",
    )
}
```

## Testing

Run tests with:
```bash
go test github.com/kwabsntim/utils-library/string-validator
```

## Dependencies

- `go.mongodb.org/mongo-driver/bson/primitive` (for ObjectID validation)

## License

MIT License

## Contributing

1. Fork the repository
2. Create your feature branch
3. Add tests for new functionality
4. Ensure all tests pass
5. Submit a pull request

## Changelog

### v1.2.0
- Added `ValidateIDString` function for string ID validation
- Added `ValidateObjectID` function for MongoDB ObjectID validation
- Replaced old `ValidateID` function with two specialized functions

### v1.1.0
- Added `ValidateID` function for UUID, ObjectID, and numeric ID validation

### v1.0.0
- Initial release with basic validation functions