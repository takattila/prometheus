package prometheus

import (
	"fmt"

	kitProm "github.com/go-kit/kit/metrics/prometheus"
	clientGo "github.com/prometheus/client_golang/prometheus"
)

// Gauge is a metric that represents a single numerical value
// that can arbitrarily go up and down.
//
// Gauges are typically used for measured values like temperatures
// or current memory usage, but also "counts" that can go up and down,
// like the number of concurrent requests.
func (o *Object) Gauge(metricName string, labels []Label, value float64) {
	if o.gauges[metricName] == nil {
		o.gauges[metricName] = kitProm.NewGaugeFrom(clientGo.GaugeOpts{
			Namespace:   o.App,
			Subsystem:   o.Env,
			Name:        metricName + "_gauge",
			Help:        fmt.Sprintf("Gauge for: %s %+v", metricName, labels),
			ConstLabels: clientGo.Labels{},
		}, getLabelNames(labels))
		o.gauges[metricName].With(makeSlice(labels)...).Add(value)
		return
	}
	o.gauges[metricName].With(makeSlice(labels)...).Set(value)
}
