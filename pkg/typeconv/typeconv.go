package typeconv

import (
	"fmt"
	"strconv"
)

func Str2Type(value string, t interface{}) (interface{}, error) {
	switch t.(type) {
		case int:
			return strconv.Atoi(value)
		case float64:
			return strconv.ParseFloat(value, 64)
		case bool:
			return strconv.ParseBool(value)
		default:
			return nil, fmt.Errorf("unsupported type")
	}
}