package outoushuugou

import (
	"time"
)
type metric struct {
	// Các metric khác
	ColdStartPods   int
	WarmPods        int
	ProcessingTime  string
	QueueingTime    string
	Incoming        string
	Outgoing        string
	// Các metric khác
	// ...
}
func (rp *ResponsePool) StartRLexporter(interval time.Duration) {
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
					// Ghi metric tại đây
					
					// Đẩy metric cho RL tại đây
					exportRL(fb)
				}
				rp.lastExportedCount = newCount
			}
			rp.PoolAppendingLock.Unlock()
		}
	}()
}

func exportRL(fb *ResponseFeedback) {
	
}

