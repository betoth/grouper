package converter

import (
	"grouper/adapter/output/model/entity"
	"grouper/application/domain"
)

func ConvertDomainToEntity(userDomain *domain.UserDomain) entity.UserEntity {

	return entity.UserEntity{
		ID:        "",
		Name:      userDomain.Name,
		Email:     userDomain.Email,
		Username:  userDomain.Username,
		Password:  userDomain.Password,
		CreatedAt: userDomain.CreatedAt,
	}

}
