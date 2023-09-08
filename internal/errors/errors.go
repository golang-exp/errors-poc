package errors

import "errors"

func (e *errorString) Error() string {
	return e.msg
}

func (e *errorWithStack) Error() string {
	return e.msg
}

func NewErr(sl StackLeveler, skip int, s string) error {
	if sl == nil {
		return &errorString{msg: s}
	}
	switch sl.Level() {
	case LevelNoStack:
		return &errorString{msg: s}
	case LevelProgramCounter:
		return &errorWithStack{
			msg: s,
			pcs:  getCaller(skip),
		}
	case LevelStackOnFirst, LevelFullStack:
		return &errorWithStack{
			msg: s,
			pcs: getCallers(skip),
		}
	default:
		return &errorString{s}
	}
}

func Unwrap(err error) error {
    return errors.Unwrap(err)
}


