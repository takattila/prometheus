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
