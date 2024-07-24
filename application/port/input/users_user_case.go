package input

import (
	"grouper/application/domain"
	"grouper/config/rest_errors"
)

type UserDomainService interface {
	CreateUserServices(domain.UserDomain) (*domain.UserDomain, *rest_errors.RestErr)
}
