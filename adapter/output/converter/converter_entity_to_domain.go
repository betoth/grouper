package converter

import (
	"grouper/adapter/output/model/entity"
	"grouper/application/domain"
)

func ConverterUserEntityToDomain(userEntity *entity.User) domain.User {

	return domain.User{
		ID:        userEntity.ID,
		Name:      userEntity.Name,
		Username:  userEntity.Username,
		Email:     userEntity.Email,
		Password:  userEntity.Password,
		CreatedAt: userEntity.CreatedAt,
	}
}

func ConverterGroupEntityToDomain(groupEntity *entity.Group) domain.Group {

	return domain.Group{
		ID:         groupEntity.ID,
		Name:       groupEntity.Name,
		UserID:     groupEntity.UserID,
		TopicID:    groupEntity.TopicID,
		SubtopicID: groupEntity.SubtopicID,
		CreatedAt:  groupEntity.CreatedAt,
	}
}

func ConverterTopicEntityToDomain(topicEntity *entity.Topic) domain.Topic {

	return domain.Topic{
		ID:        topicEntity.ID,
		Name:      topicEntity.Name,
		CreatedAt: topicEntity.CreatedAt,
	}
}
