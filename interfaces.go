package logger

import "go.uber.org/zap"

//go:generate mockgen -source=interfaces.go -destination=mocks/logger.go -package=mock_logger

type Config interface {
	CreateLoggerConfig(level string) (cfg zap.Config, err error)
}
