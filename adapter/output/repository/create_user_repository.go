package repository

import (
	"context"
	"database/sql"
	"grouper/adapter/output/converter"
	"grouper/application/domain"
	"grouper/application/port/output"
	"grouper/config/logger"
	"grouper/config/rest_errors"

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

func (ur *userRepository) CreateUser(userDomain domain.UserDomain) (*domain.UserDomain, *rest_errors.RestErr) {
	logger.Info("CreateUser repository execution started",
		zap.String("journey", "createUser"))

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

	err := row.Scan(&userEntity.ID, &userEntity.CreatedAt)
	if err != nil {
		logger.Error("Error trying to create user in database",
			err,
			zap.String("journey", "createUser"))
		return nil, rest_errors.NewInternalServerError(err.Error())

	}

	userCreatedDomain := converter.ConverterEntityToDomain(&userEntity)

	logger.Info(
		"CreateUser repository executed successfully",
		zap.String("userId", userCreatedDomain.ID),
		zap.String("journey", "createUser"))

	return &userCreatedDomain, nil

}
