package safelist

import (
	"sync"
)

type SafeList[T any] struct {
	mu    sync.Mutex
	items []T
}

func (sl *SafeList[T]) Add(item T) {
	sl.mu.Lock()
	defer sl.mu.Unlock()
	sl.items = append(sl.items, item)
}

func (sl *SafeList[T]) GetAll() []T {
	sl.mu.Lock()
	defer sl.mu.Unlock()
	return sl.items
}

func (sl *SafeList[T]) Clear() {
	sl.mu.Lock()
	defer sl.mu.Unlock()
	sl.items = []T{}
}
