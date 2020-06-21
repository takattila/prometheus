package prometheus

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type counterSuite struct {
	suite.Suite
}

func (s counterSuite) TestCounter() {
	p := New(initProm("TestCounter"))

	err := p.Counter("example_counter", 1, Labels{
		"foo1": "bar1",
		"foo2": "bar2",
	})
	s.Equal(nil, err)

	expected := `example_counter{app="TestCounter",env="test",foo1="bar1",foo2="bar2"} 1`
	actual := p.GetMetrics(p.App)

	s.Equal(true, strings.Contains(actual, expected))
	p.StopHttpServer()
}

func (s counterSuite) TestCounterError() {
	p := New(initProm("TestCounterError"))

	actual := p.Counter("example_counter_error", 1, Labels{
		"bad label foo1": "bar1",
	})

	// Incorrect label name
	expected := `metric: 'example_counter_error', error: 'descriptor Desc{fqName: "example_counter_error", help: "Counter created for example_counter_error", constLabels: {}, variableLabels: [bad label foo1 app env]} is invalid: "bad label foo1" is not a valid label name for metric "example_counter_error"', input label names: 'app, bad label foo1, env'` + "\n"
	s.Equal(expected, fmt.Sprint(actual))

	p.StopHttpServer()
}

func TestCounterSuite(t *testing.T) {
	suite.Run(t, new(counterSuite))
}
