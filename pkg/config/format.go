package config

import (
	"regexp"
	"strings"
)

// splitName splits a camel case string into a delimeted-split string.
func splitCamelCase(name string, delim string) string {
	pattern := `[A-Z][a-z]+`
	r := regexp.MustCompile(pattern)

	matches := r.FindAllStringIndex(name, -1)

	value := ""

	if len(matches) == 0 {
		return value
	}

	if matches[0][0] != 0 {
		start := 0
		end := matches[0][0]
		value += name[start:end]
	}

	for _, match := range matches {
		start, stop := match[0], match[1]
		word := name[start:stop]

		if match[0] != 0 {
			word = delim + word
		}

		value += word
	}

	return value
}

// toScreamingSnakeCase converts a delimeted string into an all caps, underscore-split string.
func toScreamingSnakeCase(name string, delim string) string {
	value := name
	value = strings.ReplaceAll(value, delim, "_")
	value = strings.ToUpper(value)

	return value
}

// toKebabcase converts a delimeted string into a lowercase, hyphen-split string.
func toKebabCase(name string, delim string) string {
	value := name
	value = strings.ReplaceAll(value, delim, "-")
	value = strings.ToLower(value)

	return value

   return ""
}
