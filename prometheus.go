// This package is a Prometheus implementation for metrics.
// In its creation, the main consideration was the ease of use.
//
// Provides the following metric types:
//   - Counter,
//   - Gauge,
//   - Histogram.
//
// It can also provide built-in statistics (optionally)
// about the system and the application:
//   - Goroutines,
//   - Memory usage,
//   - CPU usage.
package prometheus

import (
	"context"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

// Init is used for prometheus Object initialization.
type Init struct {
	Host        string
	Port        int
	Environment string
	AppName     string

	MetricEndpoint      string
	StatCountGoroutines bool
	StatMemoryUsage     bool
	StatCpuUsage        bool
}

// Object provides structure to use metric types.
type Object struct {
	Addr string
	Env  string
	App  string

	MetricsEndpoint     string
	StatCountGoroutines bool
	StatMemoryUsage     bool
	StatCpuUsage        bool

	counters   map[string]*prometheus.CounterVec
	gauges     map[string]*prometheus.GaugeVec
	histograms map[string]*prometheus.HistogramVec
	server     *http.Server
	reg        *prometheus.Registry
}

// Label used by metric types: Counter, Histogram, Gauge
type Labels prometheus.Labels

func (i Init) setMetricsEndpoint() string {
	if i.MetricEndpoint == "" {
		return "/metrics"
	}
	return i.MetricEndpoint
}

// New creates a new Object structure.
func New(i Init) *Object {
	o := &Object{
		Addr: fmt.Sprintf("%s:%d", i.Host, i.Port),
		Env:  i.Environment,
		App:  i.AppName,

		MetricsEndpoint:     i.setMetricsEndpoint(),
		StatCountGoroutines: i.StatCountGoroutines,
		StatMemoryUsage:     i.StatMemoryUsage,
		StatCpuUsage:        i.StatCpuUsage,

		counters:   make(map[string]*prometheus.CounterVec),
		gauges:     make(map[string]*prometheus.GaugeVec),
		histograms: make(map[string]*prometheus.HistogramVec),
		reg:        prometheus.NewRegistry(),
	}

	o.StartHttpServer()
	o.statCountGoroutines()
	o.statMemoryUsage()
	o.statCpuUsage()

	return o
}

// StartHttpServer starts providing metrics data
// on given host and port on the endpoint
// which set by (*Object).MetricsEndpoint.
func (o *Object) StartHttpServer() {
	o.server = o.serve()
}

// StopHttpServer ends providing metrics data.
func (o *Object) StopHttpServer() {
	_ = o.server.Shutdown(context.Background())
}

// GetMetrics queries all metrics data which matches to 'text' string.
func (o *Object) GetMetrics(text string) string {
	return Grep(text, o.getMetrics())
}
