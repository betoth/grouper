package output

import (
	"grouper/application/domain"
	"grouper/application/dto"
	"grouper/config/rest_errors"
)

type GroupPort interface {
	CreateGroup(groupDomain domain.GroupDomain) (*domain.GroupDomain, *rest_errors.RestErr)
	Join(userID, groupID string) *rest_errors.RestErr
	Leave(userID, groupID string) *rest_errors.RestErr
	GetGroups(parameters dto.GetGroupsParameter) (*[]domain.GroupDomain, *rest_errors.RestErr)
}
