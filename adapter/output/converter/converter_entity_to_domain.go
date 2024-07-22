package converter

import (
	"grouper/adapter/output/model/entity"
	"grouper/application/domain"
)

func ConverterEntityToDomain(userEntity *entity.UserEntity) domain.UserDomain {

	return domain.UserDomain{
		ID:        userEntity.ID,
		Name:      userEntity.Name,
		Username:  userEntity.Username,
		Email:     userEntity.Email,
		Password:  userEntity.Password,
		CreatedAt: userEntity.CreatedAt,
	}
}
