package prometheus

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type gaugeSuite struct {
	suite.Suite
}

func (s gaugeSuite) TestGauge() {
	p := New(initProm("TestGauge"))

	for _, value := range []float64{5, 2, 19, 77} {
		err := p.Gauge(GaugeArgs{
			MetricName: "example_gauge",
			Labels:     Labels{"foo1": "bar1"},
			Value:      value,
		})
		s.Equal(nil, err)

		expected := `example_gauge{app="TestGauge",env="test",foo1="bar1"} ` + fmt.Sprintf("%g", value)
		actual := p.GetMetrics(p.App)

		s.Contains(actual, expected)
	}

	p.StopHttpServer()
}

func (s gaugeSuite) TestGaugeError() {
	p := New(initProm("TestGaugeError"))

	err := p.Gauge(GaugeArgs{
		MetricName: "example_gauge_error",
		Labels:     Labels{"foo1": "bar1"},
		Value:      60,
	})
	s.Equal(nil, err)

	actual := p.Gauge(GaugeArgs{
		MetricName: "example_gauge_error",
		Labels:     Labels{"bad_label": "bar1"},
		Value:      50,
	})

	// Bad label name
	expected := `metric: 'example_gauge_error', error: 'label name "foo1" missing in label map', input label names: 'app, bad_label, env', correct label names: 'app, env, foo1'` + "\n"
	s.Equal(expected, fmt.Sprint(actual))

	p.StopHttpServer()
}

func TestGaugeSuite(t *testing.T) {
	suite.Run(t, new(gaugeSuite))
}
