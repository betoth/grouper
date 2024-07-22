package repository

import (
	"context"
	"database/sql"
	"fmt"
	"grouper/adapter/output/converter"
	"grouper/application/domain"
	"grouper/application/port/output"
	"grouper/configuration/logger"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func NewUserRepository(database *sql.DB) output.UserPort {
	return &userRepository{
		database,
	}

}

type userRepository struct {
	db *sql.DB
}

func (ur *userRepository) CreateUser(userDomain domain.UserDomain) (*domain.UserDomain, error) {

	userEntity := converter.ConvertDomainToEntity(&userDomain)

	query := `
        INSERT INTO users ( name, email, username, password, createdat)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, createdat
    `

	row := ur.db.QueryRowContext(
		context.Background(),
		query,
		userEntity.Name,
		userEntity.Email,
		userEntity.Username,
		userEntity.Password,
		userEntity.CreatedAt,
	)

	if err := row.Scan(&userEntity.ID, &userEntity.CreatedAt); err != nil {
		return nil, fmt.Errorf("could not insert user: %v", err)
	}

	userCreatedDomain := converter.ConverterEntityToDomain(&userEntity)

	logger.Info(
		"CreateUser repository executed successfully",
		zap.String("userId", userCreatedDomain.ID),
		zap.String("journey", "createUser"))
	return &userCreatedDomain, nil

}
