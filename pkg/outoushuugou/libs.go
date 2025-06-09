package outoushuugou // 応答集合 - おうとうしゅうごう - Response Pool

import (
	"fmt"
	"os/exec"
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

func parseRFC3339Nano(ts string) (time.Time, error) {
	if ts == "" {
		err := fmt.Errorf("chuỗi rỗng")
		bonalib.Log("[parseRFC3339Nano] ❌ Lỗi khi parse:", ts, err)
		return time.Time{}, err
	}

	t, err := time.Parse(time.RFC3339Nano, ts)
	if err != nil {
		bonalib.Log("[parseRFC3339Nano] ❌ Lỗi khi parse:", ts, err)
		return time.Time{}, err
	}
	// bonalib.Log("[parseRFC3339Nano] ✅ Parse thành công:", t)
	return t, nil
}

func parseFloat(s string) (float64, error) {
	if s == "" {
		err := fmt.Errorf("chuỗi rỗng")
		bonalib.Log("[parseFloat] ❌ Lỗi khi parse:", s, err)
		return 0, err
	}

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		bonalib.Log("[parseFloat] ❌ Lỗi khi parse:", s, err)
		return 0, err
	}
	// bonalib.Log("[parseFloat] ✅ Parse thành công:", f)
	return f, nil
}

func GetColdStartPods() int {
	// TODO: Cập nhật logic khi có metric từ miporin
	return 0
}

func appendMetricFromFeedback(fb *ResponseFeedback, data *[]*Metric) {
	// bonalib.Log("[appendMetricFromFeedback] Bắt đầu phân tích metric cho request:", fb.ID)

	// for _, h := range fb.Headers {
	// 	bonalib.Log("[appendMetricFromFeedback][Header] field:", h.Field, ", value:", h.Value)
	// }
	h := headerToMap(fb.Headers)

	// Parse các trường
	tIncomingJ, err1 := parseRFC3339Nano(h["Incoming-J-Moment-Responsed"])
	queueJLen, err2 := parseFloat(h["Queue-J-Length-Responsed"])
	tOutcomingJ, err3 := parseRFC3339Nano(h["Outcoming-J-Moment-Responsed"])
	tIncomingN, err4 := parseRFC3339Nano(h["Incoming-N-Moment"])
	tOutcomingN, err5 := parseRFC3339Nano(h["Outcoming-N-Moment"])
	tProcess, err6 := parseFloat(strings.TrimSuffix(h["Shuka-Processing-Time"], "s"))

	// Gom lỗi vào map
	errMap := map[string]error{
		"Incoming-J-Moment-Responsed":  err1,
		"Queue-J-Length-Responsed":     err2,
		"Outcoming-J-Moment-Responsed": err3,
		"Incoming-N-Moment":            err4,
		"Outcoming-N-Moment":           err5,
		"Shuka-Processing-Time":        err6,
	}

	// Check lỗi gọn gàng
	hasError := false
	for field, err := range errMap {
		if err != nil {
			hasError = true
			bonalib.Log("[appendMetricFromFeedback] ❌ Parse lỗi tại", field, ":", err)
		}
	}
	if hasError {
		bonalib.Log("[appendMetricFromFeedback] ⛔ Metric bị loại vì parse lỗi.")
		return
	}

	queueingJTime := tOutcomingJ.Sub(tIncomingJ).Seconds()
	queueingNTime := tOutcomingN.Sub(tIncomingN).Seconds()

	IP_dest, ok := fb.GetHeader("Ip-Destination-Responsed")
	if !ok || IP_dest == "" {
		bonalib.Log("[appendMetricFromFeedback] ❌ Thiếu header Ip-Destination-Responsed → metric bị loại")
		return
	}

	node := fukabunsan.IPfromNode(IP_dest)
	// bonalib.Log("[appendMetricFromFeedback] Source:", fb.SourceIP, ", Dest:", IP_dest, ", Node:", node)

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
		QueueingJLength: queueJLen,
	}

	*data = append(*data, metric)
	bonalib.Log("[appendMetricFromFeedback] ✅ Metric ghi nhận:", metric)
}

func clearScreen(mode int) {
	switch mode {
	case 1:
		cmd := exec.Command("clear")
		cmd.Stdout = nil // không in ra stdout của app
		_ = cmd.Run()
	case 2:
		fmt.Print("\033[H\033[2J")
	default:
		fmt.Println("Invalid clear mode")
	}
}

