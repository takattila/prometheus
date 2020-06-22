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
	for n := 0; n < b.N; n++ {
		since := time.Since(time.Now()).Seconds()
		units := GenerateUnits(1, 1, 5)
		_ = p.Histogram("history", since, Labels{
			"sell": "actual",
		}, units...)
	}
}
