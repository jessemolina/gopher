package config

import (
	"flag"
	"reflect"
	"strings"
)

/*
  TODO Brainstorm on how other fi.tags, such as mask, can be used
*/

// Parse unmarshalls the given cfg from os env vars, flag values, and config field tags.
func Parse(cfg interface{}, prefixes ...string) {
	prefix := ""

	if len(prefixes) > 0 {
		prefix = strings.Join(prefixes, "")
	}

	fieldInfos := makeInfo(cfg, prefix)

	for _, fi := range fieldInfos {
		fi.SetFlag()
	}

	flag.Parse()
}

// makeInfo builds a collection of fieldInfo.
func makeInfo(cfg interface{}, prefix string) []fieldInfo {
	var fieldInfos []fieldInfo

	values := reflect.ValueOf(cfg)

	if values.Kind() == reflect.Pointer {
		values = values.Elem()
	}

	for i := 0; i < values.NumField(); i++ {
		sf := values.Type().Field(i)
		v := values.Field(i)

		if v.Kind() == reflect.Struct {
			newPrefix := prefix + sf.Name
			nestedInfo := makeInfo(v.Addr().Interface(), newPrefix)
			fieldInfos = append(fieldInfos, nestedInfo...)
		} else {
			fi := NewFieldInfo(sf, v, prefix)
			fieldInfos = append(fieldInfos, fi)

		}
	}
	return fieldInfos
}
