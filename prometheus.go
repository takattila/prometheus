package prometheus

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	kitMet "github.com/go-kit/kit/metrics"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
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

// GetMetrics queries all metrics data which matches to 'text' string.
func (o *Object) GetMetrics(text string) string {
	return grep(text, o.getMetrics())
}

// ParseOutput reads 'text' as the simple and flat text-based
// exchange format and creates MetricFamily proto messages.
func ParseOutput(text string) map[string]*dto.MetricFamily {
	p := expfmt.TextParser{}
	m, _ := p.TextToMetricFamilies(strings.NewReader(text))
	return m
}

// GetLabels gives back the labels for a metric.
func GetLabels(text, metric string) (labels []Label) {
	p := ParseOutput(text)
	obj := p[metric]

	for _, m := range obj.GetMetric() {
		for _, l := range m.GetLabel() {
			labels = append(labels, Label{
				Name:  *l.Name,
				Value: *l.Value,
			})
		}
	}
	return
}
