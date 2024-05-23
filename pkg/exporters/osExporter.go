package exporters

import (
	"log"
	"prometheus-exporters/pkg/clients"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	cpuLoad1m = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_load_1m",
		Help: "Average CPU load 1 minute",
	})
	cpuLoad5m = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_load_5m",
		Help: "Average CPU load 5 minutes",
	})
	cpuLoad15m = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_load_15m",
		Help: "Average CPU load 15 minutes",
	})
	totalMem = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "total_memory_bytes",
		Help: "Total memory in bytes",
	})
	usedMem = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "used_memory_bytes",
		Help: "Used memory in bytes",
	})
	availableMem = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "available_memory_bytes",
		Help: "Available memory in bytes",
	})
)

type OSExporter struct {
	cpuInterval time.Duration
	memInterval time.Duration
	osClient    clients.OSClient
}

func NewOSExporter(client clients.OSClient, cpuInterval, memInterval time.Duration) *OSExporter {
	return &OSExporter{
		cpuInterval: cpuInterval,
		memInterval: memInterval,
		osClient:    client,
	}
}

func (e *OSExporter) Run() {
	go e.sendCPUMetrics()
	go e.sendMemMetrics()
}

func (e *OSExporter) sendCPUMetrics() {
	for {
		v, err := e.osClient.GetCPU()
		if err != nil {
			log.Fatal(err)
		}

		cpuLoad1m.Set(v.Load1 * 100)
		cpuLoad5m.Set(v.Load5 * 100)
		cpuLoad15m.Set(v.Load15 * 100)

		time.Sleep(time.Duration(e.cpuInterval) * time.Millisecond)
	}
}

func (e *OSExporter) sendMemMetrics() {
	for {
		v, err := e.osClient.GetMem()

		if err != nil {
			log.Fatal(err)
		}

		availableMem.Set(float64(v.Available))
		usedMem.Set(float64(v.Used))
		totalMem.Set(float64(v.Total))
		time.Sleep(time.Duration(e.memInterval) * time.Millisecond)
	}
}
