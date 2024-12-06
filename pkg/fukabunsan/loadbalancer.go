package fukabunsan // 負荷分散 - ふかぶんさん - Load Balancing

import (
	"reflect"

	"github.com/bonavadeur/katyusha/pkg/bonalib"
	"github.com/bonavadeur/katyusha/pkg/hashi"
)

type LoadBalancer struct {
	lbBridge *hashi.Hashi
}

func NewLoadBalancer() *LoadBalancer {
	newLoadBalancerServer := &LoadBalancer{}

	newLoadBalancerServer.lbBridge = hashi.NewHashi(
		"lbBridge",
		hashi.HASHI_TYPE_SERVER,
		BASE_PATH+"/lb-bridge",
		bonalib.Cm2Int("katyusha-threads"),
		reflect.TypeOf(LBRequest{}),
		reflect.TypeOf(LBResponse{}),
		newLoadBalancerServer.LBResponseAdapter,
	)

	return newLoadBalancerServer
}

func (lb *LoadBalancer) LBResponseAdapter(params ...interface{}) (interface{}, error) {
	lbRequest := params[0].(*LBRequest)

	response := lb.loadBalance(lbRequest)

	return response, nil
}

func (lb *LoadBalancer) loadBalance(lbRequest *LBRequest) *LBResponse {
	response := lb.LBAlgorithm(lbRequest)
	return response
}
