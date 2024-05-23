package main

import (
	"net/http"
	"prometheus-exporters/pkg/exporters"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	exporters.EmitOSMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9090", nil)
}
