package errors

import "net/http"

// NotFoundError represents a not found error indicating that the requested resource was not found.
type NotFoundError struct {
	StatusCode int
	Message    string
}

// NewNotFoundError creates a new instance of NotFoundError with the given message and status code.
func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}

// Error returns the error message for the not found error.
func (e *NotFoundError) Error() string {
	return e.Message
}

// Status returns the HTTP status code for the not found error.
func (e *NotFoundError) Status() int {
	return e.StatusCode
}
