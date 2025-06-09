package outoushuugou

import (
	"github.com/bonavadeur/katyusha/pkg/bonalib"
)

type Metric struct {
	ID              uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	SourceIP        string `protobuf:"bytes,2,opt,name=SourceIP,proto3" json:"SourceIP,omitempty"`
	Domain          string `protobuf:"bytes,3,opt,name=Domain,proto3" json:"Domain,omitempty"`
	URI             string `protobuf:"bytes,4,opt,name=URI,proto3" json:"URI,omitempty"`
	Method          string `protobuf:"bytes,5,opt,name=Method,proto3" json:"Method,omitempty"`
	DestIP          string
	NodeD           string
	ProcessingTime  float64
	QueueingNTime   float64
	QueueingJTime   float64
	QueueingJLength float64
}

func clearDataPOOL() {
	bonalib.Log("[clearDataPOOL] Bắt đầu clear POOL")

	POOL.PoolAppendingLock.Lock()
	defer POOL.PoolAppendingLock.Unlock()

	POOL.Pool = make([]*ResponseFeedback, 0)
	POOL.lastExportedCount = 0
	DATA = nil
}
