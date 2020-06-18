package prometheus

import (
	"log"
	"testing"

	"github.com/phayes/freeport"
	"github.com/stretchr/testify/suite"
)

type commonSuite struct {
	suite.Suite
}

func getFreePort() (port int) {
	port, err := freeport.GetFreePort()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func initProm(appName string) Init {
	return Init{
		Host:        "0.0.0.0",
		Port:        getFreePort(),
		Environment: "test",
		AppName:     appName,
	}
}

func (s commonSuite) TestGenerateUnits() {
	expected := []float64{0.5, 1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5, 9.5}
	actual := GenerateUnits(0.5, 1, 10)
	s.Equal(expected, actual)
}

func (s commonSuite) TestGetLabelNames() {
	expected := []string{"foo1", "foo2"}
	actual := getLabelNames([]Label{
		{Name: "foo1", Value: "bar1"},
		{Name: "foo2", Value: "bar2"},
	})
	s.Equal(expected, actual)
}

func (s commonSuite) TestMakeSlice() {
	expected := []string{"foo1", "bar1", "foo2", "bar2"}
	actual := makeSlice([]Label{
		{Name: "foo1", Value: "bar1"},
		{Name: "foo2", Value: "bar2"},
	})
	s.Equal(expected, actual)
}

func (s commonSuite) TestMakeFQDN() {
	expected := `TestCounter_test_example_counter`
	actual := makeFQDN("TestCounter", "test", "example", "counter")
	s.Equal(expected, actual)
}

func TestCommonSuite(t *testing.T) {
	suite.Run(t, new(commonSuite))
}
