package converter

import (
	resp "grouper/adapter/input/model/response"
	"grouper/application/domain"
	"grouper/application/dto"
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

func ConvertGroupDtoToResponse(groupDto *dto.GroupDTO) resp.GroupResponse {
	return resp.GroupResponse{
		ID:       groupDto.ID,
		Name:     groupDto.Name,
		UserName: groupDto.UserName,
		Topic: resp.Topic{
			ID:   groupDto.Topic.ID,
			Name: groupDto.Topic.Name,
			Subtopic: resp.Subtopic{
				ID:   groupDto.Topic.Subtopic.ID,
				Name: groupDto.Topic.Subtopic.Name,
			},
		},
		CreatedAt: groupDto.CreatedAt,
	}

}
