package text_test

import (
	"testing"

	"github.com/barlus-engineer/barlus-api/pkg/text"
)

func TestCleanEmail(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Barluscuda@gmail.com /(+=-)", "barluscuda@gmail.com"},
		{"Example@Domain.com", "example@domain.com"},
		{"Test.Email+123@Example.com", "test.email123@example.com"},
		{"UserName!@Example.com", "username@example.com"},
		{"User Name@Example.com", "username@example.com"},
	}

	for _, test := range tests {
		result := text.CleanEmail(test.input)
		if result != test.expected {
			t.Errorf("CleanEmail(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}