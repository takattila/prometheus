package prometheus

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var defaultBuckets = []float64{
	0.000025, 0.00005, 0.0025, 0.005, 0.025, 0.05,
	0.1, 0.25, 0.5, 1, 2.5, 5, 10, 20, 25, 50, 60, 120,
}

// HistogramArgs contains the necessary arguments
// of the *Object.Histogram function.
type HistogramArgs struct {
	MetricName string
	Labels     Labels
	Buckets    []float64
	Value      float64
}

// Histogram samples observations (usually things like request durations
// or response sizes) and counts them in configurable buckets.
// It also provides a sum of all observed values.
//
// A histogram with a base metric name of <basename>
// exposes multiple time series during a scrape:
//
//   - cumulative counters for the observation buckets, exposed
//     as <basename>_bucket{le="<upper inclusive bound>"}
//   - the total sum of all observed values, exposed as <basename>_sum
//   - the count of events that have been observed, exposed
//     as <basename>_count (identical to <basename>_bucket{le="+Inf"} above)
//
// Use the histogram_quantile() function to calculate quantiles
// from histograms or even aggregations of histograms.
//
// A histogram is also suitable to calculate an Apdex score.
// When operating on buckets, remember that the histogram is cumulative.
func (o *Object) Histogram(args HistogramArgs) (err error) {
	args.Labels = o.addServiceInfoToLabels(args.Labels)
	defer func() {
		if r := recover(); r != nil {
			err = o.errorHandler(r, args.MetricName, getLabelNames(args.Labels))
		}
	}()
	if o.histograms[args.MetricName] == nil {
		o.histograms[args.MetricName] = promauto.With(o.reg).NewHistogramVec(prometheus.HistogramOpts{
			Name:    args.MetricName,
			Help:    "Histogram created for " + args.MetricName,
			Buckets: makeLinearBuckets(args.Buckets),
		}, getLabelNames(args.Labels))
	}
	o.histograms[args.MetricName].With(prometheus.Labels(args.Labels)).Observe(args.Value)
	return
}

func calculateDecimalPlaces(start, width float64) (dp int) {
	s := DecimalPlaces(start)
	w := DecimalPlaces(width)
	dp = w
	if s > w {
		dp = s
	}
	return
}

// GenerateBuckets creates a float64 slice to measure request durations.
func GenerateBuckets(start, width float64, count int) []float64 {
	buckets := prometheus.LinearBuckets(start, width, count)
	dp := calculateDecimalPlaces(start, width)
	for i := range buckets {
		buckets[i] = RoundFloat(start, dp)
		start += width
	}
	return buckets
}

func makeLinearBuckets(buckets []float64) []float64 {
	if len(buckets) == 0 {
		return defaultBuckets
	}
	return buckets
}

// MeasureExecTimeArgs is used by
// Start and StopMeasureExecTime functions.
type MeasureExecTime struct {
	MetricName   string
	Labels       Labels
	Buckets      []float64
	TimeDuration time.Duration

	object *Object
	start  time.Time
}

// MeasureExecTimeArgs contains the necessary arguments
// of the *Object.StartMeasureExecTime function.
type MeasureExecTimeArgs MeasureExecTime

// StartMeasureExecTime launches the measurement of the runtime
// of a particular calculation.
//
// Use the TimeDuration field to set the unit of the elapsed time measurement:
// Minute, Second, Millisecond, Microsecond, Nanosecond.
func (o *Object) StartMeasureExecTime(m MeasureExecTimeArgs) *MeasureExecTime {
	return &MeasureExecTime{
		MetricName:   m.MetricName,
		Labels:       m.Labels,
		Buckets:      m.Buckets,
		TimeDuration: m.TimeDuration,

		object: o,
		start:  time.Now(),
	}
}

// StopMeasureExecTime ends the execution time measurement.
func (m *MeasureExecTime) StopMeasureExecTime() error {
	var executionTime float64
	switch m.TimeDuration {
	case time.Minute:
		executionTime = time.Since(m.start).Minutes()
	case time.Second:
		executionTime = time.Since(m.start).Seconds()
	case time.Millisecond:
		executionTime = float64(time.Since(m.start).Nanoseconds() / 1e6)
	case time.Microsecond:
		executionTime = float64(time.Since(m.start).Nanoseconds() / 1e3)
	case time.Nanosecond:
		executionTime = float64(time.Since(m.start).Nanoseconds())
	default:
		executionTime = float64(time.Since(m.start).Nanoseconds() / 1e6)
	}
	return m.object.Histogram(HistogramArgs{
		MetricName: m.MetricName,
		Labels:     m.Labels,
		Buckets:    m.Buckets,
		Value:      executionTime,
	})
}
