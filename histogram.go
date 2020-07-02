package prometheus

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

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
func (o *Object) Histogram(metricName string, value float64, labels Labels, units ...float64) (err error) {
	labels = o.addServiceInfoToLabels(labels)
	defer func() {
		if r := recover(); r != nil {
			err = o.errorHandler(r, metricName, getLabelNames(labels))
		}
	}()
	if o.histograms[metricName] == nil {
		o.histograms[metricName] = promauto.With(o.reg).NewHistogramVec(prometheus.HistogramOpts{
			Name:    metricName,
			Help:    "Histogram created for " + metricName,
			Buckets: makeLinearBuckets(units),
		}, getLabelNames(labels))
	}
	o.histograms[metricName].With(prometheus.Labels(labels)).Observe(value)
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

// GenerateUnits creates a float64 slice to measure request durations.
func GenerateUnits(start, width float64, count int) []float64 {
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
		return GenerateUnits(0.5, 0.5, 20)
	}
	return buckets
}

// MeasureExecTime is used by
// Start and StopMeasureExecTime functions.
type MeasureExecTime struct {
	MetricName   string
	Labels       Labels
	Units        []float64
	TimeDuration time.Duration

	object *Object
	start  time.Time
}

// StartMeasureExecTime launches the measurement of the runtime
// of a particular calculation.
//
// Use the TimeDuration field to set the unit of the elapsed time measurement:
// Minute, Second, Millisecond, Microsecond, Nanosecond.
func (o *Object) StartMeasureExecTime(m MeasureExecTime) *MeasureExecTime {
	m.object = o
	m.start = time.Now()
	return &m
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
		executionTime = float64(time.Since(m.start).Milliseconds())
	case time.Microsecond:
		executionTime = float64(time.Since(m.start).Microseconds())
	case time.Nanosecond:
		executionTime = float64(time.Since(m.start).Nanoseconds())
	default:
		executionTime = float64(time.Since(m.start).Milliseconds())
	}
	return m.object.Histogram(m.MetricName, executionTime, m.Labels, m.Units...)
}
