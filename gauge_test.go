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
		p.Gauge("example", []Label{
			{
				Name:  "foo1",
				Value: "bar1",
			},
		}, value)

		expected := `TestGauge_test_example_gauge{foo1="bar1"} ` + fmt.Sprintf("%g", value)
		actual := grep(p.App, p.getMetrics())

		s.Equal(true, strings.Contains(actual, expected))
	}

	p.StopHttpServer()
}

func TestGaugeSuite(t *testing.T) {
	suite.Run(t, new(gaugeSuite))
}
