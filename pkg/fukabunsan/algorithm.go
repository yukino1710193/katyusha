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
	bonalib.Log("LBAlgorithm : ", lbRequest)
	var node_Source = IPfromNode(strings.Split(lbRequest.SourceIP, ":")[0])

	node_Source_STT, _ := strconv.Atoi(strings.TrimPrefix(node_Source, "node"))

	var node_target = Choose(MIPORIN_matrix[node_Source_STT-1])

	var selectedTargets []string
	for _, target := range lbRequest.Targets {
		if IsPodinPodcidr(target, PODCIDRS[node_target]) {
			selectedTargets = append(selectedTargets, target)
		}
	}
	var result = rand.Intn(len(selectedTargets))
	/// end
	ret := &LBResponse{
		// Target:  selectedTargets[result].IP,
		Target:  selectedTargets[result],
		Headers: make([]*LBResponse_HeaderSchema, 0),
	}
	ret.Headers = append(ret.Headers, &LBResponse_HeaderSchema{
		Field: "LB-Momment",
		Value: time.Now().Format(time.RFC3339Nano),
	}, &LBResponse_HeaderSchema{
		Field: "Ip-Destination",
		Value: selectedTargets[result],
	})
	global.IncOutgoing()
	return ret
}
