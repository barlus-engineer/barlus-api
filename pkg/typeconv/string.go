package typeconv

import (
	"fmt"
	"reflect"
	"strconv"
)

func Str2Any(value string, t interface{}) (interface{}, error) {
	switch t {
		case reflect.String:
			return value, nil
		case reflect.Int:
			return str2Int(value)
		case reflect.Float64:
			return str2Float64(value)
		case reflect.Bool:
			return str2Bool(value)
		default:
			return nil, fmt.Errorf("unsupported type")
	}
}

// ====

func str2Int(value string) (interface{}, error) {
	if value == "" {
		value = "0"
	}
	return strconv.Atoi(value)
}

func str2Float64(value string) (interface{}, error) {
	if value == "" {
		value = "0"
	}
	return strconv.ParseFloat(value, 64)
}

func str2Bool(value string) (interface{}, error) {
	if value == "" {
		value = "false"
	}
	return strconv.ParseBool(value)
}