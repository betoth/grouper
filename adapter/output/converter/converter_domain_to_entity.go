package converter

import (
	"grouper/adapter/output/model/entity"
	"grouper/application/domain"
)

func ConvertUserDomainToEntity(userDomain *domain.User) entity.User {

	return entity.User{
		ID:        "",
		Name:      userDomain.Name,
		Email:     userDomain.Email,
		Username:  userDomain.Username,
		Password:  userDomain.Password,
		CreatedAt: userDomain.CreatedAt,
	}
}

func ConvertGroupDomainToEntity(groupDomain *domain.Group) entity.Group {

	return entity.Group{
		ID:         "",
		Name:       groupDomain.Name,
		UserID:     groupDomain.UserID,
		TopicID:    groupDomain.TopicID,
		SubtopicID: groupDomain.SubtopicID,
		CreatedAt:  groupDomain.CreatedAt,
	}
}
