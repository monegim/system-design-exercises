package metrics

import "github.com/prometheus/client_golang/prometheus"

type Collectors struct {
	Reloader *prometheus.CounterVec
}