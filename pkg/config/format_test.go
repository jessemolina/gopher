package config

import (
	"testing"
)

// Formated error message for failed tests.
var message = "\nExpected:\t %v\nResults:\t %v\n"

// Test case scenario for string formatting functions.
type FormatTest struct {
	input    string
	delim    string
	expected string
}

// Test function for splitCamelCase.
func TestSplitCamelCase(t *testing.T) {
	tests := []FormatTest{
		{"helloWorld", "_", "hello_World"},
	}

	for _, test := range tests {
		results := splitCamelCase(test.input, test.delim)

		if test.expected != results {
			t.Fatalf("\nUnexpected value" + message, test.expected, results)
		}

	}

}

// Test function toScreamingSnakeCase.
func TestToScreamingSnakeCase(t *testing.T) {
	tests := []FormatTest{
		{"hello_World", "_", "HELLO_WORLD"},
	}

	for _, test := range tests {
		results := toScreamingSnakeCase(test.input, test.delim)

		if test.expected != results {
			t.Fatalf("\nUnexpected value" + message, test.expected, results)
		}

	}

}
