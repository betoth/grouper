package controller

import (
	"encoding/json"
	"grouper/adapter/input/converter"
	"grouper/adapter/input/model/dto"
	"grouper/adapter/input/model/request"
	resp "grouper/adapter/input/model/response"
	"grouper/adapter/input/response"
	"grouper/application/port/input"
	"grouper/application/util/secutiry"
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
	GetGroups(w http.ResponseWriter, r *http.Request)
	Join(w http.ResponseWriter, r *http.Request)
	Leave(w http.ResponseWriter, r *http.Request)
}

type groupControllerInterface struct {
	service input.GroupDomainService
}

func (gc *groupControllerInterface) CreateGroup(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Init createGroup controller", zap.String("journey", "CreateGroup"))
	var groupRequest request.GroupRequest

	userID, restErr := secutiry.NewJwtToken().ExtractUserID(r)
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

	logger.Debug("finish createGroup controller", zap.String("journey", "CreateGroup"))
	groupResponse := converter.ConvertGroupDomainToResponse(domainResult)
	response.JSON(w, http.StatusCreated, groupResponse)
}

func (gc *groupControllerInterface) Join(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Init Join controller", zap.String("journey", "JoinGroup"))

	parameter := mux.Vars(r)
	groupID := parameter["groupId"]

	userID, err := secutiry.NewJwtToken().ExtractUserID(r)
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
	logger.Debug("Finish Join controller", zap.String("journey", "JoinGroup"))
	response.JSON(w, http.StatusCreated, nil)
}

func (gc *groupControllerInterface) Leave(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Init Leave controller", zap.String("journey", "LeaveGroup"))

	parameter := mux.Vars(r)
	groupID := parameter["groupId"]

	userID, err := secutiry.NewJwtToken().ExtractUserID(r)
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
	logger.Debug("Finish Leave controller", zap.String("journey", "LeaveGroup"))
}

func (uc *groupControllerInterface) GetGroups(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Init GetGroups controller", zap.String("journey", "GetGroups"))

	param := dto.GetGroupsQueryParameter{
		User:  r.URL.Query().Get("user"),
		Topic: r.URL.Query().Get("topic"),
	}

	groupsDomain, errRest := uc.service.GetGroupsService(param)
	if errRest != nil {
		logger.Error("Error trying GetGroupService", errRest, zap.String("journey", "GetGroups"))
		response.JSON(w, errRest.Code, errRest)
		return
	}

	var groups []resp.GroupResponse

	for _, groupDomain := range *groupsDomain {
		groups = append(groups, converter.ConvertGroupDomainToResponse(&groupDomain))
	}

	logger.Debug("Finish GetGroups controller", zap.String("journey", "GetGroups"))
	response.JSON(w, http.StatusOK, groups)
}
