package proxy

import (
	"sync"
	"time"

	"github.com/fagongzi/goetty"
	"github.com/fagongzi/util/task"
	"github.com/pkg/errors"
)

var (
	errConnect            = errors.New("not connected")
	defaultConnectTimeout = time.Second * 5
)

type backend struct {
	sync.RWMutex

	p    *RedisProxy
	addr string
	conn goetty.IOSession
	reqs *task.Queue
}

func newBackend(p *RedisProxy, addr string, conn goetty.IOSession) *backend {
	_ = "STUB: not implemented"
	return nil
}

func (bc *backend) isConnected() bool { _ = "STUB: not implemented"; return false }

func (bc *backend) connect() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (bc *backend) close(exit bool) { _ = "STUB: not implemented"; return }

func (bc *backend) addReq(r *req) error { _ = "STUB: not implemented"; return nil }

func (bc *backend) readLoop() { _ = "STUB: not implemented"; return }

func (bc *backend) writeLoop() { _ = "STUB: not implemented"; return }

func (p *RedisProxy) getConn(addr string) (*backend, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *RedisProxy) getConnLocked(addr string) *backend { _ = "STUB: not implemented"; return nil }

func (p *RedisProxy) createConn(addr string) *backend {
	_ = "STUB: not implemented"

	// double check
	return nil
}

// update p.bcAddrs

func (p *RedisProxy) checkConnect(addr string, bc *backend) bool {
	_ = "STUB: not implemented"
	return false
}
