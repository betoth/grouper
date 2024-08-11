package controller

import (
	"encoding/json"
	"grouper/adapter/input/converter"
	"grouper/adapter/input/model/requests"
	"grouper/adapter/input/model/responses"
	"grouper/adapter/input/response"
	"grouper/application/dto"
	"grouper/application/port/input"
	"grouper/application/util/security"
	"grouper/config/logger"
	"grouper/config/rest_errors"
	"grouper/config/validation"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func NewGroupController(serviceInterface input.GroupService) GroupController {

	return &groupController{
		service: serviceInterface,
	}
}

type GroupController interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetGroups(w http.ResponseWriter, r *http.Request)
	Join(w http.ResponseWriter, r *http.Request)
	Leave(w http.ResponseWriter, r *http.Request)
}

type groupController struct {
	service input.GroupService
}

func (ctrl *groupController) Create(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Init createGroup controller", zap.String("journey", "CreateGroup"))
	var groupRequest requests.Group

	userID, restErr := security.NewJwtToken().ExtractUserID(r)
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

	validationErr := validation.ValidateRequest(&groupRequest)
	if validationErr != nil {
		logger.Error("Validation failed", validationErr, zap.String("journey", "CreateGroup"))
		response.JSON(w, http.StatusBadRequest, validationErr)
		return
	}

	groupDomain := converter.ConvertGroupRequestToDomain(&groupRequest)
	groupDomain.UserID = userID

	domainResult, restErr := ctrl.service.CreateGroup(groupDomain)
	if restErr != nil {
		logger.Error("Error trying to call CreateGroup service", restErr, zap.String("journey", "CreateGroup"))
		response.JSON(w, restErr.Code, restErr)
		return
	}

	logger.Debug("finish createGroup controller", zap.String("journey", "CreateGroup"))
	groupResponse := converter.ConvertGroupDtoToResponse(domainResult)
	response.JSON(w, http.StatusCreated, groupResponse)
}

func (ctrl *groupController) Join(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Init Join controller", zap.String("journey", "JoinGroup"))

	parameter := mux.Vars(r)
	groupID := parameter["groupId"]

	userID, err := security.NewJwtToken().ExtractUserID(r)
	if err != nil {
		logger.Error("Error trying to extract userID from token", err, zap.String("journey", "JoinGroup"))
		response.JSON(w, http.StatusBadRequest, rest_errors.NewBadRequestError("Error trying to extrac from token"))
		return
	}

	errRest := ctrl.service.JoinGroup(userID, groupID)
	if errRest != nil {
		logger.Error(errRest.Err, err, zap.String("journey", "JoinGroup"))
		response.JSON(w, errRest.Code, errRest)
		return
	}
	logger.Debug("Finish Join controller", zap.String("journey", "JoinGroup"))
	response.JSON(w, http.StatusCreated, nil)
}

func (ctrl *groupController) Leave(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Init Leave controller", zap.String("journey", "LeaveGroup"))

	parameter := mux.Vars(r)
	groupID := parameter["groupId"]
	if groupID == "" {
		restErr := rest_errors.NewBadRequestError("groupId is required")
		logger.Error("groupId is missing", nil, zap.String("journey", "LeaveGroup"))
		response.JSON(w, restErr.Code, restErr)
		return
	}

	userID, err := security.NewJwtToken().ExtractUserID(r)
	if err != nil {
		logger.Error("Error trying to extract userID from token", err, zap.String("journey", "LeaveGroup"))
		response.JSON(w, http.StatusBadRequest, rest_errors.NewBadRequestError("Error trying to extrac from token"))
		return
	}

	errRest := ctrl.service.LeaveGroup(userID, groupID)
	if errRest != nil {
		logger.Error("Error trying LeaveService", err, zap.String("journey", "LeaveGroup"))
		response.JSON(w, errRest.Code, errRest)
		return
	}

	logger.Debug("Finish Leave controller", zap.String("journey", "LeaveGroup"))
	response.JSON(w, http.StatusCreated, nil)
}

func (ctrl *groupController) GetGroups(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Init GetGroups controller", zap.String("journey", "GetGroups"))

	param := dto.GetGroupsParameter{
		Name: r.URL.Query().Get("name"),
	}

	groupsDomain, errRest := ctrl.service.GetGroups(param)
	if errRest != nil {
		logger.Error("Error trying GetGroupService", errRest, zap.String("journey", "GetGroups"))
		response.JSON(w, errRest.Code, errRest)
		return
	}

	groups := make([]responses.Group, len(*groupsDomain))

	for i, groupDomain := range *groupsDomain {
		groups[i] = converter.ConvertGroupDtoToResponse(&groupDomain)
	}

	logger.Debug("Finish GetGroups controller", zap.String("journey", "GetGroups"))
	response.JSON(w, http.StatusOK, groups)
}