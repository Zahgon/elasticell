package proxy

import (
	"sync"

	"github.com/deepfabric/elasticell/pkg/pb/raftcmdpb"
	"github.com/fagongzi/goetty"
	"github.com/fagongzi/util/task"
)

type redisSession struct {
	sync.RWMutex

	session goetty.IOSession
	resps   *task.Queue
	addr    string

	aggLock      sync.RWMutex
	aggregations map[string]*aggregationReq
}

func newSession(session goetty.IOSession) *redisSession { _ = "STUB: not implemented"; return nil }

func (rs *redisSession) close() { _ = "STUB: not implemented"; return }

func (rs *redisSession) addAggregation(id []byte, req *aggregationReq) {
	_ = "STUB: not implemented"
	return
}

func (rs *redisSession) resp(rsp *raftcmdpb.Response) { _ = "STUB: not implemented"; return }

func (rs *redisSession) errorResp(err error) { _ = "STUB: not implemented"; return }

func (rs *redisSession) writeLoop() { _ = "STUB: not implemented"; return }

func (rs *redisSession) doResp(resp *raftcmdpb.Response, buf *goetty.ByteBuf) {
	_ = "STUB: not implemented"
	return
}
