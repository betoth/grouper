package output

import "grouper/application/domain"

type UserPort interface {
	CreateUser(userDomain domain.UserDomain) (*domain.UserDomain, error)
}
