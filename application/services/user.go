package services

import (
	"grouper/application/domain"
	"grouper/application/port/input"
	"grouper/application/port/output"
	util "grouper/application/util/secutiry"
	"grouper/config/logger"
	"grouper/config/rest_errors"

	"go.uber.org/zap"
)

func NewUserServices(userRepository output.UserPort) input.UserDomainService {
	return &userDomainService{
		userRepository,
	}
}

type userDomainService struct {
	repository output.UserPort
}

func (ud *userDomainService) CreateUserServices(userDomain domain.UserDomain) (*domain.UserDomain, *rest_errors.RestErr) {
	logger.Debug("Init CreateUser service", zap.String("journey", "CreateUser"))
	hashPassword, err := util.HashSHA256(userDomain.Password)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "createUser"))
		return nil, rest_errors.NewInternalServerError("")

	}

	userDomain.Password = hashPassword

	userDomainRepository, restErr := ud.repository.CreateUser(userDomain)
	if restErr != nil {
		logger.Error("Error trying to call repository", restErr, zap.String("journey", "createUser"))
		return nil, restErr
	}

	logger.Debug("Finish CreateUser service", zap.String("journey", "CreateUser"))
	return userDomainRepository, nil
}

func (ud *userDomainService) FindUserByUsernameServices(username string) (*[]domain.UserDomain, *rest_errors.RestErr) {
	logger.Debug("Init FindUserByName service", zap.String("journey", "FindUserByName"))
	usersDomainRepository, err := ud.repository.FindUserByUsername(username)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "FindUserByUsername"))
		return &[]domain.UserDomain{}, err
	}

	logger.Debug("Finish FindUserByName service", zap.String("journey", "FindUserByName"))
	return usersDomainRepository, err
}

func (ud *userDomainService) FindUserByEmailServices(email string) (*[]domain.UserDomain, *rest_errors.RestErr) {
	logger.Debug("Init FindUserByEmail service", zap.String("journey", "FindUserByEmail"))
	usersDomainRepository, err := ud.repository.FindUserByEmail(email)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "FindUserByEmail"))
		return &[]domain.UserDomain{}, err
	}

	logger.Debug("Finish FindUserByEmail service", zap.String("journey", "FindUserByEmail"))

	return usersDomainRepository, err
}

func (ud *userDomainService) LoginServices(userDomain domain.UserDomain) (*domain.UserDomain, *rest_errors.RestErr) {
	logger.Debug("Init LoginServices service", zap.String("journey", "Login"))
	userRepository, err := ud.repository.Login(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "Login"))
		return nil, rest_errors.NewInternalServerError("")
	}

	ok := util.VerifyPassword(userDomain.Password, userRepository.Password)
	if ok != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "Login"))
		return nil, rest_errors.NewUnauthorizedError("Invalid username or password")
	}

	logger.Debug("Finish LoginServices service", zap.String("journey", "Login"))
	return userRepository, nil
}

func (ud *userDomainService) GetUserGroupsService(userID string) (*[]domain.GroupDomain, *rest_errors.RestErr) {
	logger.Debug("Init GetUserGroups service", zap.String("journey", "GetUserGroups"))

	groups, err := ud.repository.GetUserGroups(userID)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "GetUserGroups"))
		return nil, err
	}

	logger.Debug("Finish GetUserGroups service", zap.String("journey", "GetUserGroups"))
	return groups, nil
}
