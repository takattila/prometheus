// The prometheus package provides Prometheus implementations for metrics.
package prometheus

import (
	"bufio"
	"fmt"
	"math"
	"strings"

	"github.com/phayes/freeport"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
)

// ParseOutput reads 'text' as the simple and flat text-based
// exchange format and creates MetricFamily proto messages.
func ParseOutput(text string) map[string]*dto.MetricFamily {
	p := expfmt.TextParser{}
	m, _ := p.TextToMetricFamilies(strings.NewReader(text))
	return m
}

// GetLabels gives back the labels for a metric.
func GetLabels(text, metric string) Labels {
	out := ParseOutput(text)
	obj := out[metric]

	labels := make(Labels)
	for _, m := range obj.GetMetric() {
		for _, l := range m.GetLabel() {
			labels[*l.Name] = *l.Value
		}
	}
	return labels
}

// GetFreePort asks the kernel for a free open port that is ready to use.
func GetFreePort() (port int) {
	port, _ = freeport.GetFreePort()
	return
}

// Grep processes text line by line,
// and gives back any lines which match a specified word.
func Grep(find, inText string) (result string) {
	scanner := bufio.NewScanner(strings.NewReader(inText))
	for scanner.Scan() {
		if strings.Contains(strings.ToLower(scanner.Text()), strings.ToLower(find)) {
			result += "\n" + scanner.Text()
		}
	}
	return
}

// RoundFloat truncate the decimal places of a float64 number
// by a given precision:
//   RoundFloat(1.599633154856, 2) -> 1.6
func RoundFloat(value float64, decimalPlaces int) float64 {
	pow := math.Pow(10, float64(decimalPlaces))
	return math.Round(value*pow) / pow
}

// DecimalPlaces returns the number of decimal places of a float64 number.
func DecimalPlaces(n float64) int {
	num := fmt.Sprintf("%g", n)
	if strings.Contains(num, ".") {
		num = strings.Split(num, ".")[1]
		num = strings.TrimRight(num, "0") // remove trailing 0s
		return len(num)
	}
	return 0
}
