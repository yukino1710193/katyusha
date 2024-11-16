package fukabunsan // 負荷分散 - ふかぶんさん - Load Balancing

import "github.com/bonavadeur/katyusha/pkg/bonalib"

const (
	BASE_PATH = "/katyusha/fukabunsan"
)

var (
	LB *LoadBalancer
)

func init() {
	enableKatyusha := bonalib.Cm2Bool("ikukantai-enable-katyusha")
	enableKatyushaFukabunsan := bonalib.Cm2Bool("katyusha-enable-fukabunsan")

	if enableKatyusha && enableKatyushaFukabunsan {
		bonalib.Log("Fukabunsan is enabled")
		go func() {
			LB = NewLoadBalancer()
		}()
	}
}
