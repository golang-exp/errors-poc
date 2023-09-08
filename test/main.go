package main

import (
	"log/slog"
	"os"

	"github.com/golang-exp/errors-poc/errors"
	errformatter "github.com/golang-exp/errors-poc/fmt"
)

type ErrorStackItem struct {
	Error   string                   `json:"error"`
	Details []map[string]interface{} `json:"details,omitempty"`
}

func ErrorStack(err error) []ErrorStackItem {
	if err == nil {
		return nil
	}
	errorStack := []ErrorStackItem{}
	for err != nil {
		errorStack = append(errorStack, ErrorStackItem{
			Error:   err.Error(),
			Details: errors.Details(err),
		})

		_, ok := err.(interface{ Unwrap() []error })
		if ok {
			// DO WE CLIMB THE TREE OR NOT???
			// NOT FOR NOW
			break
		}
		_, ok = err.(interface{ Unwrap() error })
		if !ok {
			// DO NOT KNOW WHAT TO DO HERE
			break
		}

		err = errors.Unwrap(err)
	}
	return errorStack
}

func main() {
	l := errors.NewStackLevelVar(errors.LevelFullStack)
	errorer := errors.NewErrorer(l)
	ff := errformatter.NewErrorFormatter(l)
	err0 := errorer.New("error 0")
	err1 := ff.Errorf("error 1 - %w", err0)
	err2 := ff.Errorf("error 2 - %w", err1)
	err3 := ff.Errorf("error 3 - %w", err2)

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, nil)))

	slog.Info(err0.Error(),
		"errorstack", ErrorStack(err0),
	)
	slog.Info(err3.Error(),
		"errorstack", ErrorStack(err3),
	)
}
