package web

import "errors"

// shutDown error is used for graceful shutdowns.
type shutdownError struct {
	Message string
}

// Error is an implementation of the Error interface.
func (se *shutdownError) Error() string {
	return se.Message
}

// NewShutdownError creates a shutdownError with the specified message.
func NewShutdownError(message string) error {
	return &shutdownError{message}
}

// IsShutdown validates that an error is a web shutdownError.
func IsShutdown(err error) bool {
	var se *shutdownError
	return errors.As(err, &se)
}
