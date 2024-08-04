package output

import (
	"grouper/application/domain"
	"grouper/config/rest_errors"
)

type UserPort interface {
	CreateUser(userDomain domain.UserDomain) (*domain.UserDomain, *rest_errors.RestErr)
	FindUserByUsername(username string) (*[]domain.UserDomain, *rest_errors.RestErr)
	FindUserByEmail(email string) (*[]domain.UserDomain, *rest_errors.RestErr)
	Login(userDomain domain.UserDomain) (*domain.UserDomain, *rest_errors.RestErr)
}
