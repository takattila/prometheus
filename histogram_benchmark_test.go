package prometheus

import (
	"testing"
	"time"
)

func BenchmarkHistogram(b *testing.B) {
	p := New(Init{
		Host:        "0.0.0.0",
		Port:        GetFreePort(),
		Environment: "test",
		AppName:     "ExampleHistogram",
	})
	buckets := GenerateBuckets(1, 1, 10)
	for n := 0; n < b.N; n++ {
		_ = p.Histogram(HistogramArgs{
			MetricName: "history",
			Labels:     Labels{"sell": "actual"},
			Buckets:    buckets,
			Value:      time.Since(time.Now()).Seconds(),
		})
	}
}
