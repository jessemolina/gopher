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
  TODO Determine how to pass fi string up the call stack for the logger.
  TODO Brainstorm on how other fi.tags, such as mask, can be used
  TODO Check for whether a field is of type nested struct
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
func Parse(cfg interface{}) {

	fieldsInfo := makeInfo(cfg)

	for _, fi := range fieldsInfo {
		fi.SetFlag()
	}

	flag.Parse()
}

// makeInfo returns an array of fieldInfo for each field in cfg.
func makeInfo(cfg interface{}) []fieldInfo {
	fieldsInfo := []fieldInfo{}

	t := reflect.TypeOf(cfg).Elem()
	v := reflect.ValueOf(cfg).Elem()

	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).PkgPath != "" {
			break
		}
		name := strings.ToLower(t.Field(i).Name)
		desc := splitCamelCase(t.Field(i).Name, " ")
		env := toScreamingSnakeCase(desc, " ")
		kind := t.Field(i).Type.Kind()
		value := v.Field(i)
		tags := make(map[string]string)

		configTag := t.Field(i).Tag.Get("config")
		configList := strings.Split(configTag, ",")

		for _, config := range configList {
			values := strings.Split(config, ":")
			if len(values) == 2 {
				tags[values[0]] = values[1]
			} else {
				tags[values[0]] = "true"
			}

		}

		info := fieldInfo{
			name:  name,
			desc:  desc,
			env:   env,
			kind:  kind,
			value: value,
			tags:  tags,
		}

		fieldsInfo = append(fieldsInfo, info)
	}

	return fieldsInfo
}
