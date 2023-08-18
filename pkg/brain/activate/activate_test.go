package activate

import (
	"testing"
)

// message is a formatted error for failed tests.
var message = "\nTest:\t\t %v\nExpected:\t %v\nResults:\t %v\n"

// Test the Linear activation function.
func TestLinear(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{10, 10},
		{-2, -2},
		{0, 0},
	}

	for _, test := range tests {
		results := Linear(test.input)
		if results != test.expected {
			t.Errorf(message, "Linear", test.expected, results)
		}
	}
}

// Test the ReLU activation function.
func TestReLU(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{10, 10},
		{-2, 0},
		{0, 0},
		{1, 1},
	}

	for _, test := range tests {
		results := ReLU(test.input)
		if results != test.expected {
			t.Errorf(message, "ReLU", test.expected, results)
		}
	}

}
