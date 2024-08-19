package services

import (
	"grouper/application/domain"
	"grouper/application/dto"
	"grouper/application/port/input"
	"grouper/application/port/output"
	"grouper/application/util/security"
	"grouper/config/logger"
	"grouper/config/rest_errors"
	"net/http"

	"go.uber.org/zap"
)

func NewUserService(userRepository output.UserPort) input.UserService {
	return &userService{
		userRepository,
	}
}

type userService struct {
	repository output.UserPort
}

func (service *userService) CreateUser(userDomain domain.User) (*domain.User, *rest_errors.RestErr) {
	logger.Debug("Init CreateUser service", zap.String("journey", "CreateUser"))
	hashPassword, err := security.HashSHA256(userDomain.Password)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "createUser"))
		return nil, rest_errors.NewInternalServerError("")

	}

	userDomain.Password = hashPassword

	userDomainRepository, restErr := service.repository.CreateUser(userDomain)
	if restErr != nil {
		logger.Error("Error trying to call repository", restErr, zap.String("journey", "createUser"))
		return nil, restErr
	}

	logger.Debug("Finish CreateUser service", zap.String("journey", "CreateUser"))
	return userDomainRepository, nil
}

func (service *userService) FindUserByUsername(username string) (*[]domain.User, *rest_errors.RestErr) {
	logger.Debug("Init FindUserByName service", zap.String("journey", "FindUserByName"))
	usersDomainRepository, err := service.repository.FindUserByUsername(username)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "FindUserByUsername"))
		return &[]domain.User{}, err
	}

	logger.Debug("Finish FindUserByName service", zap.String("journey", "FindUserByName"))
	return usersDomainRepository, err
}

func (service *userService) FindUserByEmail(email string) (*[]domain.User, *rest_errors.RestErr) {
	logger.Debug("Init FindUserByEmail service", zap.String("journey", "FindUserByEmail"))
	usersDomainRepository, err := service.repository.FindUserByEmail(email)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "FindUserByEmail"))
		return &[]domain.User{}, err
	}

	logger.Debug("Finish FindUserByEmail service", zap.String("journey", "FindUserByEmail"))

	return usersDomainRepository, err
}

func (service *userService) Login(userDomain domain.User) (*domain.User, *rest_errors.RestErr) {
	logger.Debug("Init LoginServices service", zap.String("journey", "Login"))
	userRepository, err := service.repository.Login(userDomain)

	if err != nil {
		if err.Code == http.StatusUnauthorized || err.Code == http.StatusNotFound {
			logger.Error("User or password is invalid", err, zap.String("journey", "Login"))
			return nil, rest_errors.NewUnauthorizedError("User or password is invalid")
		}

		logger.Error("Error trying to call repository", err, zap.String("journey", "Login"))
		return nil, rest_errors.NewInternalServerError("")
	}

	ok := security.VerifyPassword(userDomain.Password, userRepository.Password)
	if ok != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "Login"))
		return nil, rest_errors.NewUnauthorizedError("Invalid username or password")
	}

	logger.Debug("Finish LoginServices service", zap.String("journey", "Login"))
	return userRepository, nil
}

func (service *userService) GetUserGroups(userID string) (*[]dto.Group, *rest_errors.RestErr) {
	logger.Debug("Init GetUserGroups service", zap.String("journey", "GetUserGroups"))

	groupsRepo, err := service.repository.GetUserGroups(userID)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "GetUserGroups"))
		return nil, err
	}

	var groupsDto []dto.Group
	for _, groupRepo := range *groupsRepo {

		groupDto := dto.Group{
			ID:       groupRepo.ID,
			Name:     groupRepo.Name,
			UserName: "UserNameDTO",
			Topic: dto.GroupTopic{
				ID:   "TopicDto ID",
				Name: "TopicDto Name",
				Subtopic: dto.GroupSubtopic{
					ID:   "SubtopicDto ID",
					Name: "SubtopicDTO Name",
				},
			},
			CreatedAt: groupRepo.CreatedAt,
		}

		groupsDto = append(groupsDto, groupDto)

	}

	logger.Debug("Finish GetUserGroups service", zap.String("journey", "GetUserGroups"))
	return &groupsDto, nil
}

func (service *userService) FindByID(userID string) (*domain.User, *error) {

	return nil, nil
}
