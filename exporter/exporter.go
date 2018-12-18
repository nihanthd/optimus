package exporter

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/go-resty/resty"
	"github.com/nihanthd/optimus/collectors"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

type Exporter struct {
	statsd *statsd.Client
	log    *zap.Logger
	client *resty.Client
}

func NewExporter(log *zap.Logger, statsd *statsd.Client, client *resty.Client) *Exporter {
	return &Exporter{
		log:    log,
		statsd: statsd,
		client: client,
	}
}

func (e *Exporter) GetRegistry(labels prometheus.Labels) *prometheus.Registry {
	registry := prometheus.NewRegistry()
	var collector prometheus.Collector
	collector = collectors.NewPiCollector(labels, e.log, e.client)
	registry.MustRegister(collector)
	return registry
}
