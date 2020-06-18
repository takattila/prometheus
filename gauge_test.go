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
		err := p.Gauge("example", []Label{
			{
				Name:  "foo1",
				Value: "bar1",
			},
		}, value)
		s.Equal(nil, err)

		expected := `TestGauge_test_example_gauge{foo1="bar1"} ` + fmt.Sprintf("%g", value)
		actual := grep(p.App, p.getMetrics())

		s.Equal(true, strings.Contains(actual, expected))
	}

	p.StopHttpServer()
}

func (s gaugeSuite) TestGaugeError() {
	p := New(initProm("TestGaugeError"))

	err := p.Gauge("example", []Label{
		{
			Name:  "foo1",
			Value: "bar1",
		},
	}, 60)
	s.Equal(nil, err)

	actual := p.Gauge("example", []Label{
		{
			Name:  "bad_label",
			Value: "bar1",
		},
	}, 50)

	// Bad label name
	expected := `metric: 'TestGaugeError_test_example_gauge', error: 'label name "foo1" missing in label map', input label names: 'bad_label', correct label names: 'foo1'` + "\n"
	s.Equal(expected, fmt.Sprint(actual))

	p.StopHttpServer()
}

func TestGaugeSuite(t *testing.T) {
	suite.Run(t, new(gaugeSuite))
}
