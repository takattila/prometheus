package prometheus

import (
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func (o *Object) statCountGoroutines() {
	if o.StatCountGoroutines {
		t := time.NewTicker(time.Second)

		go func() {
			for range t.C {
				_ = o.Gauge(GaugeArgs{MetricName: "stat_goroutines:count", Value: float64(runtime.NumGoroutine())})
			}
		}()
	}
}

func (o *Object) statMemoryUsage() {
	if o.StatMemoryUsage {
		t := time.NewTicker(time.Second)

		go func() {
			for range t.C {
				// System memory usage
				m, _ := mem.VirtualMemory()
				// Total amount of RAM on this system
				_ = o.Gauge(GaugeArgs{MetricName: "stat_memory_usage:total", Value: float64(m.Total)})
				// RAM available for programs to allocate
				_ = o.Gauge(GaugeArgs{MetricName: "stat_memory_usage:avail", Value: float64(m.Available)})
				// RAM used by programs
				_ = o.Gauge(GaugeArgs{MetricName: "stat_memory_usage:used", Value: float64(m.Used)})
				// This is the kernel's notion of free memory
				_ = o.Gauge(GaugeArgs{MetricName: "stat_memory_usage:free", Value: float64(m.Free)})
				// Percentage of RAM used by programs
				_ = o.Gauge(GaugeArgs{MetricName: "stat_memory_usage:used_percent", Value: float64(m.UsedPercent)})

				// Used memory by the application
				var memory runtime.MemStats
				runtime.ReadMemStats(&memory)
				// Sys is the total bytes of memory obtained from the OS
				_ = o.Gauge(GaugeArgs{MetricName: "stat_memory_usage:sys", Value: float64(memory.Sys)})
				// Alloc is bytes of allocated heap objects
				_ = o.Gauge(GaugeArgs{MetricName: "stat_memory_usage:alloc", Value: float64(memory.Alloc)})
				// HeapSys is bytes of heap memory obtained from the OS
				_ = o.Gauge(GaugeArgs{MetricName: "stat_memory_usage:heapsys", Value: float64(memory.HeapSys)})
				// HeapInuse is bytes in in-use spans
				_ = o.Gauge(GaugeArgs{MetricName: "stat_memory_usage:heapinuse", Value: float64(memory.HeapInuse)})
				// Frees is the cumulative count of heap objects freed
				_ = o.Gauge(GaugeArgs{MetricName: "stat_memory_usage:frees", Value: float64(memory.Frees)})
				// NumGC is the number of completed GC cycles
				_ = o.Gauge(GaugeArgs{MetricName: "stat_memory_usage:numgc", Value: float64(memory.NumGC)})
			}
		}()
	}
}

type cpuP struct {
	per []float64
	err error
}

func getCpuPercent() cpuP {
	c := cpuP{}
	c.per, c.err = cpu.Percent(1*time.Second, false)
	return c
}

func (c cpuP) getFirstElement() float64 {
	if c.err != nil {
		return float64(0)
	}
	return c.per[0]
}

func (o *Object) statCpuUsage() {
	if o.StatCpuUsage {
		t := time.NewTicker(time.Second)

		go func() {
			for range t.C {
				percent := getCpuPercent().getFirstElement()
				_ = o.Gauge(GaugeArgs{MetricName: "stat_cpu_usage:percent", Value: percent})
			}
		}()
	}
}
