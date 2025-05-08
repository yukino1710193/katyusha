package junbanmachi // 順番待ち - じゅんばんまち - Queuing

import (
	"time"

	_ "github.com/bonavadeur/katyusha/pkg/bonalib"
)

func (q *ExtraQueue) HeaderModifier(p *Packet) {
	// example of adding header
	p.Headers = append(p.Headers, &PushRequest_HeaderSchema{
		Field: "OutComing-J-moment",
		Value: time.Now().Format(time.RFC3339Nano),
	})
}
