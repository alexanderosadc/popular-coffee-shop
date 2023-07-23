package app

import "errors"

type CustomError struct {
	Err        error
	StatusCode int
	TimeToWait float64
}

var (
	ErrTooManyReq = CustomError{
		Err:        errors.New("too many requests"),
		StatusCode: 429,
	}
)
