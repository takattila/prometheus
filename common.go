package prometheus

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pressly/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (o *Object) serve() *http.Server {
	r := chi.NewRouter()
	r.Handle("/*", promhttp.Handler())

	s := &http.Server{
		Addr:           o.Addr,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() { _ = s.ListenAndServe() }()

	return s
}

func (o *Object) getMetrics() string {
	var resp *http.Response
	var err error

	for {
		resp, err = http.Get("http://" + o.Addr)
		if err == nil {
			break
		}
	}

	buf, _ := ioutil.ReadAll(resp.Body)

	return string(buf)
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

func (o *Object) errorHandler(err interface{}, fqdn string, inputLabelNames []string) error {
	metric := o.GetMetrics(fqdn)
	givenLabelNames := strings.Join(inputLabelNames, ", ")
	correctLabelNames := func() (ret string) {
		ln := getLabelNames(GetLabels(metric, fqdn))
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

func (o *Object) addServiceInfoToLabels(labels []Label) []Label {
	return append(labels,
		Label{Name: "app", Value: o.App},
		Label{Name: "env", Value: o.Env})
}
