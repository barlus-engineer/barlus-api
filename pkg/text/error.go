package text

import "errors"

var (
	ErrDotenvLoad = errors.New("dotenv: Error loading .env, using only system env")

	ErrGetenvErrConvType = errors.New("getenv: field '%s', Error %v")
	ErrGetenvUnsupportType = errors.New("getenv: field '%s', unsuported type '%v'")

	ErrTypeConvUnsupportType = errors.New("typeconv: unsupported type")

	ErrConfigGetenv = errors.New("config: error getting env:\n%v")
)
