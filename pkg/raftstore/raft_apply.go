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
	"sync"

	"github.com/coreos/etcd/raft/raftpb"
	"github.com/deepfabric/elasticell/pkg/pb/metapb"
	"github.com/deepfabric/elasticell/pkg/pb/mraft"
	"github.com/deepfabric/elasticell/pkg/pb/raftcmdpb"
)

type applyMetrics struct {
	// an inaccurate difference in cell size since last reset.
	sizeDiffHint uint64
	// delete keys' count since last reset.
	deleteKeysHint uint64
	writtenBytes   uint64
	writtenKeys    uint64

	admin raftAdminMetrics
}

type asyncApplyResult struct {
	cellID           uint64
	appliedIndexTerm uint64
	applyState       mraft.RaftApplyState
	result           *execResult
	metrics          applyMetrics
}

func (res *asyncApplyResult) reset() { _ = "STUB: not implemented"; return }

func (res *asyncApplyResult) hasSplitExecResult() bool { _ = "STUB: not implemented"; return false }

type changePeer struct {
	confChange raftpb.ConfChange
	peer       metapb.Peer
	cell       metapb.Cell
}

type splitResult struct {
	left  metapb.Cell
	right metapb.Cell
}

type raftGCResult struct {
	state      mraft.RaftTruncatedState
	firstIndex uint64
}

type execResult struct {
	adminType    raftcmdpb.AdminCmdType
	changePeer   *changePeer
	splitResult  *splitResult
	raftGCResult *raftGCResult
}

type applyDelegate struct {
	sync.RWMutex

	store *Store
	ps    *peerStorage

	peerID uint64
	cell   metapb.Cell

	// if we remove ourself in ChangePeer remove, we should set this flag, then
	// any following committed logs in same Ready should be applied failed.
	pendingRemove bool

	applyState       mraft.RaftApplyState
	appliedIndexTerm uint64
	term             uint64

	pendingCMDs          []*cmd
	pendingChangePeerCMD *cmd
}

func (d *applyDelegate) clearAllCommandsAsStale() { _ = "STUB: not implemented"; return }

func (d *applyDelegate) findCB(ctx *applyContext) *cmd { _ = "STUB: not implemented"; return nil }

// Because of the lack of original RaftCmdRequest, we skip calling
// coprocessor here.

func (d *applyDelegate) appendPendingCmd(c *cmd) { _ = "STUB: not implemented"; return }

func (d *applyDelegate) setPendingChangePeerCMD(c *cmd) { _ = "STUB: not implemented"; return }

func (d *applyDelegate) getPendingChangePeerCMD() *cmd { _ = "STUB: not implemented"; return nil }

func (d *applyDelegate) popPendingCMD(raftLogEntryTerm uint64) *cmd {
	_ = "STUB: not implemented"
	return nil
}

func isChangePeerCMD(req *raftcmdpb.RaftCMDRequest) bool { _ = "STUB: not implemented"; return false }

func (d *applyDelegate) notifyStaleCMD(c *cmd) { _ = "STUB: not implemented"; return }

func (d *applyDelegate) notifyCellRemoved(c *cmd) { _ = "STUB: not implemented"; return }

func (d *applyDelegate) applyCommittedEntries(commitedEntries []raftpb.Entry) {
	_ = "STUB: not implemented"
	return
}

// This peer is about to be destroyed, skip everything.

// only release RaftCMDRequest. Header and Requests fields is pb created in Unmarshal

func (d *applyDelegate) applyEntry(ctx *applyContext, entry *raftpb.Entry) *execResult {
	_ = "STUB: not implemented"
	return nil
}

// when a peer become leader, it will send an empty entry.

// apprently, all the callbacks whose term is less than entry's term are stale.

func (d *applyDelegate) applyConfChange(ctx *applyContext, entry *raftpb.Entry) *execResult {
	_ = "STUB: not implemented"
	return nil
}

func (d *applyDelegate) destroy() { _ = "STUB: not implemented"; return }

func (d *applyDelegate) setPendingRemove() { _ = "STUB: not implemented"; return }

func (d *applyDelegate) isPendingRemove() bool { _ = "STUB: not implemented"; return false }
