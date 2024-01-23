package config

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

/*
  TODO Brainstorm on how other fi.tags, such as mask, can be used
*/

// fieldInfo represents information about a field in a struct.
type fieldInfo struct {
	name  string
	desc  string
	env   string
	kind  reflect.Kind
	value reflect.Value
	tags  map[string]string
}

// Usage returns a string that describes the field's usage.
func (fi *fieldInfo) Usage(message string) string {
	return fmt.Sprintf(message, fi.desc, fi.name, fi.env)
}

// Default returns the field's default value, prioritizing OS ENV,
// and defaulting to tag value.
func (fi *fieldInfo) Default() string {
	defaultValue, _ := fi.tags["default"]
	envValue := os.Getenv(fi.env)
	if envValue != "" {
		defaultValue = envValue
	}
	return defaultValue
}

// SetFlag sets the field's CLI flag option.
func (fi *fieldInfo) SetFlag() {
	message := "%v\n--%v/$%v"

	usage := fi.Usage(message)
	defaultString := fi.Default()

	switch fi.kind {
	case reflect.String:
		flag.StringVar(
			fi.value.Addr().Interface().(*string),
			fi.name,
			defaultString,
			usage)
	case reflect.Int:
		intValue, err := strconv.Atoi(defaultString)
		if err != nil {
			intValue = 0
		}

		flag.IntVar(
			fi.value.Addr().Interface().(*int),
			fi.name,
			intValue,
			usage)
	}

}

// Parse unmarshalls the given cfg from os env vars, flag values, and config field tags.
func Parse(cfg interface{}, prefixes ...string) {
	prefix := ""

	if len(prefixes) > 0 {
		prefix = strings.Join(prefixes, "")
	}

	fieldInfos := makeInfo(cfg,prefix)

	for _, fi := range fieldInfos {
		fi.SetFlag()
	}

	flag.Parse()
}

// makeInfo builds a collection of fieldInfo.
func makeInfo(cfg interface{}, prefix string) []fieldInfo {
	fieldInfos := []fieldInfo{}

	values := reflect.ValueOf(cfg)

	if values.Kind() == reflect.Pointer {
		values = values.Elem()
	}

	for i := 0; i < values.NumField(); i++ {
		sf := values.Type().Field(i)
		v := values.Field(i)

		if v.Kind() == reflect.Struct {
			newPrefix := prefix + sf.Name
			nestedInfo := makeInfo(v.Interface(), newPrefix)
			fieldInfos = append(fieldInfos, nestedInfo...)
		} else {
			fi := newFieldInfo(sf, v, prefix)
			fieldInfos = append(fieldInfos, fi)
		}

	}

	return fieldInfos
}

// newFieldInfo is a constructor for FieldInfo.
func newFieldInfo(sf reflect.StructField, rv reflect.Value, prefix string) fieldInfo {
	desc := splitCamelCase(prefix+sf.Name, " ")
	name := toKebabCase(desc, " ")
	env := toScreamingSnakeCase(desc, " ")
	kind := rv.Type().Kind()
	value := rv
	tags := make(map[string]string)

	configTag := sf.Tag.Get("config")
	configList := strings.Split(configTag, ",")

	for _, config := range configList {
		values := strings.Split(config, ":")
		if len(values) == 2 {
			tags[values[0]] = values[1]
		} else {
			tags[values[0]] = "true"
		}

	}

	return fieldInfo{
		name:  name,
		desc:  desc,
		env:   env,
		kind:  kind,
		value: value,
		tags:  tags,
	}

}
