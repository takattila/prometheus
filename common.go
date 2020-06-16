package prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (o *Object) serve() {
	_ = http.ListenAndServe(o.Addr, promhttp.Handler())
}

func getLabelNames(labels []Label) []string {
	var slice []string
	for _, o := range labels {
		slice = append(slice, o.Name)
	}
	return slice
}

func makeSlice(labels []Label) []string {
	var slice []string
	for _, o := range labels {
		slice = append(slice, o.Name, o.Value)
	}
	return slice
}
