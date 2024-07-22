package services

import (
	"grouper/application/domain"
	"grouper/application/port/input"
	"grouper/application/port/output"
)

func NewUserServices(userRepository output.UserPort) input.UserDomainService {
	return &userDomainService{
		userRepository,
	}
}

type userDomainService struct {
	repository output.UserPort
}

func (ud *userDomainService) CreateUserServices(userDomain domain.UserDomain) (*domain.UserDomain, error) {

	userDominRepository, err := ud.repository.CreateUser(userDomain)
	if err != nil {
		return nil, err
	}

	return userDominRepository, nil
}
