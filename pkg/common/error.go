package common

import "fmt"

// Error represents an error from a service.
type Error struct {
	Code    string
	Message string
}

// NewError makes a new error.
func NewError(code string) Error {
	return Error{
		Code: code,
	}
}

// Errorf makes a new error with a message.
func Errorf(code string, message string, v ...interface{}) Error {
	return Error{
		Code:    code,
		Message: fmt.Sprintf(message, v),
	}
}

func (e Error) Error() string {
	if e.Message == "" {
		return e.Code
	}
	return e.Code + ": " + e.Message
}
