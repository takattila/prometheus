package prometheus

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type counterSuite struct {
	suite.Suite
}

func (s counterSuite) TestCounter() {
	p := New(initProm("TestCounter"))

	err := p.Counter(CounterArgs{
		MetricName: "example_counter",
		Labels:     Labels{"foo1": "bar1", "foo2": "bar2"},
		Value:      1,
	})
	s.Equal(nil, err)

	expected := `example_counter{app="TestCounter",env="test",foo1="bar1",foo2="bar2"} 1`
	actual := p.GetMetrics(p.App)

	s.Contains(actual, expected)
	p.StopHttpServer()
}

func (s counterSuite) TestCounterError() {
	p := New(initProm("TestCounterError"))
	actual := p.Counter(CounterArgs{
		MetricName: "example_counter_error",
		Labels:     Labels{"bad label foo1": "bar1"},
		Value:      1,
	})

	// Incorrect label name
	expected := `invalid: "bad label foo1" is not a valid label name for metric "example_counter_error"`
	s.Contains(fmt.Sprint(actual), expected)

	p.StopHttpServer()
}

func TestCounterSuite(t *testing.T) {
	suite.Run(t, new(counterSuite))
}
