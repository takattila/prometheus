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
	err := p.Histogram("example_histogram", begin, Labels{
		"foo1": "bar1",
		"foo2": "bar2",
	}, GenerateUnits(0.5, 0.5, 20)...)
	s.Equal(nil, err)

	for _, unit := range GenerateUnits(0.5, 0.5, 20) {
		expected := `example_histogram_bucket{app="TestHistogram",env="test",foo1="bar1",foo2="bar2",le="` + fmt.Sprintf("%g", unit) + `"} 1`
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
	err := p.Histogram("example_histogram_error", begin, Labels{
		"foo1": "bar1",
		"foo2": "bar2",
	})
	s.Equal(nil, err)

	// Missing label name
	begin = time.Since(time.Now()).Seconds()
	err = p.Histogram("example_histogram_error", begin, Labels{
		"foo1": "bar1",
	})

	expected := `metric: 'example_histogram_error', error: 'inconsistent label cardinality: expected 4 label values but got 3 in prometheus.Labels{"app":"TestHistogramError", "env":"test", "foo1":"bar1"}', input label names: 'app, env, foo1', correct label names: 'app, env, foo1, foo2'` + "\n"
	s.Equal(expected, fmt.Sprint(err))

	p.StopHttpServer()
}

func TestHistogramSuite(t *testing.T) {
	suite.Run(t, new(histogramSuite))
}
