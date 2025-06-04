package fukabunsan

import (
	"strings"
	"sync"
	"time"

	"github.com/bonavadeur/katyusha/pkg/bonalib"
)

// PodStatus định nghĩa các trạng thái chuẩn hóa
const (
	StatusNull  = "Null"
	StatusCold  = "Cold"
	StatusWarm  = "Warm"
	StatusReady = "Ready"
)

type PodState struct {
	PodName    string
	Region     int
	IP         string
	Status     string
	LastUpdate time.Time
}

type StateManager struct {
	mu   sync.RWMutex
	pods map[string]PodState
}

type Adapter interface {
	Apply(newPod PodState, state map[string]PodState)
}

type AdapterReplace struct{}
type AdapterIgnore struct{}
type AdapterRemove struct{}

func NewStateManager() *StateManager {
	bonalib.Log("[NewStateManager] Khởi tạo StateManager")
	return &StateManager{
		pods: make(map[string]PodState),
	}
}

func (s *StateManager) UpdatePodState(p PodState) {
	s.mu.Lock()
	defer s.mu.Unlock()

	p.LastUpdate = time.Now()
	s.pods[p.PodName] = p
	bonalib.Log("[UpdatePodState] Cập nhật pod:", p.PodName, "Status:", p.Status)
}

func (s *StateManager) GetReadyPods(region int) []PodState {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []PodState
	for _, pod := range s.pods {
		if strings.EqualFold(pod.Status, StatusReady) && (region == -1 || pod.Region == region) {
			result = append(result, pod)
		}
	}

	bonalib.Log("[GetReadyPods] Found", len(result), "ready pods for region:", region)
	return result
}

// ---- Adapter implementations ----

func (a *AdapterReplace) Apply(newPod PodState, state map[string]PodState) {
	newPod.LastUpdate = time.Now()
	state[newPod.PodName] = newPod
	bonalib.Log("[AdapterReplace] Thay thế pod:", newPod.PodName)
}

func (a *AdapterIgnore) Apply(newPod PodState, state map[string]PodState) {
	bonalib.Log("[AdapterIgnore] Bỏ qua pod:", newPod.PodName)
}

func (a *AdapterRemove) Apply(newPod PodState, state map[string]PodState) {
	delete(state, newPod.PodName)
	bonalib.Log("[AdapterRemove] Xóa pod:", newPod.PodName)
}

// ---- Adapter selector ----

func SelectAdapter(newPod PodState, oldPod *PodState) Adapter {
	if oldPod == nil {
		bonalib.Log("[SelectAdapter] Chọn AdapterReplace vì oldPod là nil")
		return &AdapterReplace{}
	}

	if strings.EqualFold(newPod.Status, StatusCold) {
		bonalib.Log("[SelectAdapter] Chọn AdapterRemove vì status là Cold")
		return &AdapterRemove{}
	}

	if !strings.EqualFold(newPod.Status, oldPod.Status) {
		bonalib.Log("[SelectAdapter] Chọn AdapterReplace vì status thay đổi")
		return &AdapterReplace{}
	}

	bonalib.Log("[SelectAdapter] Chọn AdapterIgnore vì không thay đổi")
	return &AdapterIgnore{}
}

// ---- Handle state update ----

func (s *StateManager) HandleIncomingState(newPod PodState) {
	s.mu.Lock()
	defer s.mu.Unlock()

	bonalib.Log("[HandleIncomingState] Nhận pod:", newPod.PodName, "Status:", newPod.Status)

	var old *PodState
	if existing, found := s.pods[newPod.PodName]; found {
		old = &existing
	}

	adapter := SelectAdapter(newPod, old)
	adapter.Apply(newPod, s.pods)
}
