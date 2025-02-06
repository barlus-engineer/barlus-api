package setenv

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/barlus-engineer/barlus-api/pkg/typeconv"
)

var (
	ErrSetenv = errors.New("setenv: %v")
	ErrItsnotStruct = errors.New("setenv: expected a pointer to a struct")
)

func Set(key string, value any) error {
	envValue := typeconv.Any2Str(value)
	if err := os.Setenv(key, envValue); err != nil {
		return err
	}
	return nil
}

func SetStruct(cfgStruct interface{}) error {
	v := reflect.ValueOf(cfgStruct)
	if v.Type().Kind() != reflect.Struct {
		return ErrItsnotStruct
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Type.Kind() == reflect.Struct {
			v2 := v.Field(i)
			t2 := v2.Type()
			for j := 0; j < t2.NumField(); j++ {
				field2 := t2.Field(j)

				key2 := field2.Tag.Get("envkey")
				val2 := v2.Field(j).Interface()

				if err := Set(key2, val2);err != nil {
					return fmt.Errorf(ErrSetenv.Error(), err)
				}
			}
			continue
		}

		key := field.Tag.Get("envkey")
		val := v.Field(i).Interface()

		if err := Set(key, val);err != nil {
			return fmt.Errorf(ErrSetenv.Error(), err)
		}
	}
	return nil
}