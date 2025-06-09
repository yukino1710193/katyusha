package fukabunsan // 負荷分散 - ふかぶんさん - Load Balancing

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bonavadeur/katyusha/pkg/bonalib"
	"github.com/bonavadeur/katyusha/pkg/global"
)

func (lb *LoadBalancer) LBAlgorithm(lbRequest *LBRequest) *LBResponse {
	const retryDelay = 200 * time.Millisecond
	attempt := 0

	for {
		attempt++
		bonalib.Log("[LBAlgorithm] 🌀 Attempt", attempt, "Nhận request:", lbRequest)

		srcIP := strings.Split(lbRequest.SourceIP, ":")[0]
		node_Source := IPfromNode(srcIP)

		node_Source_STT, err := strconv.Atoi(strings.TrimPrefix(node_Source, "node"))
		if err != nil || node_Source_STT <= 0 {
			bonalib.Warn("[LBAlgorithm] ❌ Lỗi parse node_Source:", node_Source, "err:", err)
			gotoDelay(attempt, retryDelay)
			continue
		}

		if node_Source_STT-1 >= len(MIPORIN_matrix) {
			bonalib.Warn("[LBAlgorithm] ❌ node_Source_STT vượt giới hạn MIPORIN_matrix:",
				"STT =", node_Source_STT, "Len =", len(MIPORIN_matrix))
			gotoDelay(attempt, retryDelay)
			continue
		}

		node_target := Choose(MIPORIN_matrix[node_Source_STT-1])
		if node_target < 0 || node_target >= len(PODCIDRS) {
			bonalib.Warn("[LBAlgorithm] ❌ node_target không hợp lệ:", node_target)
			gotoDelay(attempt, retryDelay)
			continue
		}

		var selectedTargets []string
		for _, target := range lbRequest.Targets {
			if IsPodinPodcidr(target, PODCIDRS[node_target]) {
				selectedTargets = append(selectedTargets, target)
			}
		}

		if len(selectedTargets) == 0 {
			bonalib.Warn("[LBAlgorithm] ❌ Không tìm thấy target hợp lệ cho node", node_target)
			gotoDelay(attempt, retryDelay)
			continue
		}

		selected := selectedTargets[rand.Intn(len(selectedTargets))]

		ret := &LBResponse{
			Target: selected,
			Headers: []*LBResponse_HeaderSchema{
				{
					Field: "LB-Momment",
					Value: time.Now().Format(time.RFC3339Nano),
				},
				{
					Field: "Ip-Destination",
					Value: selected,
				},
			},
		}

		global.IncOutgoing()
		bonalib.Log("[LBAlgorithm] ✅ Thành công tại attempt", attempt)
		return ret
	}
}
