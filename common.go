package prometheus

import (
	"net/http"
	"time"

	"github.com/pressly/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func serve(addr string) *http.Server {
	r := chi.NewRouter()

	r.Handle("/", promhttp.Handler())

	s := &http.Server{
		Addr:           addr,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() { _ = s.ListenAndServe() }()

	return s
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
