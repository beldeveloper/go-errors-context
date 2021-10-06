package errors

import "errors"

// NewWithContext creates a new error with the context data.
func NewWithContext(text string, c Context) error {
	return WrapContext(errors.New(text), c)
}

// WrapContext wraps the existing error with the context data.
func WrapContext(err error, c Context) error {
	if err != nil {
		return ErrorWithContext{e: err, c: c}
	}
	return nil
}

// Is checks if the error wraps the target one.
func Is(err, target error) (res bool) {
	goThrough(err, func(err error, c *Context) bool {
		if errors.Is(err, target) {
			res = true
			return false // break the goThrough loop
		}
		return true // continue
	})
	return
}

func goThrough(err error, f func(err error, c *Context) bool) {
	var proceed bool
	for err != nil {
		errWithContext, ok := err.(ErrorWithContext)
		if ok {
			proceed = f(err, &errWithContext.c)
			err = errWithContext.e
		} else {
			proceed = f(err, nil)
			err = nil
		}
		if !proceed {
			break
		}
	}
}
