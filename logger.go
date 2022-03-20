package logger

import (
	"github.com/stank0s/zap-logger/builder"
	"github.com/stank0s/zap-logger/config"
	"go.uber.org/zap"
)

var build func(zap.Config) (*zap.Logger, error) = func(cfg zap.Config) (*zap.Logger, error) {
	b := builder.NewBuilder(cfg)
	return b.Build()
}

type Logger struct {
	cfg Config
}

func NewLogger() *Logger {
	return &Logger{
		cfg: config.NewConfig(zap.NewAtomicLevel()),
	}
}

func (l *Logger) Logger(level string) (*zap.Logger, error) {
	cfg, err := l.cfg.CreateLoggerConfig(level)
	if err != nil {
		return nil, err
	}

	log, err := build(cfg)
	if err != nil {
		return nil, err
	}

	return log, nil
}

func (l *Logger) SugaredLogger(level string) (sl *zap.SugaredLogger, err error) {
	var log *zap.Logger
	log, err = l.Logger(level)
	if err != nil {
		return nil, err
	}

	sl = log.Sugar()
	defer sl.Sync()

	return sl, nil
}
