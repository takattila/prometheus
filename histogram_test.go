package prometheus

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type histogramSuite struct {
	suite.Suite
}

func (s histogramSuite) TestHistogram() {
	p := New(initProm("TestHistogram"))

	err := p.Histogram("example", []Label{
		{
			Name:  "foo1",
			Value: "bar1",
		},
		{
			Name:  "foo2",
			Value: "bar2",
		},
	}, time.Since(time.Now()).Seconds(), GenerateUnits(0.5, 0.5, 20)...)
	s.Equal(nil, err)

	for _, unit := range GenerateUnits(0.5, 0.5, 20) {
		expected := `TestHistogram_test_example_histogram_bucket{foo1="bar1",foo2="bar2",le="` + fmt.Sprintf("%g", unit) + `"} 1`
		actual := p.GetMetrics(p.App)

		s.Equal(true, strings.Contains(actual, expected))
	}

	expectedSum := `TestHistogram_test_example_histogram_sum{foo1="bar1",foo2="bar2"}`
	expectedCount := `TestHistogram_test_example_histogram_count{foo1="bar1",foo2="bar2"} 1`

	actual := p.GetMetrics(p.App)

	s.Equal(true, strings.Contains(actual, expectedSum))
	s.Equal(true, strings.Contains(actual, expectedCount))

	p.StopHttpServer()
}

func (s histogramSuite) TestElapsedTime() {
	p := New(initProm("TestElapsedTime"))

	defer func(begin time.Time) {

		buckets := GenerateUnits(0.02, 0.02, 10)

		err := p.ElapsedTime("example", []Label{
			{
				Name:  "foo1",
				Value: "bar1",
			},
		}, begin, buckets...)

		s.Equal(nil, err)

		output := p.getMetrics()

		for _, unit := range buckets {
			expected := `TestElapsedTime_test_example_histogram_bucket{foo1="bar1",le="` + fmt.Sprintf("%g", unit) + `"}`
			actual := Grep(p.App, output)

			s.Equal(true, strings.Contains(actual, expected))
		}

		expectedSum := `TestElapsedTime_test_example_histogram_sum{foo1="bar1"}`
		expectedCount := `TestElapsedTime_test_example_histogram_count{foo1="bar1"} 1`

		actual := Grep(p.App, output)

		s.Equal(true, strings.Contains(actual, expectedSum))
		s.Equal(true, strings.Contains(actual, expectedCount))

		p.StopHttpServer()

	}(time.Now())

	time.Sleep(100 * time.Millisecond)
}

func (s histogramSuite) TestHistogramError() {
	p := New(initProm("TestHistogramError"))

	err := p.Histogram("example", []Label{
		{
			Name:  "foo1",
			Value: "bar1",
		},
		{
			Name:  "foo2",
			Value: "bar2",
		},
	}, time.Since(time.Now()).Seconds())
	s.Equal(nil, err)

	// Missing label name
	err = p.Histogram("example", []Label{
		{
			Name:  "foo1",
			Value: "bar1",
		},
	}, time.Since(time.Now()).Seconds())

	expected := `metric: 'TestHistogramError_test_example_histogram', error: 'inconsistent label cardinality: expected 2 label values but got 1 in prometheus.Labels{"foo1":"bar1"}', input label names: 'foo1', correct label names: 'foo1, foo2'` + "\n"
	s.Equal(expected, fmt.Sprint(err))

	p.StopHttpServer()
}

func TestHistogramSuite(t *testing.T) {
	suite.Run(t, new(histogramSuite))
}
