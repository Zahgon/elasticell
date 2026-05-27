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

// command like: bmcreate bm1 [1, 3, 3]
func (s *Store) execBMCreate(ctx *applyContext, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

// command like: bmadd bm1 1 [2 3 4]
func (s *Store) execBMAdd(ctx *applyContext, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

// command like: bmremove bm1 1 [2 3 4]
func (s *Store) execBMRemove(ctx *applyContext, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

// command like: bmclear bm1
func (s *Store) execBMClear(ctx *applyContext, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

// command like: bmdel bm1
func (s *Store) execBMDel(ctx *applyContext, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

// command like: bmrange bm1 start count
func (s *Store) execBMRange(id uint64, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

// command like: bmcount bm1
func (s *Store) execBMCount(id uint64, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}

// command like: bmcontains bm1 1
func (s *Store) execBMContains(id uint64, req *raftcmdpb.Request) *raftcmdpb.Response {
	_ = "STUB: not implemented"
	return nil
}
