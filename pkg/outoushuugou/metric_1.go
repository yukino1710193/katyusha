package outoushuugou

import (
	// "github.com/prometheus/client_golang/prometheus"
)

// var AllMetrics = []prometheus.Collector{
// 	ColdStartPods,
// 	WarmPods,
// 	ProcessingTime,
// 	QueueingTime,
// 	Incoming,
// 	Outgoing,
// 	// thêm các metric khác ở đây
// }

// // Định nghĩa các metric
// // Chú ý: Các metric này cần được khởi tạo và đăng ký với Prometheus

// var ColdStartPods = prometheus.NewGaugeVec(
// 	prometheus.GaugeOpts{
// 		Name: "cold_start_pods",
// 		Help: "Number of cold pods (just started, not ready yet)",
// 	},
// 	[]string{"app"},
// )

// var WarmPods = prometheus.NewGaugeVec(
// 	prometheus.GaugeOpts{
// 		Name: "warm_pods",
// 		Help: "Number of warm pods (ready to serve requests)",
// 	},
// 	[]string{"app"},
// )

// var ProcessingTime = prometheus.NewHistogramVec(
// 	prometheus.HistogramOpts{
// 		Name:    "shuka_processing_seconds",
// 		Help:    "Processing time of request in shuka",
// 		Buckets: prometheus.ExponentialBuckets(0.001, 2, 15),
// 	},
// 	[]string{"uri", "method", "app"},
// )

// var QueueingTime = prometheus.NewHistogramVec(
// 	prometheus.HistogramOpts{
// 		Name:    "request_queueing_seconds",
// 		Help:    "Time spent in queue before processing",
// 		Buckets: prometheus.ExponentialBuckets(0.001, 2, 15),
// 	},
// 	[]string{"uri", "method", "app"},
// )

// var Incoming = prometheus.NewHistogramVec(
// 	prometheus.HistogramOpts{
// 		Name:    "request_incoming_seconds",
// 		Help:    "Timestamp when the request entered the system",
// 		Buckets: prometheus.ExponentialBuckets(0.001, 2, 15),
// 	},
// 	[]string{"uri", "method", "app"},
// )

// var Outgoing = prometheus.NewHistogramVec(
// 	prometheus.HistogramOpts{
// 		Name:    "request_outgoing_seconds",
// 		Help:    "Timestamp when the request exited the system",
// 		Buckets: prometheus.ExponentialBuckets(0.001, 2, 15),
// 	},
// 	[]string{"uri", "method", "app"},
// )

