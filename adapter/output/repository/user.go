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

func (ur *userRepository) CreateUser(userDomain domain.UserDomain) (*domain.UserDomain, *rest_errors.RestErr) {
	logger.Debug("Init CreateUser repository", zap.String("journey", "createUser"))

	userEntity := converter.ConvertUserDomainToEntity(&userDomain)

	if err := ur.db.Create(&userEntity).Error; err != nil {
		logger.Error("Error trying to create user in database", err, zap.String("journey", "createUser"))
		return nil, rest_errors.NewInternalServerError("Internal server error")
	}

	userCreatedDomain := converter.ConverterUserEntityToDomain(&userEntity)
	logger.Debug("Finish CreateUser repository", zap.String("journey", "createUser"))
	logger.Info("User created successfully", zap.String("userId", userCreatedDomain.ID), zap.String("journey", "createUser"))

	return &userCreatedDomain, nil
}

func (ur *userRepository) FindUserByUsername(username string) (*[]domain.UserDomain, *rest_errors.RestErr) {
	logger.Debug("Init FindUserByUsername repository", zap.String("journey", "FindUserByUsername"))

	var entities []entity.User
	if err := ur.db.Where("username = ?", username).Find(&entities).Error; err != nil {
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

	var users []domain.UserDomain
	for _, entity := range entities {
		user := converter.ConverterUserEntityToDomain(&entity)
		users = append(users, user)
	}

	logger.Debug("Finish FindUserByUsername repository", zap.String("journey", "FindUserByUsername"))
	return &users, nil
}

func (ur *userRepository) FindUserByEmail(email string) (*[]domain.UserDomain, *rest_errors.RestErr) {
	logger.Debug("Init FindUserByEmail repository", zap.String("journey", "FindUserByEmail"))

	var entities []entity.User
	if err := ur.db.Where("email = ?", email).Find(&entities).Error; err != nil {
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

	var users []domain.UserDomain
	for _, entity := range entities {
		user := converter.ConverterUserEntityToDomain(&entity)
		users = append(users, user)
	}

	logger.Debug("Finish FindUserByEmail repository", zap.String("journey", "FindUserByEmail"))
	return &users, nil
}

func (ur *userRepository) Login(userDomain domain.UserDomain) (*domain.UserDomain, *rest_errors.RestErr) {
	logger.Debug("Init Login repository", zap.String("journey", "Login"))

	var userEntity entity.User
	err := ur.db.Where("email = ?", userDomain.Email).First(&userEntity).Error

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

func (ur *userRepository) GetUserGroups(userId string) (*[]domain.GroupDomain, *rest_errors.RestErr) {
	logger.Debug("Init GetUserGroups repository", zap.String("journey", "GetUserGroups"))

	var groups []domain.GroupDomain

	if err := ur.db.Table("user_groups").
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
