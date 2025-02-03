package getenv_test

import (
	"os"
	"testing"

	"github.com/barlus-engineer/barlus-api/pkg/getenv"
)

func TestGet(t *testing.T) {
	key := "TEST_GETENV"
	expected := "test"
	os.Setenv(key, expected)
	value := getenv.Get(key, expected)

	if value != expected {
		t.Errorf("%s != %s", value, expected)
	}
}

type DataTest struct {
	User User
	Score float64 `envkey:"TEST_GETENV_SCORE"`
}

type User struct {
	Name string `envkey:"TEST_GETENV_NAME" envdef:"Barlus"`
	Age  int    `envkey:"TEST_GETENV_AGE" envdef:"15"`
}

var DefaultUserExpected = DataTest{
	User: User{
		Name: "Barlus",
		Age:  15,
	},
	Score: 0.0,
}

var UserExpected = DataTest{
	User: User{
		Name: "Mouk",
		Age:  14,
	},
	Score: 3.7,
}

func TestGetStructDefault(t *testing.T) {
	var user DataTest
	if err := getenv.GetStruct(&user); err != nil {
		t.Error(err)
	}	
	if user != DefaultUserExpected {
		t.Errorf("Expected %+v, but got %+v", DefaultUserExpected, user)
	}
}

// func TestGetStruct(t *testing.T) {
// 	var user DataTest
// 	if err := getenv.GetStruct(&user); err != nil {
// 		t.Error(err)
// 	}
// 	if user != UserExpected {
// 		t.Errorf("Expected %+v, but got %+v", DefaultUserExpected, user)
// 	}
// }

// func TestGetStruct(t *testing.T) {
// 	var user DataTest
// 	if err := getenv.GetStruct(&user); err != nil {
// 		t.Error(err)
// 	}
// 	if user != UserExpected {
// 		t.Errorf("Expected %+v, but got %+v", DefaultUserExpected, user)
// 	}
// }
