package converter

import (
	"grouper/adapter/input/model/request"
	"grouper/application/domain"
	"time"
)

func ConvertRequestToDomain(userRequest *request.UserRequest) domain.UserDomain {
	return domain.UserDomain{
		ID:        "",
		Name:      userRequest.Name,
		Email:     userRequest.Email,
		Username:  userRequest.Username,
		Password:  userRequest.Password,
		CreatedAt: time.Now(),
	}

}
