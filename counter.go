package prometheus

import (
	"fmt"

	kitProm "github.com/go-kit/kit/metrics/prometheus"
	clientGo "github.com/prometheus/client_golang/prometheus"
)

// Counter is a cumulative metric that represents a single monotonically increasing counter
// whose value can only increase or be reset to zero on restart.
// For example, you can use a counter to represent the number
// of requests served, tasks completed, or errors.
func (o *Object) Counter(metricName string, labels []Label, delta float64) (err error) {
	labels = o.addServiceInfoToLabels(labels)
	labelNames := getLabelNames(labels)

	defer func() {
		if r := recover(); r != nil {
			err = o.errorHandler(r, metricName, labelNames)
		}
	}()

	if o.counters[metricName] == nil {
		o.counters[metricName] = kitProm.NewCounterFrom(clientGo.CounterOpts{
			Name:        metricName,
			Help:        fmt.Sprintf("Counter for: %s", metricName),
			ConstLabels: clientGo.Labels{},
		}, labelNames)
	}

	o.counters[metricName].With(makeSlice(labels)...).Add(delta)
	return
}
