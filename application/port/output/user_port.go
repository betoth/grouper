package output

import (
	"grouper/application/domain"
	"grouper/config/rest_errors"
)

type UserPort interface {
	CreateUser(userDomain domain.User) (*domain.User, *rest_errors.RestErr)
	FindUserByUsername(username string) (*[]domain.User, *rest_errors.RestErr)
	FindUserByEmail(email string) (*[]domain.User, *rest_errors.RestErr)
	Login(userDomain domain.User) (*domain.User, *rest_errors.RestErr)
	GetUserGroups(userID string) (*[]domain.Group, *rest_errors.RestErr)
}
