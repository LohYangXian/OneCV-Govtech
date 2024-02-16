package errors

import "net/http"

// BadRequestError represents a bad request error indicating missing or invalid request data.
type BadRequestError struct {
	StatusCode int
	Message    string
}

// NewBadRequestError creates a new instance of BadRequestError with the given message and status code.
func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}

// Error returns the error message for the bad request error.
func (e *BadRequestError) Error() string {
	return e.Message
}

// Status returns the HTTP status code for the bad request error.
func (e *BadRequestError) Status() int {
	return e.StatusCode
}
