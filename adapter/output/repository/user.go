package repository

import (
	"context"
	"database/sql"
	"grouper/adapter/output/converter"
	"grouper/adapter/output/model/entity"
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
	logger.Info("CreateUser repository execution started", zap.String("journey", "createUser"))

	userEntity := converter.ConvertUserDomainToEntity(&userDomain)

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
		logger.Error("Error trying to create user in database", err, zap.String("journey", "createUser"))
		return nil, rest_errors.NewInternalServerError(err.Error())

	}

	userCreatedDomain := converter.ConverterUserEntityToDomain(&userEntity)

	logger.Info("CreateUser repository executed successfully", zap.String("userId", userCreatedDomain.ID), zap.String("journey", "createUser"))

	return &userCreatedDomain, nil
}

func (ur *userRepository) FindUserByUsername(username string) (*[]domain.UserDomain, *rest_errors.RestErr) {
	logger.Info("Start repository user find by username", zap.String("journey", "findUserByUsername"))

	query := "SELECT id, name, email, username, createdat FROM users WHERE username = $1"

	rows, err := ur.db.Query(query, username)
	if err != nil {
		logger.Error("Error trying to search user in database", err, zap.String("journey", "FindUserByUsername"))
		return nil, rest_errors.NewInternalServerError("")
	}
	defer rows.Close()

	var users []domain.UserDomain

	for rows.Next() {
		var userEntity entity.UserEntity

		err := rows.Scan(
			&userEntity.ID,
			&userEntity.Name,
			&userEntity.Email,
			&userEntity.Username,
			&userEntity.CreatedAt,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				logger.Error(
					"No user with this username", err, zap.String("journey", "FindUserByUsername"))
				return nil, rest_errors.NewNotFoundError("User not found")
			}
			logger.Error(
				"Error trying to scan user in database", err, zap.String("journey", "FindUserByUsername"))
			return nil, rest_errors.NewInternalServerError("")
		}
		user := converter.ConverterUserEntityToDomain(&userEntity)
		users = append(users, user)
	}

	return &users, nil
}

func (ur *userRepository) FindUserByEmail(email string) (*[]domain.UserDomain, *rest_errors.RestErr) {
	logger.Info("Start repository user find by email", zap.String("journey", "FindUserByEmail"))

	query := "SELECT id, name, email, username, createdat FROM users WHERE email = $1"

	rows, err := ur.db.Query(query, email)
	if err != nil {
		logger.Error("Error trying to search user in database", err, zap.String("journey", "FindUserByEmail"))
		return nil, rest_errors.NewInternalServerError("")
	}

	defer rows.Close()

	var users []domain.UserDomain

	for rows.Next() {
		var userEntity entity.UserEntity

		err := rows.Scan(
			&userEntity.ID,
			&userEntity.Name,
			&userEntity.Email,
			&userEntity.Username,
			&userEntity.CreatedAt,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				logger.Error(
					"No user with this username", err, zap.String("journey", "FindUserByEmail"))
				return nil, rest_errors.NewNotFoundError("User not found")
			}
			logger.Error(
				"Error trying to scan user in database", err, zap.String("journey", "FindUserByEmail"))
			return nil, rest_errors.NewInternalServerError("")
		}
		user := converter.ConverterUserEntityToDomain(&userEntity)
		users = append(users, user)
	}

	return &users, nil
}

func (ur *userRepository) Login(userDomain domain.UserDomain) (*domain.UserDomain, *rest_errors.RestErr) {
	userEntity := converter.ConvertUserDomainToEntity(&userDomain)

	logger.Info("Start repository Login",
		zap.String("journey", "Login"))

	query := "SELECT password, id FROM users WHERE email = $1"

	row, err := ur.db.Query(query, userEntity.Email)
	if err != nil {
		logger.Error("Error trying to search user in database", err, zap.String("journey", "Login"))
		return nil, rest_errors.NewInternalServerError("")
	}

	if row.Next() {
		err = row.Scan(&userEntity.Password, &userEntity.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				logger.Error("No user with this email", err, zap.String("journey", "Login"))
				return nil, rest_errors.NewNotFoundError("User not found")
			}
			logger.Error("Error trying to scan user in database", err, zap.String("journey", "Login"))
			return nil, rest_errors.NewInternalServerError("")
		}
		userDomain = converter.ConverterUserEntityToDomain(&userEntity)
	}

	return &userDomain, nil
}
