package output

import (
	"grouper/application/domain"
	"grouper/config/rest_errors"
)

type UserPort interface {
	CreateUser(userDomain domain.UserDomain) (*domain.UserDomain, *rest_errors.RestErr)
}
