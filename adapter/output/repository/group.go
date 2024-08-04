package repository

import (
	"context"
	"database/sql"
	"grouper/adapter/output/converter"
	"grouper/application/domain"
	"grouper/application/port/output"
	"grouper/config/logger"
	"grouper/config/rest_errors"

	"github.com/lib/pq"
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
	logger.Info("CreateGroup repository execution started", zap.String("journey", "CreateGroup"))

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
		logger.Error("Error trying to create group in database", err, zap.String("journey", "CreateGroup"))
		return nil, rest_errors.NewInternalServerError(err.Error())

	}

	groupCreatedDomain := converter.ConverterGroupEntityToDomain(&groupEntity)

	logger.Info("CreateGroup repository executed successfully", zap.String("groupId", groupCreatedDomain.ID), zap.String("journey", "CreateGroup"))

	return &groupCreatedDomain, nil
}

func (gr *groupRepository) Join(userID, groupID string) *rest_errors.RestErr {
	logger.Info("JoinGroup repository execution started", zap.String("journey", "JoinGroup"))

	query := `INSERT INTO public.user_groups
	(id, user_id, group_id, joined_at)
	VALUES(gen_random_uuid(), $1, $2, now())
	RETURNING id;`

	var insertID string

	err := gr.db.QueryRow(query, userID, groupID).Scan(&insertID)
	if err != nil {
		if isForeignKeyViolation(err) {
			logger.Error("Group ID does not exist", err, zap.String("journey", "JoinGroup"))
			return rest_errors.NewNotFoundError("Group not found")
		}
		logger.Error("Error while trying to join group", err, zap.String("journey", "JoinGroup"))
		return rest_errors.NewInternalServerError("Failed to join group")
	}

	logger.Info("Successfully joined the group", zap.String("journey", "JoinGroup"))
	return nil
}

// isForeignKeyViolation checks if the error is a foreign key violation
func isForeignKeyViolation(err error) bool {
	// Example for PostgreSQL, adjust if using a different database
	pgErr, ok := err.(*pq.Error)
	return ok && pgErr.Code == "23503" // Foreign key violation error code
}

func (gr *groupRepository) Leave(userID, groupID string) *rest_errors.RestErr {
	logger.Info("LeaveGroup repository execution started", zap.String("journey", "LeaveGroup"))

	query := `DELETE FROM public.user_groups 
	WHERE user_id = $1 AND group_id = $2
	RETURNING id;`

	var deletedID string

	err := gr.db.QueryRow(query, userID, groupID).Scan(&deletedID)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("User is not a member of this group", err, zap.String("journey", "LeaveGroup"))
			return rest_errors.NewNotFoundError("User is not a member of this group")
		}
		logger.Error("Error while trying to leave group", err, zap.String("journey", "LeaveGroup"))
		return rest_errors.NewInternalServerError("Internal server error")
	}

	logger.Info("Successfully left the group", zap.String("journey", "LeaveGroup"))
	return nil
}
