// Copyright 2016 DeepFabric, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package raftstore

import (
	"github.com/deepfabric/elasticell/pkg/pb/raftcmdpb"
)

func (s *Store) execLIndex(id uint64, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) execLLEN(id uint64, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) execLRange(id uint64, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) execLInsert(ctx *applyContext, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) execLPop(ctx *applyContext, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) execLPush(ctx *applyContext, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) execLPushX(ctx *applyContext, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) execLRem(ctx *applyContext, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) execLSet(ctx *applyContext, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) execLTrim(ctx *applyContext, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) execRPop(ctx *applyContext, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) execRPush(ctx *applyContext, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) execRPushX(ctx *applyContext, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}
