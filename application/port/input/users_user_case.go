package input

import (
	"grouper/application/domain"
	"grouper/application/dto"
	"grouper/config/rest_errors"
)

type UserDomainService interface {
	CreateUserServices(domain.UserDomain) (*domain.UserDomain, *rest_errors.RestErr)
	FindUserByUsernameServices(username string) (*[]domain.UserDomain, *rest_errors.RestErr)
	FindUserByEmailServices(email string) (*[]domain.UserDomain, *rest_errors.RestErr)
	LoginServices(domain.UserDomain) (*domain.UserDomain, *rest_errors.RestErr)
	GetUserGroupsService(userID string) (*[]dto.GroupDTO, *rest_errors.RestErr)
}
