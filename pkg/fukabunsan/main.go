package fukabunsan // 負荷分散 - ふかぶんさん - Load Balancing

import "github.com/bonavadeur/katyusha/pkg/bonalib"

const (
	BASE_PATH = "/katyusha/fukabunsan"
	MIPORIN_URL = "http://miporin.knative-serving.svc.cluster.local/api/weight/okasan/okaasan/kodomo/hello"
)

var (
	LB *LoadBalancer
	MIPORIN_matrix [][]int32
)

func init() {
	enableKatyusha := bonalib.Cm2Bool("ikukantai-enable-katyusha")
	enableKatyushaFukabunsan := bonalib.Cm2Bool("katyusha-enable-fukabunsan")

	if enableKatyusha && enableKatyushaFukabunsan {
		bonalib.Log("Fukabunsan is enabled")
		go startPeriodicTask()
		go func() {
			LB = NewLoadBalancer()
		}()
	}
}
