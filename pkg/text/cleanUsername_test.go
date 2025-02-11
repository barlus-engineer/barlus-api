package text_test

import (
	"testing"

	"github.com/barlus-engineer/barlus-api/pkg/text"
)

func TestCleanUsername(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"UserName123", "username123"},
		{"User.Name_123", "user.name_123"},
		{"User@Name!123", "username123"},
		{"User Name 123", "username123"},
		{"User-Name_123", "username_123"},
		{"UserName!@#$%^&*()123", "username123"},
	}

	for _, test := range tests {
		result := text.CleanUsername(test.input)
		if result != test.expected {
			t.Errorf("CleanUsername(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}