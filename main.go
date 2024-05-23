package main

import (
	"net/http"
	"prometheus-exporters/pkg/clients"
	"prometheus-exporters/pkg/exporters"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	osExporter := exporters.NewOSExporter(
		clients.NewMacOSClient(),
		time.Second*10, time.Second*10,
	)
	go osExporter.Run()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9090", nil)
}
