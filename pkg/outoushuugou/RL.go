package outoushuugou

import (
	"time"
	"strconv"
	"strings"
	// "github.com/bonavadeur/katyusha/pkg/bonalib"
	_ "github.com/bonavadeur/katyusha/pkg/global"
	"github.com/bonavadeur/katyusha/pkg/fukabunsan"
)

func (rp *ResponsePool) RL(interval time.Duration) {
	go func() {
		for {
			time.Sleep(interval)
			// Tính toán RL_p
			QueueingNTime := make([]float64, len(fukabunsan.POD_READY))
			ProcessingTime := make([]float64, len(fukabunsan.POD_READY))
			for _ , s := range DATA {
				node_Dest_STT, _ := strconv.Atoi(strings.TrimPrefix(s.NodeD, "node"))
				node_Dest_STT = node_Dest_STT - 1
				QueueingNTime[node_Dest_STT] += float64(s.QueueingNTime)
			}
			for _, s := range DATA {
				node_Dest_STT, _ := strconv.Atoi(strings.TrimPrefix(s.NodeD, "node"))
				node_Dest_STT = node_Dest_STT - 1
				ProcessingTime[node_Dest_STT] += float64(s.ProcessingTime)
			}


			// và cập nhật lại giá trị cho RL_p

		}
	}()
}