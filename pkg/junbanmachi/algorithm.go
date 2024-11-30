package junbanmachi // 順番待ち - じゅんばんまち - Queuing

func (q *ExtraQueue) SortAlgorithm(p *Packet) {
	// example of adding header
	p.Headers = append(p.Headers, &PushRequest_HeaderSchema{
		Field: "J-naniField",
		Value: "J-naniValue",
	})

	q.Queue = append([]*Packet{p}, q.Queue...)
}
