package config

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// fieldInfo represents information about a field in a struct.
type fieldInfo struct {
	name  string
	desc  string
	env   string
	kind  reflect.Kind
	value reflect.Value
}

// Parse unmarshalls the given cfg from os env vars, flag values, and config field tags.
func Parse(cfg interface{}) {

	fieldsInfo := makeInfo(cfg)

	for _, fi := range fieldsInfo {
		msg := "%v\n--%v/$%v"
		usage := fmt.Sprintf(msg, fi.desc, fi.name, fi.env)
		envValue := os.Getenv(fi.env)

		switch fi.kind {
		case reflect.String:
			flag.StringVar(
				fi.value.Addr().Interface().(*string),
				fi.name,
				envValue,
				usage)
		case reflect.Int:
			intValue, err := strconv.Atoi(envValue)
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

	flag.Parse()
}

// makeInfo returns an array of fieldInfo for each field in cfg.
func makeInfo(cfg interface{}) []fieldInfo {
	fieldsInfo := []fieldInfo{}

	t := reflect.TypeOf(cfg).Elem()
	v := reflect.ValueOf(cfg).Elem()

	for i := 0; i < t.NumField(); i++ {
		name := strings.ToLower(t.Field(i).Name)
		desc := splitCamelCase(t.Field(i).Name, " ")
		env := screamingSnakeCase(desc, " ")
		kind := t.Field(i).Type.Kind()
		value := v.Field(i)

		info := fieldInfo{
			name:  name,
			desc:  desc,
			env:   env,
			kind:  kind,
			value: value,
		}

		fieldsInfo = append(fieldsInfo, info)
	}

	return fieldsInfo
}

// setEnv checks for a field's environment variable and sets if it exists.
func setEnv(fi fieldInfo) {
	osEnv := os.Getenv(fi.env)

	// TODO Determine how to convernt osEnv strings to non-string values (i.e. int)
	if osEnv != "" {
		fi.value.SetString(osEnv)
	}

}

// setFlag defines a cli flag for each fieldInfo.
func setFlags(fi fieldInfo) {

	msg := "%v\n--%v/$%v"
	usage := fmt.Sprintf(msg, fi.desc, fi.name, fi.env)

	switch fi.kind {
	case reflect.String:
		flag.String(fi.name, "", usage)
	case reflect.Int:
		flag.Int(fi.name, 0, usage)
	}
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
