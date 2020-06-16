package prometheus

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	clientGo "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/suite"
)

type histogramSuite struct {
	suite.Suite
}

func (s histogramSuite) TestHistogram() {
	p := New(Init{
		Host:        "0.0.0.0",
		Port:        8080,
		Environment: "test",
		AppName:     "TestHistogram",
	})

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

	ts := httptest.NewServer(promhttp.HandlerFor(clientGo.DefaultGatherer, promhttp.HandlerOpts{}))
	defer ts.Close()

	responseBody := func() string {
		resp, _ := http.Get(ts.URL)
		buf, _ := ioutil.ReadAll(resp.Body)
		return string(buf)
	}

	for _, unit := range GenerateUnits(0.5, 0.5, 20) {
		expected := `TestHistogram_test_example_histogram_bucket{foo1="bar1",foo2="bar2",le="` + fmt.Sprintf("%g", unit) + `"} 1`
		actual := grep(p.App, responseBody())

		s.Equal(true, strings.Contains(actual, expected))
	}

	expectedSum := `TestHistogram_test_example_histogram_sum{foo1="bar1",foo2="bar2"}`
	expectedCount := `TestHistogram_test_example_histogram_count{foo1="bar1",foo2="bar2"} 1`

	actual := grep(p.App, responseBody())

	s.Equal(true, strings.Contains(actual, expectedSum))
	s.Equal(true, strings.Contains(actual, expectedCount))

	if err := p.server.Shutdown(context.Background()); err != nil {
		s.T().Fatal(err)
	}
}

func (s histogramSuite) TestElapsedTime() {
	p := New(Init{
		Host:        "0.0.0.0",
		Port:        8080,
		Environment: "test",
		AppName:     "TestElapsedTime",
	})

	p.ElapsedTime("example", []Label{
		{
			Name:  "foo1",
			Value: "bar1",
		},
	}, time.Now())

	ts := httptest.NewServer(promhttp.HandlerFor(clientGo.DefaultGatherer, promhttp.HandlerOpts{}))
	defer ts.Close()

	responseBody := func() string {
		resp, _ := http.Get(ts.URL)
		buf, _ := ioutil.ReadAll(resp.Body)
		return string(buf)
	}

	expectedSum := `TestElapsedTime_test_example_histogram_sum{foo1="bar1"}`
	expectedCount := `TestElapsedTime_test_example_histogram_count{foo1="bar1"} 1`

	actual := grep(p.App, responseBody())

	s.Equal(true, strings.Contains(actual, expectedSum))
	s.Equal(true, strings.Contains(actual, expectedCount))

	if err := p.server.Shutdown(context.Background()); err != nil {
		s.T().Fatal(err)
	}
}

func TestHistogramSuite(t *testing.T) {
	suite.Run(t, new(histogramSuite))
}
