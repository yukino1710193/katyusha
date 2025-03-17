package outoushuugou // 応答集合 - おうとうしゅうごう - Response Pool

import "github.com/yukino1710193/katyusha/pkg/bonalib"

const (
	BASE_PATH = "/katyusha/outoushuugou"
)

var (
	POOL *ResponsePool
)

func init() {
	enableKatyusha := bonalib.Cm2Bool("ikukantai-enable-katyusha")
	enableKatyushaOutoushuugou := bonalib.Cm2Bool("katyusha-enable-outoushuugou")

	if enableKatyusha && enableKatyushaOutoushuugou {
		bonalib.Log("Outoushuugou is enabled")
		go func() {
			POOL = NewResponsePool()
		}()
	}
}
