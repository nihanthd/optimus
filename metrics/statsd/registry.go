package statsd

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

func WrapRegistry(prefix string, reg *prometheus.Registry, client *statsd.Client, tags []string) prometheus.Gatherer {
	family, err := reg.Gather()
	for _, f := range family {
		switch *f.Type {
		case dto.MetricType_GAUGE:
			for _, v := range f.Metric {
				client.Gauge(prefix+*f.Name, v.Gauge.GetValue(), append(toTags(v.Label), tags...), 1.0)
			}
		case dto.MetricType_COUNTER:
			for _, v := range f.Metric {
				client.Count(prefix+*f.Name, int64(v.Counter.GetValue()), append(toTags(v.Label), tags...), 1.0)
			}
		case dto.MetricType_HISTOGRAM:
			for _, v := range f.Metric {
				client.Histogram(prefix+*f.Name, v.Counter.GetValue(), append(toTags(v.Label), tags...), 1.0)
			}
		}
	}
	return statsdRegistry{metricsFamily: family, err: err}
}

type statsdRegistry struct {
	metricsFamily []*dto.MetricFamily
	err           error
}

func (r statsdRegistry) Gather() ([]*dto.MetricFamily, error) {
	return r.metricsFamily, r.err
}

func toTags(labels []*dto.LabelPair) []string {
	tags := make([]string, len(labels))
	for _, lp := range labels {
		tags = append(tags, *lp.Name+":"+*lp.Value)
	}
	return tags
}
