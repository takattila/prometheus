package prometheus_test

import (
	"fmt"
	"log"
	"time"

	"github.com/takattila/prometheus"
)

func ExampleObject_Histogram() {
	p := prometheus.New(prometheus.Init{
		Host:        "0.0.0.0",
		Port:        prometheus.GetFreePort(),
		Environment: "test",
		AppName:     "ExampleHistogram",
	})

	since := time.Since(time.Now()).Seconds()
	units := prometheus.GenerateUnits(1, 1, 5)

	err := p.Histogram("history", []prometheus.Label{
		{
			Name:  "sell",
			Value: "actual",
		},
	}, since, units...)

	fmt.Println()
	fmt.Println(p.GetMetrics("history_bucket"))
	fmt.Println("Error:", err)

	// Output:
	// history_bucket{app="ExampleHistogram",env="test",sell="actual",le="1"} 1
	// history_bucket{app="ExampleHistogram",env="test",sell="actual",le="2"} 1
	// history_bucket{app="ExampleHistogram",env="test",sell="actual",le="3"} 1
	// history_bucket{app="ExampleHistogram",env="test",sell="actual",le="4"} 1
	// history_bucket{app="ExampleHistogram",env="test",sell="actual",le="5"} 1
	// history_bucket{app="ExampleHistogram",env="test",sell="actual",le="+Inf"} 1
	// Error: <nil>
}

func ExampleObject_ElapsedTime() {
	p := prometheus.New(prometheus.Init{
		Host:        "0.0.0.0",
		Port:        prometheus.GetFreePort(),
		Environment: "test",
		AppName:     "ExampleElapsedTime",
	})

	start := time.Now()

	defer func(begin time.Time) {
		units := prometheus.GenerateUnits(0.05, 0.05, 5)

		err := p.ElapsedTime("get_stat", []prometheus.Label{
			{
				Name:  "handler",
				Value: "purchases",
			},
		}, begin, units...)

		if err != nil {
			log.Fatal(err)
		}
	}(start)

	time.Sleep(100 * time.Millisecond)

	// Output example:
	// # HELP get_stat Histogram for: get_stat
	// # TYPE get_stat histogram
	// get_stat_bucket{app="ExampleElapsedTime",env="test",handler="purchases",le="0.05"} 0
	// get_stat_bucket{app="ExampleElapsedTime",env="test",handler="purchases",le="0.1"} 0
	// get_stat_bucket{app="ExampleElapsedTime",env="test",handler="purchases",le="0.15000000000000002"} 1
	// get_stat_bucket{app="ExampleElapsedTime",env="test",handler="purchases",le="0.2"} 1
	// get_stat_bucket{app="ExampleElapsedTime",env="test",handler="purchases",le="0.25"} 1
	// get_stat_bucket{app="ExampleElapsedTime",env="test",handler="purchases",le="+Inf"} 1
	// get_stat_sum{app="ExampleElapsedTime",env="test",handler="purchases"} 0.100132995
	// get_stat_count{app="ExampleElapsedTime",env="test",handler="purchases"} 1
}

func ExampleGenerateUnits() {
	units := prometheus.GenerateUnits(1, 1.5, 8)
	fmt.Println(units)

	units = prometheus.GenerateUnits(2, 4, 10)
	fmt.Println(units)

	// Output:
	// [1 2.5 4 5.5 7 8.5 10 11.5]
	// [2 6 10 14 18 22 26 30 34 38]
}
