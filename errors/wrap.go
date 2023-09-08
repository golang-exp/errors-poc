package errors

import (
	"errors"

	errorsinternal "github.com/golang-exp/errors-poc/internal/errors"
)

// Wraps standard errors.Is
func Is(err error, target error) bool {
	return errors.Is(err, target)
}

// Wraps standard errors.As
func As(err error, target any) bool {
	return errors.As(err, target)
}

func Unwrap(err error) error {
	return errorsinternal.Unwrap(err)
}
