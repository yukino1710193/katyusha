package junbanmachi // 順番待ち - じゅんばんまち - Queuing

import "github.com/yukino1710193/katyusha/pkg/bonalib"

func (q *ExtraQueue) HeaderModifier(p *Packet) {
	bonalib.Info("HeaderModifier", "Packet", p)
	// example of adding header
	p.Headers = append(p.Headers, &PushRequest_HeaderSchema{
		Field: "Queue-J-Field-2",
		Value: "Trò chơi kết thúc",
	})
}
