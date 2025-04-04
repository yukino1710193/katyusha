package global

import "sync/atomic"

var INCOMING int32
var OUTCOMING int32

func IncIncoming() {
	atomic.AddInt32(&INCOMING, 1)
}

func IncOutgoing() {
	atomic.AddInt32(&OUTCOMING, 1)
}

func GetIncoming() int32 {
	return atomic.LoadInt32(&INCOMING)
}

func GetOutgoing() int32 {
	return atomic.LoadInt32(&OUTCOMING)
}

func ResetCounters() {
	atomic.StoreInt32(&INCOMING, 0)
	atomic.StoreInt32(&OUTCOMING, 0)
}