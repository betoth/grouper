package controller

import (
	"encoding/json"
	"grouper/adapter/input/controller/response"
	"grouper/adapter/input/converter"
	"grouper/adapter/input/model/request"
	"grouper/application/port/input"
	util "grouper/application/util/secutiry"
	"grouper/config/logger"
	"grouper/config/rest_errors"
	"grouper/config/validation"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func NewGroupControllerInterface(serviceInterface input.GroupDomainService) GroupControllerInterface {

	return &groupControllerInterface{
		service: serviceInterface,
	}
}

type GroupControllerInterface interface {
	CreateGroup(w http.ResponseWriter, r *http.Request)
	Join(w http.ResponseWriter, r *http.Request)
	Leave(w http.ResponseWriter, r *http.Request)
}

type groupControllerInterface struct {
	service input.GroupDomainService
}

func (gc *groupControllerInterface) CreateGroup(w http.ResponseWriter, r *http.Request) {
	logger.Info("Init createGroup controller", zap.String("journey", "CreateGroup"))
	var groupRequest request.GroupRequest

	userID, restErr := util.NewJwtToken().ExtractUserID(r)
	if restErr != nil {
		logger.Error("Error trying to extract userID from token", restErr, zap.String("journey", "CreateGroup"))
		response.JSON(w, restErr.Code, rest_errors.NewBadRequestError("Error trying to extrac from token"))
		return
	}

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
	groupDomain.UserID = userID

	domainResult, restErr := gc.service.CreateGroupService(groupDomain)
	if restErr != nil {
		logger.Error("Error trying to call CreateGroup service", restErr, zap.String("journey", "CreateGroup"))
		response.JSON(w, restErr.Code, restErr)
		return
	}

	groupResponse := converter.ConvertGroupDomainToResponse(domainResult)
	response.JSON(w, http.StatusCreated, groupResponse)
}

func (gc *groupControllerInterface) Join(w http.ResponseWriter, r *http.Request) {
	logger.Info("Init Join controller", zap.String("journey", "JoinGroup"))

	parameter := mux.Vars(r)
	groupID := parameter["groupId"]

	userID, err := util.NewJwtToken().ExtractUserID(r)
	if err != nil {
		logger.Error("Error trying to extract userID from token", err, zap.String("journey", "JoinGroup"))
		response.JSON(w, http.StatusBadRequest, rest_errors.NewBadRequestError("Error trying to extrac from token"))
		return
	}

	errRest := gc.service.JoinService(userID, groupID)
	if errRest != nil {
		logger.Error(errRest.Err, err, zap.String("journey", "JoinGroup"))
		response.JSON(w, errRest.Code, errRest)
		return
	}
	response.JSON(w, http.StatusCreated, nil)
}

func (gc *groupControllerInterface) Leave(w http.ResponseWriter, r *http.Request) {
	logger.Info("Init Leave controller", zap.String("journey", "LeaveGroup"))

	parameter := mux.Vars(r)
	groupID := parameter["groupId"]

	userID, err := util.NewJwtToken().ExtractUserID(r)
	if err != nil {
		logger.Error("Error trying to extract userID from token", err, zap.String("journey", "LeaveGroup"))
		response.JSON(w, http.StatusBadRequest, rest_errors.NewBadRequestError("Error trying to extrac from token"))
		return
	}

	errRest := gc.service.LeaveService(userID, groupID)
	if errRest != nil {
		logger.Error("Error trying LeaveService", err, zap.String("journey", "LeaveGroup"))
		response.JSON(w, errRest.Code, errRest)
		return
	}
	response.JSON(w, http.StatusCreated, nil)
}
