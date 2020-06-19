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

	err := p.Counter("example", []Label{
		{
			Name:  "foo1",
			Value: "bar1",
		},
		{
			Name:  "foo2",
			Value: "bar2",
		},
	}, 1)
	s.Equal(nil, err)

	expected := `TestCounter_test_example_counter{foo1="bar1",foo2="bar2"} 1`
	actual := p.GetMetrics(p.App)

	s.Equal(true, strings.Contains(actual, expected))

	p.StopHttpServer()
}

func (s counterSuite) TestCounterError() {
	p := New(initProm("TestCounterError"))

	actual := p.Counter("example", []Label{
		{
			Name:  "bad label foo1",
			Value: "bar1",
		},
	}, 1)

	// Incorrect label name
	expected := `metric: 'TestCounterError_test_example_counter', error: 'descriptor Desc{fqName: "TestCounterError_test_example_counter", help: "Counter for: example", constLabels: {}, variableLabels: [bad label foo1]} is invalid: "bad label foo1" is not a valid label name for metric "TestCounterError_test_example_counter"', input label names: 'bad label foo1'` + "\n"
	s.Equal(expected, fmt.Sprint(actual))

	p.StopHttpServer()
}

func TestCounterSuite(t *testing.T) {
	suite.Run(t, new(counterSuite))
}
