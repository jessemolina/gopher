package main

import (
	"testing"
)

var message = "\nResults:\t %v\nExpected:\t %v\n"

type testConfig struct {
	params   []vector // array of vectors to cross multiply
	expected float64  // product of vector multiplication is a scalar
}

func TestDotProduct(t *testing.T) {
	tests := []testConfig{
		{
			params:   []vector{{1, 2, 3}, {1, 2, 3}}, // Basic
			expected: 14,
		},
		{
			params:   []vector{{8}, {9}}, // Single value scalar
			expected: 72,
		},
		{
			params:   []vector{{-7, -1, 0}, {-5, -6, -3}}, // Negative scalars
			expected: 41,
		},
		{
			params:   []vector{{1026, 1600}, {768, 1200}}, // Large scalars
			expected: 2707968,
		},
	}

	for _, test := range tests {
		results := DotProduct(test.params[0], test.params[1])
		if results != test.expected {
			t.Errorf(message, results, test.expected)
		}
	}
}
