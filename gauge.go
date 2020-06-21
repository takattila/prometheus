package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Gauge is a metric that represents a single numerical value
// that can arbitrarily go up and down.
//
// Gauges are typically used for measured values like temperatures
// or current memory usage, but also "counts" that can go up and down,
// like the number of concurrent requests.
func (o *Object) Gauge(metricName string, value float64, labels Labels) (err error) {
	labels = o.addServiceInfoToLabels(labels)
	defer func() {
		if r := recover(); r != nil {
			err = o.errorHandler(r, metricName, getLabelNames(labels))
		}
	}()
	if o.gauges[metricName] == nil {
		o.gauges[metricName] = promauto.With(o.reg).NewGaugeVec(prometheus.GaugeOpts{
			Name: metricName,
			Help: "Gauge created for " + metricName,
		}, getLabelNames(labels))
		o.gauges[metricName].With(prometheus.Labels(labels)).Add(value)
		return
	}
	o.gauges[metricName].With(prometheus.Labels(labels)).Set(value)
	return
}
