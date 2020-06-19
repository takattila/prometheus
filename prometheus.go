// The prometheus package provides Prometheus implementations for metrics.
package prometheus

import (
	"context"
	"fmt"
	"net/http"

	kitMet "github.com/go-kit/kit/metrics"
)

// Init is used for prometheus Object initialization.
type Init struct {
	Host        string
	Port        int
	Environment string
	AppName     string

	StatCountGoroutines bool
	StatMemoryUsage     bool
}

// Label used by metric types: Counter, Histogram, Gauge
type Label struct {
	Name  string
	Value string
}

// Object provides structure to use metric types.
type Object struct {
	Addr string
	Env  string
	App  string

	StatCountGoroutines bool
	StatMemoryUsage     bool

	counters   map[string]kitMet.Counter
	histograms map[string]kitMet.Histogram
	gauges     map[string]kitMet.Gauge
	server     *http.Server
}

// New creates a new Object structure.
func New(i Init) *Object {
	o := &Object{
		Addr: fmt.Sprintf("%s:%d", i.Host, i.Port),
		Env:  i.Environment,
		App:  i.AppName,

		StatCountGoroutines: i.StatCountGoroutines,
		StatMemoryUsage:     i.StatMemoryUsage,

		counters:   make(map[string]kitMet.Counter),
		histograms: make(map[string]kitMet.Histogram),
		gauges:     make(map[string]kitMet.Gauge),
	}

	o.StartHttpServer()
	o.statCountGoroutines()
	o.statMemoryUsage()

	return o
}

// StartHttpServer starts providing metrics data
// on given host and port on all route.
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
