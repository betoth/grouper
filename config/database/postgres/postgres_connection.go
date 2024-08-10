package postgres

import (
	"fmt"
	"grouper/config"
	"grouper/config/logger"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection(cfg *config.Config) (*gorm.DB, error) {
	logger.Debug("Init NewPostgresConnection ", zap.String("journey", "Bootstrap"))
	ConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=America/Sao_Paulo",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)
	db, err := gorm.Open(postgres.Open(ConnStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil

}
