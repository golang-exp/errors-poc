package errors

import (
	"sync/atomic"

	errorsinternal "github.com/golang-exp/errors-poc/internal/errors"
)

type errorer struct {
	StackLevel errorsinternal.StackLeveler
}

var defaultErrorer atomic.Value

func init() {
	defaultErrorer.Store(new(errorer))
}

// Default returns the default Logger.
func Default() *errorer { return defaultErrorer.Load().(*errorer) }

func NewErrorer(s errorsinternal.StackLeveler) *errorer {
	e := new(errorer)
	e.StackLevel = s
	return e
}

func SetDefault(e *errorer) {
	defaultErrorer.Store(e)
}

func (e *errorer) New(s string) error {
	return errorsinternal.NewErr(e.StackLevel, 4, s)
}
