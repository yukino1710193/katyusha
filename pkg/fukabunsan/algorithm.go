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
	// random region Edge or cloud

	// 	percentice := RL_RATE
	// 	var region_target = Choose(percentice)
	// 	bonalib.Log("Request will be forward to :",region_target,"region")

	// //
	// 	var selectedTargets []PodState = STATE.GetReadyPods(region_target)
	// 	var result = rand.Intn(len(selectedTargets))
	// random node by rate from MIPORIN
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
		Value: time.Now().Format(time.RFC3339),
	})
	global.IncOutgoing()
	return ret
}
