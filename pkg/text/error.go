package text

import "errors"

var (
	ErrDotenvLoad = errors.New("dotenv: Error loading .env, using only system env")
	ErrReadConfigFile = errors.New("config: error reading config file")
	ErrDecodeConfig   = errors.New("config: error decoding into struct")
)