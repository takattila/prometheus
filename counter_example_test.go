package prometheus_test

import (
	"fmt"

	"github.com/takattila/prometheus"
)

func ExampleObject_Counter() {
	p := prometheus.New(prometheus.Init{
		Host:        "0.0.0.0",
		Port:        prometheus.GetFreePort(),
		Environment: "test",
		AppName:     "ExampleCounter",
	})

	err := p.Counter("response", []prometheus.Label{
		{
			Name:  "handler",
			Value: "MyHandler1",
		},
		{
			Name:  "statuscode",
			Value: "200",
		},
	}, 1)

	fmt.Println()
	fmt.Println(p.GetMetrics("ExampleCounter"))
	fmt.Println("Error:", err)

	// Output:
	// # HELP ExampleCounter_test_response_counter Counter for: response
	// # TYPE ExampleCounter_test_response_counter counter
	// ExampleCounter_test_response_counter{handler="MyHandler1",statuscode="200"} 1
	// Error: <nil>
}
