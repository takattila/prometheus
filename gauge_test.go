package prometheus

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	clientGo "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/suite"
)

type gaugeSuite struct {
	suite.Suite
}

func (s gaugeSuite) TestGauge() {
	p := New(Init{
		Host:        "0.0.0.0",
		Port:        8080,
		Environment: "test",
		AppName:     "TestGauge",
	})

	for _, value := range []float64{5, 2, 19, 77} {
		p.Gauge("example", []Label{
			{
				Name:  "foo1",
				Value: "bar1",
			},
		}, value)

		ts := httptest.NewServer(promhttp.HandlerFor(clientGo.DefaultGatherer, promhttp.HandlerOpts{}))
		defer ts.Close()

		responseBody := func() string {
			resp, _ := http.Get(ts.URL)
			buf, _ := ioutil.ReadAll(resp.Body)
			return string(buf)
		}

		expected := `TestGauge_test_example_gauge{foo1="bar1"} ` + fmt.Sprintf("%g", value)
		actual := grep(p.App, responseBody())

		s.Equal(true, strings.Contains(actual, expected))
	}

	if err := p.server.Shutdown(context.Background()); err != nil {
		s.T().Fatal(err)
	}
}

func TestGaugeSuite(t *testing.T) {
	suite.Run(t, new(gaugeSuite))
}
