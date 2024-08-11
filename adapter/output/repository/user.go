package repository

import (
	"grouper/adapter/output/converter"
	"grouper/adapter/output/model/entity"
	"grouper/application/domain"
	"grouper/application/port/output"
	"grouper/config/logger"
	"grouper/config/rest_errors"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewUserRepository(database *gorm.DB) output.UserPort {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	db *gorm.DB
}

func (repo *userRepository) CreateUser(userDomain domain.User) (*domain.User, *rest_errors.RestErr) {
	logger.Debug("Init CreateUser repository", zap.String("journey", "createUser"))

	userEntity := converter.ConvertUserDomainToEntity(&userDomain)

	if err := repo.db.Create(&userEntity).Error; err != nil {
		logger.Error("Error trying to create user in database", err, zap.String("journey", "createUser"))
		return nil, rest_errors.NewInternalServerError("Internal server error")
	}

	userCreatedDomain := converter.ConverterUserEntityToDomain(&userEntity)
	logger.Debug("Finish CreateUser repository", zap.String("journey", "createUser"))
	logger.Info("User created successfully", zap.String("userId", userCreatedDomain.ID), zap.String("journey", "createUser"))

	return &userCreatedDomain, nil
}

func (repo *userRepository) FindUserByUsername(username string) (*[]domain.User, *rest_errors.RestErr) {
	logger.Debug("Init FindUserByUsername repository", zap.String("journey", "FindUserByUsername"))

	var entities []entity.User
	if err := repo.db.Where("username = ?", username).Find(&entities).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Error("No user with this username", err, zap.String("journey", "FindUserByUsername"))
			return nil, rest_errors.NewNotFoundError("User not Found")
		}
		logger.Error("Error trying to search user in database", err, zap.String("journey", "FindUserByUsername"))
		return nil, rest_errors.NewInternalServerError("Internal server error")
	}

	if len(entities) == 0 {
		err := rest_errors.NewNotFoundError("User not Found")
		logger.Error(err.Message, nil, zap.String("journey", "FindUserByUsername"))
		return nil, err
	}

	var users []domain.User
	for _, entity := range entities {
		user := converter.ConverterUserEntityToDomain(&entity)
		users = append(users, user)
	}

	logger.Debug("Finish FindUserByUsername repository", zap.String("journey", "FindUserByUsername"))
	return &users, nil
}

func (repo *userRepository) FindUserByEmail(email string) (*[]domain.User, *rest_errors.RestErr) {
	logger.Debug("Init FindUserByEmail repository", zap.String("journey", "FindUserByEmail"))

	var entities []entity.User
	if err := repo.db.Where("email = ?", email).Find(&entities).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Error("No user with this email", err, zap.String("journey", "FindUserByEmail"))
			return nil, rest_errors.NewNotFoundError("User not found")
		}
		logger.Error("Error trying to search user in database", err, zap.String("journey", "FindUserByEmail"))
		return nil, rest_errors.NewInternalServerError("Internal server error")
	}

	if len(entities) == 0 {
		err := rest_errors.NewNotFoundError("User not found")
		logger.Error(err.Message, nil, zap.String("journey", "FindUserByEmail"))
		return nil, err
	}

	var users []domain.User
	for _, entity := range entities {
		user := converter.ConverterUserEntityToDomain(&entity)
		users = append(users, user)
	}

	logger.Debug("Finish FindUserByEmail repository", zap.String("journey", "FindUserByEmail"))
	return &users, nil
}

func (repo *userRepository) Login(userDomain domain.User) (*domain.User, *rest_errors.RestErr) {
	logger.Debug("Init Login repository", zap.String("journey", "Login"))

	var userEntity entity.User
	err := repo.db.Where("email = ?", userDomain.Email).First(&userEntity).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Error("No user with this email", err, zap.String("journey", "Login"))
			return nil, rest_errors.NewNotFoundError("User not found")
		}
		logger.Error("Error trying to search user in database", err, zap.String("journey", "Login"))
		return nil, rest_errors.NewInternalServerError("Internal server error")
	}

	userDomain = converter.ConverterUserEntityToDomain(&userEntity)

	logger.Debug("Finish Login repository", zap.String("journey", "Login"))
	return &userDomain, nil
}

func (repo *userRepository) GetUserGroups(userId string) (*[]domain.Group, *rest_errors.RestErr) {
	logger.Debug("Init GetUserGroups repository", zap.String("journey", "GetUserGroups"))

	var groups []domain.Group

	if err := repo.db.Table("user_groups").
		Select("g.id, g.name, g.created_at").
		Joins("INNER JOIN groups g ON user_groups.group_id = g.id").
		Where("user_groups.user_id = ?", userId).
		Scan(&groups).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Error("User is not a member of any group", err, zap.String("journey", "GetUserGroups"))
			return nil, rest_errors.NewNotFoundError("User is not a member of any group")
		}
		logger.Error("Error while GetUserGroups", err, zap.String("journey", "GetUserGroups"))
		return nil, rest_errors.NewInternalServerError("Internal server error")
	}

	logger.Debug("Finish GetUserGroups repository", zap.String("journey", "GetUserGroups"))
	return &groups, nil
}
