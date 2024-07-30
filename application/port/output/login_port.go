package output

import (
	"grouper/application/domain"
	"grouper/config/rest_errors"
)

type LoginPort interface {
	Login(loginDomain domain.LoginDomain) (*domain.LoginDomain, *rest_errors.RestErr)
}
