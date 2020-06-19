package prometheus

import (
	"fmt"
	"runtime"
	"time"

	kitProm "github.com/go-kit/kit/metrics/prometheus"
	clientGo "github.com/prometheus/client_golang/prometheus"
)

// Gauge is a metric that represents a single numerical value
// that can arbitrarily go up and down.
//
// Gauges are typically used for measured values like temperatures
// or current memory usage, but also "counts" that can go up and down,
// like the number of concurrent requests.
func (o *Object) Gauge(metricName string, labels []Label, value float64) (err error) {
	labels = o.addServiceInfoToLabels(labels)
	labelNames := getLabelNames(labels)

	defer func() {
		if r := recover(); r != nil {
			err = o.errorHandler(r, metricName, labelNames)
		}
	}()

	if o.gauges[metricName] == nil {
		o.gauges[metricName] = kitProm.NewGaugeFrom(clientGo.GaugeOpts{
			Name:        metricName,
			Help:        fmt.Sprintf("Gauge for: %s", metricName),
			ConstLabels: clientGo.Labels{},
		}, labelNames)

		o.gauges[metricName].With(makeSlice(labels)...).Add(value)
		return
	}

	o.gauges[metricName].With(makeSlice(labels)...).Set(value)
	return
}

func (o *Object) statCountGoroutines() {
	t := time.NewTicker(time.Second)
	go func() {
		for range t.C {
			_ = o.Gauge("stat_count_goroutines", []Label{}, float64(runtime.NumGoroutine()))
		}
	}()
}
