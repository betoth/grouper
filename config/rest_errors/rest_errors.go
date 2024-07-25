package rest_errors

import "net/http"

type RestErr struct {
	Message string   `json:"message,omitempty"`
	Err     string   `json:"error,omitempty"`
	Code    int      `json:"code,omitempty"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad request",
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorizedRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Internal server error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Not found",
		Code:    http.StatusNotFound,
	}
}

func NewUnprocessableEntityError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Unprocessable entity",
		Code:    http.StatusUnprocessableEntity,
	}
}

func NewConflictError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Conflict error",
		Code:    http.StatusConflict,
	}
}
