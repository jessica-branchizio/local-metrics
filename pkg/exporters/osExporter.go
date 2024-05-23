package exporters

import (
	"log"
	"prometheus-exporters/pkg/clients"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	cpuLoad1m    prometheus.Gauge
	cpuLoad5m    prometheus.Gauge
	cpuLoad15m   prometheus.Gauge
	totalMem     prometheus.Gauge
	usedMem      prometheus.Gauge
	availableMem prometheus.Gauge
	client       clients.Client
}

func NewMetrics(cpuIteration, memIteration int) *metrics {
	return &metrics{
		cpuLoad1m: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "cpu_load_1m",
			Help: "Average CPU load 1 minute",
		}),
		cpuLoad5m: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "cpu_load_5m",
			Help: "Average CPU load 5 minutes",
		}),
		cpuLoad15m: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "cpu_load_15m",
			Help: "Average CPU load 15 minutes",
		}),
		totalMem: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "total_mem_bytes",
			Help: "Total memory in bytes",
		}),
		usedMem: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "used_memory_bytes",
			Help: "Used memory in bytes",
		}),
		availableMem: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "available_memory_bytes",
			Help: "Available memory in bytes",
		}),
		client: *clients.NewOSClient(cpuIteration, memIteration),
	}
}

func EmitOSMetrics() {
	m := NewMetrics(10, 10)
	prometheus.MustRegister(m.cpuLoad1m, m.cpuLoad5m, m.cpuLoad15m, m.totalMem, m.usedMem, m.availableMem)
	go m.sendCPUMetrics()
	go m.sendMemMetrics()
}

func (m *metrics) sendCPUMetrics() {
	for {
		v, err := m.client.GetCPU()
		if err != nil {
			log.Fatal(err)
			continue
		}

		m.cpuLoad1m.Set(v.Load1 * 100)
		m.cpuLoad5m.Set(v.Load5 * 100)
		m.cpuLoad15m.Set(v.Load15 * 100)

		time.Sleep(time.Duration(m.client.CpuIteration) * time.Millisecond)
	}
}

func (m *metrics) sendMemMetrics() {

	for {
		v, err := m.client.GetMem()

		if err != nil {
			log.Fatal(err)
			continue
		}

		m.availableMem.Set(float64(v.Available))
		m.usedMem.Set(float64(v.Used))
		m.totalMem.Set(float64(v.Total))
		time.Sleep(time.Duration(m.client.MemIteration) * time.Millisecond)
	}
}
