package controller

import (
	"encoding/json"
	"fmt"
	"grouper/adapter/input/controller/response"
	"grouper/adapter/input/converter"
	"grouper/adapter/input/model/request"
	"grouper/application/port/input"
	"grouper/configuration/logger"
	"net/http"

	"go.uber.org/zap"
)

func NewUserControllerInterface(serviceInterface input.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type userControllerInterface struct {
	service input.UserDomainService
}

func (uc *userControllerInterface) CreateUser(w http.ResponseWriter, r *http.Request) {
	logger.Info("Init createUser controller",
		zap.String("journey", "CreateUSer"),
	)
	var userRequest request.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		logger.Error("Error trying to marshal object", err)
		fmt.Fprint(w, err)
		return
	}

	userDomain := converter.ConvertRequestToDomain(&userRequest)
	domainResult, err := uc.service.CreateUserServices(userDomain)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	userResponse := converter.ConvertDomainToResponse(domainResult)

	response.JSON(w, http.StatusCreated, userResponse)

}
