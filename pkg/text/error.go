package text

import "errors"

var (
	ErrEncodeJson = errors.New("unable to encode data to JSON")
	ErrDecodeJson = errors.New("unable to decode JSON to data")
)