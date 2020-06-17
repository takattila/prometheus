package prometheus

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type counterSuite struct {
	suite.Suite
}

func (s counterSuite) TestCounter() {
	p := New(initProm("TestCounter"))

	p.Counter("example", []Label{
		{
			Name:  "foo1",
			Value: "bar1",
		},
		{
			Name:  "foo2",
			Value: "bar2",
		},
	}, 1)

	expected := `TestCounter_test_example_counter{foo1="bar1",foo2="bar2"} 1`
	actual := grep(p.App, p.getMetrics())

	s.Equal(true, strings.Contains(actual, expected))

	p.StopHttpServer()
}

func TestCounterSuite(t *testing.T) {
	suite.Run(t, new(counterSuite))
}
