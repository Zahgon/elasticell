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
	"github.com/deepfabric/elasticell/pkg/storage"
	"github.com/fagongzi/util/task"
)

const (
	// When we create a region peer, we should initialize its log term/index > 0,
	// so that we can force the follower peer to sync the snapshot first.
	raftInitLogTerm  = 5
	raftInitLogIndex = 5

	maxSnapTryCnt = 5
)

type snapshotState int

var (
	relax        = snapshotState(1)
	generating   = snapshotState(2)
	applying     = snapshotState(3)
	applyAborted = snapshotState(4)
)

const (
	pending = iota
	running
	cancelling
	cancelled
	finished
	failed
)

type peerStorage struct {
	store *Store
	cell  metapb.Cell

	lastTerm         uint64
	appliedIndexTerm uint64
	lastReadyIndex   uint64
	lastCompactIndex uint64
	raftState        mraft.RaftLocalState
	applyState       mraft.RaftApplyState

	snapTriedCnt     int
	genSnapJob       *task.Job
	applySnapJob     *task.Job
	applySnapJobLock sync.RWMutex

	pendingReads *readIndexQueue
}

func newPeerStorage(store *Store, cell metapb.Cell) (*peerStorage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ps *peerStorage) initRaftState() error { _ = "STUB: not implemented"; return nil }

func (ps *peerStorage) initApplyState() error { _ = "STUB: not implemented"; return nil }

func (ps *peerStorage) initLastTerm() error { _ = "STUB: not implemented"; return nil }

func (ps *peerStorage) isApplyComplete() bool { _ = "STUB: not implemented"; return false }

func (ps *peerStorage) setApplyState(applyState *mraft.RaftApplyState) {
	_ = "STUB: not implemented"
	return
}

func (ps *peerStorage) getApplyState() *mraft.RaftApplyState { _ = "STUB: not implemented"; return nil }

func (ps *peerStorage) getAppliedIndex() uint64 { _ = "STUB: not implemented"; return 0 }

func (ps *peerStorage) getCommittedIndex() uint64 { _ = "STUB: not implemented"; return 0 }

func (ps *peerStorage) getTruncatedIndex() uint64 { _ = "STUB: not implemented"; return 0 }

func (ps *peerStorage) getTruncatedTerm() uint64 { _ = "STUB: not implemented"; return 0 }

func (ps *peerStorage) getAppliedIndexTerm() uint64 { _ = "STUB: not implemented"; return 0 }

func (ps *peerStorage) setAppliedIndexTerm(appliedIndexTerm uint64) {
	_ = "STUB: not implemented"
	return
}

func (ps *peerStorage) validateSnap(snap *raftpb.Snapshot) bool {
	_ = "STUB: not implemented"
	return false
}

// stale snapshot, should generate again.

func (ps *peerStorage) isInitialized() bool { _ = "STUB: not implemented"; return false }

func (ps *peerStorage) isApplyingSnapshot() bool { _ = "STUB: not implemented"; return false }

func (ps *peerStorage) getCell() metapb.Cell { _ = "STUB: not implemented"; return *new(metapb.Cell) }

func (ps *peerStorage) setCell(cell metapb.Cell) { _ = "STUB: not implemented"; return }

func (ps *peerStorage) checkRange(low, high uint64) error { _ = "STUB: not implemented"; return nil }

func (ps *peerStorage) loadLogEntry(index uint64) (*raftpb.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ps *peerStorage) loadCellLocalState(job *task.Job) (*mraft.CellLocalState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ps *peerStorage) applySnapshot(job *task.Job) error { _ = "STUB: not implemented"; return nil }

func (ps *peerStorage) loadApplyState() (*mraft.RaftApplyState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ps *peerStorage) unmarshal(v []byte, expectIndex uint64) (*raftpb.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// / Delete all data belong to the region.
// / If return Err, data may get partial deleted.
func (ps *peerStorage) clearData() error { _ = "STUB: not implemented"; return nil }

// Delete all data that is not covered by `newCell`.
func (ps *peerStorage) clearExtraData(newCell metapb.Cell) error {
	_ = "STUB: not implemented"
	return nil
}

func (ps *peerStorage) updatePeerState(cell metapb.Cell, state mraft.PeerState, wb storage.WriteBatch) error {
	_ = "STUB: not implemented"
	return nil
}

func (ps *peerStorage) writeInitialState(cellID uint64, wb storage.WriteBatch) error {
	_ = "STUB: not implemented"
	return nil
}

func (ps *peerStorage) deleteAllInRange(start, end []byte, job *task.Job) error {
	_ = "STUB: not implemented"
	return nil
}

func compactRaftLog(cellID uint64, state *mraft.RaftApplyState, compactIndex, compactTerm uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// we don't actually delete the logs now, we add an async task to do it.

func loadCellLocalState(cellID uint64, driver storage.Driver, allowNotFound bool) (*mraft.CellLocalState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
