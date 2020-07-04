package prometheus

import "testing"

func BenchmarkGenerateBuckets(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateBuckets(0.5, 1, 10)
	}
}

func BenchmarkGetLabelNames(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getLabelNames(Labels{
			"foo1": "bar1",
			"foo2": "bar2",
		})
	}
}
