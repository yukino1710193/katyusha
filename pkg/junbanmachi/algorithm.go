package junbanmachi // 順番待ち - じゅんばんまち - Queuing

import "github.com/bonavadeur/katyusha/pkg/bonalib"

func (q *ExtraQueue) SortAlgorithm(p *Packet) {
	bonalib.Info("SortAlgorithm", "Packet", p)
	// example of adding header
	p.Headers = append(p.Headers, &PushRequest_HeaderSchema{
		Field: "Queue-J-field-1",
		Value: "Nước mắt em rơi",
	})

	q.Queue = append([]*Packet{p}, q.Queue...)
}
