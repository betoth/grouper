package controller

import (
	"encoding/json"
	"grouper/adapter/input/controller/response"
	"grouper/adapter/input/converter"
	"grouper/adapter/input/model/request"
	"grouper/application/port/input"
	"grouper/config/logger"
	"grouper/config/rest_errors"
	"grouper/config/validation"
	"net/http"

	"go.uber.org/zap"
)

func NewGroupControllerInterface(serviceInterface input.GroupDomainService) GroupControllerInterface {

	return &groupControllerInterface{
		service: serviceInterface,
	}
}

type GroupControllerInterface interface {
	CreateGroup(w http.ResponseWriter, r *http.Request)
}

type groupControllerInterface struct {
	service input.GroupDomainService
}

func (gc *groupControllerInterface) CreateGroup(w http.ResponseWriter, r *http.Request) {

	var groupRequest request.GroupRequest

	err := json.NewDecoder(r.Body).Decode(&groupRequest)
	if err != nil {
		logger.Error("Error trying to validate group info", err, zap.String("journey", "CreateGroup"))
		response.JSON(w, http.StatusBadRequest, rest_errors.NewBadRequestError("Error trying to validate user info"))
		return
	}

	validationErr := validation.ValidateUserRequest(&groupRequest)
	if validationErr != nil {
		logger.Error("Validation failed", validationErr, zap.String("journey", "CreateGroup"))
		response.JSON(w, http.StatusBadRequest, validationErr)
		return
	}

	groupDomain := converter.ConvertGroupRequestToDomain(&groupRequest)

	domainResult, restErr := gc.service.CreateGroupService(groupDomain)
	if restErr != nil {
		logger.Error("Error trying to call CreateGroup service", restErr, zap.String("journey", "CreateGroup"))
		response.JSON(w, restErr.Code, restErr)
		return
	}

	groupResponse := converter.ConvertGroupDomainToResponse(domainResult)
	response.JSON(w, http.StatusCreated, groupResponse)
}
