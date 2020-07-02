package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// GaugeArgs contains the necessary arguments
// of the *Object.Gauge function.
type GaugeArgs struct {
	MetricName string
	Labels     Labels
	Value      float64
}

// Gauge is a metric that represents a single numerical value
// that can arbitrarily go up and down.
//
// Gauges are typically used for measured values like temperatures
// or current memory usage, but also "counts" that can go up and down,
// like the number of concurrent requests.
func (o *Object) Gauge(args GaugeArgs) (err error) {
	args.Labels = o.addServiceInfoToLabels(args.Labels)
	defer func() {
		if r := recover(); r != nil {
			err = o.errorHandler(r, args.MetricName, getLabelNames(args.Labels))
		}
	}()
	if o.gauges[args.MetricName] == nil {
		o.gauges[args.MetricName] = promauto.With(o.reg).NewGaugeVec(prometheus.GaugeOpts{
			Name: args.MetricName,
			Help: "Gauge created for " + args.MetricName,
		}, getLabelNames(args.Labels))
		o.gauges[args.MetricName].With(prometheus.Labels(args.Labels)).Add(args.Value)
		return
	}
	o.gauges[args.MetricName].With(prometheus.Labels(args.Labels)).Set(args.Value)
	return
}
