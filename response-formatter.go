package formatter

import "net/http"

type ResponseFormatter struct {
	Status  int
	Message string
	Data    interface{}
}

// Add message
const (
	successMessage = "Success"
	failMessage    = "Failed"
)

// Custom response formatter
func CustomFormatResponses(status int, message string, data interface{}) ResponseFormatter {
	return ResponseFormatter{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

// Success response formatter
func SuccessResponse(data interface{}) ResponseFormatter {
	return ResponseFormatter{
		Status:  http.StatusOK,
		Message: successMessage,
		Data:    data,
	}
}

// Created response formatter
func CreatedResponse(data interface{}) ResponseFormatter {
	return ResponseFormatter{
		Status:  http.StatusCreated,
		Message: successMessage,
		Data:    data,
	}
}

// Bad request response formatter
func BadRequestResponse(data interface{}) ResponseFormatter {
	return ResponseFormatter{
		Status:  http.StatusBadRequest,
		Message: failMessage,
		Data:    data,
	}
}

// Unauthorized response formatter
func UnauthorizedResponse(data interface{}) ResponseFormatter {
	return ResponseFormatter{
		Status:  http.StatusUnauthorized,
		Message: failMessage,
		Data:    data,
	}
}

// Forbidden response formatter
func ForbiddenResponse(data interface{}) ResponseFormatter {
	return ResponseFormatter{
		Status:  http.StatusForbidden,
		Message: failMessage,
		Data:    data,
	}
}

// Not found response formatter
func NotFoundResponse(data interface{}) ResponseFormatter {
	return ResponseFormatter{
		Status:  http.StatusNotFound,
		Message: failMessage,
		Data:    data,
	}
}

// Internal server error response formatter
func InternalServerErrorResponse(data interface{}) ResponseFormatter {
	return ResponseFormatter{
		Status:  http.StatusInternalServerError,
		Message: failMessage,
		Data:    data,
	}
}

// Bad gateway response formatter
func BadGatewayResponse(data interface{}) ResponseFormatter {
	return ResponseFormatter{
		Status:  http.StatusBadGateway,
		Message: failMessage,
		Data:    data,
	}
}
