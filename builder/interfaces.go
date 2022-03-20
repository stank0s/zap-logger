package builder

import "go.uber.org/zap"

//go:generate mockgen -source=interfaces.go -destination=mocks/builder.go -package=mock_builder

type Config interface {
	Build(opts ...zap.Option) (*zap.Logger, error)
}
