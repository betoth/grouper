package output

import (
	"grouper/application/domain"
	"grouper/application/dto"
)

type GroupPort interface {
	CreateGroup(groupDomain domain.Group) (*domain.Group, error)
	JoinGroup(userID, groupID string) error
	LeaveGroup(userID, groupID string) error
	GetGroups(parameters dto.GetGroupsParameter) (*[]domain.Group, error)
	FindByID(groupID string) (*domain.Group, error)
}
