package input

import (
	"grouper/application/domain"
	"grouper/config/rest_errors"
)

type LoginDomainService interface {
	LoginServices(domain.LoginDomain) (string, *rest_errors.RestErr)
}
