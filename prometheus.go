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

	counters   map[string]kitMet.Counter
	histograms map[string]kitMet.Histogram
	gauges     map[string]kitMet.Gauge
	server     *http.Server
}

// New creates a new Object structure.
func New(i Init) *Object {
	addr := fmt.Sprintf("%s:%d", i.Host, i.Port)
	o := &Object{
		Addr: addr,
		Env:  i.Environment,
		App:  i.AppName,

		counters:   make(map[string]kitMet.Counter),
		histograms: make(map[string]kitMet.Histogram),
		gauges:     make(map[string]kitMet.Gauge),
	}
	o.StartHttpServer()
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
