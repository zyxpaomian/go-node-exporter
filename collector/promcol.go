package collector

import (
	"sync"
	"github.com/prometheus/client_golang/prometheus"
	"prom_exporter/collector/countercol"
	"prom_exporter/collector/gaugecol"
)

// 指标结构体
type Metrics struct {
	metrics map[string]*prometheus.Desc
	mutex   sync.Mutex
}

// 创建监控指标描述符, 命名方式为namespace+"_"+metricname 
func newGlobalMetric(namespace string, metricName string, docString string, labels []string) *prometheus.Desc {
	return prometheus.NewDesc(namespace + "_" + metricName, docString, labels, nil)
}


// 返回监控的metric的结构体
func NewMetrics(namespace string) *Metrics {
	return &Metrics{
		metrics: map[string]*prometheus.Desc{
			"cpu_counter_metric": newGlobalMetric(namespace, "cpu_counter_metric", "cpu_counter_metric 模板", []string{"cpu"}),
			"mem_gauge_metric": newGlobalMetric(namespace, "mem_gauge_metric","mem_gauge_metric 模板", []string{"memory"}),
		},
	}
}


// Describe的Interface， 实现描述符到channel
func (c *Metrics) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range c.metrics {
		ch <- m
	}
}

// Collect的Interface，实现具体的收集任务
func (c *Metrics) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock() 
	defer c.mutex.Unlock()

	// counter value
	for label, curValue := range countercol.GetCounterData() {
		ch <-prometheus.MustNewConstMetric(c.metrics["cpu_counter_metric"], prometheus.CounterValue, float64(curValue), label)
	}

	// gauge value
	for label, curValue := range gaugecol.GetGaugeData() {
		ch <-prometheus.MustNewConstMetric(c.metrics["mem_gauge_metric"], prometheus.GaugeValue, float64(curValue), label)
	}
}