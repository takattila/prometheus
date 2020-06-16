package prometheus

import (
	"fmt"

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
}

// New creates a new Object structure.
func New(i Init) *Object {
	o := &Object{
		Addr: fmt.Sprintf("%s:%d", i.Host, i.Port),
		Env:  i.Environment,
		App:  i.AppName,

		counters:   make(map[string]kitMet.Counter),
		histograms: make(map[string]kitMet.Histogram),
		gauges:     make(map[string]kitMet.Gauge),
	}
	go o.serve()
	return o
}
