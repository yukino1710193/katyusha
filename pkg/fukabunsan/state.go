package fukabunsan

import (
	"sync"
	"time"

	_ "github.com/bonavadeur/katyusha/pkg/bonalib"
)

type PodState struct {
	PodName    string
	Region     int
	IP         string
	Status     string // "Null", "Cold", "Warm", "Ready", ...
	LastUpdate time.Time
}

type StateManager struct {
	mu   sync.RWMutex
	pods map[string]PodState // key là PodName
}

type Adapter interface {
	Apply(newPod PodState, state map[string]PodState) map[string]PodState
}

type AdapterReplace struct{}
type AdapterIgnore struct{}
type AdapterRemove struct{}


func NewStateManager() *StateManager {
	return &StateManager{
		pods: make(map[string]PodState),
	}
}

func (s *StateManager) UpdatePodState(p PodState) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p.LastUpdate = time.Now()
	s.pods[p.PodName] = p
}

func (s *StateManager) GetReadyPods(region int) []PodState {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []PodState
	if region == -1 {
		for _, pod := range s.pods {
			if pod.Status == "Ready" {
				result = append(result, pod)
			}
		}
	} else {
		for _, pod := range s.pods {
			if pod.Region == region && pod.Status == "Ready" {
				result = append(result, pod)
			}
		}
	}
	return result
}

func (a *AdapterReplace) Apply(newPod PodState, state map[string]PodState) map[string]PodState {
	newPod.LastUpdate = time.Now()
	state[newPod.PodName] = newPod
	return state
}

func (a *AdapterIgnore) Apply(newPod PodState, state map[string]PodState) map[string]PodState {
	return state
}

func (a *AdapterRemove) Apply(newPod PodState, state map[string]PodState) map[string]PodState {
	delete(state, newPod.PodName)
	return state
}

func SelectAdapter(newPod PodState, oldPod *PodState) Adapter {
	if oldPod == nil {
		return &AdapterReplace{}
	}

	if newPod.Status == "cold" {
		return &AdapterRemove{}
	}

	if oldPod.Status != newPod.Status {
		return &AdapterReplace{}
	}

	return &AdapterIgnore{}
}

func (pst *StateManager) HandleIncomingState(newPod PodState) {
	pst.mu.Lock()
	defer pst.mu.Unlock()

	var old *PodState
	if existing, found := pst.pods[newPod.PodName]; found {
		old = &existing // tạo bản sao để dùng trong SelectAdapter
	}

	adapter := SelectAdapter(newPod, old)
	pst.pods = adapter.Apply(newPod, pst.pods)
}
