package shared

import "sync"

type AtomicString struct {
	value string
	mutex sync.RWMutex
}

func (as *AtomicString) Store(val string) {
	as.mutex.Lock()
	defer as.mutex.Unlock()

	as.value = val
}

func (as *AtomicString) Load() string {
	as.mutex.RLock()
	defer as.mutex.RUnlock()

	return as.value
}
