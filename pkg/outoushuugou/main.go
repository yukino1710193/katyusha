package outoushuugou // 応答集合 - おうとうしゅうごう - Response Pool

import (
	"github.com/bonavadeur/katyusha/pkg/bonalib"
	_ "github.com/bonavadeur/katyusha/pkg/global"
)

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
