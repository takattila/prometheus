package prometheus

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type histogramSuite struct {
	suite.Suite
}

func (s histogramSuite) TestHistogram() {
	p := New(initProm("TestHistogram"))

	begin := time.Since(time.Now()).Seconds()
	err := p.Histogram(HistogramArgs{
		MetricName: "example_histogram",
		Labels:     Labels{"foo1": "bar1", "foo2": "bar2"},
		Buckets:    GenerateBuckets(0.5, 0.5, 20),
		Value:      begin,
	})
	s.Equal(nil, err)

	for _, bucket := range GenerateBuckets(0.5, 0.5, 20) {
		expected := `example_histogram_bucket{app="TestHistogram",env="test",foo1="bar1",foo2="bar2",le="` + fmt.Sprintf("%g", bucket) + `"} 1`
		actual := p.GetMetrics(p.App)

		s.Contains(actual, expected)
	}

	expectedSum := `example_histogram_sum{app="TestHistogram",env="test",foo1="bar1",foo2="bar2"}`
	expectedCount := `example_histogram_count{app="TestHistogram",env="test",foo1="bar1",foo2="bar2"} 1`

	actual := p.GetMetrics(p.App)

	s.Contains(actual, expectedSum)
	s.Contains(actual, expectedCount)

	p.StopHttpServer()
}

func (s histogramSuite) TestHistogramError() {
	p := New(initProm("TestHistogramError"))

	begin := time.Since(time.Now()).Seconds()
	err := p.Histogram(HistogramArgs{
		MetricName: "example_histogram_error",
		Labels:     Labels{"foo1": "bar1", "foo2": "bar2"},
		Value:      begin,
	})
	s.Equal(nil, err)

	// Missing label name
	begin = time.Since(time.Now()).Seconds()
	err = p.Histogram(HistogramArgs{
		MetricName: "example_histogram_error",
		Labels:     Labels{"foo1": "bar1"},
		Value:      begin,
	})

	expected := `metric: 'example_histogram_error', error: 'inconsistent label cardinality: expected 4 label values but got 3 in prometheus.Labels{"app":"TestHistogramError", "env":"test", "foo1":"bar1"}', input label names: 'app, env, foo1', correct label names: 'app, env, foo1, foo2'` + "\n"
	s.Equal(expected, fmt.Sprint(err))

	p.StopHttpServer()
}

func (s histogramSuite) TestGenerateBuckets() {
	expected := []float64{0.5, 1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5, 9.5}
	actual := GenerateBuckets(0.5, 1, 10)
	s.Equal(expected, actual)
}

func (s histogramSuite) TestStartStopMeasureExecTime() {
	p := New(initProm("TestStartStopMeasureExecTime"))

	for _, t := range []struct {
		Sleep        time.Duration
		MetricName   string
		Buckets      []float64
		TimeDuration time.Duration
		Expected     string
	}{
		{
			Sleep:        5000 * time.Nanosecond,
			MetricName:   "execution_time_nano_sec",
			Buckets:      GenerateBuckets(50000, 100000, 10),
			TimeDuration: time.Nanosecond,
			Expected:     `execution_time_nano_sec_bucket{app="TestStartStopMeasureExecTime",env="test",foo1="bar1",le="950000"} 1`,
		},
		{
			Sleep:        100 * time.Microsecond,
			MetricName:   "execution_time_micro_sec",
			Buckets:      GenerateBuckets(50, 50, 10),
			TimeDuration: time.Microsecond,
			Expected:     `execution_time_micro_sec_bucket{app="TestStartStopMeasureExecTime",env="test",foo1="bar1",le="500"} 1`,
		},
		{
			Sleep:        10 * time.Millisecond,
			MetricName:   "execution_time_milli_sec",
			Buckets:      GenerateBuckets(5, 5, 10),
			TimeDuration: time.Millisecond,
			Expected:     `execution_time_milli_sec_bucket{app="TestStartStopMeasureExecTime",env="test",foo1="bar1",le="50"} 1`,
		},
		{
			Sleep:        1 * time.Second,
			MetricName:   "execution_time_seconds",
			Buckets:      GenerateBuckets(0.5, 0.5, 10),
			TimeDuration: time.Second,
			Expected:     `execution_time_seconds_bucket{app="TestStartStopMeasureExecTime",env="test",foo1="bar1",le="5"} 1`,
		},
		{
			Sleep:        1 * time.Second,
			MetricName:   "execution_time_minutes",
			Buckets:      GenerateBuckets(0.005, 0.005, 10),
			TimeDuration: time.Minute,
			Expected:     `execution_time_minutes_bucket{app="TestStartStopMeasureExecTime",env="test",foo1="bar1",le="0.05"} 1`,
		},
		{
			// TimeDuration: not given,
			MetricName: "execution_time_default_milli_sec",
			Buckets:    GenerateBuckets(5, 5, 10),
			Sleep:      10 * time.Millisecond,
			Expected:   `execution_time_default_milli_sec_bucket{app="TestStartStopMeasureExecTime",env="test",foo1="bar1",le="50"} 1`,
		},
	} {
		// Measure start
		met := p.StartMeasureExecTime(MeasureExecTimeArgs{
			MetricName:   t.MetricName,
			Buckets:      t.Buckets,
			TimeDuration: t.TimeDuration,
			Labels:       Labels{"foo1": "bar1"},
		})

		time.Sleep(t.Sleep)

		err := met.StopMeasureExecTime()
		// Measure end

		s.Equal(nil, err)

		actual := p.GetMetrics(t.MetricName)
		s.Contains(actual, t.Expected)
	}

	p.StopHttpServer()
}

func TestHistogramSuite(t *testing.T) {
	suite.Run(t, new(histogramSuite))
}
