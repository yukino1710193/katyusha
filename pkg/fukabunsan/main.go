package fukabunsan // 負荷分散 - ふかぶんさん - Load Balancing

import (
	"github.com/bonavadeur/katyusha/pkg/bonalib"
	_ "github.com/bonavadeur/katyusha/pkg/global"
)

const (
	BASE_PATH   = "/katyusha/fukabunsan"
	MIPORIN_URL = "http://miporin.knative-serving.svc.cluster.local/api/weight/okasan/okaasan/kodomo/hello"
	ALL         = -1
)

var (
	KUBECONFIG     = GetKubeconfig()
	CLIENTSET      = GetClientSet()
	NODENAMES      = GetNodenames()
	PODCIDRS       = GetPodsCIDRs()
	STATE          = NewStateManager()
	POD_READY      = STATE.GetReadyPods(ALL)
	LB             *LoadBalancer
	MIPORIN_matrix [][]int32
	RL_RATE        []int32
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
