package getenv

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/barlus-engineer/barlus-api/pkg/typeconv"
	"github.com/joho/godotenv"
)

var (
	Default = ""

	ErrConvType = errors.New("getenv: field '%s', Error %v")
	ErrItsnotStruct = errors.New("getenv: expected a pointer to a struct")
	ErrUnsupportType = errors.New("getenv: field '%s', unsuported type '%v'")
)

func Get(key string, deValue string) string {
	godotenv.Load()

	value := os.Getenv(key)
	if value == "" {
		value = deValue
	}

	return value
}

func GetStruct(cfgStruct interface{}) error {
	v := reflect.ValueOf(cfgStruct).Elem()
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
				def2 := field2.Tag.Get("envdef")

				env2 := Get(key2, def2)

				value2, err := typeconv.Str2Any(env2, field2.Type.Kind())
				if err != nil {
					if err != ErrUnsupportType {
						return fmt.Errorf(ErrConvType.Error(), field.Name, err)
					}
					return fmt.Errorf(ErrUnsupportType.Error(), field.Name, reflect.TypeOf(env2).Kind())
				}

				if reflect.TypeOf(value2) != field2.Type {
					panic("getenv: Type not match")
				}

				v2.Field(j).Set(reflect.ValueOf(value2))
			}
			continue
		}

		key := field.Tag.Get("envkey")
		def := field.Tag.Get("envdef")

		env := Get(key, def)

		value, err := typeconv.Str2Any(env, field.Type.Kind())
		if err != nil {
			if err != ErrUnsupportType {
				return fmt.Errorf("getenv: field '%s', Error %v", field.Name, err)
			}
			return fmt.Errorf(ErrUnsupportType.Error(), reflect.TypeOf(env).Kind())
		}

		if reflect.TypeOf(value) != field.Type {
			panic("getenv: Type not match")
		}

		v.Field(i).Set(reflect.ValueOf(value))
	}
	return nil
}