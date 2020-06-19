# GO - Prometheus

[![Test](https://github.com/takattila/prometheus/workflows/Test/badge.svg?branch=master)](https://github.com/takattila/prometheus/actions?query=workflow:Test)
[![Coverage Status](https://coveralls.io/repos/github/takattila/prometheus/badge.svg?branch=master)](https://coveralls.io/github/takattila/prometheus?branch=master)
[![GOdoc](https://img.shields.io/badge/godoc-reference-orange)](https://godoc.org/github.com/takattila/prometheus)
[![Version](https://img.shields.io/badge/dynamic/json.svg?label=version&url=https://api.github.com/repos/takattila/prometheus/releases/latest&query=tag_name)](https://github.com/takattila/prometheus/releases)

The prometheus package provides Prometheus implementations for metrics.
It provides statistics as well:

- Goroutines (count)
- Memory usage (bytes)
- CPU usage (percentage)

![prometheus screenshot](./img/screenshot-01.png)

## Table of contents

* [Example usage](#example-usage)
   * [Initialization](#initialization)
      * [Example code](#example-code)
      * [Example output](#example-output)
   * [Counter](#counter)
      * [Example code](#example-code-1)
      * [Example output](#example-output-1)
   * [Gauge](#gauge)
      * [Example code](#example-code-2)
      * [Example output](#example-output-2)
   * [Histogram](#histogram)
      * [Example code](#example-code-3)
      * [Example output](#example-output-3)
   * [Elapsed time](#elapsed-time)
      * [Example code](#example-code-4)
      * [Example output](#example-output-4)
   * [Other examples](#other-examples)

## Example usage

### Initialization

#### Example code

```go
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
```

[Back to top](#table-of-contents)

#### Example output

```json
{
  "Addr": "0.0.0.0:38033",
  "Env": "test",
  "App": "ExampleService",
  "StatCountGoroutines": true,
  "StatMemoryUsage": true,
  "StatCpuUsage": true
}
```

[Back to top](#table-of-contents)

### Counter

#### Example code

```go
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

fmt.Println(p.GetMetrics("response_status"))
```

[Back to top](#table-of-contents)

#### Example output

```bash
# HELP response_status Counter for: response_status
# TYPE response_status counter
response_status{app="ExampleCounter",env="test",handler="MyHandler1",statuscode="200"} 1
```

[Back to top](#table-of-contents)

### Gauge

#### Example code

```go
err := p.Gauge("cpu_usage", []prometheus.Label{
    {
        Name:  "core",
        Value: "0",
    },
}, 15)

fmt.Println(p.GetMetrics("cpu_usage"))
```

[Back to top](#table-of-contents)

#### Example output

```bash
# HELP cpu_usage Gauge for: cpu_usage
# TYPE cpu_usage gauge
cpu_usage{app="ExampleGauge",core="0",env="test"} 15
```

[Back to top](#table-of-contents)

### Histogram

#### Example code

```go
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
```

[Back to top](#table-of-contents)

#### Example output

```bash
history_bucket{app="ExampleHistogram",env="test",sell="actual",le="1"} 1
history_bucket{app="ExampleHistogram",env="test",sell="actual",le="2"} 1
history_bucket{app="ExampleHistogram",env="test",sell="actual",le="3"} 1
history_bucket{app="ExampleHistogram",env="test",sell="actual",le="4"} 1
history_bucket{app="ExampleHistogram",env="test",sell="actual",le="5"} 1
history_bucket{app="ExampleHistogram",env="test",sell="actual",le="+Inf"} 1
```

[Back to top](#table-of-contents)

### Elapsed time

#### Example code

```go
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
```

[Back to top](#table-of-contents)

#### Example output

```bash
# HELP get_stat Histogram for: get_stat
# TYPE get_stat histogram
get_stat_bucket{app="ExampleElapsedTime",env="test",handler="purchases",le="0.05"} 0
get_stat_bucket{app="ExampleElapsedTime",env="test",handler="purchases",le="0.1"} 0
get_stat_bucket{app="ExampleElapsedTime",env="test",handler="purchases",le="0.15000000000000002"} 1
get_stat_bucket{app="ExampleElapsedTime",env="test",handler="purchases",le="0.2"} 1
get_stat_bucket{app="ExampleElapsedTime",env="test",handler="purchases",le="0.25"} 1
get_stat_bucket{app="ExampleElapsedTime",env="test",handler="purchases",le="+Inf"} 1
get_stat_sum{app="ExampleElapsedTime",env="test",handler="purchases"} 0.100132995
get_stat_count{app="ExampleElapsedTime",env="test",handler="purchases"} 1
```

[Back to top](#table-of-contents)

### Other examples

For more examples, please visit: [godoc page](https://godoc.org/github.com/takattila/prometheus#pkg-examples) .

[Back to top](#table-of-contents)
