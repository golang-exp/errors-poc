// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors

// Names for common levels.
const (
	LevelNoStack        StackLevel = 0
	LevelProgramCounter StackLevel = 1
	LevelStackOnFirst   StackLevel = 2
	LevelFullStack      StackLevel = 3
)

type StackLevel int

func (s StackLevel) Level() StackLevel {
	return s
}

type StackLeveler interface {
	Level() StackLevel
}
