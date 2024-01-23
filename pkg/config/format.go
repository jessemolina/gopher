package config

import (
	"regexp"
	"strings"
)

// splitCamelCase splits a camel case string into a delimeted-split string.
func splitCamelCase(input string, delim string) string {
	// Regular expression to match the position before a lowercase letter that follows uppercase letters
	re := regexp.MustCompile(`([a-z0-9])([A-Z])|([A-Z])([A-Z][a-z])`)

	// Replace matched positions with a delimeter
	return re.ReplaceAllStringFunc(input, func(str string) string {
		return str[:len(str)/2] + delim + str[len(str)/2:]
	})
}

// toScreamingSnakeCase converts a delimeted string into an all caps, underscore-split string.
func toScreamingSnakeCase(input string, delim string) string {
	value := input
	value = strings.ReplaceAll(value, delim, "_")
	value = strings.ToUpper(value)

	return value
}

// toKebabcase converts a delimeted string into a lowercase, hyphen-split string.
func toKebabCase(input string, delim string) string {
	value := input
	value = strings.ReplaceAll(value, delim, "-")
	value = strings.ToLower(value)

	return value
}
