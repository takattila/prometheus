package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// GenerateUnits creates a float64 slice to measure request durations.
func GenerateUnits(start, width float64, count int) []float64 {
	return prometheus.LinearBuckets(start, width, count)
}

func makeLinearBuckets(buckets []float64) []float64 {
	if len(buckets) == 0 {
		return GenerateUnits(0.5, 0.5, 20)
	}
	return buckets
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
