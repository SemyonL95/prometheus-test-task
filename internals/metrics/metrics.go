package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	MetricsIPName = "unique_ip_addresses"
	MetricsIpHelp = "The total number of unique ip adresses"
)

type Metrics struct {
	UniqueIPCount prometheus.Counter
}

func New() *Metrics {
	return &Metrics{
		UniqueIPCount: promauto.NewCounter(prometheus.CounterOpts{
			Name: MetricsIPName,
			Help: MetricsIpHelp,
		}),
	}
}

func (m *Metrics) Inc() {
	m.UniqueIPCount.Inc()
}
