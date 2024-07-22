package input

import "grouper/application/domain"

type UserDomainService interface {
	CreateUserServices(domain.UserDomain) (*domain.UserDomain, error)
}
