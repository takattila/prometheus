package prometheus

import (
	"testing"
)

func BenchmarkGauge(b *testing.B) {
	p := New(Init{
		Host:        "0.0.0.0",
		Port:        GetFreePort(),
		Environment: "test",
		AppName:     "ExampleGauge",
	})
	for n := 0; n < b.N; n++ {
		_ = p.Gauge("cpu_usage_example", 15, Labels{
			"core": "0",
		})
	}
}
