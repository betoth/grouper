package controller

import (
	"encoding/json"
	"fmt"
	"grouper/adapter/input/controller/response"
	"grouper/adapter/input/converter"
	"grouper/adapter/input/model/request"
	"grouper/application/port/input"
	"grouper/config/logger"
	"grouper/config/rest_errors"
	"net/http"

	"go.uber.org/zap"
)

func NewLoginControllerInterface(serviceInterface input.LoginDomainService) LoginControllerInterface {
	return &loginControllerInterface{
		service: serviceInterface,
	}
}

type LoginControllerInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type loginControllerInterface struct {
	service input.LoginDomainService
}

func (lg *loginControllerInterface) Login(w http.ResponseWriter, r *http.Request) {

	var LoginRequest request.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&LoginRequest)
	if err != nil {
		logger.Error("Error trying to validate login info", err, zap.String("journey", "Login"))
		response.JSON(w, http.StatusBadRequest, rest_errors.NewBadRequestError("Error trying to validate login info"))
		return
	}

	LoginDomain := converter.ConvertLoginRequestToDomain(&LoginRequest)

	domainResult, restErr := lg.service.LoginServices(LoginDomain)
	if restErr != nil {
		logger.Error("Error trying to call Login service", restErr, zap.String("journey", "Login"))
		response.JSON(w, restErr.Code, restErr)
		return
	}

	fmt.Fprint(w, domainResult)

}
