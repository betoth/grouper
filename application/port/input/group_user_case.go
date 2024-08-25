package input

import (
	"grouper/application/domain"
	"grouper/application/dto"
)

type GroupService interface {
	CreateGroup(groupDomain domain.Group) (*dto.Group, error)
	JoinGroup(userID, groupID string) error
	LeaveGroup(userID, groupID string) error
	GetGroups(parameter dto.GetGroupsParameter) (*[]dto.Group, error)
	FindByID(userID string) (*dto.Group, error)
}
