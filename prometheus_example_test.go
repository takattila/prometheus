package prometheus_test

import (
	"encoding/json"
	"fmt"

	"github.com/takattila/prometheus"
)

func ExampleNew() {
	p := prometheus.New(prometheus.Init{
		// Obligatory fields
		Host:        "0.0.0.0",
		Port:        prometheus.GetFreePort(),
		Environment: "test",
		AppName:     "ExampleService",

		// Optional fields
		StatCountGoroutines: true, // default: false
		StatMemoryUsage:     true, // default: false
		StatCpuUsage:        true, // default: false
	})

	b, _ := json.MarshalIndent(p, "", "  ")
	fmt.Println(string(b))

	// Output example:
	// {
	//   "Addr": "0.0.0.0:38033",
	//   "Env": "test",
	//   "App": "ExampleService",
	//   "StatCountGoroutines": true,
	//   "StatMemoryUsage": true,
	//   "StatCpuUsage": true
	// }
}

func ExampleObject_StartHttpServer() {
	p := prometheus.New(prometheus.Init{
		Host:        "0.0.0.0",
		Port:        prometheus.GetFreePort(),
		Environment: "test",
		AppName:     "ExampleStartHttpServer",
	})

	p.StartHttpServer()
}

func ExampleObject_StopHttpServer() {
	p := prometheus.New(prometheus.Init{
		Host:        "0.0.0.0",
		Port:        prometheus.GetFreePort(),
		Environment: "test",
		AppName:     "ExampleStopHttpServer",
	})

	p.StopHttpServer()
}

func ExampleObject_GetMetrics() {
	p := prometheus.New(prometheus.Init{
		Host:        "0.0.0.0",
		Port:        prometheus.GetFreePort(),
		Environment: "test",
		AppName:     "ExampleGetMetrics",
	})

	fmt.Println(p.GetMetrics("go_"))

	// Output example:
	// # HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
	// # TYPE go_gc_duration_seconds summary
	// go_gc_duration_seconds{quantile="0"} 1.5162e-05
	// go_gc_duration_seconds{quantile="0.25"} 1.9539e-05
	// go_gc_duration_seconds{quantile="0.5"} 3.6708e-05
	// go_gc_duration_seconds{quantile="0.75"} 9.2103e-05
	// go_gc_duration_seconds{quantile="1"} 0.00023626
	// go_gc_duration_seconds_sum 0.000506999
	// go_gc_duration_seconds_count 7
	// # HELP go_goroutines Number of goroutines that currently exist.
	// # TYPE go_goroutines gauge
	// go_goroutines 24
}
