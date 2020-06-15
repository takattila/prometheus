package prometheus

import (
	"fmt"
	"net/http"

	kitMet "github.com/go-kit/kit/metrics"
	kitProm "github.com/go-kit/kit/metrics/prometheus"
	clientGo "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Init struct {
	Host        string
	Port        int
	Environment string
	AppName     string
}

type Label struct {
	Name  string
	Value string
}

type Object struct {
	Addr string
	Env  string
	App  string

	counters   map[string]kitMet.Counter
	histograms map[string]kitMet.Histogram
	gauges     map[string]kitMet.Gauge
}

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

func (o *Object) serve() {
	_ = http.ListenAndServe(o.Addr, promhttp.Handler())
}

func getLabelNames(labels []Label) []string {
	var slice []string
	for _, o := range labels {
		slice = append(slice, o.Name)
	}
	return slice
}

func makeSlice(labels []Label) []string {
	var slice []string
	for _, o := range labels {
		slice = append(slice, o.Name, o.Value)
	}
	return slice
}

func (o *Object) Counter(metricName string, labels []Label, delta float64) {
	if o.counters[metricName] == nil {
		o.counters[metricName] = kitProm.NewCounterFrom(clientGo.CounterOpts{
			Namespace:   o.App,
			Subsystem:   o.Env,
			Name:        metricName + "_counter",
			Help:        fmt.Sprintf("Counter for: %s %+v", metricName, labels),
			ConstLabels: clientGo.Labels{},
		}, getLabelNames(labels))
	}
	o.counters[metricName].With(makeSlice(labels)...).Add(delta)
}
