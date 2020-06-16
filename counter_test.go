package prometheus

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	clientGo "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/suite"
)

type counterSuite struct {
	suite.Suite
}

func (s counterSuite) TestCounter() {
	p := New(Init{
		Host:        "0.0.0.0",
		Port:        8080,
		Environment: "test",
		AppName:     "TestCounter",
	})

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

	ts := httptest.NewServer(promhttp.HandlerFor(clientGo.DefaultGatherer, promhttp.HandlerOpts{}))
	defer ts.Close()

	responseBody := func() string {
		resp, _ := http.Get(ts.URL)
		buf, _ := ioutil.ReadAll(resp.Body)
		return string(buf)
	}

	expected := `TestCounter_test_example_counter{foo1="bar1",foo2="bar2"} 1`
	actual := grep(p.App, responseBody())

	s.Equal(true, strings.Contains(actual, expected))

	if err := p.server.Shutdown(context.Background()); err != nil {
		s.T().Fatal(err)
	}
}

func TestCounterSuite(t *testing.T) {
	suite.Run(t, new(counterSuite))
}
