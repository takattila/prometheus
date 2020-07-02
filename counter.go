package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// CounterArgs contains the necessary arguments
// of the *Object.Counter function.
type CounterArgs struct {
	MetricName string
	Labels     Labels
	Value      float64
}

// Counter is a cumulative metric that represents a single monotonically increasing counter
// whose value can only increase or be reset to zero on restart.
// For example, you can use a counter to represent the number
// of requests served, tasks completed, or errors.
func (o *Object) Counter(args CounterArgs) (err error) {
	args.Labels = o.addServiceInfoToLabels(args.Labels)
	defer func() {
		if r := recover(); r != nil {
			err = o.errorHandler(r, args.MetricName, getLabelNames(args.Labels))
		}
	}()
	if o.counters[args.MetricName] == nil {
		o.counters[args.MetricName] = promauto.With(o.reg).NewCounterVec(prometheus.CounterOpts{
			Name: args.MetricName,
			Help: "Counter created for " + args.MetricName,
		}, getLabelNames(args.Labels))

	}
	o.counters[args.MetricName].With(prometheus.Labels(args.Labels)).Add(args.Value)
	return
}
