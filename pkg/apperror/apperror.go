package apperror

import (
	"errors"
)

type AppError struct {
	Code    ErrorCode    `json:"code"`
	Message ErrorMessage `json:"message"`
	Errors  interface{}  `json:"errors"`
}

type ErrorModel struct {
	Field   string       `json:"field,omitempty"`
	Message ErrorMessage `json:"message"`
}

func (e AppError) Error() string {
	return string(e.Message)
}

func NewErrorMessage(v string) ErrorMessage {
	return ErrorMessage(v)
}

func IsAppError(err error) (e AppError, ok bool) {
	ok = errors.As(err, &e)
	return
}

func NewBadRequestError() error {
	code := ErrBadRequestCode
	return AppError{
		Code:    code,
		Message: Message[code],
		Errors:  []ErrorModel{},
	}
}

func NewUnavailableError() error {
	code := ErrBadRequestCode
	return AppError{
		Code:    code,
		Message: Message[code],
		Errors:  []ErrorModel{},
	}
}

func NewBadRequestWithFieldError(errs []ErrorModel) error {
	code := ErrBadRequestCode
	return AppError{
		Code:    code,
		Message: Message[code],
		Errors:  errs,
	}
}

func NewInternalServerError() error {
	code := ErrInternalServerError
	return AppError{
		Code:    code,
		Message: Message[code],
		Errors:  []ErrorModel{},
	}
}
