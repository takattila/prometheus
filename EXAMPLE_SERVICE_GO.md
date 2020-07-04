# Example Sevice - GO Source Code

[Back to README.md](./README.md#table-of-contents)

```go
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/go-chi/chi"
	"github.com/jedib0t/go-pretty/table"
	"github.com/takattila/prometheus"
)

var serviceAddr = "localhost:6060"

func init() {
	go func() {
		t := time.NewTicker(10 * time.Millisecond)
		for range t.C {
			for _, endpoint := range []string{"1", "2"} {
				resp, _ := http.Get("http://" + serviceAddr + "/" + endpoint)
				_ = resp.Body.Close()
			}
		}
	}()
}

func main() {
	p := prometheus.New(prometheus.Init{
		Host:        "0.0.0.0",
		Port:        8080,
		Environment: "test",
		AppName:     "exampleService",

		StatCountGoroutines: true,
		StatMemoryUsage:     true,
		StatCpuUsage:        true,
		EnablePprof:         true,
	})

	r := chi.NewRouter()

	r.Get("/1", MyHandler1(p))
	r.Get("/2", MyHandler2(p))

	s := &http.Server{
		Addr:           serviceAddr,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fatalIfErr(s.ListenAndServe())
}

func MyHandler1(p *prometheus.Object) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		printMemStats()
		handlerName := "MyHandler1"

		// Elapsed time - Histogram
		defer func(begin time.Time) {
			fatalIfErr(p.Histogram(prometheus.HistogramArgs{
				MetricName: "response_time:sec",
				Labels:     prometheus.Labels{"handler": handlerName},
				Units:      prometheus.GenerateUnits(1, 1, 10),
				Value:      time.Since(begin).Seconds(),
			}))
		}(time.Now())

		// Response status - Counter
		fatalIfErr(p.Counter(prometheus.CounterArgs{
			MetricName: "response_status",
			Labels: prometheus.Labels{
				"handler":    handlerName,
				"statuscode": randomdata.StringSample("200", "302", "200", "200", "404"),
			},
			Value: 1,
		}))

		// Response Size - Gauge
		fatalIfErr(p.Gauge(prometheus.GaugeArgs{
			MetricName: "response_size",
			Labels:     prometheus.Labels{"handler": handlerName},
			Value:      float64(rand.Intn(100)),
		}))

		// MeasureExecTime - Nanoseconds
		ns := p.StartMeasureExecTime(prometheus.MeasureExecTimeArgs{
			MetricName:   "execution_time:nano_sec",
			Labels:       prometheus.Labels{"handler": handlerName},
			Units:        prometheus.GenerateUnits(5000, 5000, 20),
			TimeDuration: time.Nanosecond,
		})
		calculateSomething(1000, time.Nanosecond)
		fatalIfErr(ns.StopMeasureExecTime())

		// MeasureExecTime - Microseconds
		µs := p.StartMeasureExecTime(prometheus.MeasureExecTimeArgs{
			MetricName:   "execution_time:micro_sec",
			Labels:       prometheus.Labels{"handler": handlerName},
			Units:        prometheus.GenerateUnits(50, 50, 20),
			TimeDuration: time.Microsecond,
		})
		calculateSomething(1000, time.Microsecond)
		fatalIfErr(µs.StopMeasureExecTime())

		// MeasureExecTime - Milliseconds
		ms := p.StartMeasureExecTime(prometheus.MeasureExecTimeArgs{
			MetricName:   "execution_time:milli_sec",
			Labels:       prometheus.Labels{"handler": handlerName},
			Units:        prometheus.GenerateUnits(5, 5, 20),
			TimeDuration: time.Millisecond,
		})
		calculateSomething(100, time.Millisecond)
		fatalIfErr(ms.StopMeasureExecTime())

		// MeasureExecTime - Seconds
		s := p.StartMeasureExecTime(prometheus.MeasureExecTimeArgs{
			MetricName:   "execution_time:seconds",
			Labels:       prometheus.Labels{"handler": handlerName},
			Units:        prometheus.GenerateUnits(0.5, 0.5, 10),
			TimeDuration: time.Second,
		})
		calculateSomething(5, time.Second)
		fatalIfErr(s.StopMeasureExecTime())

		// MeasureExecTime - Minutes
		m := p.StartMeasureExecTime(prometheus.MeasureExecTimeArgs{
			MetricName:   "execution_time:minutes",
			Labels:       prometheus.Labels{"handler": handlerName},
			Units:        prometheus.GenerateUnits(0.005, 0.005, 20),
			TimeDuration: time.Minute,
		})
		calculateSomething(5, time.Second)
		fatalIfErr(m.StopMeasureExecTime())
	}
}

func MyHandler2(p *prometheus.Object) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		printMemStats()
		handlerName := "MyHandler2"

		// Elapsed time - Histogram
		defer func(begin time.Time) {
			fatalIfErr(p.Histogram(prometheus.HistogramArgs{
				MetricName: "response_time:milli_sec",
				Labels:     prometheus.Labels{"handler": handlerName},
				Units:      prometheus.GenerateUnits(0.05, 0.05, 10),
				Value:      time.Since(begin).Seconds(),
			}))
		}(time.Now())

		// Response Size - Gauge
		fatalIfErr(p.Gauge(prometheus.GaugeArgs{
			MetricName: "response_size",
			Labels:     prometheus.Labels{"handler": handlerName},
			Value:      float64(rand.Intn(100)),
		}))

		// Response status - Counter
		fatalIfErr(p.Counter(prometheus.CounterArgs{
			MetricName: "response_status",
			Labels: prometheus.Labels{
				"handler":    handlerName,
				"statuscode": randomdata.StringSample("200", "302", "200", "200", "200", "404"),
			},
			Value: 1,
		}))

		calculateSomething(300, time.Millisecond)
	}
}

func fatalIfErr(err error) {
	if err != nil {
		log.Fatal("[FATAL] ", err)
	}
}

func calculateSomething(num int, duration time.Duration) {
	time.Sleep(time.Duration(rand.Intn(num)) * duration)
}

func printMemStats() {
	formatBytes := func(b uint64) string {
		const unit = 1000
		if b < unit {
			return fmt.Sprintf("%d B", b)
		}
		div, exp := int64(unit), 0
		for n := b / unit; n >= unit; n /= unit {
			div *= unit
			exp++
		}
		return fmt.Sprintf("%.1f %cB",
			float64(b)/float64(div), "kMGTPE"[exp])
	}

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Stat", "Value", "Help"})
	t.AppendRows([]table.Row{
		{"Sys", formatBytes(mem.Sys), "Sys is the total bytes of memory obtained from the OS"},
		{"Alloc", formatBytes(mem.Alloc), "Alloc is bytes of allocated heap objects"},
		{"HeapSys", formatBytes(mem.HeapSys), "HeapSys is bytes of heap memory obtained from the OS"},
		{"HeapInuse", formatBytes(mem.HeapInuse), "HeapInuse is bytes in in-use spans"},
		{"Frees", formatBytes(mem.Frees), "Frees is the cumulative count of heap objects freed"},
		{"NumGC", uint64(mem.NumGC), "NumGC is the number of completed GC cycles"},
	})
	t.Render()
}
```
[Back to README.md](./README.md#table-of-contents)
