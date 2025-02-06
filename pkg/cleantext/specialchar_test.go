package cleantext_test

import (
	"testing"

	"github.com/barlus-engineer/barlus-api/pkg/cleantext"
)

func TestSpecialChar(t *testing.T) {
	data := "/Barlus#@%^_cuda._1234"
	excepted := "barlus_cuda_1234"

	result := cleantext.SpecialChar(data)

	if result != excepted {
		t.Errorf("expected %s but got %s", excepted, result)
	}
}