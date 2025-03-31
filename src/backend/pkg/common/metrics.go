package common

import "github.com/prometheus/client_golang/prometheus"

const (
	bucketMultiplier = 1
	bucketFactor     = 2
	bucketCount      = 15
)

type Metrics struct {
	RequestCounter  *prometheus.CounterVec
	RequestDuration *prometheus.HistogramVec
}

func NewMetrics() *Metrics {
	metrics := &Metrics{
		RequestCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total number of HTTP requests",
			},
			[]string{"handler", "method", "code"},
		),
		RequestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_request_duration_milliseconds",
				Help:    "Histogram of response latency (milliseconds) of HTTP requests.",
				Buckets: prometheus.ExponentialBuckets(bucketMultiplier, bucketFactor, bucketCount),
			},
			[]string{"handler", "method"},
		),
	}

	prometheus.MustRegister(metrics.RequestCounter)
	prometheus.MustRegister(metrics.RequestDuration)

	return metrics
}
