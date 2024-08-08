package input

import (
	"grouper/adapter/input/model/dto"
	"grouper/application/domain"
	"grouper/config/rest_errors"
)

type GroupDomainService interface {
	CreateGroupService(groupDomain domain.GroupDomain) (*domain.GroupDomain, *rest_errors.RestErr)
	JoinService(userID, groupID string) *rest_errors.RestErr
	LeaveService(userID, groupID string) *rest_errors.RestErr
	GetGroupsService(parameter dto.GetGroupsQueryParameter) (*[]domain.GroupDomain, *rest_errors.RestErr)
}
