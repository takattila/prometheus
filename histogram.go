package prometheus

import (
	"fmt"
	"time"

	kitProm "github.com/go-kit/kit/metrics/prometheus"
	clientGo "github.com/prometheus/client_golang/prometheus"
)

// GenerateUnits creates a float64 slice to measure request durations.
func GenerateUnits(start, width float64, count int) []float64 {
	return clientGo.LinearBuckets(start, width, count)
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
func (o *Object) Histogram(metricName string, labels []Label, since float64, units ...float64) (err error) {
	labels = o.addServiceInfoToLabels(labels)
	labelNames := getLabelNames(labels)

	defer func() {
		if r := recover(); r != nil {
			err = o.errorHandler(r, metricName, labelNames)
		}
	}()

	if o.histograms[metricName] == nil {
		o.histograms[metricName] = kitProm.NewHistogramFrom(clientGo.HistogramOpts{
			Name:        metricName,
			Help:        fmt.Sprintf("Histogram for: %s", metricName),
			Buckets:     makeLinearBuckets(units),
			ConstLabels: clientGo.Labels{},
		}, labelNames)
	}

	o.histograms[metricName].With(makeSlice(labels)...).Observe(since)
	return
}

// ElapsedTime is a histogram for request duration,
// exported via a Prometheus summary with dynamically-computed quantiles.
func (o *Object) ElapsedTime(metricName string, labels []Label, since time.Time, units ...float64) (err error) {
	func(begin time.Time) {
		err = o.Histogram(metricName, labels, time.Since(begin).Seconds(), units...)
	}(since)
	return
}
