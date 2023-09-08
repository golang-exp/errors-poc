package errors

type errorString struct {
	msg string
}

type errorWithStack struct {
	msg string
	pcs []uintptr
}

type wrapError struct {
	err error
	msg string
}

type wrapErrors struct {
	msg  string
	errs []error
}

type wrapErrorWithStack struct {
	err error
	msg string
	pcs []uintptr
}

type wrapErrorsWithStack struct {
	msg  string
	errs []error
	pcs  []uintptr
}
