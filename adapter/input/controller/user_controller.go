package controller

import (
	"encoding/json"
	"grouper/adapter/input/converter"
	"grouper/adapter/input/model/request"
	resp "grouper/adapter/input/model/response"
	"grouper/adapter/input/response"
	"grouper/application/port/input"
	"grouper/application/util/secutiry"
	"grouper/config/logger"
	"grouper/config/rest_errors"
	"grouper/config/validation"
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
	Login(w http.ResponseWriter, r *http.Request)
}

type userControllerInterface struct {
	service input.UserDomainService
}

func (uc *userControllerInterface) CreateUser(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Init createUser controller", zap.String("journey", "CreateUSer"))

	var userRequest request.UserRequest

	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		restErr := rest_errors.NewBadRequestError(err.Error())
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		response.JSON(w, restErr.Code, restErr)
		return
	}

	validationErr := validation.ValidateUserRequest(&userRequest)
	if validationErr != nil {
		logger.Error("Validation failed", validationErr, zap.String("journey", "createUser"))
		response.JSON(w, http.StatusBadRequest, validationErr)
		return
	}

	userDomain := converter.ConvertUserRequestToDomain(&userRequest)

	restErr := checkUsernameAvailability(uc, userDomain.Username)
	if restErr != nil {
		logger.Error("Username validation failed", restErr, zap.String("journey", "createUser"))
		response.JSON(w, restErr.Code, err)
		return
	}

	restErr = checkEmailAvailability(uc, userDomain.Email)
	if restErr != nil {
		logger.Error("Email validation failed", err, zap.String("journey", "createUser"))
		response.JSON(w, restErr.Code, restErr)
		return
	}

	domainResult, restErr := uc.service.CreateUserServices(userDomain)
	if restErr != nil {
		logger.Error("Error trying to call CreateUser service", restErr, zap.String("journey", "createUser"))
		response.JSON(w, restErr.Code, restErr)
		return
	}

	logger.Debug("Finish createUser controller", zap.String("journey", "CreateUSer"))
	userResponse := converter.ConvertUserDomainToResponse(domainResult)
	response.JSON(w, http.StatusCreated, userResponse)
}

func (uc *userControllerInterface) Login(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Init Login controller", zap.String("journey", "Login"))
	var LoginRequest request.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&LoginRequest)
	if err != nil {
		restErr := rest_errors.NewBadRequestError(err.Error())
		logger.Error("Error trying to validate user info", err, zap.String("journey", "Login"))
		response.JSON(w, restErr.Code, restErr)
		return
	}

	userDomain := converter.ConvertLoginRequestToUserDomain(&LoginRequest)

	domainResult, restErr := uc.service.LoginServices(userDomain)
	if restErr != nil {
		logger.Error("Error trying to call Login service", restErr, zap.String("journey", "Login"))
		response.JSON(w, restErr.Code, restErr)
		return
	}

	token, restErr := secutiry.NewJwtToken().GenerateToken(domainResult.ID)
	if restErr != nil {
		restErr := rest_errors.NewInternalServerError("Error trying to generate token")
		logger.Error("Error trying to generate token", restErr, zap.String("journey", "Login"))
		response.JSON(w, restErr.Code, restErr)
		return
	}

	logger.Debug("Finish Login controller", zap.String("journey", "login"))
	var tokenResp resp.LoginResponse
	tokenResp.Token = token
	response.JSON(w, http.StatusOK, tokenResp)
}

func checkUsernameAvailability(uc *userControllerInterface, username string) *rest_errors.RestErr {
	logger.Debug("Init check username availability controller", zap.String("journey", "checkUsernameAvailability"))
	usersWithUsername, restErr := uc.service.FindUserByUsernameServices(username)
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

func checkEmailAvailability(uc *userControllerInterface, email string) *rest_errors.RestErr {
	logger.Debug("Init check email availability controller", zap.String("journey", "checkEmailAvailability"))
	usersWithEmail, restErr := uc.service.FindUserByEmailServices(email)
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
