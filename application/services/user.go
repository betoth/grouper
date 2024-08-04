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

	logger.Info("CreateUser service executed successfully", zap.String("userId", userDomainRepository.ID), zap.String("journey", "createUser"))

	return userDomainRepository, nil
}

func (ud *userDomainService) FindUserByUsernameServices(username string) (*[]domain.UserDomain, *rest_errors.RestErr) {
	usersDomainRepository, err := ud.repository.FindUserByUsername(username)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "FindUserByUsername"))
		return &[]domain.UserDomain{}, err
	}

	logger.Info("Find user by username executed successfully", zap.String("journey", "FindUserByUsername"))

	return usersDomainRepository, err
}

func (ud *userDomainService) FindUserByEmailServices(email string) (*[]domain.UserDomain, *rest_errors.RestErr) {
	usersDomainRepository, err := ud.repository.FindUserByEmail(email)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "FindUserByEmail"))
		return &[]domain.UserDomain{}, err
	}

	logger.Info("Find user by email executed successfully", zap.String("journey", "FindUserByEmail"))

	return usersDomainRepository, err
}

func (ud *userDomainService) LoginServices(userDomain domain.UserDomain) (*domain.UserDomain, *rest_errors.RestErr) {
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

	return userRepository, nil
}
