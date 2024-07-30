package converter

import (
	"grouper/adapter/output/model/entity"
	"grouper/application/domain"
)

func ConverterUserEntityToDomain(userEntity *entity.UserEntity) domain.UserDomain {

	return domain.UserDomain{
		ID:        userEntity.ID,
		Name:      userEntity.Name,
		Username:  userEntity.Username,
		Email:     userEntity.Email,
		Password:  userEntity.Password,
		CreatedAt: userEntity.CreatedAt,
	}
}

func ConverterGroupEntityToDomain(groupEntity *entity.GroupEntity) domain.GroupDomain {

	return domain.GroupDomain{
		ID:        groupEntity.ID,
		Name:      groupEntity.Name,
		UserID:    groupEntity.UserID,
		CreatedAt: groupEntity.CreatedAt,
	}
}

func ConverterLoginEntityToDomain(LoginEntity *entity.LoginEntity) domain.LoginDomain {

	return domain.LoginDomain{
		Email:    LoginEntity.Email,
		Password: LoginEntity.Password,
	}
}
