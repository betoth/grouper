package controller

import (
	"encoding/json"
	"grouper/adapter/input/controller/response"
	"grouper/adapter/input/converter"
	"grouper/adapter/input/model/request"
	"grouper/application/port/input"
	"grouper/config/logger"
	"grouper/config/rest_errors"
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
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		response.JSON(w, http.StatusBadRequest, rest_errors.NewBadRequestError("Error trying to validate user info"))
		return
	}

	userDomain := converter.ConvertRequestToDomain(&userRequest)

	domainResult, restErr := uc.service.CreateUserServices(userDomain)
	if restErr != nil {
		logger.Error(
			"Error trying to call CreateUser service",
			restErr,
			zap.String("journey", "createUser"))
		response.JSON(w, restErr.Code, restErr)
		return
	}

	userResponse := converter.ConvertDomainToResponse(domainResult)

	response.JSON(w, http.StatusCreated, userResponse)

}
