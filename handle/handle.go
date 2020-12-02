package handle

import (
	"prom_exporter/server"
	"net/http"
    "prom_exporter/collector"
    "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


func InitHandle(r *server.WWWMux) {
	// 初始化
	initBasicMapping(r)
}

func initBasicMapping(r *server.WWWMux) {
	metrics := collector.NewMetrics("default")
	registry := prometheus.NewRegistry()
	registry.MustRegister(metrics)


    r.GetRouter().Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	// 默认路由
	r.GetRouter().NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
						<head><title>A Prometheus Exporter</title></head>
						<body>
						<h1>A Prometheus Exporter</h1>
						<p><a href='/metrics'>Metrics</a></p>
					</body>
					</html>`))
	})
}