package proxy

import (
	"sync"
)

const (
	bucketSize = 128
	bucketM    = 127
)

type routingMap struct {
	sync.RWMutex
	m map[string]*req
}

func (m *routingMap) put(key string, value *req) { _ = "STUB: not implemented"; return }

func (m *routingMap) delete(key string) *req { _ = "STUB: not implemented"; return nil }

type routing struct {
	rms []*routingMap
}

func newRouting() *routing { _ = "STUB: not implemented"; return nil }

func (r *routing) put(uuid []byte, value *req) { _ = "STUB: not implemented"; return }

func (r *routing) delete(uuid []byte) *req { _ = "STUB: not implemented"; return nil }

func getIndex(key []byte) int { _ = "STUB: not implemented"; return 0 }
