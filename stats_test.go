package prometheus

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type statsSuite struct {
	suite.Suite
}

func (s statsSuite) TestStatCountGoroutines() {
	p := New(initProm("TestStatCountGoroutines"))

	expected := "stat_goroutines:count"
	actual := ""
	for {
		actual = p.GetMetrics(expected)
		if strings.Contains(actual, expected) {
			break
		}
	}

	s.Contains(actual, expected)
	p.StopHttpServer()
}

func (s statsSuite) TestStatMemoryUsage() {
	p := New(initProm("TestStatMemoryUsage"))

	for _, t := range []string{"total", "avail", "used", "free", "used_percent"} {

		expected := "stat_memory_usage:" + t
		actual := ""
		for {
			actual = p.GetMetrics(expected)
			if strings.Contains(actual, expected) {
				break
			}
		}

		s.Contains(actual, expected)

	}
	p.StopHttpServer()
}

func (s statsSuite) TestStatCpuUsage() {
	p := New(initProm("TestStatCpuUsage"))

	expected := "stat_cpu_usage:percent"
	actual := ""
	for {
		actual = p.GetMetrics(expected)
		if strings.Contains(actual, expected) {
			break
		}
	}

	s.Equal(true, strings.Contains(actual, expected))
	p.StopHttpServer()
}

func (s statsSuite) TestHandleCpuPercentError() {
	c := cpuP{
		per: nil,
		err: fmt.Errorf("received two CPU counts: %d != %d", 1, 2),
	}

	expected := float64(0)
	actual := c.getFirstElement()

	s.Equal(expected, actual)
}

func TestStatsSuite(t *testing.T) {
	suite.Run(t, new(statsSuite))
}
