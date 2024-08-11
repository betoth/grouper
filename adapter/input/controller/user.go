package controller

import (
	"encoding/json"
	"grouper/adapter/input/converter"
	"grouper/adapter/input/model/requests"
	"grouper/adapter/input/model/responses"
	"grouper/adapter/input/response"
	"grouper/application/port/input"
	"grouper/application/util/security"
	"grouper/config/logger"
	"grouper/config/rest_errors"
	"grouper/config/validation"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func NewUserController(serviceInterface input.UserService) UserController {
	return &userController{
		service: serviceInterface,
	}
}

type UserController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	GetGroups(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	service input.UserService
}

func (ctrl *userController) Create(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Init CreateUser controller", zap.String("journey", "CreateUser"))

	var userRequest requests.User

	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		restErr := rest_errors.NewBadRequestError(err.Error())
		logger.Error("Error trying to validate user info", err, zap.String("journey", "CreateUser"))
		response.JSON(w, restErr.Code, restErr)
		return
	}

	validationErr := validation.ValidateRequest(&userRequest)
	if validationErr != nil {
		logger.Error("Validation failed", validationErr, zap.String("journey", "CreateUser"))
		response.JSON(w, http.StatusBadRequest, validationErr)
		return
	}

	userDomain := converter.ConvertUserRequestToDomain(&userRequest)

	restErr := checkUsernameAvailability(ctrl, userDomain.Username)
	if restErr != nil {
		logger.Error("Username validation failed", restErr, zap.String("journey", "CreateUser"))
		response.JSON(w, restErr.Code, err)
		return
	}

	restErr = checkEmailAvailability(ctrl, userDomain.Email)
	if restErr != nil {
		logger.Error("Email validation failed", restErr, zap.String("journey", "CreateUser"))
		response.JSON(w, restErr.Code, restErr)
		return
	}

	domainResult, restErr := ctrl.service.CreateUser(userDomain)
	if restErr != nil {
		logger.Error("Error trying to call CreateUser service", restErr, zap.String("journey", "CreateUser"))
		response.JSON(w, restErr.Code, restErr)
		return
	}
	userResponse := converter.ConvertUserDomainToResponse(domainResult)

	logger.Debug("Finish CreateUser controller", zap.String("journey", "CreateUser"))
	response.JSON(w, http.StatusCreated, userResponse)
}

func (ctrl *userController) Login(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Init Login controller", zap.String("journey", "Login"))

	var LoginRequest requests.Login

	err := json.NewDecoder(r.Body).Decode(&LoginRequest)
	if err != nil {
		restErr := rest_errors.NewBadRequestError(err.Error())
		logger.Error("Error trying to validate user info", err, zap.String("journey", "Login"))
		response.JSON(w, restErr.Code, restErr)
		return
	}

	validationErr := validation.ValidateRequest(&LoginRequest)
	if validationErr != nil {
		logger.Error("Validation failed", validationErr, zap.String("journey", "Login"))
		response.JSON(w, http.StatusBadRequest, validationErr)
		return
	}
	userDomain := converter.ConvertLoginRequestToUserDomain(&LoginRequest)

	domainResult, restErr := ctrl.service.Login(userDomain)
	if restErr != nil {
		logger.Error("Error trying to call Login service", restErr, zap.String("journey", "Login"))
		response.JSON(w, restErr.Code, restErr)
		return
	}

	token, restErr := security.NewJwtToken().GenerateToken(domainResult.ID)
	if restErr != nil {
		restErr := rest_errors.NewInternalServerError("Error trying to generate token")
		logger.Error("Error trying to generate token", restErr, zap.String("journey", "Login"))
		response.JSON(w, restErr.Code, restErr)
		return
	}
	tokenResp := responses.Login{Token: token}

	logger.Debug("Finish Login controller", zap.String("journey", "login"))
	response.JSON(w, http.StatusOK, tokenResp)
}

// TODO:Passar esse metodo pro service
func checkUsernameAvailability(ctrl *userController, username string) *rest_errors.RestErr {
	logger.Debug("Init check username availability controller", zap.String("journey", "checkUsernameAvailability"))

	usersWithUsername, restErr := ctrl.service.FindUserByUsername(username)

	if restErr != nil && restErr.Code != http.StatusNotFound {
		logger.Error("Error validating username uniqueness", restErr, zap.String("username", username), zap.String("journey", "UserCheck"))
		return rest_errors.NewInternalServerError("Error validating username uniqueness: " + restErr.Error())
	}

	if usersWithUsername != nil && len(*usersWithUsername) > 0 {
		logger.Error("Username is already in use", nil, zap.String("username", username), zap.String("journey", "UserCheck"))
		return rest_errors.NewConflictError("Username is already in use")
	}
	return nil
}

// TODO:Passar esse metodo pro service
func checkEmailAvailability(ctrl *userController, email string) *rest_errors.RestErr {
	logger.Debug("Init check email availability controller", zap.String("journey", "checkEmailAvailability"))

	usersWithEmail, restErr := ctrl.service.FindUserByEmail(email)

	if restErr != nil && restErr.Code != http.StatusNotFound {
		logger.Error("Error validating email uniqueness", restErr, zap.String("email", email), zap.String("journey", "UserCheck"))
		return rest_errors.NewInternalServerError("Error validating email uniqueness: " + restErr.Error())
	}

	if usersWithEmail != nil && len(*usersWithEmail) > 0 {
		logger.Error("Email is already in use", nil, zap.String("email", email), zap.String("journey", "UserCheck"))
		return rest_errors.NewConflictError("Email is already in use")
	}
	return nil
}

func (ctrl *userController) GetGroups(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Init GetGroups controller", zap.String("journey", "GetUserGroups"))

	param := mux.Vars(r)
	userID := param["userId"]

	if userID == "" {
		restErr := rest_errors.NewBadRequestError("userId is required")
		logger.Error("UserID is missing", nil, zap.String("journey", "GetUserGroups"))
		response.JSON(w, restErr.Code, restErr)
		return
	}

	groupsDto, errRest := ctrl.service.GetUserGroups(userID)
	if errRest != nil {
		logger.Error("Error trying GetGroupService", errRest, zap.String("journey", "GetUserGroups"))
		response.JSON(w, errRest.Code, errRest)
		return
	}

	var groups []responses.Group

	for _, groupDomain := range *groupsDto {
		groups = append(groups, converter.ConvertGroupDtoToResponse(&groupDomain))
	}

	logger.Debug("Finish GetUserGroups controller", zap.String("journey", "GetUserGroups"))
	response.JSON(w, http.StatusOK, groups)
}
