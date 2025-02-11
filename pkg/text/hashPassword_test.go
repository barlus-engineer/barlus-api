package text_test

import (
	"testing"

	"github.com/barlus-engineer/barlus-api/pkg/text"
)

func TestHashPassword(t *testing.T) {
	password := "password_1234"
	hashedPassword := text.HashPassword(password)

	if hashedPassword == "" {
		t.Error("Expected hashed password to be non-empty")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "password_1234"
	hashedPassword := text.HashPassword(password)

	if !text.CheckPasswordHash(password, hashedPassword) {
		t.Error("Expected password to match hashed password")
	}

	wrongPassword := "wrongpassword"
	if text.CheckPasswordHash(wrongPassword, hashedPassword) {
		t.Error("Expected wrong password to not match hashed password")
	}
}
