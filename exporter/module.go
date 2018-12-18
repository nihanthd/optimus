package exporter

import "go.uber.org/fx"

var Module = fx.Provide(
	NewExporter,
	NewMetricsHandler,
)
