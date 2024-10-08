package controller

import (
	"encoding/json"
	"grouper/adapter/input/converter"
	"grouper/adapter/input/httperror"
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
	FindByID(w http.ResponseWriter, r *http.Request)
}

type groupController struct {
	service input.GroupService
}

func (ctrl *groupController) Create(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Init createGroup controller", zap.String("journey", "CreateGroup"))
	var groupRequest requests.Group

	// Extração do userID do token JWT
	userID, restErr := security.NewJwtToken().ExtractUserID(r)
	if restErr != nil {
		logger.Error("Error trying to extract userID from token", restErr, zap.String("journey", "CreateGroup"))
		httperror.MapAndRespond(w, restErr, "CreateGroup")
		return
	}

	// Decodificação do corpo da requisição para a struct Group
	err := json.NewDecoder(r.Body).Decode(&groupRequest)
	if err != nil {
		logger.Error("Error trying to decode group info", err, zap.String("journey", "CreateGroup"))
		httperror.MapAndRespond(w, err, "CreateGroup")
		return
	}

	// Validação dos dados da requisição
	validationErr := validation.ValidateRequest(&groupRequest)
	if validationErr != nil {
		logger.Error("Validation failed", validationErr, zap.String("journey", "CreateGroup"))
		httperror.MapAndRespond(w, validationErr, "CreateGroup")
		return
	}

	// Conversão da requisição para o domínio
	groupDomain := converter.ConvertGroupRequestToDomain(&groupRequest)
	groupDomain.UserID = userID

	// Chamada do serviço para criar o grupo
	domainResult, appErr := ctrl.service.CreateGroup(groupDomain)
	if appErr != nil {
		logger.Error("Error trying to call CreateGroup service", appErr, zap.String("journey", "CreateGroup"))
		httperror.ErrorToErrorResponse(w, appErr, "CreateGroup")
		return
	}

	// Resposta de sucesso
	logger.Debug("Finish createGroup controller", zap.String("journey", "CreateGroup"))
	groupResponse := converter.ConvertGroupDtoToResponse(domainResult)
	response.JSON(w, http.StatusCreated, groupResponse)
}

func (ctrl *groupController) Join(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Init Join controller", zap.String("journey", "JoinGroup"))

	// Extrair o parâmetro groupId da URL
	parameter := mux.Vars(r)
	groupID := parameter["groupId"]

	// Extrair userID do token JWT
	userID, err := security.NewJwtToken().ExtractUserID(r)
	if err != nil {
		logger.Error("Error trying to extract userID from token", err, zap.String("journey", "JoinGroup"))
		httperror.MapAndRespond(w, err, "JoinGroup")
		return
	}

	// Chamar o serviço para associar o usuário ao grupo
	errRest := ctrl.service.JoinGroup(userID, groupID)
	if errRest != nil {
		logger.Error("Error trying to call JoinGroup service", errRest, zap.String("journey", "JoinGroup"))
		httperror.ErrorToErrorResponse(w, errRest, "JoinGroup")
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

	userID, restErr := security.NewJwtToken().ExtractUserID(r)
	if restErr != nil {
		logger.Error("Error trying to extract userID from token", restErr, zap.String("journey", "LeaveGroup"))
		response.JSON(w, http.StatusBadRequest, rest_errors.NewBadRequestError("Error trying to extrac from token"))
		return
	}

	err := ctrl.service.LeaveGroup(userID, groupID)
	if err != nil {
		logger.Error("Error trying LeaveService", err, zap.String("journey", "LeaveGroup"))
		httperror.ErrorToErrorResponse(w, err, "LeaveGroup")
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

	groupsDomain, err := ctrl.service.GetGroups(param)
	if err != nil {
		logger.Error("Error trying GetGroupService", err, zap.String("journey", "GetGroups"))
		httperror.ErrorToErrorResponse(w, err, "LeaveGroup")
		return
	}

	groups := make([]responses.Group, len(*groupsDomain))

	for i, groupDomain := range *groupsDomain {
		groups[i] = converter.ConvertGroupDtoToResponse(&groupDomain)
	}

	logger.Debug("Finish GetGroups controller", zap.String("journey", "GetGroups"))
	response.JSON(w, http.StatusOK, groups)
}

func (ctrl *groupController) FindByID(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Init FindByID controller", zap.String("journey", "FindByID"))
	parameter := mux.Vars(r)
	groupID := parameter["groupId"]

	domainResult, err := ctrl.service.FindByID(groupID)
	if err != nil {
		logger.Error("Error trying to call FindByID service", err, zap.String("journey", "FindByID"))
		httperror.MapAndRespond(w, err, "FindByID")
		return
	}
	groupResponse := converter.ConvertGroupDtoToResponse(domainResult)
	logger.Debug("Finish FindByID controller", zap.String("journey", "FindByID"))
	response.JSON(w, http.StatusOK, groupResponse)
}
