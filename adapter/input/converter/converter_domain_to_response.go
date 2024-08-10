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

func ConvertGroupDomainToResponse(groupDomain *domain.GroupDomain) resp.GroupResponse {
	return resp.GroupResponse{
		ID:         groupDomain.ID,
		Name:       groupDomain.Name,
		UserID:     groupDomain.UserID,
		TopicID:    groupDomain.TopicID,
		SubtopicID: groupDomain.SubtopicID,
		CreatedAt:  groupDomain.CreatedAt,
	}

}

func ConvertGroupDtoToResponse(groupDto *dto.GroupDTO) resp.GroupResponse2 {
	return resp.GroupResponse2{
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
