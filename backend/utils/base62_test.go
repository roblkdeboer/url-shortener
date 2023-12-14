package utils

import (
	"testing"
)

func TestToBase62(t *testing.T) {
	// Test cases
	testCases := []struct {
		input    int64
		expected string
	}{
		{0, "0"},
		{10, "A"},
		{61, "z"},
		{62, "10"},
		{100, "1c"},
	}

	// Run tests
	for _, tc := range testCases {
		result := toBase62(tc.input)
		if result != tc.expected {
			t.Errorf("toBase62(%v) = %v, expected %v", tc.input, result, tc.expected)
		}
	}
}

func TestGenerateRandomNumber(t *testing.T) {
	// Call the function multiple times and ensure randomness
	for i := 0; i < 10; i++ {
		result := generateRandomNumber()
		if result < 0 || result > 255 {
			t.Errorf("generateRandomNumber() = %v, expected a number between 0 and 255", result)
		}
	}
}
