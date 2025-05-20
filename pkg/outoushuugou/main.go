package outoushuugou // 応答集合 - おうとうしゅうごう - Response Pool

import (
	"net/http"
	"time"
	"github.com/bonavadeur/katyusha/pkg/fukabunsan"
	"github.com/bonavadeur/katyusha/pkg/bonalib"
	_ "github.com/bonavadeur/katyusha/pkg/global"
	"github.com/labstack/echo/v4"
)

const (
	BASE_PATH       = "/katyusha/outoushuugou"
	SCRAPE_INTERVAL = 10 * time.Second
)

var (
	POOL *ResponsePool
	DATA []*Metric
	RL_p []float64
)

func init() {
	RL_p = make([]float64,len(fukabunsan.NODENAMES))

	enableKatyusha := bonalib.Cm2Bool("ikukantai-enable-katyusha")
	enableKatyushaOutoushuugou := bonalib.Cm2Bool("katyusha-enable-outoushuugou")

	if enableKatyusha && enableKatyushaOutoushuugou {
		bonalib.Log("Outoushuugou is enabled")
		
		go func() {
			POOL = NewResponsePool()
		}()
		go POOL.RL(SCRAPE_INTERVAL)
		go server()
	}
}

func server() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	e.POST("/clear", func(c echo.Context) error {
		clearDataPOOL()
		return c.String(http.StatusOK, "Pool and DATA has been cleared")
	})
	e.GET("/RL", func(c echo.Context) error {
		return c.JSON(http.StatusOK, RL_p);
	})

	e.Logger.Fatal(e.Start(":19090"))
}

// func monitor() {
// 	for POOL == nil {
// 		time.Sleep(100 * time.Millisecond)
// 	}
// 	POOL.StartPrometheusExporter(SCRAPE_INTERVAL)
// 	StartPrometheusServer()
// }
