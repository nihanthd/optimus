package log

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type (
	Config struct {
		Enabled bool   `yaml:enabled`
		Level   string `yaml:"level"`
	}
)

func NewDevLogger(lc fx.Lifecycle) (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return logger.Sync()
		},
	})
	return logger, nil
}

func NewProdLogger(lc fx.Lifecycle) (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return logger.Sync()
		},
	})
	return logger, nil
}

func NewConfigLogger(lc fx.Lifecycle, config *Config) (*zap.Logger, error) {

	// Logger
	zapConfig := zap.NewProductionConfig()
	zapConfig.OutputPaths = []string{"stdout"}

	// No Sampling
	zapConfig.Sampling = nil

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}

	// Set Log level from config
	if config.Level != "" {
		if err := zapConfig.Level.UnmarshalText([]byte(strings.ToLower(config.Level))); err != nil {
			return nil, errors.Errorf("Unable to parse log level from config %s: %s", config.Level, err.Error())
		}
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return logger.Sync()
		},
	})

	return logger, nil
}
