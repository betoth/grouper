package repository

import (
	"database/sql"
	"grouper/adapter/output/converter"
	"grouper/application/domain"
	"grouper/application/port/output"
	"grouper/config/logger"
	"grouper/config/rest_errors"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func NewLoginRepository(database *sql.DB) output.LoginPort {
	return &loginRepository{
		database,
	}
}

type loginRepository struct {
	db *sql.DB
}

func (lr *loginRepository) Login(loginDomain domain.LoginDomain) (*domain.LoginDomain, *rest_errors.RestErr) {
	loginEntity := converter.ConvertLoginDomainToEntity(&loginDomain)

	logger.Info("Start repository Login",
		zap.String("journey", "Login"))

	query := "SELECT password FROM users WHERE email = $1"

	row, err := lr.db.Query(query, loginEntity.Email)
	if err != nil {
		logger.Error(
			"Error trying to search user in database",
			err,
			zap.String("journey", "Login"))
		return nil, rest_errors.NewInternalServerError("")
	}

	if row.Next() {

		err = row.Scan(&loginEntity.Password)
		if err != nil {
			if err == sql.ErrNoRows {
				logger.Error(
					"No user with this email",
					err,
					zap.String("journey", "Login"))
				return nil, rest_errors.NewNotFoundError("User not found")
			}
			logger.Error(
				"Error trying to scan user in database",
				err,
				zap.String("journey", "Login"))
			return nil, rest_errors.NewInternalServerError("")
		}

		loginDomain = converter.ConverterLoginEntityToDomain(&loginEntity)

	}

	return &loginDomain, nil
}
