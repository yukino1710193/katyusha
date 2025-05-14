package outoushuugou // 応答集合 - おうとうしゅうごう - Response Pool

import (
	"strconv"
	"time"

	// "github.com/bonavadeur/katyusha/pkg/global"
)

func (p *ResponseFeedback) GetHeader(key string) (string, bool) {
	for _, h := range p.Headers {
		if h.Field == key {
			return h.Value, true
		}
	}
	return "", false
}

func headerToMap(headers []*ResponseFeedback_HeaderSchema) map[string]string {
	m := make(map[string]string)
	for _, h := range headers {
		m[h.Field] = h.Value
	}
	return m
}

func parseRFC3339Nano(ts string) time.Time {
	t, err := time.Parse(time.RFC3339Nano, ts)
	if err != nil {
		return time.Time{}
	}
	return t
}

func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func GetColdStartPods() int{
	// Lấy giá trị cold-start-pods từ miporin
	return 0
}

func appendMetricFromFeedback(fb *ResponseFeedback, data *[]*Metric) {
	h := headerToMap(fb.Headers)


	// Parse timestamp từ header
	tIncomingJMomentResponsed := parseRFC3339Nano(h["Incoming-J-Moment-Responsed"])
	QueueJLengthResponsed := parseFloat(h["Queue-J-Length-Responsed"])
	tOutcomingJMomentResponsed := parseRFC3339Nano(h["Outcoming-J-Moment-Responsed"])
	// tLbMommentResponsed := parseRFC3339Nano(h["Lb-Momment-Responsed"])
	tIncomingNMoment := parseRFC3339Nano(h["Incoming-N-Moment"])
	tOutcomingNMoment := parseRFC3339Nano(h["Outcoming-N-Moment"])
	
	tProcess := parseFloat(h["Shuka-Processing-Time"])


	// Tính toán (thời gian chênh lệch, đơn vị giây)
	queueingJTime := tOutcomingJMomentResponsed.Sub(tIncomingJMomentResponsed).Seconds()
	queueingNTime := tOutcomingNMoment.Sub(tIncomingNMoment).Seconds()
	// totaltime := time.Now().Sub(tLbMommentResponsed).Seconds()



	metric := &Metric{
		ProcessingTime: tProcess,
		QueueingNTime:   queueingNTime,
		QueueingJTime: queueingJTime,
		QueueingJLength: QueueJLengthResponsed,
		// Incoming:       global.GetIncoming(),
		// Outgoing:       global.GetOutgoing(),
		
	}

	*data = append(*data, metric)
}