package outoushuugou // 応答集合 - おうとうしゅうごう - Response Pool

import (
	"github.com/bonavadeur/katyusha/pkg/bonalib"
	_ "github.com/bonavadeur/katyusha/pkg/global"
	"time"

)

const (
	BASE_PATH = "/katyusha/outoushuugou"
	SCRAPE_INTERVAL = 10 * time.Second
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
		// go monitor()
	}
}

// func monitor() {
// 	for POOL == nil {
// 		time.Sleep(100 * time.Millisecond)
// 	}
// 	POOL.StartPrometheusExporter(SCRAPE_INTERVAL)
// 	StartPrometheusServer()
// }