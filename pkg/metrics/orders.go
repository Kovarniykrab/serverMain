package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	orderCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "created_order_total",
			Help: "Total orders created.",
		}, []string{"type"},
	)
)

func IncOrderLaw() {
	orderCounter.WithLabelValues("law").Inc()
}

func IncOrderPhis() {
	orderCounter.WithLabelValues("phisic").Inc()
}
