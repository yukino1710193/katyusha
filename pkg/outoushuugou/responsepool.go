package outoushuugou // 応答集合 - おうとうしゅうごう - Response Pool

import (
	reflect "reflect"
	"sync"

	"github.com/bonavadeur/katyusha/pkg/bonalib"
	"github.com/bonavadeur/katyusha/pkg/hashi"
)

type ResponsePool struct {
	responseBridge    *hashi.Hashi
	Pool              []*ResponseFeedback
	PoolAppendingLock *sync.Mutex
}

func NewResponsePool() *ResponsePool {
	newResponsePool := &ResponsePool{
		Pool:              make([]*ResponseFeedback, 0),
		PoolAppendingLock: &sync.Mutex{},
	}

	newResponsePool.responseBridge = hashi.NewHashi(
		"responsePoolBridge",
		hashi.HASHI_TYPE_SERVER,
		BASE_PATH+"/response-pool-bridge",
		bonalib.Cm2Int("katyusha-threads"),
		reflect.TypeOf(ResponseFeedback{}),
		reflect.TypeOf(ResponseConfirm{}),
		newResponsePool.ResponsePoolAdapter,
	)

	return newResponsePool
}

func (rp *ResponsePool) ResponsePoolAdapter(params ...interface{}) (interface{}, error) {
	feedback := params[0].(*ResponseFeedback)

	rp.PoolAppendingLock.Lock()
	bonalib.Info("ResponsePoolAdapter", feedback)
	rp.Pool = append([]*ResponseFeedback{feedback}, rp.Pool...)
	rp.PoolAppendingLock.Unlock()

	return &ResponseConfirm{SymbolizeResponse: Status_Success}, nil
}
