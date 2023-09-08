package errors

import errorsinternal "github.com/golang-exp/errors-poc/internal/errors"

func StackTrace(err error) []uintptr {
	if e, ok := err.(errorsinternal.StackTracer); ok {
		return e.StackTrace()
	}
	return nil
}
