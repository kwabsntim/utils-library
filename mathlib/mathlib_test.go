package mathlib_test

import (
	"fmt"
	"testing"

	"github.com/kwabsntim/utils-library/mathlib"
)

//------------ EXAMPLE TESTS ----------------------

// function for addition
func ExampleAdd() {
	sum := mathlib.Add(5, 6)
	fmt.Println(sum)
	// Output:11
}

// fuction test for subtraction
func ExampleSubtract() {
	diff := mathlib.Subtract(8, 6)
	fmt.Println(diff)
	// Output:2

}

// ------------- UNIT TESTS -----------------
func TestAdd(t *testing.T) {
	result := mathlib.Add(5, 6)
	if result != 11 {
		t.Errorf("Expected 11, got %d", result)
	}
}

func TestSubtract(t *testing.T) {
	result := mathlib.Subtract(8, 6)
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}

func TestMultiply(t *testing.T) {
	result := mathlib.Multiply(3, 4)
	if result != 12 {
		t.Errorf("Expected 12, got %d", result)
	}
}

func TestDivide(t *testing.T) {
	result, err := mathlib.Divide(10, 2)
	if err != nil || result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := mathlib.Divide(10, 0)
	if err == nil {
		t.Error("Expected error for division by zero")
	}
}
