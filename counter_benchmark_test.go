package prometheus

import "testing"

func BenchmarkCounter(b *testing.B) {
	p := New(Init{
		Host:        "0.0.0.0",
		Port:        GetFreePort(),
		Environment: "test",
		AppName:     "ExampleCounter",
	})
	for n := 0; n < b.N; n++ {
		_ = p.Counter("response_status", 1, Labels{
			"handler":    "MyHandler1",
			"statuscode": "200",
		})
	}
}
