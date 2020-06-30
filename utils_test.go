package prometheus

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type utilsSuite struct {
	suite.Suite
}

func (s utilsSuite) TestRoundFloat() {
	float64Num := 1.56569355485

	for _, t := range []struct {
		decimalPlaces int
		expected      float64
	}{
		{
			decimalPlaces: 0,
			expected:      2,
		},
		{
			decimalPlaces: 1,
			expected:      1.6,
		},
		{
			decimalPlaces: 2,
			expected:      1.57,
		},
		{
			decimalPlaces: 3,
			expected:      1.566,
		},
		{
			decimalPlaces: 4,
			expected:      1.5657,
		},
		{
			decimalPlaces: 5,
			expected:      1.56569,
		},
		{
			decimalPlaces: 6,
			expected:      1.565694,
		},
		{
			decimalPlaces: 7,
			expected:      1.5656936,
		},
		{
			decimalPlaces: 8,
			expected:      1.56569355,
		},
		{
			decimalPlaces: 9,
			expected:      1.565693555,
		},
		{
			decimalPlaces: 10,
			expected:      1.5656935549,
		},
	} {
		actual := RoundFloat(float64Num, t.decimalPlaces)
		s.Equal(t.expected, actual)
	}
}

func (s utilsSuite) TestDecimalPlaces() {

	for _, t := range []struct {
		expected int
		number   float64
	}{
		{
			expected: 0,
			number:   1,
		},
		{
			expected: 0,
			number:   1.0,
		},
		{
			expected: 1,
			number:   1.1,
		},
		{
			expected: 2,
			number:   1.11,
		},
		{
			expected: 3,
			number:   1.111,
		},
		{
			expected: 4,
			number:   1.1111,
		},
		{
			expected: 5,
			number:   1.11111,
		},
		{
			expected: 6,
			number:   1.111111,
		},
		{
			expected: 7,
			number:   1.1111111,
		},
		{
			expected: 8,
			number:   1.11111111,
		},
		{
			expected: 9,
			number:   1.111111111,
		},
		{
			expected: 10,
			number:   1.1111111111,
		},
		{
			expected: 11,
			number:   1.11111111111,
		},
		{
			expected: 12,
			number:   1.111111111111,
		},
		{
			expected: 13,
			number:   1.1111111111111,
		},
		{
			expected: 14,
			number:   1.11111111111111,
		},
		{
			expected: 15,
			number:   1.111111111111111,
		},
		{
			expected: 15,
			number:   1.0111111111111111,
		},
		{
			expected: 16,
			number:   1.1111111111111111,
		},
		{
			expected: 16,
			number:   1.11111111111111111,
		},
		{
			expected: 16,
			number:   1.11111111111111110,
		},
	} {
		actual := DecimalPlaces(t.number)
		s.Equal(t.expected, actual)
	}
}

func TestUtilsSuite(t *testing.T) {
	suite.Run(t, new(utilsSuite))
}
