package config

import (
	"grouper/config/env"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBName     string
	DBPassword string
	DBSSLMode  string
	LOGOutput  string
	LOGLevel   string
	APIPort    string
	SECRET_KEY string
}

func NewConfig() *Config {
	env.LoadEnv()

	return &Config{
		DBHost:     env.GetEnv("DB_HOST"),
		DBPort:     env.GetEnv("DB_PORT"),
		DBUser:     env.GetEnv("DB_USER"),
		DBName:     env.GetEnv("DB_NAME"),
		DBPassword: env.GetEnv("DB_PASSWORD"),
		DBSSLMode:  env.GetEnv("DB_SSL_MODE"),
		LOGOutput:  env.GetEnv("LOG_OUTPUT"),
		LOGLevel:   env.GetEnv("LOG_LEVEL"),
		APIPort:    env.GetEnv("API_PORT"),
	}
}
