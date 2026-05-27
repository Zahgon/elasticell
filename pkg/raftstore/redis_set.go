package raftstore

import (
	"github.com/deepfabric/elasticell/pkg/pb/raftcmdpb"
)

func (s *Store) execSAdd(ctx *applyContext, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) execSRem(ctx *applyContext, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) execSCard(id uint64, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) execSMembers(id uint64, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) execSIsMember(id uint64, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) execSPop(id uint64, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}
