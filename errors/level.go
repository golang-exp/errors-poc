// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors

import (
	"sync/atomic"

	errorsinternal "github.com/golang-exp/errors-poc/internal/errors"
)

// Names for common levels.
const (
	LevelNoStack        errorsinternal.StackLevel = errorsinternal.LevelNoStack
	LevelProgramCounter errorsinternal.StackLevel = errorsinternal.LevelProgramCounter
	LevelStackOnFirst   errorsinternal.StackLevel = errorsinternal.LevelStackOnFirst
	LevelFullStack      errorsinternal.StackLevel = errorsinternal.LevelFullStack
)

type stackLevelVar struct {
	val atomic.Int32
}

func NewStackLevelVar(sl ...errorsinternal.StackLevel) *stackLevelVar {
	slv := new(stackLevelVar)
	if len(sl) == 0 {
		return slv
	}
	slv.Set(sl[0])
	return slv
}

// Level returns v's level.
func (v *stackLevelVar) Level() errorsinternal.StackLevel {
	return errorsinternal.StackLevel(int(v.val.Load()))
}

// Set sets v's level to l.
func (v *stackLevelVar) Set(l errorsinternal.StackLevel) {
	v.val.Store(int32(l))
}
