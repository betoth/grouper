package input

import (
	"grouper/application/domain"
	"grouper/application/dto"
	"grouper/config/rest_errors"
)

type UserService interface {
	CreateUser(domain.User) (*domain.User, *rest_errors.RestErr)
	FindUserByUsername(username string) (*[]domain.User, *rest_errors.RestErr)
	FindUserByEmail(email string) (*[]domain.User, *rest_errors.RestErr)
	Login(domain.User) (*domain.User, *rest_errors.RestErr)
	GetUserGroups(userID string) (*[]dto.Group, *rest_errors.RestErr)
}
