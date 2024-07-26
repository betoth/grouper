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

func NewGroupRepository(database *sql.DB) output.GroupPort {
	return &groupRepository{
		database,
	}
}

type groupRepository struct {
	db *sql.DB
}

func (gr *groupRepository) CreateGroup(groupDomain domain.GroupDomain) (*domain.GroupDomain, *rest_errors.RestErr) {
	logger.Info("CreateGroup repository execution started",
		zap.String("journey", "CreateGroup"))

	groupEntity := converter.ConvertGroupDomainToEntity(&groupDomain)

	query := `
	INSERT INTO groups (id, name, user_id, created_at) 
	VALUES (gen_random_uuid(), $1, $2, $3)
	RETURNING id, created_at;
    `

	row := gr.db.QueryRowContext(
		context.Background(),
		query,
		groupEntity.Name,
		groupEntity.UserID,
		groupEntity.CreatedAt,
	)

	err := row.Scan(&groupEntity.ID, &groupEntity.CreatedAt)
	if err != nil {
		logger.Error("Error trying to create group in database",
			err,
			zap.String("journey", "CreateGroup"))
		return nil, rest_errors.NewInternalServerError(err.Error())

	}

	groupCreatedDomain := converter.ConverterGroupEntityToDomain(&groupEntity)

	logger.Info(
		"CreateGroup repository executed successfully",
		zap.String("groupId", groupCreatedDomain.ID),
		zap.String("journey", "CreateGroup"))

	return &groupCreatedDomain, nil
}
