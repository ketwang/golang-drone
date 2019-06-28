package httpd

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"time"
	"util/pkg/singal"
)

var (
	HttpCmd = &cobra.Command{
		Use:   "serve",
		Short: "start a http server",
		RunE:  serve,
	}

	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temp_cal",
		Help: "cpu temp cal.",
		ConstLabels: map[string]string{
			"cpu_num": "all",
		},
	})

	hdFailure = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "hdd_failure_counter",
			Help: "show hdd failure counter",
		},
		[]string{"hd_name", "size"},
	)
)

func serve(cmd *cobra.Command, args []string) error {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("/backend/golang-drone/cmd/agent/config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	port := viper.GetString("listen.port")
	fmt.Println(port)

	prometheus.MustRegister(hdFailure)
	prometheus.MustRegister(cpuTemp)

	var errChan chan error
	ctx := singal.WithSignalsContext(context.Background())

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

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":9999", nil); err != nil {
			errChan <- err
		}
	}()

	select {
	case <-ctx.Done():
	case <-errChan:
	}

	return nil
}

/*
source:
  prometheus server
  push gateway
  jobs/exporter

data model:
  metric name:
  label:
  sample:

<metric name>{<label name>=<label value>, …}

counter:
  counter: 累加型metric
  Gauge:   常规metric、可以任意加减
  Histogram: 柱状图histogram
  Summary: 类似于histigram，但是提供count和sum功能；童工百分位功能
*/
