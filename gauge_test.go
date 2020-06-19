package prometheus

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type gaugeSuite struct {
	suite.Suite
}

func (s gaugeSuite) TestGauge() {
	p := New(initProm("TestGauge"))

	for _, value := range []float64{5, 2, 19, 77} {
		err := p.Gauge("example_gauge", []Label{
			{
				Name:  "foo1",
				Value: "bar1",
			},
		}, value)
		s.Equal(nil, err)

		expected := `example_gauge{app="TestGauge",env="test",foo1="bar1"} ` + fmt.Sprintf("%g", value)
		actual := p.GetMetrics(p.App)

		s.Equal(true, strings.Contains(actual, expected))
	}

	p.StopHttpServer()
}

func (s gaugeSuite) TestGaugeError() {
	p := New(initProm("TestGaugeError"))

	err := p.Gauge("example_gauge_error", []Label{
		{
			Name:  "foo1",
			Value: "bar1",
		},
	}, 60)
	s.Equal(nil, err)

	actual := p.Gauge("example_gauge_error", []Label{
		{
			Name:  "bad_label",
			Value: "bar1",
		},
	}, 50)

	// Bad label name
	expected := `metric: 'example_gauge_error', error: 'label name "foo1" missing in label map', input label names: 'bad_label, app, env', correct label names: 'app, env, foo1'` + "\n"
	s.Equal(expected, fmt.Sprint(actual))

	p.StopHttpServer()
}

func TestGaugeSuite(t *testing.T) {
	suite.Run(t, new(gaugeSuite))
}
