package errors

type StackTracer interface {
    StackTrace() []uintptr
}

func (e *errorWithStack) StackTrace() []uintptr {
	return e.pcs
}

func (e *wrapErrorWithStack) StackTrace() []uintptr {
	return e.pcs
}

func (e *wrapErrorsWithStack) StackTrace() []uintptr {
	return e.pcs
}
