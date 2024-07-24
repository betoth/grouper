package logger

import (
	"grouper/config"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func Init(cfg *config.Config) {
	logConfig := zap.Config{
		OutputPaths: []string{getOutputLogs(cfg.LOGOutput)},
		Level:       zap.NewAtomicLevelAt(getLevelLogs(cfg.LOGLevel)),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfig.Build()
}

func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	errSync := log.Sync()
	if errSync != nil {
		return
	}
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Info(message, tags...)
	errSync := log.Sync()
	if errSync != nil {
		return
	}
}

func getOutputLogs(logOutput string) string {
	output := strings.ToLower(strings.TrimSpace(logOutput))
	if output == "" {
		return "stdout"
	}

	return output
}

func getLevelLogs(logLevel string) zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(logLevel)) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}
