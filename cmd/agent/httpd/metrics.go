package httpd

import "github.com/prometheus/client_golang/prometheus"

var (
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

	httpAction = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "request_path_sum",
			Help: "show request time by path",
		},
		[]string{"method", "path", "status_code"},
	)
)
