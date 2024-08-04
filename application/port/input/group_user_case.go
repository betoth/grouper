package input

import (
	"grouper/application/domain"
	"grouper/config/rest_errors"
)

type GroupDomainService interface {
	CreateGroupService(groupDomain domain.GroupDomain) (*domain.GroupDomain, *rest_errors.RestErr)
	JoinService(userID, groupID string) *rest_errors.RestErr
	LeaveService(userID, groupID string) *rest_errors.RestErr
}
