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

	err := p.Gauge("cpu_usage", []prometheus.Label{
		{
			Name:  "core",
			Value: "0",
		},
	}, 15)

	fmt.Println()
	fmt.Println(p.GetMetrics("cpu_usage"))
	fmt.Println("Error:", err)

	// Output:
	// # HELP cpu_usage Gauge for: cpu_usage
	// # TYPE cpu_usage gauge
	// cpu_usage{app="ExampleGauge",core="0",env="test"} 15
	// Error: <nil>
}
