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
		_ = p.Gauge(GaugeArgs{
			MetricName: "cpu_usage_example",
			Labels:     Labels{"core": "0"},
			Value:      15,
		})
	}
}
