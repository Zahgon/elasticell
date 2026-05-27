package proxy

import (
	"github.com/deepfabric/elasticell/pkg/pb/raftcmdpb"
	"github.com/fagongzi/goetty/protocol/redis"
	"github.com/pilosa/pilosa/roaring"
)

const (
	optionWithBM = "WITHBM"
)

// bmand [withbm|start count]  bm1 bm2 [bm3 bm4]
func (p *RedisProxy) doBMAnd(rs *redisSession, cmd redis.Command) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// bmor [withbm|start count]  bm1 bm2 [bm3 bm4]
func (p *RedisProxy) doBMOr(rs *redisSession, cmd redis.Command) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// bmxor [withbm|start count]  bm1 bm2 [bm3 bm4]
func (p *RedisProxy) doBMXor(rs *redisSession, cmd redis.Command) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// bmandnot [withbm|start count]  bm1 bm2 [bm3 bm4]
func (p *RedisProxy) doBMAndNot(rs *redisSession, cmd redis.Command) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *RedisProxy) doBMAggregation(rs *redisSession, cmd redis.Command, mergeFn func([][]byte, ...*raftcmdpb.Response) *raftcmdpb.Response) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *RedisProxy) doBMAndMerge(args [][]byte, rsps ...*raftcmdpb.Response) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

func (p *RedisProxy) doBMOrMerge(args [][]byte, rsps ...*raftcmdpb.Response) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

func (p *RedisProxy) doBMXorMerge(args [][]byte, rsps ...*raftcmdpb.Response) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

// and not A - (A and B)
func (p *RedisProxy) doBMAndNotMerge(args [][]byte, rsps ...*raftcmdpb.Response) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

func (p *RedisProxy) buildResult(bm *roaring.Bitmap, args [][]byte) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}
