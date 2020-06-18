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

	err := p.Gauge("cpu", []prometheus.Label{
		{
			Name:  "usage",
			Value: "percentage",
		},
	}, 15)

	fmt.Println()
	fmt.Println(p.GetMetrics("ExampleGauge"))
	fmt.Println("Error:", err)

	// Output:
	// # HELP ExampleGauge_test_cpu_gauge Gauge for: cpu
	// # TYPE ExampleGauge_test_cpu_gauge gauge
	// ExampleGauge_test_cpu_gauge{usage="percentage"} 15
	// Error: <nil>
}
