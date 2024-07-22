package converter

import (
	"grouper/adapter/input/model/response"
	"grouper/application/domain"
)

func ConvertDomainToResponse(userDomain *domain.UserDomain) response.UserResponse {
	return response.UserResponse{
		ID:        userDomain.ID,
		Name:      userDomain.Name,
		Email:     userDomain.Email,
		Username:  userDomain.Username,
		CreatedAt: userDomain.CreatedAt,
	}

}
