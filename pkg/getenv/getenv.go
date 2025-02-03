package getenv

import (
	"os"
	"reflect"

	"github.com/barlus-engineer/barlus-api/pkg/logger"
	"github.com/barlus-engineer/barlus-api/pkg/typeconv"
	"github.com/joho/godotenv"
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
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		key := field.Tag.Get("env")
		def := field.Tag.Get("def")
		
		env := Get(key, def)

		value, err := typeconv.Str2Type(env, field.Type)
		if err != nil {
			logger.Fatalf("getenv: Unsupport type '%v'", reflect.TypeOf(env))
		}

		if reflect.TypeOf(value) != field.Type {
			panic("getenv: Type not match")
		}

		v.Field(i).Set(reflect.ValueOf(value))
	}
	return nil
}

// ======

// func isStruct(strc interface{}) bool {
// 	t := reflect.TypeOf(strc)

// 	if t.Kind() == reflect.Ptr {
// 		t = t.Elem()
// 	}

// 	return t.Kind() == reflect.Struct
// }