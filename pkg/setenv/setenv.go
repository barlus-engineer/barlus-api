package setenv

import (
	"os"

	"github.com/barlus-engineer/barlus-api/pkg/typeconv"
)

func Set(key string, value any) {
	envValue := typeconv.Any2Str(value)
	os.Setenv(key, envValue)
}

// ===