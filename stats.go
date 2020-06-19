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
				_ = o.Gauge("stat_goroutines:count", []Label{}, float64(runtime.NumGoroutine()))
			}
		}()
	}
}

func (o *Object) statMemoryUsage() {
	if o.StatMemoryUsage {
		t := time.NewTicker(time.Second)

		//byteToMB := func(b uint64) uint64 {
		//	return b / 1024 / 1024
		//}

		go func() {
			for range t.C {
				// var m runtime.MemStats
				// runtime.ReadMemStats(&m)
				m, _ := mem.VirtualMemory()
				_ = o.Gauge("stat_memory_usage:total", []Label{}, float64(m.Total))
				_ = o.Gauge("stat_memory_usage:avail", []Label{}, float64(m.Available))
				_ = o.Gauge("stat_memory_usage:used", []Label{}, float64(m.Used))
				_ = o.Gauge("stat_memory_usage:free", []Label{}, float64(m.Free))
				_ = o.Gauge("stat_memory_usage:used_percent", []Label{}, float64(m.UsedPercent))
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
