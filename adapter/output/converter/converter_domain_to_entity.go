package converter

import (
	"grouper/adapter/output/model/entity"
	"grouper/application/domain"
)

func ConvertUserDomainToEntity(userDomain *domain.UserDomain) entity.UserEntity {

	return entity.UserEntity{
		ID:        "",
		Name:      userDomain.Name,
		Email:     userDomain.Email,
		Username:  userDomain.Username,
		Password:  userDomain.Password,
		CreatedAt: userDomain.CreatedAt,
	}

}

func ConvertGroupDomainToEntity(groupDomain *domain.GroupDomain) entity.GroupEntity {

	return entity.GroupEntity{
		ID:        "",
		Name:      groupDomain.Name,
		UserID:    groupDomain.UserID,
		CreatedAt: groupDomain.CreatedAt,
	}

}

func ConvertLoginDomainToEntity(LoginDomain *domain.LoginDomain) entity.LoginEntity {

	return entity.LoginEntity{
		Email:    LoginDomain.Email,
		Password: LoginDomain.Password,
	}

}
