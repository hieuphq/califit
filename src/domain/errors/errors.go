package errors

import (
	"net/http"
)

// Error interface for error
type Error interface {
	StatusCode() int
	Error() string
}

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errInternalServerError{}

	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errNotFound{}

	// ErrConflict will throw if the current action already exists
	ErrConflict = errConflict{}

	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errBadParamInput{}
)

type errInternalServerError struct{}

func (errInternalServerError) StatusCode() int {
	return http.StatusInternalServerError
}

func (errInternalServerError) Error() string {
	return "Internal Server Error"
}

type errNotFound struct{}

func (errNotFound) StatusCode() int {
	return http.StatusNotFound
}

func (errNotFound) Error() string {
	return "Your requested Item is not found"
}

type errConflict struct{}

func (errConflict) StatusCode() int {
	return http.StatusBadRequest
}

func (errConflict) Error() string {
	return "Your Item already exist"
}

type errBadParamInput struct{}

func (errBadParamInput) StatusCode() int {
	return http.StatusBadRequest
}

func (errBadParamInput) Error() string {
	return "Given Param is not valid"
}

// StringError error for string
type StringError struct {
	err        string
	statusCode int
}

// NewStringError new a new string err
func NewStringError(err string, statusCode int) StringError {
	return StringError{
		err:        err,
		statusCode: statusCode,
	}
}

// StatusCode implementation for error interface
func (e StringError) StatusCode() int {
	return e.statusCode
}

func (e StringError) Error() string {
	return e.err
}
