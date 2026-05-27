package proxy

import (
	"context"
	"sync"

	"github.com/deepfabric/elasticell/pkg/pb/metapb"
	"github.com/deepfabric/elasticell/pkg/pb/pdpb"
	"github.com/deepfabric/elasticell/pkg/pb/raftcmdpb"
	"github.com/deepfabric/elasticell/pkg/pd"
	"github.com/deepfabric/elasticell/pkg/util"
	"github.com/fagongzi/goetty"
	"github.com/fagongzi/goetty/protocol/redis"
	"github.com/fagongzi/util/task"
)

const (
	batch = 64
)

var (
	pingReq = &raftcmdpb.Request{
		Cmd: [][]byte{[]byte("ping")},
	}
)

type req struct {
	rs      *redisSession
	raftReq *raftcmdpb.Request
	retries int
}

func newReqUUID(id []byte, cmd redis.Command, rs *redisSession) *req {
	_ = "STUB: not implemented"
	return nil
}

func newReq(cmd redis.Command, rs *redisSession) *req { _ = "STUB: not implemented"; return nil }

func newID() []byte { _ = "STUB: not implemented"; return nil }

func (r *req) errorDone(err error) { _ = "STUB: not implemented"; return }

func (r *req) done(rsp *raftcmdpb.Response) { _ = "STUB: not implemented"; return }

// RedisProxy is a redis proxy
type RedisProxy struct {
	sync.RWMutex

	cfg             *Cfg
	svr             *goetty.Server
	pdClient        *pd.Client
	watcher         *pd.Watcher
	aggregationCmds map[string]func(*redisSession, redis.Command) (bool, error)
	supportCmds     map[string]struct{}
	keyConvertFun   func([]byte, func([]byte) metapb.Cell) metapb.Cell
	ranges          *util.CellTree
	stores          map[uint64]*metapb.Store
	cellLeaderAddrs map[uint64]string   // cellid -> leader peer store addr
	bcs             map[string]*backend // store addr -> netconn
	routing         *routing            // uuid -> session
	syncEpoch       uint64
	reqs            []*task.Queue
	retries         *task.Queue
	pings           chan string
	bcAddrs         []string // store addrs
	rrNext          int64    // round robin of bcAddrs

	ctx      context.Context
	cancel   context.CancelFunc
	stopOnce sync.Once
	stopWG   sync.WaitGroup
	stopC    chan struct{}
}

// NewRedisProxy returns a redisp proxy
func NewRedisProxy(cfg *Cfg) *RedisProxy { _ = "STUB: not implemented"; return nil }

// Start starts the proxy
func (p *RedisProxy) Start() error { _ = "STUB: not implemented"; return nil }

// Stop stop the proxy
func (p *RedisProxy) Stop() { _ = "STUB: not implemented"; return }

func (p *RedisProxy) listenToStop() { _ = "STUB: not implemented"; return }

func (p *RedisProxy) init() {
	p.ctx, p.cancel = context.WithCancel(context.TODO())

	p.initKeyConvert()
	p.initSupportCMDs()
	p.initWatcher()
	p.initQueues()
}

func (p *RedisProxy) doStop() { _ = "STUB: not implemented"; return }

func (p *RedisProxy) initWatcher() { _ = "STUB: not implemented"; return }

func (p *RedisProxy) initQueues() { _ = "STUB: not implemented"; return }

func (p *RedisProxy) doConnection(session goetty.IOSession) error {
	_ = "STUB: not implemented"
	return nil
}

// every client has 2 goroutines, read, write

func (p *RedisProxy) initKeyConvert() { _ = "STUB: not implemented"; return }

func (p *RedisProxy) initSupportCMDs() { _ = "STUB: not implemented"; return }

// kv

// bitmap

func (p *RedisProxy) refreshStores() { _ = "STUB: not implemented"; return }

func (p *RedisProxy) refreshRanges() { _ = "STUB: not implemented"; return }

func (p *RedisProxy) refreshRange(r *pdpb.Range) { _ = "STUB: not implemented"; return }

func (p *RedisProxy) doRefreshRange(r *pdpb.Range) { _ = "STUB: not implemented"; return }

func (p *RedisProxy) clean() { _ = "STUB: not implemented"; return }

func (p *RedisProxy) getSyncEpoch() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *RedisProxy) addToPing(target string) { _ = "STUB: not implemented"; return }

func (p *RedisProxy) retry(r *req) { _ = "STUB: not implemented"; return }

func (p *RedisProxy) addToForward(r *req) { _ = "STUB: not implemented"; return }

func (p *RedisProxy) readyToHandleReq(ctx context.Context) { _ = "STUB: not implemented"; return }

func (p *RedisProxy) handleReq(r *req) {
	_ = "STUB: not implemented"

	// If epoch is not stale, wait next
	return
}

func (p *RedisProxy) forwardTo(addr string, r *req) error { _ = "STUB: not implemented"; return nil }

func (p *RedisProxy) onResp(rsp *raftcmdpb.Response) { _ = "STUB: not implemented"; return }

func (p *RedisProxy) search(value []byte) metapb.Cell {
	_ = "STUB: not implemented"
	return *new(metapb.Cell)
}

func (p *RedisProxy) getLeaderStoreAddr(key []byte) (string, uint64) {
	_ = "STUB: not implemented"
	return "", 0
}

func (p *RedisProxy) getRandomStoreAddr() string { _ = "STUB: not implemented"; return "" }

func (p *RedisProxy) getRRStoreAddr() (target string) { _ = "STUB: not implemented"; return "" }
