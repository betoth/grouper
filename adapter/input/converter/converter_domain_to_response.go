package converter

import (
	"grouper/adapter/input/model/responses"
	"grouper/application/domain"
	"grouper/application/dto"
)

func ConvertUserDomainToResponse(userDomain *domain.User) responses.User {
	return responses.User{
		ID:        userDomain.ID,
		Name:      userDomain.Name,
		Email:     userDomain.Email,
		Username:  userDomain.Username,
		CreatedAt: userDomain.CreatedAt,
	}

}

func ConvertGroupDtoToResponse(groupDto *dto.Group) responses.Group {
	return responses.Group{
		ID:       groupDto.ID,
		Name:     groupDto.Name,
		UserName: groupDto.UserName,
		Topic: responses.Topic{
			ID:   groupDto.Topic.ID,
			Name: groupDto.Topic.Name,
			Subtopic: responses.Subtopic{
				ID:   groupDto.Topic.Subtopic.ID,
				Name: groupDto.Topic.Subtopic.Name,
			},
		},
		CreatedAt: groupDto.CreatedAt,
	}

}
