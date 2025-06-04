package outoushuugou // 応答集合 - おうとうしゅうごう - Response Pool

import (
	"strconv"
	"strings"
	"time"

	"github.com/bonavadeur/katyusha/pkg/bonalib"
	"github.com/bonavadeur/katyusha/pkg/fukabunsan"
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
		bonalib.Log("[parseRFC3339Nano] Lỗi khi parse:", ts, err)
		return time.Time{}
	}
	return t
}

func parseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		bonalib.Log("[parseFloat] Lỗi khi parse:", s, err)
	}
	return f
}

func GetColdStartPods() int {
	// TODO: Cập nhật logic khi có metric từ miporin
	return 0
}

func appendMetricFromFeedback(fb *ResponseFeedback, data *[]*Metric) {
	bonalib.Log("[appendMetricFromFeedback] Bắt đầu phân tích metric cho request:", fb.ID)

	// In toàn bộ headers
	for _, h := range fb.Headers {
		bonalib.Log("[appendMetricFromFeedback][Header] field:", h.Field, ", value:", h.Value)
	}
	h := headerToMap(fb.Headers)

	tIncomingJMomentResponsed := parseRFC3339Nano(h["Incoming-J-Moment-Responsed"])
	QueueJLengthResponsed := parseFloat(h["Queue-J-Length-Responsed"])
	tOutcomingJMomentResponsed := parseRFC3339Nano(h["Outcoming-J-Moment-Responsed"])
	tIncomingNMoment := parseRFC3339Nano(h["Incoming-N-Moment"])
	tOutcomingNMoment := parseRFC3339Nano(h["Outcoming-N-Moment"])
	tProcess := parseFloat(strings.TrimSuffix(h["Shuka-Processing-Time"], "s"))

	queueingJTime := tOutcomingJMomentResponsed.Sub(tIncomingJMomentResponsed).Seconds()
	queueingNTime := tOutcomingNMoment.Sub(tIncomingNMoment).Seconds()

	IP_dest, ok := fb.GetHeader("Ip-Destination-Responsed")
	if !ok || IP_dest == "" {
		bonalib.Log("[appendMetricFromFeedback] Thiếu header Ip-Destination-Responsed")
	}

	node := fukabunsan.IPfromNode(IP_dest)
	bonalib.Log("[appendMetricFromFeedback] Source:", fb.SourceIP, ", Dest:", IP_dest, ", Node:", node)

	metric := &Metric{
		ID:              fb.ID,
		SourceIP:        fb.SourceIP,
		Domain:          fb.Domain,
		URI:             fb.URI,
		Method:          fb.Method,
		DestIP:          IP_dest,
		NodeD:           node,
		ProcessingTime:  tProcess,
		QueueingNTime:   queueingNTime,
		QueueingJTime:   queueingJTime,
		QueueingJLength: QueueJLengthResponsed,
	}

	*data = append(*data, metric)
	bonalib.Log("[appendMetricFromFeedback] Metric ghi nhận:", metric)
}
