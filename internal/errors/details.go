package errors

import "runtime"

type Detailer interface {
	Details() []map[string]interface{}
}

func (e *errorWithStack) Details() []map[string]interface{} {
	if e == nil || len(e.pcs) == 0 {
		return nil
	}
	return defaultDetails(e.pcs)
}

func (e *wrapErrorWithStack) Details() []map[string]interface{} {
    	if e == nil || len(e.pcs) == 0 {
		return nil
	}
	return defaultDetails(e.pcs)
}

func (e *wrapErrorsWithStack) Details() []map[string]interface{} {
	if e == nil || len(e.pcs) == 0 {
		return nil
	}
	return defaultDetails(e.pcs)
}

func defaultDetails(pcs []uintptr) []map[string]interface{} {
	var details []map[string]interface{}
	fs := runtime.CallersFrames(pcs)
    for f, more := fs.Next(); more; f, more = fs.Next() {
		details = append(details, map[string]interface{}{
			"func": f.Func.Name(),
			"file": f.File,
			"line": f.Line,
		})
	}
	return details
}
