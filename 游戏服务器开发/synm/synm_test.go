package synm

import (
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {
	sm := sync.Map{}
	sm.Load("")
	sm.Delete("")
	rwMutex := sync.RWMutex{}
	rwMutex.RLock()
}
