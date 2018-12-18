package collectors

import (
	"github.com/go-resty/resty"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

const (
	One float64 = 1.0
)

var (
	requests = prometheus.NewDesc(
		"input_requests",
		"Number of requests received by PI",
		nil, nil)
)

//implements prometheus.Collector interface
type piCollector struct {
	url    string
	labels prometheus.Labels
	logger *zap.Logger
	client *resty.Client
}

func (p *piCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- requests
}

func (p *piCollector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(requests, prometheus.CounterValue, One)
}

func NewPiCollector(labels prometheus.Labels, logger *zap.Logger, client *resty.Client) prometheus.Collector {
	// Parse url from labels
	url := labels["scheme"] + "://" + labels["target"] + labels["path"]
	return &piCollector{url, labels, logger, client}
}
