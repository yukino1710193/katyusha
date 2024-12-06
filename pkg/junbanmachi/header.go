package junbanmachi // 順番待ち - じゅんばんまち - Queuing

import "github.com/bonavadeur/katyusha/pkg/bonalib"

func (q *ExtraQueue) HeaderModifier(p *Packet) {
	bonalib.Info("HeaderModifier", "Packet", p)
	// example of adding header
	p.Headers = append(p.Headers, &PushRequest_HeaderSchema{
		Field: "Katyusha-J-Field-2",
		Value: "Katyusha-J-Value-2",
	})
}
