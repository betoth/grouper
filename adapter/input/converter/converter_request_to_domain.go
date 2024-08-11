package converter

import (
	"grouper/adapter/input/model/requests"
	"grouper/application/domain"
	"time"
)

func ConvertUserRequestToDomain(userRequest *requests.User) domain.User {
	return domain.User{
		ID:        "",
		Name:      userRequest.Name,
		Email:     userRequest.Email,
		Username:  userRequest.Username,
		Password:  userRequest.Password,
		CreatedAt: time.Now(),
	}
}

func ConvertGroupRequestToDomain(groupRequest *requests.Group) domain.Group {
	return domain.Group{
		ID:         "",
		Name:       groupRequest.Name,
		TopicID:    groupRequest.TopicID,
		SubtopicID: groupRequest.SubtopicID,
		CreatedAt:  time.Now(),
	}
}

func ConvertLoginRequestToUserDomain(loginRequest *requests.Login) domain.User {
	return domain.User{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}
}
