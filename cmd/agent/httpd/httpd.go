package httpd

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"net/http"
	"net/http/pprof"
	"time"
)

func NewRestfulServer(addr string, logger *zap.Logger) http.Server {
	prometheus.MustRegister(hdFailure)
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(httpAction)

	cpuTicker := time.NewTicker(1 * time.Second)
	defer cpuTicker.Stop()
	go func() {
		for {
			<-cpuTicker.C
			cpuTemp.Inc()
		}
	}()

	hdTicker := time.NewTicker(1 * time.Second)
	defer hdTicker.Stop()
	go func() {
		for {
			<-hdTicker.C
			hdFailure.With(prometheus.Labels{
				"hd_name": "/dev/xxx1",
				"size":    "6T",
			}).Set(100)
		}
	}()

	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	wrapperMux := NewHttpHandler(mux, logger)

	return http.Server{Addr: addr, Handler: wrapperMux}
}
