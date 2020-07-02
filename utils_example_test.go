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

	err := p.Counter(prometheus.CounterArgs{
		MetricName: "parse_output_example",
		Labels:     prometheus.Labels{"parsed": "true"},
		Value:      1,
	})
	fmt.Println(err)

	metric := p.GetMetrics("parse_output_example")
	parsed := prometheus.ParseOutput(metric)

	fmt.Println(metric)
	fmt.Println(parsed)

	// Output:
	// <nil>
	//
	// # HELP parse_output_example Counter created for parse_output_example
	// # TYPE parse_output_example counter
	// parse_output_example{app="ExampleParseOutput",env="test",parsed="true"} 1
	// map[parse_output_example:name:"parse_output_example" help:"Counter created for parse_output_example" type:COUNTER metric:<label:<name:"app" value:"ExampleParseOutput" > label:<name:"env" value:"test" > label:<name:"parsed" value:"true" > > ]
}

func ExampleGetLabels() {
	p := prometheus.New(prometheus.Init{
		Host:        "0.0.0.0",
		Port:        prometheus.GetFreePort(),
		Environment: "test",
		AppName:     "ExampleGetLabels",
	})

	metric := "promhttp_metric_handler_requests_total"

	err := p.Counter(prometheus.CounterArgs{
		MetricName: metric,
		Labels:     prometheus.Labels{"code": "200"},
		Value:      1,
	})
	fmt.Println(err)

	output := p.GetMetrics(p.App)
	fmt.Println(prometheus.GetLabels(output, metric))

	// Output:
	// <nil>
	// map[app:ExampleGetLabels code:200 env:test]
}

func ExampleRoundFloat() {
	float := prometheus.RoundFloat(1.559633154856, 2)
	fmt.Println(float)

	// Output:
	// 1.56
}

func ExampleDecimalPlaces() {
	dp := prometheus.DecimalPlaces(1.559633154856)
	fmt.Println(dp)

	// Output:
	// 12
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
