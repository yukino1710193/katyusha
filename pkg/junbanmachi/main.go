package junbanmachi // 順番待ち - じゅんばんまち - Queuing

import "github.com/bonavadeur/katyusha/pkg/bonalib"

const (
	BASE_PATH = "/katyusha/junbanmachi"
)

var (
	QUEUE *ExtraQueue
)

func init() {
	enableKatyusha := bonalib.Cm2Bool("ikukantai-enable-katyusha")
	enableKatyushaJunbanmachi := bonalib.Cm2Bool("katyusha-enable-junbanmachi")

	if enableKatyusha && enableKatyushaJunbanmachi {
		bonalib.Log("Junbanmachi is enabled")
		go func() {
			QUEUE = NewExtraQueue()
		}()
	}
}
