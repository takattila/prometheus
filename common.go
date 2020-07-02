package prometheus

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/pprof"
	"sort"
	"strings"
	"time"

	"github.com/pressly/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (o *Object) serve() *http.Server {
	r := chi.NewRouter()

	s := &http.Server{
		Addr:           o.Addr,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	r.Handle(o.MetricsEndpoint, promhttp.HandlerFor(
		o.reg, promhttp.HandlerOpts{
			DisableCompression:  true,
			MaxRequestsInFlight: 1,
			EnableOpenMetrics:   true,
		}))

	r.HandleFunc("/debug/pprof/", pprof.Index)

	for _, pp := range []string{
		"allocs",
		"block",
		"cmdline",
		"goroutine",
		"heap",
		"mutex",
		"profile",
		"threadcreate",
		"trace",
	} {
		r.Handle("/debug/pprof/"+pp+"", pprof.Handler((pp)))
	}

	go func() { _ = s.ListenAndServe() }()

	return s
}

func (o *Object) getMetrics() string {
	var resp *http.Response
	var err error

	for {
		resp, err = http.Get("http://" + o.Addr + o.MetricsEndpoint)
		if err == nil {
			break
		}
	}

	defer func() { _ = resp.Body.Close() }()
	buf, _ := ioutil.ReadAll(resp.Body)

	return string(buf)
}

func getLabelNames(labels Labels) []string {
	var slice []string
	for k := range labels {
		slice = append(slice, k)
	}
	return slice
}

func (o *Object) errorHandler(err interface{}, fqdn string, inputLabelNames []string) error {
	metric := o.GetMetrics(fqdn)
	sort.Strings(inputLabelNames)
	givenLabelNames := strings.Join(inputLabelNames, ", ")
	correctLabelNames := func() (ret string) {
		ln := getLabelNames(GetLabels(metric, fqdn))
		sort.Strings(ln)
		if len(ln) != 0 {
			ret = fmt.Sprintf(", correct label names: '%s'", strings.Join(ln, ", "))
		}
		return
	}

	return fmt.Errorf("metric: '%s', error: '%s', input label names: '%s'%s\n",
		fqdn,
		err,
		givenLabelNames,
		correctLabelNames())
}

func (o *Object) addServiceInfoToLabels(labels Labels) Labels {
	if labels == nil {
		labels = Labels{}
	}
	labels["app"] = o.App
	labels["env"] = o.Env
	return labels
}
