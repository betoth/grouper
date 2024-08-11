package repository

import (
	"database/sql"
	"fmt"
	"grouper/adapter/output/converter"
	"grouper/adapter/output/model/entity"
	"grouper/application/domain"
	"grouper/application/dto"
	"grouper/application/port/output"
	"grouper/config/logger"
	"grouper/config/rest_errors"

	"github.com/lib/pq"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewGroupRepository(database *gorm.DB) output.GroupPort {
	return &groupRepository{
		database,
	}
}

type groupRepository struct {
	db *gorm.DB
}

// isForeignKeyViolation checks if the error is a foreign key violation
func isForeignKeyViolation(err error) bool {
	pgErr, ok := err.(*pq.Error)
	return ok && pgErr.Code == "23503" // Foreign key violation error code
}

func (repo *groupRepository) CreateGroup(groupDomain domain.Group) (*domain.Group, *rest_errors.RestErr) {
	logger.Debug("Init CreateGroup repository", zap.String("journey", "CreateGroup"))

	groupEntity := converter.ConvertGroupDomainToEntity(&groupDomain)

	result := repo.db.Create(&groupEntity)
	if result.Error != nil {
		logger.Error("Error trying to create group in database", result.Error, zap.String("journey", "CreateGroup"))
		return nil, rest_errors.NewInternalServerError(result.Error.Error())
	}

	groupCreatedDomain := converter.ConverterGroupEntityToDomain(&groupEntity)

	logger.Debug("Finish CreateGroup repository", zap.String("journey", "CreateGroup"))
	logger.Info("Group created successfully", zap.String("groupId", groupCreatedDomain.ID), zap.String("journey", "CreateGroup"))

	return &groupCreatedDomain, nil
}

func (repo *groupRepository) JoinGroup(userID, groupID string) *rest_errors.RestErr {
	logger.Debug("Init JoinGroup repository", zap.String("journey", "JoinGroup"))

	userGroup := entity.UserGroup{
		UserID:  userID,
		GroupID: groupID,
	}
	result := repo.db.Create(&userGroup)

	if result.Error != nil {
		if isForeignKeyViolation(result.Error) {
			logger.Error("Group ID does not exist", result.Error, zap.String("journey", "JoinGroup"))
			return rest_errors.NewNotFoundError("Group not found")
		}
		logger.Error("Error while trying to join group", result.Error, zap.String("journey", "JoinGroup"))
		return rest_errors.NewInternalServerError("Failed to join group")
	}
	logger.Debug("Finish JoinGroup repository", zap.String("journey", "JoinGroup"))
	logger.Info("Successfully joined the group", zap.String("user_id", userID), zap.String("group_id", groupID), zap.String("journey", "JoinGroup"))

	return nil
}

func (repo *groupRepository) LeaveGroup(userID, groupID string) *rest_errors.RestErr {
	logger.Debug("Init LeaveGroup repository", zap.String("journey", "LeaveGroup"))

	var userGroup entity.UserGroup

	result := repo.db.Where("user_id = ? AND group_id = ?", userID, groupID).Delete(&userGroup)
	if result.Error != nil {
		if result.Error == sql.ErrNoRows {
			logger.Error("User is not a member of this group", result.Error, zap.String("journey", "LeaveGroup"))
			return rest_errors.NewNotFoundError("User is not a member of this group")
		}
		logger.Error("Error while trying to leave group", result.Error, zap.String("journey", "LeaveGroup"))
		return rest_errors.NewInternalServerError("Internal server error")
	}
	logger.Debug("Finish LeaveGroup repository", zap.String("journey", "LeaveGroup"))
	logger.Info("Successfully left the group", zap.String("user_id", userID), zap.String("group_id", groupID), zap.String("journey", "LeaveGroup"))
	return nil
}

func (repo *groupRepository) GetGroups(parameters dto.GetGroupsParameter) (*[]domain.Group, *rest_errors.RestErr) {
	logger.Debug("Init GetGroups repository", zap.String("journey", "GetGroups"))

	var groupsEntity []entity.Group
	dbQuery := repo.db.Model(&entity.Group{})

	if parameters.Name != "" {
		dbQuery = dbQuery.Where("LOWER(name) LIKE ?", fmt.Sprintf("%%%s%%", parameters.Name))
	}

	err := dbQuery.Find(&groupsEntity).Error
	if err != nil {
		logger.Error("Error while get groups", err, zap.String("journey", "GetGroups"))
		return nil, rest_errors.NewInternalServerError("Internal server error")
	}

	if len(groupsEntity) == 0 {
		err := rest_errors.NewNotFoundError("No group meets the search criteria")
		logger.Error(err.Message, nil, zap.String("journey", "GetGroups"))
		return nil, err
	}

	var groupsDomain []domain.Group
	for _, groupEntity := range groupsEntity {
		groupDomain := converter.ConverterGroupEntityToDomain(&groupEntity)
		groupsDomain = append(groupsDomain, groupDomain)
	}

	logger.Debug("Finish GetGroups repository", zap.String("journey", "GetGroups"))
	return &groupsDomain, nil
}
