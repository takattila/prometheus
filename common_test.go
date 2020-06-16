package prometheus

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

func grep(grep, text string) (result string) {
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		if strings.Contains(strings.ToLower(scanner.Text()), strings.ToLower(grep)) {
			result += "\n" + scanner.Text()
		}
	}
	return
}

type commonSuite struct {
	suite.Suite
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