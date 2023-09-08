package errors

import (
	"runtime"
)

func getCaller(skip int) []uintptr {
	var pcs [1]uintptr
	runtime.Callers(skip, pcs[:])
	if len(pcs) > 0 {
        return pcs[:]
	}
	return nil
}

func getCallers(skip int) []uintptr {
	var pcs [32]uintptr
    n := runtime.Callers(skip, pcs[:])
	if len(pcs) > 0 {
        return pcs[:n]
	}
	return nil
}
