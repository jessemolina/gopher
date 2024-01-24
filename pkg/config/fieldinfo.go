package config

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

const delimeter = " "

// fieldInfo represents information about a field in a struct.
type fieldInfo struct {
	name    string
	pointer interface{}
	tags    map[string]string
}

// Description parses CamelCase into a space delimeted string.
func (fi fieldInfo) Description() string {
	return parseCamelCase(fi.name, delimeter)
}

// OSEnv formats the fieldInfo screaming snake case,
// a format typically used for OS ENV variables.
func (fi fieldInfo) OSEnv() string {
	return toScreamingSnakeCase(fi.Description(), delimeter)
}

// Option formats the fieldInfo name to kebab-case,
// a format typically used in Unix flag options.
func (fi fieldInfo) Option() string {
	return toKebabCase(fi.Description(), delimeter)
}

// Usage returns a string that describes the field's usage.
func (fi fieldInfo) Usage() string {
	message := "%v\n--%v/$%v"
	return fmt.Sprintf(message,
		fi.Description(),
		fi.Option(),
		fi.OSEnv())
}

// Default returns the field's default value, prioritizing OS ENV,
// and defaulting to tag value.
func (fi fieldInfo) Default() string {
	defaultValue, _ := fi.tags["default"]
	envValue := os.Getenv(fi.OSEnv())
	if envValue != "" {
		defaultValue = envValue
	}
	return defaultValue
}

// SetFlag sets the field's CLI flag option.
func (fi *fieldInfo) SetFlag() {
	switch pointer := fi.pointer.(type) {
	case *string:
		flag.StringVar(
			pointer,
			fi.Option(),
			fi.Default(),
			fi.Usage())

	case *int:
		defaultInt, err := strconv.Atoi(fi.Default())
		if err != nil {
			defaultInt = 0
		}

		flag.IntVar(
			pointer,
			fi.Option(),
			defaultInt,
			fi.Usage())
	}
}

// newFieldInfo is a constructor for FieldInfo.
func NewFieldInfo(sf reflect.StructField, rv reflect.Value, prefix string) fieldInfo {
	name := prefix + sf.Name
	tags := parseTags(sf.Tag.Get("config"))

	var pointer interface{}

	if !rv.CanAddr() {
		message := fmt.Sprintf("Unable to attain address for field %v", name)
		panic(message)
	}

	if rv.CanAddr() {
		switch rv.Kind() {
		case reflect.String:
			pointer = rv.Addr().Interface().(*string)
		case reflect.Int:
			pointer = rv.Addr().Interface().(*int)

		}
	}

	return fieldInfo{
		name:    name,
		pointer: pointer,
		tags:    tags,
	}

}
