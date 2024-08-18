package output

import (
	"grouper/application/domain"
	"grouper/application/dto"
	"grouper/config/rest_errors"
)

type GroupPort interface {
	CreateGroup(groupDomain domain.Group) (*domain.Group, error)
	JoinGroup(userID, groupID string) error
	LeaveGroup(userID, groupID string) *rest_errors.RestErr
	GetGroups(parameters dto.GetGroupsParameter) (*[]domain.Group, *rest_errors.RestErr)
	FindByID(groupID string) (*domain.Group, error)
}
