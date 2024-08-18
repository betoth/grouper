package http_errors

import (
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewRestError(message string, status int, err string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  status,
		Error:   err,
	}
}

func NewBadRequestError(message string) *RestErr {
	return NewRestError(message, http.StatusBadRequest, "bad_request")
}

func NewNotFoundError(message string) *RestErr {
	return NewRestError(message, http.StatusNotFound, "not_found")
}

func NewInternalServerError(message string, err error) *RestErr {
	return NewRestError(message, http.StatusInternalServerError, "internal_server_error")
}

func NewConflictError(message string) *RestErr {
	return NewRestError(message, http.StatusConflict, "conflict")
}
