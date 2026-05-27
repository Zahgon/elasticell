package proxy

import (
	"errors"

	"github.com/deepfabric/elasticell/pkg/pb/raftcmdpb"
	"github.com/fagongzi/goetty/protocol/redis"
)

var (
	errInvalidCommand = errors.New("invalid command")
)

func (p *RedisProxy) doMGet(rs *redisSession, cmd redis.Command) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *RedisProxy) doMGetMerge(args [][]byte, rsps ...*raftcmdpb.Response) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}
