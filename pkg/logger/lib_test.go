package logger_test

import (
	"testing"

	"github.com/barlus-engineer/barlus-api/pkg/logger"
)

func TestLoggerValuesJoin(t *testing.T) {
	text := []string{"Hello, ", "Barlus"}
	expected := "Hello, Barlus"
	tany := make([]any, len(text))
	for i, v := range text {
		tany[i] = v
	}
	result := logger.ValuesJoin(tany...)
	if result != expected {
		t.Fatalf("Expected %q but got %q", expected, result)
	}
}

func TestLoggerValuesJoinf(t *testing.T) {
	format := "Hello, %s"
	name := "Barlus"
	expected := "Hello, Barlus"
	result := logger.ValuesJoinf(format, name)
	if result != expected {
		t.Fatalf("Expected %q but got %q", expected, result)
	}
}
