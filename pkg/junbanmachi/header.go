package junbanmachi // 順番待ち - じゅんばんまち - Queuing

func (q *ExtraQueue) HeaderModifier(p *Packet) {
	// example of adding header
	p.Headers = append(p.Headers, &PushRequest_HeaderSchema{
		Field: "J-naniField-Pop",
		Value: "J-naniValue-Pop",
	})
}
