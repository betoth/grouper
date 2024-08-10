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
	logger.Debug("Init CreateUser repository", zap.String("journey", "createUser"))

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
	logger.Debug("Finish CreateUser repository", zap.String("journey", "createUser"))
	logger.Info("User created successfully", zap.String("userId", userCreatedDomain.ID), zap.String("journey", "createUser"))

	return &userCreatedDomain, nil
}

func (ur *userRepository) FindUserByUsername(username string) (*[]domain.UserDomain, *rest_errors.RestErr) {
	logger.Debug("Init FindUserByUsername repository", zap.String("journey", "FindUserByUsername"))
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
	logger.Debug("Finish FindUserByUsername repository", zap.String("journey", "FindUserByUsername"))
	return &users, nil
}

func (ur *userRepository) FindUserByEmail(email string) (*[]domain.UserDomain, *rest_errors.RestErr) {
	logger.Debug("Init FindUserByEmail repository", zap.String("journey", "FindUserByEmail"))

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
	logger.Debug("Finish FindUserByEmail repository", zap.String("journey", "FindUserByEmail"))
	return &users, nil
}

func (ur *userRepository) Login(userDomain domain.UserDomain) (*domain.UserDomain, *rest_errors.RestErr) {
	logger.Debug("Init Login repository", zap.String("journey", "Login"))

	userEntity := converter.ConvertUserDomainToEntity(&userDomain)

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
	logger.Debug("Finish Login repository", zap.String("journey", "Login"))
	return &userDomain, nil
}

func (ur *userRepository) GetUserGroups(userId string) (*[]domain.GroupDomain, *rest_errors.RestErr) {

	logger.Debug("Init GetUserGroups repository", zap.String("journey", "GetUserGroups"))
	var groups []domain.GroupDomain
	query := `SELECT g.id, g."name", t."name", st."name", g.created_at FROM user_groups ug 
INNER JOIN "groups" g ON ug.group_id = g.id 
inner join topic t on t.id = g.topic_id 
inner join subtopic st on st.id = g.subtopic_id 
WHERE ug.user_id = $1`

	rows, err := ur.db.Query(query, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("User is not a member of any group", err, zap.String("journey", "GetUserGroups"))
			return nil, rest_errors.NewNotFoundError("User is not a member of any group")
		}
		logger.Error("Error while GetUserGroups", err, zap.String("journey", "GetUserGroups"))
		return nil, rest_errors.NewInternalServerError("Internal server error")
	}
	defer rows.Close()

	for rows.Next() {
		var group domain.GroupDomain
		var topicName, subtopicName string
		if err := rows.Scan(&group.ID, &group.Name, &topicName, &subtopicName, &group.CreatedAt); err != nil {
			logger.Error("Error trying to get group in database", err, zap.String("journey", "GetUserGroups"))
			return nil, rest_errors.NewInternalServerError(err.Error())
		}
		groups = append(groups, group)
	}
	logger.Debug("Finish GetUserGroups repository", zap.String("journey", "GetUserGroups"))

	return &groups, nil
}
