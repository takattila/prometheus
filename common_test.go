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

		StatCountGoroutines: true,
		StatMemoryUsage:     true,
		StatCpuUsage:        true,
	}
}

func (s commonSuite) TestGenerateUnits() {
	expected := []float64{0.5, 1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5, 9.5}
	actual := GenerateUnits(0.5, 1, 10)
	s.Equal(expected, actual)
}

func (s commonSuite) TestGetLabelNames() {
	expected := []string{"foo1", "foo2"}
	actual := getLabelNames(Labels{
		"foo1": "bar1",
		"foo2": "bar2",
	})
	s.Equal(expected, actual)
}

func TestCommonSuite(t *testing.T) {
	suite.Run(t, new(commonSuite))
}
