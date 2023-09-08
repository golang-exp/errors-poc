package errors

import (
	"sync/atomic"

	errorsinternal "github.com/golang-exp/errors-poc/internal/errors"
)

var defaultDetailer atomic.Value

func init() {
	d := new(detailer)
	d.detailFunc = func(err error) []map[string]interface{} {
		e, ok := err.(errorsinternal.Detailer)
		if ok {
			return e.Details()
		}
		return nil
	}
	defaultDetailer.Store(d)
}

type detailer struct {
	detailFunc func(err error) []map[string]interface{}
}

// Default returns the default Logger.
func DefaultDetailer() *detailer { return defaultDetailer.Load().(*detailer) }

func NewDetailer(f func(err error) []map[string]interface{}) *detailer {
	d := new(detailer)
	d.detailFunc = f
	return d
}

func SetDefaultDetailer(e *detailer) {
	defaultDetailer.Store(e)
}

func Details(err error) []map[string]interface{} {
	return DefaultDetailer().detailFunc(err)
}

func (d *detailer) Details(err error) []map[string]interface{} {
	return d.detailFunc(err)
}
