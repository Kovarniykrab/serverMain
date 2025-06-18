package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	connectRabbitCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "reconnect_rabbit_total",
			Help: "Total rabbit reconnect.",
		}, []string{"type"},
	)
)

func IncRabbitLaw() {
	connectRabbitCounter.WithLabelValues("law").Inc()
}

func IncRabbitPhis() {
	connectRabbitCounter.WithLabelValues("phisic").Inc()
}
