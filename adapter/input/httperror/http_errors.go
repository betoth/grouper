package httperror

import (
	"net/http"
)

type restErrDetais struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type RestErr struct {
	Error restErrDetais `json:"error"`
}

func NewRestError(message string, status int) *RestErr {
	return &RestErr{
		Error: restErrDetais{
			Message: message,
			Status:  status,
		},
	}
}

func NewBadRequestError(message string) *RestErr {
	return NewRestError(message, http.StatusBadRequest)
}

func NewNotFoundError(message string) *RestErr {
	return NewRestError(message, http.StatusNotFound)
}

func NewInternalServerError(message string, err error) *RestErr {
	return NewRestError(message, http.StatusInternalServerError)
}

func NewConflictError(message string) *RestErr {
	return NewRestError(message, http.StatusConflict)
}
