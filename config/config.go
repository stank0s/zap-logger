package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	atom zap.AtomicLevel
	lvl  Level
}

func NewConfig(atom zap.AtomicLevel) *Config {
	return &Config{
		atom: atom,
		lvl:  &atom,
	}
}

func (c *Config) CreateLoggerConfig(level string, colored bool) (cfg zap.Config, err error) {
	encodeLevel := zapcore.CapitalLevelEncoder
	if colored {
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}

	cfg = zap.Config{
		Encoding:    "json",
		Level:       c.atom,
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "severity",
			EncodeLevel: encodeLevel,
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
		},
	}

	return cfg, c.lvl.UnmarshalText([]byte(level))
}
