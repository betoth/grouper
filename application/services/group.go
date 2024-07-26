package services

import (
	"grouper/application/domain"
	"grouper/application/port/input"
	"grouper/application/port/output"
	"grouper/config/logger"
	"grouper/config/rest_errors"

	"go.uber.org/zap"
)

func NewGroupServices(groupRepository output.GroupPort) input.GroupDomainService {
	return &groupDomainService{
		groupRepository,
	}
}

type groupDomainService struct {
	repository output.GroupPort
}

func (gd *groupDomainService) CreateGroupService(groupDomain domain.GroupDomain) (*domain.GroupDomain, *rest_errors.RestErr) {

	groupDomainRepository, err := gd.repository.CreateGroup(groupDomain)

	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info(
		"CreateUser service executed successfully",
		zap.String("userId", groupDomainRepository.ID),
		zap.String("journey", "CreateGroup"))

	return groupDomainRepository, nil
}
