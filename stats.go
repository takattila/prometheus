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
				_ = o.Gauge("stat_goroutines:count", float64(runtime.NumGoroutine()), Labels{})
			}
		}()
	}
}

func (o *Object) statMemoryUsage() {
	if o.StatMemoryUsage {
		t := time.NewTicker(time.Second)

		go func() {
			for range t.C {
				m, _ := mem.VirtualMemory()
				_ = o.Gauge("stat_memory_usage:total", float64(m.Total), Labels{})
				_ = o.Gauge("stat_memory_usage:avail", float64(m.Available), Labels{})
				_ = o.Gauge("stat_memory_usage:used", float64(m.Used), Labels{})
				_ = o.Gauge("stat_memory_usage:free", float64(m.Free), Labels{})
				_ = o.Gauge("stat_memory_usage:used_percent", float64(m.UsedPercent), Labels{})

				var memory runtime.MemStats
				runtime.ReadMemStats(&memory)
				// Sys is the total bytes of memory obtained from the OS
				_ = o.Gauge("stat_memory_usage:sys", float64(memory.Sys), Labels{})
				// Alloc is bytes of allocated heap objects
				_ = o.Gauge("stat_memory_usage:alloc", float64(memory.Alloc), Labels{})
				// HeapSys is bytes of heap memory obtained from the OS
				_ = o.Gauge("stat_memory_usage:heapsys", float64(memory.HeapSys), Labels{})
				// HeapInuse is bytes in in-use spans
				_ = o.Gauge("stat_memory_usage:heapinuse", float64(memory.HeapInuse), Labels{})
				// Frees is the cumulative count of heap objects freed
				_ = o.Gauge("stat_memory_usage:frees", float64(memory.Frees), Labels{})
				// NumGC is the number of completed GC cycles
				_ = o.Gauge("stat_memory_usage:numgc", float64(memory.NumGC), Labels{})
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
	if o.StatCountGoroutines {
		t := time.NewTicker(time.Second)

		go func() {
			for range t.C {
				percent := getCpuPercent().getFirstElement()
				_ = o.Gauge("stat_cpu_usage:percent", percent, Labels{})
			}
		}()
	}
}
