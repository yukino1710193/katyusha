package junbanmachi // 順番待ち - じゅんばんまち - Queuing

import (
	"time"
	"strconv"
	"github.com/bonavadeur/katyusha/pkg/bonalib"
	"github.com/bonavadeur/katyusha/pkg/global"
)

func (q *ExtraQueue) SortAlgorithm(p *Packet) {
	bonalib.Info("SortAlgorithm", "Packet", p)
	// example of adding header
	p.Headers = append(p.Headers, &PushRequest_HeaderSchema{
		Field: "InComing-J-moment",
		Value: time.Now().Format(time.RFC3339Nano),
	},&PushRequest_HeaderSchema{
		Field: "Queue-J-length",
		Value: strconv.Itoa(len(q.Queue)),
	})

	q.Queue = append([]*Packet{p}, q.Queue...)
	global.IncIncoming()
}
