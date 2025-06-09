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
	// bonalib.Log("[LBAlgorithm] Nhận request:", lbRequest)

	srcIP := strings.Split(lbRequest.SourceIP, ":")[0]
	node_Source := IPfromNode(srcIP)
	// bonalib.Log("[LBAlgorithm] Source IP:", srcIP, "-> Node:", node_Source)

	node_Source_STT, err := strconv.Atoi(strings.TrimPrefix(node_Source, "node"))
	if err != nil || node_Source_STT <= 0 {
		bonalib.Log("[LBAlgorithm] ❌ Lỗi khi parse node_Source:", node_Source, "err:", err)
		return nil
	}

	if node_Source_STT-1 >= len(MIPORIN_matrix) {
		bonalib.Log("[LBAlgorithm] ❌ node_Source_STT vượt giới hạn MIPORIN_matrix")
		bonalib.Log("[LBAlgorithm] node_Source_STT:", node_Source_STT, "MIPORIN_matrix length:", len(MIPORIN_matrix))
		return nil
	}

	node_target := Choose(MIPORIN_matrix[node_Source_STT-1])
	if node_target < 0 || node_target >= len(PODCIDRS) {
		bonalib.Log("[LBAlgorithm] ❌ node_target không hợp lệ:", node_target)
		return nil
	}
	// bonalib.Log("[LBAlgorithm] node_target chọn:", node_target)

	var selectedTargets []string
	for _, target := range lbRequest.Targets {
		if IsPodinPodcidr(target, PODCIDRS[node_target]) {
			selectedTargets = append(selectedTargets, target)
		}
	}

	if len(selectedTargets) == 0 {
		bonalib.Log("[LBAlgorithm] ❌ Không tìm thấy target phù hợp cho node", node_target)
		return nil
	}

	result := rand.Intn(len(selectedTargets))
	selected := selectedTargets[result]
	// bonalib.Log("[LBAlgorithm] Chọn target:", selected)

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
	// bonalib.Log("[LBAlgorithm] Hoàn tất request")
	return ret
}
