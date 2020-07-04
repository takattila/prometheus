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
		err := p.Histogram(prometheus.HistogramArgs{
			MetricName: "get_stat",
			Labels:     prometheus.Labels{"handler": "purchases"},
			Buckets:    prometheus.GenerateBuckets(0.5, 0.05, 5),
			Value:      time.Since(begin).Seconds(),
		})

		if err != nil {
			log.Fatal(err)
		}
	}(start)

	time.Sleep(100 * time.Millisecond)

	// Output example:
	// # HELP get_stat Histogram created for get_stat
	// # TYPE get_stat histogram
	// get_stat_bucket{app="ExampleHistogram",env="test",handler="purchases",le="0.05"} 0
	// get_stat_bucket{app="ExampleHistogram",env="test",handler="purchases",le="0.1"} 0
	// get_stat_bucket{app="ExampleHistogram",env="test",handler="purchases",le="0.15"} 1
	// get_stat_bucket{app="ExampleHistogram",env="test",handler="purchases",le="0.2"} 1
	// get_stat_bucket{app="ExampleHistogram",env="test",handler="purchases",le="0.25"} 1
	// get_stat_bucket{app="ExampleHistogram",env="test",handler="purchases",le="+Inf"} 1
	// get_stat_sum{app="ExampleHistogram",env="test",handler="purchases"} 0.100233303
	// get_stat_count{app="ExampleHistogram",env="test",handler="purchases"} 1
}

func ExampleGenerateBuckets() {
	buckets := prometheus.GenerateBuckets(1, 1.5, 8)
	fmt.Println(buckets)

	buckets = prometheus.GenerateBuckets(2, 4, 10)
	fmt.Println(buckets)

	// Output:
	// [1 2.5 4 5.5 7 8.5 10 11.5]
	// [2 6 10 14 18 22 26 30 34 38]
}

func ExampleObject_StartMeasureExecTime() {
	functionName := "calculate"

	p := prometheus.New(prometheus.Init{
		Host:        "0.0.0.0",
		Port:        prometheus.GetFreePort(),
		Environment: "test",
		AppName:     "StartMeasureExecTime",
	})

	// Nanoseconds start
	ns := p.StartMeasureExecTime(prometheus.MeasureExecTimeArgs{
		MetricName:   "execution_time_nano_sec",
		Labels:       prometheus.Labels{"function": functionName},
		Buckets:      prometheus.GenerateBuckets(5000, 10000, 10),
		TimeDuration: time.Nanosecond,
	})
	time.Sleep(5000 * time.Nanosecond)

	err := ns.StopMeasureExecTime()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p.GetMetrics("execution_time_nano_sec"))
	// Nanoseconds end

	// Microseconds start
	µs := p.StartMeasureExecTime(prometheus.MeasureExecTimeArgs{
		MetricName:   "execution_time_micro_sec",
		Labels:       prometheus.Labels{"function": functionName},
		Buckets:      prometheus.GenerateBuckets(50, 50, 10),
		TimeDuration: time.Microsecond,
	})
	time.Sleep(100 * time.Microsecond)

	err = µs.StopMeasureExecTime()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p.GetMetrics("execution_time_micro_sec"))
	// Microseconds end

	// Milliseconds start
	ms := p.StartMeasureExecTime(prometheus.MeasureExecTimeArgs{
		MetricName:   "execution_time_milli_sec",
		Labels:       prometheus.Labels{"function": functionName},
		Buckets:      prometheus.GenerateBuckets(5, 5, 10),
		TimeDuration: time.Millisecond,
	})
	time.Sleep(10 * time.Millisecond)

	err = ms.StopMeasureExecTime()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p.GetMetrics("execution_time_milli_sec"))
	// Milliseconds end

	// Seconds start
	s := p.StartMeasureExecTime(prometheus.MeasureExecTimeArgs{
		MetricName:   "execution_time_seconds",
		Labels:       prometheus.Labels{"function": functionName},
		Buckets:      prometheus.GenerateBuckets(0.5, 0.5, 10),
		TimeDuration: time.Second,
	})
	time.Sleep(1 * time.Second)

	err = s.StopMeasureExecTime()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p.GetMetrics("execution_time_seconds"))
	// Seconds end

	// Minutes start
	m := p.StartMeasureExecTime(prometheus.MeasureExecTimeArgs{
		MetricName:   "execution_time_minutes",
		Labels:       prometheus.Labels{"function": functionName},
		Buckets:      prometheus.GenerateBuckets(0.005, 0.005, 10),
		TimeDuration: time.Minute,
	})
	time.Sleep(1 * time.Second)

	err = m.StopMeasureExecTime()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p.GetMetrics("execution_time_minutes"))
	// Minutes end

	// Output example:
	// # HELP execution_time_nano_sec Histogram created for execution_time_nano_sec
	// # TYPE execution_time_nano_sec histogram
	// execution_time_nano_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="5000"} 0
	// execution_time_nano_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="15000"} 0
	// execution_time_nano_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="25000"} 0
	// execution_time_nano_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="35000"} 0
	// execution_time_nano_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="45000"} 1
	// execution_time_nano_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="55000"} 1
	// execution_time_nano_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="65000"} 1
	// execution_time_nano_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="75000"} 1
	// execution_time_nano_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="85000"} 1
	// execution_time_nano_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="95000"} 1
	// execution_time_nano_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="+Inf"} 1
	// execution_time_nano_sec_sum{app="StartMeasureExecTime",env="test",function="calculate"} 43280
	// execution_time_nano_sec_count{app="StartMeasureExecTime",env="test",function="calculate"} 1
	//
	// # HELP execution_time_micro_sec Histogram created for execution_time_micro_sec
	// # TYPE execution_time_micro_sec histogram
	// execution_time_micro_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="50"} 0
	// execution_time_micro_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="100"} 0
	// execution_time_micro_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="150"} 0
	// execution_time_micro_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="200"} 0
	// execution_time_micro_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="250"} 1
	// execution_time_micro_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="300"} 1
	// execution_time_micro_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="350"} 1
	// execution_time_micro_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="400"} 1
	// execution_time_micro_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="450"} 1
	// execution_time_micro_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="500"} 1
	// execution_time_micro_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="+Inf"} 1
	// execution_time_micro_sec_sum{app="StartMeasureExecTime",env="test",function="calculate"} 236
	// execution_time_micro_sec_count{app="StartMeasureExecTime",env="test",function="calculate"} 1
	//
	// # HELP execution_time_milli_sec Histogram created for execution_time_milli_sec
	// # TYPE execution_time_milli_sec histogram
	// execution_time_milli_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="5"} 0
	// execution_time_milli_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="10"} 1
	// execution_time_milli_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="15"} 1
	// execution_time_milli_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="20"} 1
	// execution_time_milli_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="25"} 1
	// execution_time_milli_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="30"} 1
	// execution_time_milli_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="35"} 1
	// execution_time_milli_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="40"} 1
	// execution_time_milli_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="45"} 1
	// execution_time_milli_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="50"} 1
	// execution_time_milli_sec_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="+Inf"} 1
	// execution_time_milli_sec_sum{app="StartMeasureExecTime",env="test",function="calculate"} 10
	// execution_time_milli_sec_count{app="StartMeasureExecTime",env="test",function="calculate"} 1
	//
	// # HELP execution_time_seconds Histogram created for execution_time_seconds
	// # TYPE execution_time_seconds histogram
	// execution_time_seconds_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="0.5"} 0
	// execution_time_seconds_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="1"} 0
	// execution_time_seconds_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="1.5"} 1
	// execution_time_seconds_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="2"} 1
	// execution_time_seconds_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="2.5"} 1
	// execution_time_seconds_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="3"} 1
	// execution_time_seconds_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="3.5"} 1
	// execution_time_seconds_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="4"} 1
	// execution_time_seconds_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="4.5"} 1
	// execution_time_seconds_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="5"} 1
	// execution_time_seconds_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="+Inf"} 1
	// execution_time_seconds_sum{app="StartMeasureExecTime",env="test",function="calculate"} 1.000324369
	// execution_time_seconds_count{app="StartMeasureExecTime",env="test",function="calculate"} 1
	//
	// # HELP execution_time_minutes Histogram created for execution_time_minutes
	// # TYPE execution_time_minutes histogram
	// execution_time_minutes_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="0.005"} 0
	// execution_time_minutes_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="0.01"} 0
	// execution_time_minutes_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="0.015"} 0
	// execution_time_minutes_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="0.02"} 1
	// execution_time_minutes_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="0.025"} 1
	// execution_time_minutes_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="0.03"} 1
	// execution_time_minutes_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="0.035"} 1
	// execution_time_minutes_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="0.04"} 1
	// execution_time_minutes_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="0.045"} 1
	// execution_time_minutes_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="0.05"} 1
	// execution_time_minutes_bucket{app="StartMeasureExecTime",env="test",function="calculate",le="+Inf"} 1
	// execution_time_minutes_sum{app="StartMeasureExecTime",env="test",function="calculate"} 0.016671208216666667
	// execution_time_minutes_count{app="StartMeasureExecTime",env="test",function="calculate"} 1
}

func ExampleMeasureExecTime_StopMeasureExecTime() {
	p := prometheus.New(prometheus.Init{
		Host:        "0.0.0.0",
		Port:        prometheus.GetFreePort(),
		Environment: "test",
		AppName:     "StopMeasureExecTime",
	})

	ms := p.StartMeasureExecTime(prometheus.MeasureExecTimeArgs{
		MetricName:   "execution_time_milli_sec",
		Labels:       prometheus.Labels{"function": "calculate"},
		Buckets:      prometheus.GenerateBuckets(5, 5, 10),
		TimeDuration: time.Millisecond,
	})

	time.Sleep(10 * time.Millisecond)

	err := ms.StopMeasureExecTime()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p.GetMetrics("execution_time_milli_sec"))

	// Output example:
	// # HELP execution_time_milli_sec Histogram created for execution_time_milli_sec
	// # TYPE execution_time_milli_sec histogram
	// execution_time_milli_sec_bucket{app="StopMeasureExecTime",env="test",function="calculate",le="5"} 0
	// execution_time_milli_sec_bucket{app="StopMeasureExecTime",env="test",function="calculate",le="10"} 1
	// execution_time_milli_sec_bucket{app="StopMeasureExecTime",env="test",function="calculate",le="15"} 1
	// execution_time_milli_sec_bucket{app="StopMeasureExecTime",env="test",function="calculate",le="20"} 1
	// execution_time_milli_sec_bucket{app="StopMeasureExecTime",env="test",function="calculate",le="25"} 1
	// execution_time_milli_sec_bucket{app="StopMeasureExecTime",env="test",function="calculate",le="30"} 1
	// execution_time_milli_sec_bucket{app="StopMeasureExecTime",env="test",function="calculate",le="35"} 1
	// execution_time_milli_sec_bucket{app="StopMeasureExecTime",env="test",function="calculate",le="40"} 1
	// execution_time_milli_sec_bucket{app="StopMeasureExecTime",env="test",function="calculate",le="45"} 1
	// execution_time_milli_sec_bucket{app="StopMeasureExecTime",env="test",function="calculate",le="50"} 1
	// execution_time_milli_sec_bucket{app="StopMeasureExecTime",env="test",function="calculate",le="+Inf"} 1
	// execution_time_milli_sec_sum{app="StopMeasureExecTime",env="test",function="calculate"} 10
	// execution_time_milli_sec_count{app="StopMeasureExecTime",env="test",function="calculate"} 1

}
