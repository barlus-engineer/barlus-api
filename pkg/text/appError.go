package text

import "errors"

type AppError struct {
	svcErr error
	appErr error
}

func NewAppErr(err error, appErr string) AppErrorStruct {
	return AppError{
		svcErr: err,
		appErr: errors.New(appErr),
	}
}