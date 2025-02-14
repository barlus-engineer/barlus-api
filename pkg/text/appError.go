package text

import "errors"

type AppError struct {
	SvcErr error
	AppErr error
}

func NewAppErr(err error, appErr string) AppError {
	return AppError{
		SvcErr: err,
		AppErr: errors.New(appErr),
	}
}