package prometheus

import (
	"testing"

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

func TestCommonSuite(t *testing.T) {
	suite.Run(t, new(commonSuite))
}
