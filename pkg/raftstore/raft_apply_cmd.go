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
	"github.com/deepfabric/elasticell/pkg/pb/mraft"
	"github.com/deepfabric/elasticell/pkg/pb/raftcmdpb"
	"github.com/deepfabric/elasticell/pkg/storage"
	"github.com/fagongzi/util/protoc"
)

type applyContext struct {
	// raft state write batch
	wb storage.WriteBatch

	// data
	kvBatch     *redisKVBatch
	bitmapBatch *bitmapBatch

	applyState mraft.RaftApplyState
	req        *raftcmdpb.RaftCMDRequest
	index      uint64
	term       uint64
	metrics    applyMetrics
}

func newApplyContext() *applyContext { _ = "STUB: not implemented"; return nil }

func (ctx *applyContext) reset() { _ = "STUB: not implemented"; return }

func (d *applyDelegate) checkEpoch(req *raftcmdpb.RaftCMDRequest) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *applyDelegate) doApplyRaftCMD(ctx *applyContext) *execResult {
	_ = "STUB: not implemented"
	return nil
}

// resp client

func (d *applyDelegate) execAdminRequest(ctx *applyContext) (*raftcmdpb.RaftCMDResponse, *execResult, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (d *applyDelegate) doExecChangePeer(ctx *applyContext) (*raftcmdpb.RaftCMDResponse, *execResult, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Remove ourself, we will destroy all cell data later.
// So we need not to apply following logs.

// remove pending snapshots

// confChange set by applyConfChange

func (d *applyDelegate) doExecSplit(ctx *applyContext) (*raftcmdpb.RaftCMDResponse, *execResult, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// splitKey < cell.Startkey

// After split, the origin cell key range is [start_key, split_key),
// the new split cell is [split_key, end).

func (d *applyDelegate) doExecRaftGC(ctx *applyContext) (*raftcmdpb.RaftCMDResponse, *execResult, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (d *applyDelegate) execWriteRequest(ctx *applyContext) *raftcmdpb.RaftCMDResponse {
	_ = "STUB: not implemented"
	return nil
}

func (pr *PeerReplicate) doExecReadCmd(c *cmd) { _ = "STUB: not implemented"; return }

func newAdminRaftCMDResponse(adminType raftcmdpb.AdminCmdType, subRsp protoc.PB) *raftcmdpb.RaftCMDResponse {
	_ = "STUB: not implemented"
	return nil
}
