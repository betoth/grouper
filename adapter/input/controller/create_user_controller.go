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
	logger.Info("Init createUser controller", zap.String("journey", "CreateUSer"))

	var userRequest request.UserRequest

	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		response.JSON(w, http.StatusBadRequest, rest_errors.NewBadRequestError("Error trying to validate user info"))
		return
	}

	validationErr := validation.ValidateUserRequest(&userRequest)
	if validationErr != nil {
		logger.Error("Validation failed", validationErr, zap.String("journey", "createUser"))
		response.JSON(w, http.StatusBadRequest, validationErr)
		return
	}

	userDomain := converter.ConvertRequestToDomain(&userRequest)

	if err := checkUsernameAvailability(uc, userDomain.Username); err != nil {
		logger.Error("Username validation failed", err, zap.String("journey", "createUser"))
		response.JSON(w, err.Code, err)
		return
	}

	if err := checkEmailAvailability(uc, userDomain.Email); err != nil {
		logger.Error("Email validation failed", err, zap.String("journey", "createUser"))
		response.JSON(w, err.Code, err)
		return
	}

	domainResult, restErr := uc.service.CreateUserServices(userDomain)
	if restErr != nil {
		logger.Error("Error trying to call CreateUser service", restErr, zap.String("journey", "createUser"))
		response.JSON(w, restErr.Code, restErr)
		return
	}

	userResponse := converter.ConvertDomainToResponse(domainResult)
	response.JSON(w, http.StatusCreated, userResponse)
}

func checkUsernameAvailability(uc *userControllerInterface, username string) *rest_errors.RestErr {
	usersWithUsername, restErr := uc.service.FindUserByUsernameServices(username)
	if restErr != nil && restErr.Code != http.StatusNotFound {
		return rest_errors.NewInternalServerError("Error validating username uniqueness")
	}

	if usersWithUsername != nil && len(*usersWithUsername) > 0 {
		return rest_errors.NewConflictError("Username is already in use")
	}
	return nil
}

func checkEmailAvailability(uc *userControllerInterface, email string) *rest_errors.RestErr {
	usersWithEmail, restErr := uc.service.FindUserByEmailServices(email)
	if restErr != nil && restErr.Code != http.StatusNotFound {
		return rest_errors.NewInternalServerError("Error validating email uniqueness")
	}

	if usersWithEmail != nil && len(*usersWithEmail) > 0 {
		return rest_errors.NewConflictError("Email is already in use")
	}
	return nil
}
