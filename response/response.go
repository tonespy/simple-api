package response

// GenericResponse :- Means to construct a response with generic data
func GenericResponse(statusCode int, message string, data interface{}) *APIResponse {
	return &APIResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}
}
