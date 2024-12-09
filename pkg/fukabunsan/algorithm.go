package fukabunsan // 負荷分散 - ふかぶんさん - Load Balancing

import (
	"github.com/bonavadeur/katyusha/pkg/bonalib"
	// "time"
)

func (lb *LoadBalancer) LBAlgorithm(lbRequest *LBRequest) *LBResponse {
	bonalib.Info("LBAlgorithm", "lbRequest", lbRequest)

	ret := &LBResponse{
		Target:  lbRequest.Targets[0],
		Headers: make([]*LBResponse_HeaderSchema, 0),
	}
	ret.Headers = append(ret.Headers, &LBResponse_HeaderSchema{
		Field: "LB-F-Field",
		Value: "Trái",
	})

	// time.Sleep(0*time.Second) // mô phỏng thời gian xử lý
	return ret
}
