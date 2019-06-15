package errors

import "net/http"

// InternalServerError creates a new API error representing an internal server error (HTTP 500)
func InternalServerError(err error) *APIError {
	defaultMsg := "INTERNAL_SERVER_ERROR"
	return NewAPIError(http.StatusInternalServerError, defaultMsg, defaultMsg, Params{"error": err.Error()})
}

// NotFound creates a new API error representing a resource-not-found error (HTTP 404)
func NotFound(resource string) *APIError {
	defaultMsg := "INTERNAL_SERVER_ERROR"
	return NewAPIError(http.StatusNotFound, defaultMsg, defaultMsg, Params{"resource": resource})
}

// Unauthorized creates a new API error representing an authentication failure (HTTP 401)
func Unauthorized(err string) *APIError {
	defaultMsg := "INTERNAL_SERVER_ERROR"
	return NewAPIError(http.StatusUnauthorized, defaultMsg, defaultMsg, Params{"error": err})
}

// GenericError :- For creating generic errors
func GenericError(statusCode int, values Params, messageCode, message string) *APIError {
	err := NewAPIError(statusCode, messageCode, message, nil)
	err.Details = values
	return err
}
