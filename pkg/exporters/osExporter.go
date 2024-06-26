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
	usedMem = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "used_memory_bytes",
		Help: "Used memory in bytes",
	})
	availableMem = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "available_memory_bytes",
		Help: "Available memory in bytes",
	})
	totalMem = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "total_memory_bytes",
		Help: "Total memory in bytes",
	})
	freeMem = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "free_memory_bytes",
		Help: "Free memory in bytes",
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
	cpuTicker := time.NewTicker(e.cpuInterval)
	memTicker := time.NewTicker(e.memInterval)

	for {
		select {
		case <-cpuTicker.C:
			e.updateCPUMetrics()
		case <-memTicker.C:
			e.updateMemMetrics()
		}
	}
}

func (e *OSExporter) updateCPUMetrics() {
	v, err := e.osClient.GetCPU()
	if err != nil {
		log.Fatal(err)
	}

	cpuLoad1m.Set(v.Load1 * 100)
	cpuLoad5m.Set(v.Load5 * 100)
	cpuLoad15m.Set(v.Load15 * 100)
}

func (e *OSExporter) updateMemMetrics() {
	v, err := e.osClient.GetMem()

	if err != nil {
		log.Fatal(err)
	}

	usedMem.Set(float64(v.Used))
	availableMem.Set(float64(v.Available))
	freeMem.Set(float64(v.Free))
	totalMem.Set(float64(v.Total))
}
