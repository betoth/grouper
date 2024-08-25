package postgres

import (
	"fmt"
	"grouper/config"
	"grouper/config/logger"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gLogger "gorm.io/gorm/logger"
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

	gormLogger := gLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		gLogger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  gLogger.Error,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(postgres.Open(ConnStr), &gorm.Config{
		TranslateError: true,
		Logger:         gormLogger,
	})
	if err != nil {
		return nil, err
	}

	return db, nil

}
