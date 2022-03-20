package builder

import "go.uber.org/zap"

type Builder struct {
	cfg Config
}

func NewBuilder(cfg Config) *Builder {
	return &Builder{cfg: cfg}
}

func (b *Builder) Build() (*zap.Logger, error) {
	return b.cfg.Build()
}
