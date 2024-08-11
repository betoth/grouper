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
	logger.Info("Init NewPostgresConnection ", zap.String("journey", "Bootstrap"))
	ConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
		cfg.TimeZone,
	)
	db, err := gorm.Open(postgres.Open(ConnStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil

}
