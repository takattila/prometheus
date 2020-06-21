package prometheus_test

import (
	"fmt"

	"github.com/takattila/prometheus"
)

func ExampleObject_Gauge() {
	p := prometheus.New(prometheus.Init{
		Host:        "0.0.0.0",
		Port:        prometheus.GetFreePort(),
		Environment: "test",
		AppName:     "ExampleGauge",
	})

	err := p.Gauge("cpu_usage_example", 15, prometheus.Labels{
		"core": "0",
	})

	fmt.Println()
	fmt.Println(p.GetMetrics("cpu_usage_example"))
	fmt.Println("Error:", err)

	// Output:
	// # HELP cpu_usage_example Gauge created for cpu_usage_example
	// # TYPE cpu_usage_example gauge
	// cpu_usage_example{app="ExampleGauge",core="0",env="test"} 15
	// Error: <nil>
}
