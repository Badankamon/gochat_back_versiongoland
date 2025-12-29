package errors

import "net/http"

type AppError struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	HttpCode int    `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

func New(code int, message string, httpCode int) *AppError {
	return &AppError{
		Code:     code,
		Message:  message,
		HttpCode: httpCode,
	}
}

// Common Errors
var (
	ErrInternalServer = New(1001, "Internal Server Error", http.StatusInternalServerError)
	ErrBadRequest     = New(1002, "Bad Request", http.StatusBadRequest)
	ErrNotFound       = New(1003, "Not Found", http.StatusNotFound)
	ErrUnauthorized   = New(1004, "Unauthorized", http.StatusUnauthorized)
	ErrForbidden      = New(1005, "Forbidden", http.StatusForbidden)
)
