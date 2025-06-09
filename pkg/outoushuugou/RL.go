package outoushuugou

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bonavadeur/katyusha/pkg/bonalib"
	"github.com/bonavadeur/katyusha/pkg/fukabunsan"
)

func (rp *ResponsePool) RL(interval time.Duration) {
	for {
		time.Sleep(interval)

		// bonalib.Log("[RL] Bắt đầu vòng tính toán RL")

		numNodes := len(fukabunsan.NODENAMES)
		if numNodes == 0 {
			bonalib.Log("[RL][WARN] Không có NODE nào, bỏ qua vòng tính toán.")
			continue
		}

		QueueingNTime := make([]float64, numNodes)
		ProcessingTime := make([]float64, numNodes)

		for _, s := range DATA {
			nodeRaw := s.NodeD
			nodeStr := strings.TrimPrefix(nodeRaw, "node")
			nodeIndex, err := strconv.Atoi(nodeStr)

			if err != nil {
				bonalib.Log(fmt.Sprintf("[RL][WARN] Invalid NodeD format: '%s', err: %v", nodeRaw, err))
				continue
			}

			nodeIndex -= 1
			if nodeIndex < 0 || nodeIndex >= numNodes {
				bonalib.Log(fmt.Sprintf("[RL][WARN] Node index out of range: %d (from '%s'), NODE size = %d", nodeIndex, nodeRaw, numNodes))
				continue
			}

			QueueingNTime[nodeIndex] += s.QueueingNTime
			ProcessingTime[nodeIndex] += s.ProcessingTime

			// DEBUG chi tiết, có thể bật nếu cần
			// bonalib.Log(fmt.Sprintf("[RL][DEBUG] node=%s index=%d QN+=%.4f PT+=%.4f",
			// 	nodeRaw, nodeIndex, s.QueueingNTime, s.ProcessingTime))
		}

		// Log tổng kết cho mỗi node
		for i := 0; i < numNodes; i++ {
			bonalib.Info(fmt.Sprintf("[RL][SUMMARY] node%d: QueueingN=%.4f, Processing=%.4f",
				i+1, QueueingNTime[i], ProcessingTime[i]))
		}

		bonalib.Log("[RL] Kết thúc vòng tính toán RL")
		clearDataPOOL()
	}
}
