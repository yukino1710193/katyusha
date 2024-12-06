package fukabunsan // 負荷分散 - ふかぶんさん - Load Balancing

import (
	"github.com/bonavadeur/katyusha/pkg/bonalib"
)

func (lb *LoadBalancer) LBAlgorithm(lbRequest *LBRequest) *LBResponse {
	bonalib.Info("LBAlgorithm", "lbRequest", lbRequest)
	// return lbRequest.Targets[0]
	ret := &LBResponse{
		Target:  lbRequest.Targets[0],
		Headers: make([]*LBResponse_HeaderSchema, 0),
	}
	ret.Headers = append(ret.Headers, &LBResponse_HeaderSchema{
		Field: "Katyusha-F-Field",
		Value: "Katyusha-F-Field",
	})

	return ret
}
