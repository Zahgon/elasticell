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
	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/deepfabric/elasticell/pkg/pb/metapb"
	"github.com/deepfabric/elasticell/pkg/pb/mraft"
	"github.com/deepfabric/elasticell/pkg/storage"
	"github.com/fagongzi/util/task"
)

type readyContext struct {
	raftState  mraft.RaftLocalState
	applyState mraft.RaftApplyState
	lastTerm   uint64
	snap       *mraft.SnapshotMessage
	wb         storage.WriteBatch
}

func (ctx *readyContext) reset() { _ = "STUB: not implemented"; return }

type splitCheckResult struct {
	cellID   uint64
	epoch    metapb.CellEpoch
	splitKey []byte
}

type applySnapResult struct {
	prevCell metapb.Cell
	cell     metapb.Cell
}

type readIndexQueue struct {
	cellID   uint64
	reads    []*cmd
	readyCnt int32
}

func (q *readIndexQueue) push(c *cmd) { _ = "STUB: not implemented"; return }

func (q *readIndexQueue) pop() *cmd { _ = "STUB: not implemented"; return nil }

func (q *readIndexQueue) incrReadyCnt() int32 { _ = "STUB: not implemented"; return 0 }

func (q *readIndexQueue) decrReadyCnt() int32 { _ = "STUB: not implemented"; return 0 }

func (q *readIndexQueue) resetReadyCnt() { _ = "STUB: not implemented"; return }

func (q *readIndexQueue) getReadyCnt() int32 { _ = "STUB: not implemented"; return 0 }

func (q *readIndexQueue) size() int { _ = "STUB: not implemented"; return 0 }

// ====================== raft ready handle methods
func (ps *peerStorage) doAppendSnapshot(ctx *readyContext, snap raftpb.Snapshot) error {
	_ = "STUB: not implemented"
	return nil
}

// The snapshot only contains log which index > applied index, so
// here the truncate state's (index, term) is in snapshot metadata.

// doAppendEntries the given entries to the raft log using previous last index or self.last_index.
// Return the new last index for later update. After we commit in engine, we can set last_index
// to the return one.
func (ps *peerStorage) doAppendEntries(ctx *readyContext, entries []raftpb.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete any previously appended log entries which never committed.

func (pr *PeerReplicate) doSaveRaftState(ctx *readyContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (pr *PeerReplicate) doSaveApplyState(ctx *readyContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (pr *PeerReplicate) doApplySnap(ctx *readyContext, rd *raft.Ready) *applySnapResult {
	_ = "STUB: not implemented"
	return nil
}

// If we apply snapshot ok, we should update some infos like applied index too.

// cleanup data before apply snap job

// No need panic here, when applying snapshot, the deletion will be tried
// again. But if the cell range changes, like [a, c) -> [a, b) and [b, c),
// [b, c) will be kept in rocksdb until a covered snapshot is applied or
// store is restarted.

// remove pending snapshots for sending

func (pr *PeerReplicate) applyCommittedEntries(rd *raft.Ready) { _ = "STUB: not implemented"; return }

func (pr *PeerReplicate) doPropose(c *cmd, isConfChange bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (pr *PeerReplicate) doSplitCheck(epoch metapb.CellEpoch, startKey, endKey []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (pr *PeerReplicate) doAskSplit(cell metapb.Cell, peer metapb.Peer, splitKey []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (pr *PeerReplicate) doPostApply(result *asyncApplyResult) { _ = "STUB: not implemented"; return }

func (s *Store) doPostApplyResult(result *asyncApplyResult) { _ = "STUB: not implemented"; return }

func (s *Store) doApplyConfChange(cellID uint64, cp *changePeer) { _ = "STUB: not implemented"; return }

// Apply failed, skip.

// Notify pd immediately.

// Add this peer to cache.

// Remove this peer from cache.

// We only care remove itself now.

func (s *Store) doApplySplit(cellID uint64, result *splitResult) { _ = "STUB: not implemented"; return }

// add new cell peers to cache

// If the store received a raft msg with the new region raft group
// before splitting, it will creates a uninitialized peer.
// We can remove this uninitialized peer directly.

// peer information is already written into db, can't recover.
// there is probably a bug.

// If this peer is the leader of the cell before split, it's intuitional for
// it to become the leader of new split cell.
// The ticks are accelerated here, so that the peer for the new split cell
// comes to campaign earlier than the other follower peers. And then it's more
// likely for this peer to become the leader of the new split cell.
// If the other follower peers applies logs too slowly, they may fail to vote the
// `MsgRequestVote` from this peer on its campaign.
// In this worst case scenario, the new split raft group will not be available
// since there is no leader established during one election timeout after the split.

func (s *Store) doApplyRaftLogGC(cellID uint64, result *raftGCResult) {
	_ = "STUB: not implemented"
	return
}

func (pr *PeerReplicate) doApplyReads(rd *raft.Ready) { _ = "STUB: not implemented"; return }

// Note that only after handle read_states can we identify what requests are
// actually stale.

// all uncommitted reads will be dropped silently in raft.

// we are not leader now, so all writes in the batch is actually stale

func (pr *PeerReplicate) updateKeyRange(result *applySnapResult) { _ = "STUB: not implemented"; return }

// we have already initialized the peer, so it must exist in cell_ranges.

func (pr *PeerReplicate) readyToHandleRead() bool {
	_ = "STUB: not implemented"
	// If applied_index_term isn't equal to current term, there may be some values that are not
	// applied by this leader yet but the old leader.
	return false
}

// ======================raft storage interface method
func (ps *peerStorage) InitialState() (raftpb.HardState, raftpb.ConfState, error) {
	_ = "STUB: not implemented"
	return *new(raftpb.HardState), *new(raftpb.ConfState), nil
}

func (ps *peerStorage) Entries(low, high, maxSize uint64) ([]raftpb.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If election happens in inactive cells, they will just try
// to fetch one empty log.

// May meet gap or has been compacted.

// If we get the correct number of entries the total size exceeds max_size, returns.

func (ps *peerStorage) Term(idx uint64) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (ps *peerStorage) LastIndex() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (ps *peerStorage) FirstIndex() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (ps *peerStorage) Snapshot() (raftpb.Snapshot, error) {
	_ = "STUB: not implemented"
	return *new(raftpb.Snapshot), nil
}

// snapshot failure, we will continue try do snapshot

func (ps *peerStorage) setGenSnapJob(job *task.Job) { _ = "STUB: not implemented"; return }

func (ps *peerStorage) setApplySnapJob(job *task.Job) { _ = "STUB: not implemented"; return }

func (s *Store) notifySplitCellIndex(leftCellID uint64, rightCellID uint64) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) notifyDestroyCellIndex(cell *metapb.Cell) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) notifyRebuildCellIndex(cell *metapb.Cell) (err error) {
	_ = "STUB: not implemented"
	return nil
}
