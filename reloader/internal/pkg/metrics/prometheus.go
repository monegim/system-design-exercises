package metrics

import "github.com/prometheus/client_golang/prometheus"

type Collectors struct {
	Reloader *prometheus.CounterVec
}

func NewCollectors() Collectors {

	reloaded := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "reloader",
			Name:      "reload_executed_total",
			Help:      "Counter of reloads executed by Reloader.",
		},
		[]string{"success"},
	)
	reloaded.With(prometheus.Labels{"success": "true"}).Add(0)
	reloaded.With(prometheus.Labels{"success": "false"}).Add(0)
	return Collectors{
		Reloader: reloaded,
	}
}