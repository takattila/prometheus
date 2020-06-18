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
			Value: "2020",
		},
	}, since, units...)

	fmt.Println()
	fmt.Println(p.GetMetrics("ExampleHistogram_test_history_histogram_bucket"))
	fmt.Println("Error:", err)

	// Output:
	// ExampleHistogram_test_history_histogram_bucket{sell="2020",le="1"} 1
	// ExampleHistogram_test_history_histogram_bucket{sell="2020",le="2"} 1
	// ExampleHistogram_test_history_histogram_bucket{sell="2020",le="3"} 1
	// ExampleHistogram_test_history_histogram_bucket{sell="2020",le="4"} 1
	// ExampleHistogram_test_history_histogram_bucket{sell="2020",le="5"} 1
	// ExampleHistogram_test_history_histogram_bucket{sell="2020",le="+Inf"} 1
	// Error: <nil>
}

func ExampleObject_ElapsedTime() {
	p := prometheus.New(prometheus.Init{
		Host:        "0.0.0.0",
		Port:        prometheus.GetFreePort(),
		Environment: "test",
		AppName:     "ExampleElapsedTime",
	})

	defer func(begin time.Time) {
		units := prometheus.GenerateUnits(0.02, 0.02, 5)

		err := p.ElapsedTime("response", []prometheus.Label{
			{
				Name:  "handler",
				Value: "purchases",
			},
		}, begin, units...)

		if err != nil {
			log.Fatal(err)
		}
	}(time.Now())

	fmt.Println()
	fmt.Println(p.GetMetrics("ExampleElapsedTime"))

	// Output example:
	// # HELP ExampleElapsedTime_test_response_histogram Histogram for: response
	// # TYPE ExampleElapsedTime_test_response_histogram histogram
	// ExampleElapsedTime_test_response_histogram_bucket{handler="purchases",le="0.02"} 1
	// ExampleElapsedTime_test_response_histogram_bucket{handler="purchases",le="0.04"} 1
	// ExampleElapsedTime_test_response_histogram_bucket{handler="purchases",le="0.06"} 1
	// ExampleElapsedTime_test_response_histogram_bucket{handler="purchases",le="0.08"} 1
	// ExampleElapsedTime_test_response_histogram_bucket{handler="purchases",le="0.1"} 1
	// ExampleElapsedTime_test_response_histogram_bucket{handler="purchases",le="+Inf"} 1
	// ExampleElapsedTime_test_response_histogram_sum{handler="purchases"} 6.675e-06
	// ExampleElapsedTime_test_response_histogram_count{handler="purchases"} 1
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
