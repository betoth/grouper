package input

import (
	"grouper/application/domain"
	"grouper/application/dto"
	"grouper/config/rest_errors"
)

type GroupService interface {
	CreateGroup(groupDomain domain.Group) (*dto.Group, *rest_errors.RestErr)
	JoinGroup(userID, groupID string) *rest_errors.RestErr
	LeaveGroup(userID, groupID string) *rest_errors.RestErr
	GetGroups(parameter dto.GetGroupsParameter) (*[]dto.Group, *rest_errors.RestErr)
}
