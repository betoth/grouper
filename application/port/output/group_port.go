package output

import (
	"grouper/application/domain"
	"grouper/application/dto"
	"grouper/config/rest_errors"
)

type GroupPort interface {
	CreateGroup(groupDomain domain.Group) (*domain.Group, *rest_errors.RestErr)
	JoinGroup(userID, groupID string) *rest_errors.RestErr
	LeaveGroup(userID, groupID string) *rest_errors.RestErr
	GetGroups(parameters dto.GetGroupsParameter) (*[]domain.Group, *rest_errors.RestErr)
}
