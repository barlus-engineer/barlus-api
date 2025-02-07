package text_test

import (
	"testing"

	"github.com/barlus-engineer/barlus-api/pkg/text"
)

func TestCleanEmail(t *testing.T) {
	data := "Barluscuda@gmail.com /(+=-)"
	excepted := "barluscuda@gmail.com"

	result := text.CleanEmail(data)

	if result != excepted {
		t.Errorf("expected %s but got %s", excepted, result)
	}
}