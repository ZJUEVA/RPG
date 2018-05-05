package ecs

import (
	"sync"
	"sync/atomic"
)

var (
	uidCount    uint64
	counterLock sync.Mutex
)

type Identity struct {
	uid uint64
}

func (i Identity) Uid() uint64 {
	return i.uid
}
func NewIdentity() Identity {
	return Identity{uid: atomic.AddUint64(&uidCount, 1)}
}
