package prometheus

import (
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
)

type statsSuite struct {
	suite.Suite
}

func (s statsSuite) TestStatCountGoroutines() {
	i := initProm("TestStatCountGoroutines")
	i.StatCountGoroutines = true
	p := New(i)

	expected := "stat_goroutines:count"
	actual := ""

	var wg sync.WaitGroup
	wg.Add(1)
	for {
		actual = p.GetMetrics(expected)
		if strings.Contains(actual, expected) {
			wg.Done()
			break
		}
	}
	wg.Wait()

	s.Contains(actual, expected)
	p.StopHttpServer()
}

func (s statsSuite) TestStatMemoryUsage() {
	i := initProm("TestStatMemoryUsage")
	i.StatMemoryUsage = true
	p := New(i)

	for _, t := range []string{"total", "avail", "used", "free", "used_percent"} {

		expected := "stat_memory_usage:" + t
		actual := ""

		var wg sync.WaitGroup
		wg.Add(1)
		for {
			actual = p.GetMetrics(expected)
			if strings.Contains(actual, expected) {
				wg.Done()
				break
			}
		}
		wg.Wait()

		s.Contains(actual, expected)

	}
	p.StopHttpServer()
}

func (s statsSuite) TestStatCpuUsage() {
	i := initProm("TestStatCpuUsage")
	i.StatCpuUsage = true
	p := New(i)

	expected := "stat_cpu_usage:percent"
	actual := ""

	var wg sync.WaitGroup
	wg.Add(1)
	for {
		actual = p.GetMetrics(expected)
		if strings.Contains(actual, expected) {
			wg.Done()
			break
		}
	}
	wg.Wait()

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
