package exporter

import (
	"github.com/labstack/echo"
	"github.com/nihanthd/optimus/metrics/statsd"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsHandler struct {
	exporter *Exporter
}

func NewMetricsHandler(exporter *Exporter) *MetricsHandler {
	return &MetricsHandler{
		exporter: exporter,
	}
}

func (h *MetricsHandler) GetMetrics(c echo.Context) error {
	labels := prometheus.Labels{}
	registry := h.exporter.GetRegistry(labels)
	if registry != nil {
		tags := []string{}
		gatherer := statsd.WrapRegistry("pi"+"_", registry, h.exporter.statsd, tags)
		promhttp.HandlerFor(gatherer, promhttp.HandlerOpts{}).
			ServeHTTP(c.Response().Writer, c.Request())
	}
	return nil
}
