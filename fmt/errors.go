package fmt

import (
	"sync/atomic"

	errorsinternal "github.com/golang-exp/errors-poc/internal/errors"
)

type errorFormatter struct {
	StackLevel errorsinternal.StackLeveler
}

var defaultErrorFormatter atomic.Value

func init() {
	defaultErrorFormatter.Store(new(errorFormatter))
}

func Default() *errorFormatter { return defaultErrorFormatter.Load().(*errorFormatter) }

func NewErrorFormatter(s errorsinternal.StackLeveler) *errorFormatter {
	e := new(errorFormatter)
	e.StackLevel = s
	return e
}

func SetDefault(e *errorFormatter) {
	defaultErrorFormatter.Store(e)
}

// THIS IS NOT USABLE - JUST FOR TESTING
func Errorf(format string, a ...any) error {
	return errorsinternal.Errorf(Default().StackLevel, 4, format, a...)
}

// THIS IS NOT USABLE - JUST FOR TESTING
func (ef *errorFormatter) Errorf(format string, a ...any) error {
	return errorsinternal.Errorf(ef.StackLevel, 4, format, a...)
}

