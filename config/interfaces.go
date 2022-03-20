package config

//go:generate mockgen -source=interfaces.go -destination=mocks/config.go -package=mock_config

type Level interface {
	UnmarshalText(text []byte) error
}
