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

	p.Histogram("example", []Label{
		{
			Name:  "foo1",
			Value: "bar1",
		},
		{
			Name:  "foo2",
			Value: "bar2",
		},
	}, time.Since(time.Now()).Seconds(), GenerateUnits(0.5, 0.5, 20)...)

	for _, unit := range GenerateUnits(0.5, 0.5, 20) {
		expected := `TestHistogram_test_example_histogram_bucket{foo1="bar1",foo2="bar2",le="` + fmt.Sprintf("%g", unit) + `"} 1`
		actual := grep(p.App, p.getMetrics())

		s.Equal(true, strings.Contains(actual, expected))
	}

	expectedSum := `TestHistogram_test_example_histogram_sum{foo1="bar1",foo2="bar2"}`
	expectedCount := `TestHistogram_test_example_histogram_count{foo1="bar1",foo2="bar2"} 1`

	actual := grep(p.App, p.getMetrics())

	s.Equal(true, strings.Contains(actual, expectedSum))
	s.Equal(true, strings.Contains(actual, expectedCount))

	p.StopHttpServer()
}

func (s histogramSuite) TestElapsedTime() {
	p := New(initProm("TestElapsedTime"))

	p.ElapsedTime("example", []Label{
		{
			Name:  "foo1",
			Value: "bar1",
		},
	}, time.Now())

	expectedSum := `TestElapsedTime_test_example_histogram_sum{foo1="bar1"}`
	expectedCount := `TestElapsedTime_test_example_histogram_count{foo1="bar1"} 1`

	actual := grep(p.App, p.getMetrics())

	s.Equal(true, strings.Contains(actual, expectedSum))
	s.Equal(true, strings.Contains(actual, expectedCount))

	p.StopHttpServer()
}

func TestHistogramSuite(t *testing.T) {
	suite.Run(t, new(histogramSuite))
}
