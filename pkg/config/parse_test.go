package config

import (
	"testing"
)

// Formated error message for failed tests.
var message = "\nExpected:\t %v\nResults:\t %v\n"

// Test case scenario for string formatting functions.
type ParseTest struct {
	input    string
	delim    string
	expected string
}

// Test function for splitCamelCase.
func TestParseCamelCase(t *testing.T) {
	tests := []ParseTest{
		{"debugPort", "_", "debug_Port"},
		{"APIPort", " ", "API Port"},
		{"ServiceAPIPort", "-", "Service-API-Port"},
		{"RESTEndpoint", "_", "REST_Endpoint"},
	}

	for _, test := range tests {
		results := parseCamelCase(test.input, test.delim)

		if test.expected != results {
			t.Fatalf("\nUnexpected value"+message, test.expected, results)
		}

	}

}

// Test function toScreamingSnakeCase.
func TestToScreamingSnakeCase(t *testing.T) {
	tests := []ParseTest{
		{"hello_World", "_", "HELLO_WORLD"},
	}

	for _, test := range tests {
		results := toScreamingSnakeCase(test.input, test.delim)

		if test.expected != results {
			t.Fatalf("\nUnexpected value"+message, test.expected, results)
		}

	}

}

