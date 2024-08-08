package postgres

import (
	"database/sql"
	"fmt"
	"grouper/config"
	"grouper/config/logger"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func NewPostgresConnection(cfg *config.Config) (*sql.DB, error) {
	logger.Debug("Init NewPostgresConnection ", zap.String("journey", "Bootstrap"))
	ConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil

}
