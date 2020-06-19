package prometheus

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type prometheusSuite struct {
	suite.Suite
}

func (s prometheusSuite) TestStartHttpServer() {
	// 1. Create new object and stert the HTTP server paralel.
	p := New(initProm("TestStartHttpServer"))

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

	expected := `TestStartHttpServer_test_example_counter{foo1="bar1",foo2="bar2"} 1`
	actual := p.GetMetrics(p.App)

	s.Equal(true, strings.Contains(actual, expected))

	// 2. Restart the HTTP server.
	p.StopHttpServer()
	p.StartHttpServer()

	expected = `TestStartHttpServer_test_example_counter{foo1="bar1",foo2="bar2"} 1`
	actual = p.GetMetrics(p.App)

	s.Equal(true, strings.Contains(actual, expected))

	// 3. Everything is alright, shutting down the server.
	p.StopHttpServer()
}

func TestPrometheusSuite(t *testing.T) {
	suite.Run(t, new(prometheusSuite))
}
