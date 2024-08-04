package converter

import (
	"grouper/adapter/input/model/request"
	"grouper/application/domain"
	"time"
)

func ConvertUserRequestToDomain(userRequest *request.UserRequest) domain.UserDomain {
	return domain.UserDomain{
		ID:        "",
		Name:      userRequest.Name,
		Email:     userRequest.Email,
		Username:  userRequest.Username,
		Password:  userRequest.Password,
		CreatedAt: time.Now(),
	}
}

func ConvertGroupRequestToDomain(groupRequest *request.GroupRequest) domain.GroupDomain {
	return domain.GroupDomain{
		ID:        "",
		Name:      groupRequest.Name,
		UserID:    groupRequest.UserID,
		CreatedAt: time.Now(),
	}
}

func ConvertLoginRequestToUserDomain(loginRequest *request.LoginRequest) domain.UserDomain {
	return domain.UserDomain{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}
}
