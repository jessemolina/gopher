package main

import (
	"fmt"
	"testing"
)

var message = "Results:\t %v\nExpected:\t %v\n"

type testConfig struct {
	params   []vector // array of vectors to cross multiply
	expected float64  // product of vector multiplication is a scalar
}

func TestDotProduct(t *testing.T) {
	tests := []testConfig{
		{
			params:   []vector{{-7, 2, 3}, {5, 2, 3}},
			expected: -22,
		},
	}

	for _, test := range tests {
		results := DotProduct(test.params[0], test.params[1])
		if results != test.expected {
			fmt.Printf(message, results, test.expected)
		}
	}
}
