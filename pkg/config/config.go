package config

import (
	"flag"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func Parse(cfg interface{}) {
	setFlags(cfg)
}

func setFlags(cfg interface{}) {
	cfgType := reflect.TypeOf(cfg).Elem()
	for i := 0; i < cfgType.NumField(); i++ {
		field := cfgType.Field(i)

		split := splitCamelCase(field.Name, " ")
		snake := screamingSnakeCase(split, " ")

		name := strings.ToLower(field.Name)
		usage := fmt.Sprintf("--%v/$%v", name, snake)

		switch field.Type.Kind() {
		case reflect.Int:
			flag.Int(name, 0, usage)
		case reflect.String:
			flag.String(name, "zero", usage)
		}
	}

	flag.Parse()

}

func parseDefault(tag reflect.StructTag) string {
	return ""
}

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

// screamingSnakeCase converts a delimeted string into an all caps, underscore-split string.
func screamingSnakeCase(name string, delim string) string {
	value := name
	value = strings.ReplaceAll(value, delim, "_")
	value = strings.ToUpper(value)

	return value
}
