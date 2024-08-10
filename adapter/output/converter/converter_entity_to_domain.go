package converter

import (
	"grouper/adapter/output/model/entity"
	"grouper/application/domain"
)

func ConverterUserEntityToDomain(userEntity *entity.User) domain.UserDomain {

	return domain.UserDomain{
		ID:        userEntity.ID,
		Name:      userEntity.Name,
		Username:  userEntity.Username,
		Email:     userEntity.Email,
		Password:  userEntity.Password,
		CreatedAt: userEntity.CreatedAt,
	}
}

func ConverterGroupEntityToDomain(groupEntity *entity.Group) domain.GroupDomain {

	return domain.GroupDomain{
		ID:         groupEntity.ID,
		Name:       groupEntity.Name,
		UserID:     groupEntity.UserID,
		TopicID:    groupEntity.TopicID,
		SubtopicID: groupEntity.SubtopicID,
		CreatedAt:  groupEntity.CreatedAt,
	}
}
