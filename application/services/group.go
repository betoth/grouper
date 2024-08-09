package services

import (
	"grouper/adapter/input/model/dto"
	dtoOut "grouper/adapter/output/model/dto"
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
	logger.Debug("Init CreateGroup service", zap.String("journey", "CreateGroup"))
	groupDomainRepository, err := gd.repository.CreateGroup(groupDomain)

	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Debug("Finish CreateGroup service", zap.String("journey", "CreateGroup"))
	return groupDomainRepository, nil
}

func (gd *groupDomainService) JoinService(userID, groupID string) *rest_errors.RestErr {
	logger.Debug("Init Join service", zap.String("journey", "JoinGroup"))
	err := gd.repository.Join(userID, groupID)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "JoinGroup"))
		return err
	}
	logger.Debug("Finish Join service", zap.String("journey", "JoinGroup"))
	return nil
}

func (gd *groupDomainService) LeaveService(userID, groupID string) *rest_errors.RestErr {
	logger.Debug("Init Leave service", zap.String("journey", "LeaveGroup"))
	err := gd.repository.Leave(userID, groupID)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "LeaveGroup"))
		return err
	}
	logger.Debug("Init Leave service", zap.String("journey", "LeaveGroup"))
	return nil
}

func (gd *groupDomainService) GetGroupsService(parameter dto.GetGroupsQueryParameter) (*[]domain.GroupDomain, *rest_errors.RestErr) {
	logger.Debug("Init GetGroups service", zap.String("journey", "GetGroups"))
	queryParameter := dtoOut.GetGroupsQuery{
		Name: parameter.Name,
	}

	groups, err := gd.repository.GetGroups(queryParameter)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "GetGroups"))
		return nil, err
	}

	if groups == nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "GetGroups"))
		return nil, err
	}
	logger.Debug("Finish GetGroups service", zap.String("journey", "GetGroups"))
	return groups, nil
}
