package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "pvz_requests_total",
			Help: "Total number of requests received by the PVZ service.",
		},
		[]string{"method"},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "pvz_request_duration_seconds",
			Help:    "Histogram of response durations for the PVZ service.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)

	CreatedPVZs = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "pvz_created_total",
			Help: "Total number of created PVZs.",
		},
	)

	CreatedReceptions = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "receptions_created_total",
			Help: "Total number of created receptions.",
		},
	)

	AddedProducts = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "products_added_total",
			Help: "Total number of products added.",
		},
	)
)

func init() {
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(RequestDuration)
	prometheus.MustRegister(CreatedPVZs)
	prometheus.MustRegister(CreatedReceptions)
	prometheus.MustRegister(AddedProducts)
}

func IncrementRequestCount(method string) {
	RequestCount.WithLabelValues(method).Inc()
}

func ObserveRequestDuration(method string, duration float64) {
	RequestDuration.WithLabelValues(method).Observe(duration)
}

func IncrementCreatedPVZs() {
	CreatedPVZs.Inc()
}

func IncrementCreatedReceptions() {
	CreatedReceptions.Inc()
}

func IncrementAddedProducts() {
	AddedProducts.Inc()
}
