package outoushuugou

import (
	"time"

	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	responseCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "response_count",
			Help: "Count of responses processed by URI and Method",
		},
		[]string{"uri", "method"},
	)
)

// func init() {
// 	prometheus.MustRegister(responseCount)
// }

func (rp *ResponsePool) StartPrometheusExporter(interval time.Duration) {
	go func() {
		for {
			time.Sleep(interval)

			rp.PoolAppendingLock.Lock()

			newCount := len(rp.Pool)
			if newCount > rp.lastExportedCount {
				// Lấy phần tử mới (prepend nên lấy [0:delta])
				delta := newCount - rp.lastExportedCount
				newItems := rp.Pool[:delta]

				for _, fb := range newItems {
					// Ghi Prometheus metric tại đây
					// Ví dụ:
					responseCount.WithLabelValues(fb.URI, fb.Method).Inc()
				}

				rp.lastExportedCount = newCount
			}

			rp.PoolAppendingLock.Unlock()
		}
	}()
}

func StartPrometheusServer() {
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":2112", nil) // hoặc bất kỳ port nào bạn muốn
}
