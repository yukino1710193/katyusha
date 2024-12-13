package fukabunsan // 負荷分散 - ふかぶんさん - Load Balancing

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/bonavadeur/katyusha/pkg/bonalib"
	// "github.com/bonavadeur/katyusha/pkg/bonalib"
	// "time"
)

func (lb *LoadBalancer) LBAlgorithm(lbRequest *LBRequest) *LBResponse {
	bonalib.Log("LBAlgorithm : ", lbRequest)
	var node_Source = IPfromNode(strings.Split(lbRequest.SourceIP, ":")[0])
	// bonalib.Info("Node gửi: ", node_Source)
	node_Source_STT, _ := strconv.Atoi(strings.TrimPrefix(node_Source, "node"))
	// bonalib.Line()
	// bonalib.Info("Node gửi (STT): ", node_Source_STT)
	// bonalib.Line()
	// bonalib.Log("MIPORIN_matrix: ", MIPORIN_matrix)
	// bonalib.Info("Trọng số bắn tải: ", MIPORIN_matrix[node_Source_STT-1])
	// Gacha node để bắn tải - theo trọng số
	random := rand.Intn(100)

	var node_target = gachaNodeTarget(random, MIPORIN_matrix[node_Source_STT-1])
	// đã chọn xong node và lưu node được chọn vào nodeTarget
	// bonalib.Log("Node được chọn: ", node_target+1)

	// Chia các targets của lbRequest vào các node
	var selectedTargets []string
	for _, target := range lbRequest.Targets {
		if IsPodinPodcidr(target, PODCIDRS[node_target]) {
			selectedTargets = append(selectedTargets, target)
		}
	}
	random = rand.Intn(len(selectedTargets))
	// bonalib.Log("Các target thuộc node được chọn: ", selectedTargets)
	// bonalib.Log("Số lượng target thuộc node được chọn: ", len(selectedTargets), " - Random: ", random)
	// bonalib.Line()
	// bonalib.Log("Targets được chọn: ", selectedTargets[random])

	ret := &LBResponse{
		Target:  selectedTargets[random],
		Headers: make([]*LBResponse_HeaderSchema, 0),
	}
	ret.Headers = append(ret.Headers, &LBResponse_HeaderSchema{
		Field: "LB-F-Field",
		Value: "Trái",
	})

	// time.Sleep(0*time.Second) // mô phỏng thời gian xử lý
	return ret
}
