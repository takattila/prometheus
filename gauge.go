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
func (o *Object) Gauge(metricName string, labels []Label, value float64) (err error) {
	fqdn := makeFQDN(o.App, o.Env, metricName, "gauge")
	labelNames := getLabelNames(labels)

	defer func() {
		if r := recover(); r != nil {
			err = o.errorHandler(r, fqdn, labelNames)
		}
	}()

	if o.gauges[fqdn] == nil {
		o.gauges[fqdn] = kitProm.NewGaugeFrom(clientGo.GaugeOpts{
			Name:        fqdn,
			Help:        fmt.Sprintf("Gauge for: %s", metricName),
			ConstLabels: clientGo.Labels{},
		}, labelNames)

		o.gauges[fqdn].With(makeSlice(labels)...).Add(value)
		return
	}

	o.gauges[fqdn].With(makeSlice(labels)...).Set(value)
	return
}
