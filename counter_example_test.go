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

	err := p.Counter("response_status", []prometheus.Label{
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
	fmt.Println(p.GetMetrics("response_status"))
	fmt.Println("Error:", err)

	// Output:
	// # HELP response_status Counter for: response_status
	// # TYPE response_status counter
	// response_status{app="ExampleCounter",env="test",handler="MyHandler1",statuscode="200"} 1
	// Error: <nil>
}
