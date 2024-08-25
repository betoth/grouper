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
		Topic: responses.GroupTopic{
			ID:   groupDto.Topic.ID,
			Name: groupDto.Topic.Name,
			SubTopic: responses.GroupSubtopic{
				ID:   groupDto.Topic.Subtopic.ID,
				Name: groupDto.Topic.Subtopic.Name,
			},
		},
		CreatedAt: groupDto.CreatedAt,
	}

}

func ConvertTopicDomainToResponse(topicDomain *domain.Topic) responses.Topic {
	return responses.Topic{
		ID:   topicDomain.ID,
		Name: topicDomain.Name,
	}

}

func ConvertSubtopicDomainToResponse(subtopicDomain *domain.Subtopic) responses.Subtopic {
	return responses.Subtopic{
		ID:   subtopicDomain.ID,
		Name: subtopicDomain.Name,
		Topic: responses.Topic{
			ID:   subtopicDomain.TopicID,
			Name: "",
		},
	}
}
