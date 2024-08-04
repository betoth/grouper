package rest_errors

import "net/http"

type RestErr struct {
	Code    int      `json:"code,omitempty"`
	Err     string   `json:"error,omitempty"`
	Message string   `json:"message,omitempty"`
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
		Code:    http.StatusBadRequest,
		Err:     "Bad request",
		Message: message,
	}
}

func NewUnauthorizedRequestError(message string) *RestErr {
	return &RestErr{
		Code:    http.StatusUnauthorized,
		Err:     "Unauthorized",
		Message: message,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Code:    http.StatusBadRequest,
		Err:     "Bad request",
		Message: message,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Code:    http.StatusInternalServerError,
		Err:     "Internal server error",
		Message: message,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Code:    http.StatusNotFound,
		Err:     "Not found",
		Message: message,
	}
}

func NewUnprocessableEntityError(message string) *RestErr {
	return &RestErr{
		Code:    http.StatusUnprocessableEntity,
		Err:     "Unprocessable entity",
		Message: message,
	}
}

func NewConflictError(message string) *RestErr {
	return &RestErr{
		Code:    http.StatusConflict,
		Err:     "Conflict error",
		Message: message,
	}
}

func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Code:    http.StatusUnauthorized,
		Err:     "Unauthorized",
		Message: message,
	}
}
