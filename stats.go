package prometheus

import (
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func (o *Object) statCountGoroutines() {
	if o.StatCountGoroutines {
		t := time.NewTicker(time.Second)
		go func() {
			for range t.C {
				_ = o.Gauge("stat_goroutines:count", []Label{}, float64(runtime.NumGoroutine()))
			}
		}()
	}
}

func (o *Object) statMemoryUsage() {
	if o.StatMemoryUsage {
		t := time.NewTicker(time.Second)

		byteToMB := func(b uint64) uint64 {
			return b / 1024 / 1024
		}

		m := runtime.MemStats{}

		go func() {
			for range t.C {
				runtime.ReadMemStats(&m)
				_ = o.Gauge("stat_memory_usage:alloc", []Label{}, float64(byteToMB(m.Alloc)))
			}
		}()

		go func() {
			for range t.C {
				runtime.ReadMemStats(&m)
				_ = o.Gauge("stat_memory_usage:total", []Label{}, float64(byteToMB(m.TotalAlloc)))
			}
		}()

		go func() {
			for range t.C {
				runtime.ReadMemStats(&m)
				_ = o.Gauge("stat_memory_usage:sys", []Label{}, float64(byteToMB(m.Sys)))
			}
		}()

		go func() {
			for range t.C {
				runtime.ReadMemStats(&m)
				_ = o.Gauge("stat_memory_usage:gc", []Label{}, float64(m.NumGC))
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
				_ = o.Gauge("stat_cpu_usage:percent", []Label{}, percent)
			}
		}()
	}
}
