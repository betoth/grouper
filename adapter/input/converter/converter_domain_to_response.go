package converter

import (
	resp "grouper/adapter/input/model/response"
	"grouper/application/domain"
)

func ConvertUserDomainToResponse(userDomain *domain.UserDomain) resp.UserResponse {
	return resp.UserResponse{
		ID:        userDomain.ID,
		Name:      userDomain.Name,
		Email:     userDomain.Email,
		Username:  userDomain.Username,
		CreatedAt: userDomain.CreatedAt,
	}

}

func ConvertGroupDomainToResponse(groupDomain *domain.GroupDomain) resp.GroupResponse {
	return resp.GroupResponse{
		ID:        groupDomain.ID,
		Name:      groupDomain.Name,
		UserID:    groupDomain.UserID,
		CreatedAt: groupDomain.CreatedAt,
	}

}
