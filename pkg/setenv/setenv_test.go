package setenv_test

import (
	"testing"

	"github.com/barlus-engineer/barlus-api/pkg/getenv"
	"github.com/barlus-engineer/barlus-api/pkg/setenv"
	"github.com/barlus-engineer/barlus-api/pkg/typeconv"
)

func TestSet(t *testing.T) {
	name := "Barlus"
	age := 15

	setenv.Set("TEST_NAME", name)
	setenv.Set("TEST_AGE", age)

	envName := getenv.Get("TEST_NAME", getenv.Default)
	envAge := getenv.Get("TEST_AGE", getenv.Default)

	if name != envName {
		expected := name
		actual := envName
		t.Errorf("expected %s, but got %s", expected, actual)
	}

	if typeconv.Any2Str(age) != envAge {
		expected := age
		actual := envAge
		t.Errorf("expected %d, but got %s", expected, actual)
	}
}
