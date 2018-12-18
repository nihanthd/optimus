package statsd

import (
	"context"

	"github.com/DataDog/datadog-go/statsd"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type (
	Config struct {
		Host    string `yaml:"host"`
		Enabled bool   `yaml:enabled`
	}
)

func NewClient(lc fx.Lifecycle, logger *zap.Logger, config *Config) (*statsd.Client, error) {

	var err error
	var stats *statsd.Client

	if config.Enabled {
		logger.Info("Starting statsd client", zap.String("host", config.Host))
		stats, err = statsd.New(config.Host)
		if err != nil {
			return nil, err
		}
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			logger.Sugar().Info("Stopping statsd client.")
			if stats != nil {
				return stats.Close()
			}
			return nil
		},
	})

	return stats, nil
}
