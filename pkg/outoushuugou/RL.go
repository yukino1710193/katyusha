package outoushuugou

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bonavadeur/katyusha/pkg/fukabunsan"
)

func (rp *ResponsePool) RL(interval time.Duration) {
	for {
		time.Sleep(interval)

		fmt.Println("[RL] Start RL computation loop")

		numPods := len(fukabunsan.POD_READY)
		if numPods == 0 {
			fmt.Println("[RL][WARN] Không có POD_READY nào, bỏ qua vòng tính toán.")
			continue
		}

		QueueingNTime := make([]float64, numPods)
		ProcessingTime := make([]float64, numPods)

		for _, s := range DATA {
			nodeRaw := s.NodeD
			nodeStr := strings.TrimPrefix(nodeRaw, "node")
			nodeIndex, err := strconv.Atoi(nodeStr)

			if err != nil {
				fmt.Printf("[RL][WARN] Invalid NodeD format: '%s', err: %v\n", nodeRaw, err)
				continue
			}

			nodeIndex -= 1
			if nodeIndex < 1 || nodeIndex >= numPods {
				fmt.Printf("[RL][WARN] Node index out of range: %d (from '%s'), POD_READY size = %d\n", nodeIndex, nodeRaw, numPods)
				continue
			}

			QueueingNTime[nodeIndex] += s.QueueingNTime
			ProcessingTime[nodeIndex] += s.ProcessingTime

			fmt.Printf("[RL][DEBUG] node=%s index=%d QN+=%.4f PT+=%.4f\n",
				nodeRaw, nodeIndex, s.QueueingNTime, s.ProcessingTime)
		}

		// Log tổng kết cho mỗi node
		for i := 0; i < numPods; i++ {
			fmt.Printf("[RL][SUMMARY] node%d: QueueingN=%.4f, Processing=%.4f\n",
				i+1, QueueingNTime[i], ProcessingTime[i])
		}

		fmt.Println("[RL] RL computation loop done.")
	}
}
