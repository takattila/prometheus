package prometheus

import (
	"io/ioutil"
	"net/http"
	"testing"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/stretchr/testify/suite"
)

type commonSuite struct {
	suite.Suite
}

func initProm(appName string) Init {
	return Init{
		Host:        "0.0.0.0",
		Port:        GetFreePort(),
		Environment: "test",
		AppName:     appName,

		StatCountGoroutines: false,
		StatMemoryUsage:     false,
		StatCpuUsage:        false,
	}
}

func (s commonSuite) TestGetLabelNames() {
	expected := []string{"foo1", "foo2"}
	actual := getLabelNames(Labels{
		"foo1": "bar1",
		"foo2": "bar2",
	})
	for _, e := range expected {
		s.Contains(actual, e)
	}
}

func (s commonSuite) TestEnablePprof() {
	var (
		pprofEndpoint = "/debug/pprof/"
		resp          *http.Response
		err           error
	)

	i := initProm("TestEnablePprof")
	i.EnablePprof = true
	p := New(i)

	count := 0
	maxLoops := 10

	for {
		if count == maxLoops {
			s.T().Fatalf("loop of TestEnablePprof is reached the maximum value: %d\n", maxLoops)
		}
		resp, err = http.Get("http://" + p.Addr + pprofEndpoint)
		if err == nil {
			break
		}
		count++
	}

	defer func() {
		err = resp.Body.Close()
		s.Equal(nil, err)
	}()

	b, err := ioutil.ReadAll(resp.Body)
	s.Equal(nil, err)

	stripped := strip.StripTags(string(b))
	s.Contains(stripped, pprofEndpoint)
	s.Contains(stripped, "Types of profiles available")
	s.Contains(stripped, "allocs")
	s.Contains(stripped, "block")
	s.Contains(stripped, "cmdline")
	s.Contains(stripped, "goroutine")
	s.Contains(stripped, "heap")
	s.Contains(stripped, "mutex")
	s.Contains(stripped, "profile")
	s.Contains(stripped, "threadcreate")
	s.Contains(stripped, "trace")
}

func TestCommonSuite(t *testing.T) {
	suite.Run(t, new(commonSuite))
}
