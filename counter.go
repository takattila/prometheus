package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Counter is a cumulative metric that represents a single monotonically increasing counter
// whose value can only increase or be reset to zero on restart.
// For example, you can use a counter to represent the number
// of requests served, tasks completed, or errors.
func (o *Object) Counter(metricName string, value float64, labels Labels) (err error) {
	labels = o.addServiceInfoToLabels(labels)
	defer func() {
		if r := recover(); r != nil {
			err = o.errorHandler(r, metricName, getLabelNames(labels))
		}
	}()
	if o.counters[metricName] == nil {
		o.counters[metricName] = promauto.With(o.reg).NewCounterVec(prometheus.CounterOpts{
			Name: metricName,
			Help: "Counter created for " + metricName,
		}, getLabelNames(labels))

	}
	o.counters[metricName].With(prometheus.Labels(labels)).Add(value)
	return
}
