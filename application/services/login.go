package services

import (
	"fmt"
	"grouper/application/domain"
	"grouper/application/port/input"
	"grouper/application/port/output"
	util "grouper/application/util/secutiry"
	"grouper/config/logger"
	"grouper/config/rest_errors"

	"go.uber.org/zap"
)

func NewLoginServices(loginRepository output.LoginPort) input.LoginDomainService {
	return &loginDomainService{
		loginRepository,
	}
}

type loginDomainService struct {
	repository output.LoginPort
}

func (lg *loginDomainService) LoginServices(loginDomain domain.LoginDomain) (string, *rest_errors.RestErr) {

	loginRepository, err := lg.repository.Login(loginDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "Login"))
		return "", rest_errors.NewInternalServerError("")
	}
	fmt.Println(loginRepository.Password)
	fmt.Println(loginDomain.Password)
	ok := util.VerifyPassword(loginDomain.Password, loginRepository.Password)
	fmt.Println(ok)
	if ok != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "Login"))
		return "", rest_errors.NewUnauthorizedError("Invalid username or password")
	}

	return "OK", nil

}
