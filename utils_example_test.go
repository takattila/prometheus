package prometheus_test

import (
	"fmt"

	"github.com/takattila/prometheus"
)

func ExampleParseOutput() {
	p := prometheus.New(prometheus.Init{
		Host:        "0.0.0.0",
		Port:        prometheus.GetFreePort(),
		Environment: "test",
		AppName:     "ExampleParseOutput",
	})

	metric := p.GetMetrics("go_goroutines")
	parsed := prometheus.ParseOutput(metric)

	fmt.Println(metric)
	fmt.Println(parsed)

	// Output example:
	// # HELP go_goroutines Number of goroutines that currently exist.
	// # TYPE go_goroutines gauge
	// go_goroutines 9
	// map[go_goroutines:name:"go_goroutines" help:"Number of goroutines that currently exist." type:GAUGE metric:<> ]
}

func ExampleGetLabels() {
	p := prometheus.New(prometheus.Init{
		Host:        "0.0.0.0",
		Port:        prometheus.GetFreePort(),
		Environment: "test",
		AppName:     "ExampleGetLabels",
	})

	output := p.GetMetrics("promhttp_metric_handler_requests_total")
	metric := "promhttp_metric_handler_requests_total"

	fmt.Println(prometheus.GetLabels(output, metric))

	// Output:
	// [{code 200} {code 500} {code 503}]
}

func ExampleGetFreePort() {
	port := prometheus.GetFreePort()
	fmt.Println(port)

	// Output example:
	// 45689
}

func ExampleGrep() {
	text := `
Lorem ipsum dolor sit amet, consectetur adipiscing elit,
sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris
nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in
reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
culpa qui officia deserunt mollit anim id est laborum.
`
	fmt.Println(prometheus.Grep("dolore", text))

	// Output:
	// sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
	// reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
}