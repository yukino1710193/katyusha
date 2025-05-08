package outoushuugou // 応答集合 - おうとうしゅうごう - Response Pool

func (p *ResponseFeedback) GetHeader(key string) (string, bool) {
	for _, h := range p.Headers {
		if h.Field == key {
			return h.Value, true
		}
	}
	return "", false
}