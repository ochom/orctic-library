package errors

// CustomError is a custom error type
type CustomError struct {
	Code    int
	Message string
}

// Error returns the error message
func (e *CustomError) Error() string {
	return e.Message
}

// NotFoundError is returned when a resource is not found
func NotFoundError(message string) *CustomError {
	return &CustomError{
		Code:    404,
		Message: message,
	}
}

// BadRequestError is returned when a request is invalid
func BadRequestError(message string) *CustomError {
	return &CustomError{
		Code:    400,
		Message: message,
	}
}

// UnauthorizedError is returned when a request is unauthorized
func UnauthorizedError(message string) *CustomError {
	return &CustomError{
		Code:    401,
		Message: message,
	}
}

// ForbiddenError is returned when a request is forbidden
func ForbiddenError(message string) *CustomError {
	return &CustomError{
		Code:    403,
		Message: message,
	}
}

// ConflictError is returned when a conflict occurs
func ConflictError(message string) *CustomError {
	return &CustomError{
		Code:    409,
		Message: message,
	}
}

// InternalServerError is returned when an internal server error occurs
func InternalServerError(message string) *CustomError {
	return &CustomError{
		Code:    500,
		Message: message,
	}
}

// ServiceUnavailableError is returned when a service is unavailable
func ServiceUnavailableError(message string) *CustomError {
	return &CustomError{
		Code:    503,
		Message: message,
	}
}
