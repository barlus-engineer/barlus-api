package getenv_test

import (
	"os"
	"testing"

	"github.com/barlus-engineer/barlus-api/pkg/getenv"
)

func TestGet(t *testing.T) {
	expected := "test"
	os.Setenv("TESTTT_GETTT_ENVVV", expected)
	value := getenv.Get("TESTTT_GETTT_ENVVV", "test")

	if value != expected {
		t.Errorf("%s != %s", value, expected)
	}
}