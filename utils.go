// The prometheus package provides Prometheus implementations for metrics.
package prometheus

import (
	"bufio"
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
func GetLabels(text, metric string) (labels []Label) {
	out := ParseOutput(text)
	obj := out[metric]

	for _, m := range obj.GetMetric() {
		for _, l := range m.GetLabel() {
			labels = append(labels, Label{
				Name:  *l.Name,
				Value: *l.Value,
			})
		}
	}
	return
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
