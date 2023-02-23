package main

import (
	"errors"
	"testing"
)

// Formated error message for failed tests.
var message = "\nExpected:\t %v\nResults:\t %v\n"

// Configuration for tests.
type testConfig struct {
	params   []vector // array of vectors to cross multiply
	expected float64  // product of vector multiplication is a scalar
	err      error    // expected error
}

func TestDotProduct(t *testing.T) {
	// Collection of test configurations
	tests := []testConfig{
		{
			params:   []vector{{1, 2, 3}, {1, 2, 3}}, // Simple
			expected: 14,
			err:      nil,
		},
		{
			params:   []vector{{-7, -1, 0}, {-5, -6, -3}}, // Negative scalars
			expected: 41,
			err: nil,
		},
		{
			params:   []vector{{1026, 1600}, {768, 1200}}, // Large scalar values
			expected: 2707968,
			err: nil,
		},
		{
			params:   []vector{{9, 8, 7}, {1, 2}}, // Mismatch in dimmensions
			expected: 0,
			err: errors.New("Vector dimmensions don't match"),
		},
	}

	// Loop through test and confirm results match expected.
	for _, test := range tests {
		results, err := DotProduct(test.params[0], test.params[1])
		// No error was expected but an error occured.
		if test.err == nil && err != nil {
			t.Fatalf("\nUnknown error" + message, nil, err)

		}
		// Expected error does not match error generated.
		if test.err != nil && test.err.Error() != err.Error() {
			t.Fatalf("\nUnexpected error" + message, test.err, err)
		}
		// Expected value did not match results without errors.
		if test.expected != results && err != nil {
			t.Errorf("\nUnexpected value" + message, test.expected, results)
		}
	}
}
