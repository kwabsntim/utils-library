//Package mathlib provides functions for mathematics operations
//It includes fucntions for additon,subtraction,division and multiplication

package mathlib

import (
	"errors"
)

// Add adds two numbers and returns the result.
func Add(a, b int) int {
	return a + b
}

// Subtract subtracts the second number from the first and returns the result.
func Subtract(a, b int) int {
	return a - b
}

// Multiply multiplies two numbers and returns the result.
func Multiply(a, b int) int {
	return a * b
}

// Divide divides the first number by the second and returns the result.
// It returns an error if the second number is zero.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}
