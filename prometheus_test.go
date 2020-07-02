package prometheus

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type prometheusSuite struct {
	suite.Suite
}

func (s prometheusSuite) TestStartHttpServer() {
	// 1. Create new object and stert the HTTP server paralel.
	p := New(initProm("TestStartHttpServer"))

	err := p.Counter(CounterArgs{
		MetricName: "example_start_http_server",
		Labels:     Labels{"foo1": "bar1", "foo2": "bar2"},
		Value:      1,
	})
	s.Equal(nil, err)

	expected := `example_start_http_server{app="TestStartHttpServer",env="test",foo1="bar1",foo2="bar2"} 1`
	actual := p.GetMetrics(p.App)

	s.Contains(actual, expected)

	// 2. Restart the HTTP server.
	p.StopHttpServer()
	p.StartHttpServer()

	expected = `example_start_http_server{app="TestStartHttpServer",env="test",foo1="bar1",foo2="bar2"} 1`
	actual = p.GetMetrics(p.App)

	s.Contains(actual, expected)

	// 3. Everything is alright, shutting down the server.
	p.StopHttpServer()
}

func (s prometheusSuite) TestSetMetricsEndpoint() {
	init := initProm("TestSetMetricsEndpoint")
	init.MetricEndpoint = "/metrics"
	p := New(init)

	err := p.Counter(CounterArgs{
		MetricName: "example_set_metrics_endpoint",
		Labels:     Labels{"foo1": "bar1", "foo2": "bar2"},
		Value:      1,
	})
	s.Equal(nil, err)

	expected := `example_set_metrics_endpoint{app="TestSetMetricsEndpoint",env="test",foo1="bar1",foo2="bar2"} 1`
	actual := p.GetMetrics(p.App)

	s.Contains(actual, expected)
	p.StopHttpServer()
}

func TestPrometheusSuite(t *testing.T) {
	suite.Run(t, new(prometheusSuite))
}
