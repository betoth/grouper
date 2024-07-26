package converter

import (
	"grouper/adapter/input/model/response"
	"grouper/application/domain"
)

func ConvertUserDomainToResponse(userDomain *domain.UserDomain) response.UserResponse {
	return response.UserResponse{
		ID:        userDomain.ID,
		Name:      userDomain.Name,
		Email:     userDomain.Email,
		Username:  userDomain.Username,
		CreatedAt: userDomain.CreatedAt,
	}

}

func ConvertGroupDomainToResponse(groupDomain *domain.GroupDomain) response.GroupResponse {
	return response.GroupResponse{
		ID:        groupDomain.ID,
		Name:      groupDomain.Name,
		UserID:    groupDomain.UserID,
		CreatedAt: groupDomain.CreatedAt,
	}

}
