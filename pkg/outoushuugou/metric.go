package outoushuugou

import (
	"time"
	"github.com/bonavadeur/katyusha/pkg/bonalib"
)

type Metric struct {
	// Các metric khác
	ProcessingTime  float64
	QueueingNTime   float64
	// Incoming        int32
	// Outgoing        int32
	QueueingJTime   float64
	QueueingJLength float64
	// Các metric khác
}

func (rp *ResponsePool) StartRLexporter(interval time.Duration) {
	go func() {
		for {
			time.Sleep(interval)

			rp.PoolAppendingLock.Lock()
			// Lấy số lượng phần tử trong Pool
			newCount := len(rp.Pool)
			if newCount > rp.lastExportedCount {
				// Lấy phần tử mới (prepend nên lấy [0:delta])
				delta := newCount - rp.lastExportedCount
				newItems := rp.Pool[:delta]
				// cho Data về null
				DATA = nil
				for _, fb := range newItems {
					// Ghi metric tại đây
					appendMetricFromFeedback(fb, &DATA)
					// Đẩy metric cho RL tại đây
				}
				rp.lastExportedCount = newCount
			}
			rp.PoolAppendingLock.Unlock()
		}
	}()
}

func clearDataPOOL() {
	// Xóa dữ liệu trong POOL
	POOL.PoolAppendingLock.Lock()
	defer POOL.PoolAppendingLock.Unlock()

	POOL.Pool = make([]*ResponseFeedback, 0)
	POOL.lastExportedCount = 0
	DATA = nil
	// Ghi log
	bonalib.Log("Outoushuugou", "Data cleared")
	// Ghi log
	bonalib.Log("Outoushuugou", "Data cleared")
}
