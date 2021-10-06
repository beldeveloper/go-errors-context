package errors

import (
	"fmt"
	"strings"
)

// Params defines the type for storing the context parameters.
type Params map[string]interface{}

// Context defines the context of the error.
type Context struct {
	Path   string
	Params Params
}

// ErrorWithContext implements the error with the context.
type ErrorWithContext struct {
	e error
	c Context
}

// Error returns the error message with attached context data.
func (err ErrorWithContext) Error() string {
	var initErr error
	path := make([]string, 0, 4)
	params := make(Params)
	goThrough(err, func(err error, c *Context) bool {
		if c != nil {
			path = append(path, c.Path)
			for k, v := range c.Params {
				params[k] = v
			}
		} else {
			initErr = err
		}
		return true
	})
	parts := make([]string, 0, 2+len(params))
	if initErr != nil {
		parts = append(parts, initErr.Error())
	}
	for k, v := range params {
		parts = append(parts, fmt.Sprintf("%s=%v", k, v))
	}
	if len(path) > 0 {
		parts = append(parts, strings.Join(path, " -> "))
	}
	return strings.Join(parts, "; ")
}
