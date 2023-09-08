package errors

import (
	"fmt"
)

func (e *wrapError) Error() string {
	return e.msg
}

func (e *wrapErrors) Error() string {
	return e.msg
}

func (e *wrapErrorWithStack) Error() string {
	return e.msg
}

func (e *wrapErrorsWithStack) Error() string {
	return e.msg
}

func (e *wrapError) Unwrap() error {
	return e.err
}

func (e *wrapErrors) Unwrap() []error {
	return e.errs
}

func (e *wrapErrorWithStack) Unwrap() error {
	return e.err
}

func (e *wrapErrorsWithStack) Unwrap() []error {
	return e.errs
}

// USABLE AND UGLY IMPLEMENTATION - JUST FOR TESTING
func Errorf(sl StackLeveler, skip int, format string, a ...any) error {
	fmtErrs := fmt.Errorf(format, a...)
	if Unwrap(fmtErrs) == nil {
		return NewErr(sl, skip+1, fmt.Sprint(format, a))
	}
	if e, ok := fmtErrs.(interface {
		Unwrap() error
		Error() string
	}); ok {
		msg := fmtErrs.Error()
		switch sl.Level() {
		case LevelNoStack:
			return fmtErrs
		case LevelProgramCounter, LevelStackOnFirst:
			return &wrapErrorWithStack{
				msg: msg,
				err: e.Unwrap(),
				pcs: getCaller(skip),
			}
		case LevelFullStack:
			return &wrapErrorWithStack{
				msg: msg,
				err: e.Unwrap(),
				pcs: getCallers(skip),
			}
		}
	}
	if e, ok := fmtErrs.(interface {
		Unwrap() []error
		Error() string
	}); ok {
		msg := fmtErrs.Error()
		switch sl.Level() {
		case LevelNoStack:
			return fmtErrs
		case LevelProgramCounter, LevelStackOnFirst:
			return &wrapErrorsWithStack{
				msg:  msg,
				errs: e.Unwrap(),
				pcs:  getCaller(skip),
			}
		case LevelFullStack:
			return &wrapErrorsWithStack{
				msg:  msg,
				errs: e.Unwrap(),
				pcs:  getCallers(skip),
			}
		}
	}
	return fmtErrs
}
