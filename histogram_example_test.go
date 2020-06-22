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

	start := time.Now()

	// Elapsed time to measure the computation time
	// of a given function, handler, etc...
	defer func(begin time.Time) {
		units := prometheus.GenerateUnits(0.05, 0.05, 5)
		since := time.Since(begin).Seconds()

		err := p.Histogram("get_stat", since, prometheus.Labels{
			"handler": "purchases",
		}, units...)

		if err != nil {
			log.Fatal(err)
		}
	}(start)

	time.Sleep(100 * time.Millisecond)

	// Output example:
	// # HELP get_stat Histogram created for get_stat
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
