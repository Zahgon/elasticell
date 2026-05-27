package proxy

import (
	"github.com/deepfabric/elasticell/pkg/pb/raftcmdpb"
)

var (
	cmdSet = []byte("SET")
	cmdGet = []byte("GET")
)

type aggregationReq struct {
	reply   int
	parts   []*raftcmdpb.Response
	mergeFn func([][]byte, ...*raftcmdpb.Response) *raftcmdpb.Response
	args    [][]byte
}

func newAggregationReq(n int, mergeFn func([][]byte, ...*raftcmdpb.Response) *raftcmdpb.Response, args [][]byte) *aggregationReq {
	_ = "STUB: not implemented"
	return nil
}

func (req *aggregationReq) addPart(index int, rsp *raftcmdpb.Response) bool {
	_ = "STUB: not implemented"
	return false
}

func (req *aggregationReq) merge() *raftcmdpb.Response { _ = "STUB: not implemented"; return nil }

func isAggregationPart(id []byte) bool { _ = "STUB: not implemented"; return false }

func parseAggregationPart(id []byte) ([]byte, int) { _ = "STUB: not implemented"; return nil, 0 }

func parseStrInt64(data []byte) int { _ = "STUB: not implemented"; return 0 }
