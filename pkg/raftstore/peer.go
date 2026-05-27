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
	"time"

	"github.com/coreos/etcd/raft"
	"github.com/deepfabric/elasticell/pkg/pb/metapb"
	"github.com/deepfabric/elasticell/pkg/pb/mraft"
	"github.com/deepfabric/elasticell/pkg/pb/pdpb"
	"github.com/deepfabric/elasticell/pkg/pb/raftcmdpb"
	"github.com/fagongzi/util/task"
)

type action int

const (
	checkSplit = iota
	checkCompact
	doCampaign
)

// PeerReplicate is the cell's peer replicate. Every cell replicate has a PeerReplicate.
type PeerReplicate struct {
	cellID            uint64
	peer              metapb.Peer
	rn                *raft.RawNode
	store             *Store
	ps                *peerStorage
	batch             *proposeBatch
	events            *task.RingBuffer
	ticks             *task.Queue
	steps             *task.Queue
	reports           *task.Queue
	applyResults      *task.Queue
	requests          *task.Queue
	actions           *task.Queue
	stopRaftTick      bool
	peerHeartbeatsMap *peerHeartbeatsMap
	pendingReads      *readIndexQueue
	lastHBJob         *task.Job
	writtenKeys       uint64
	writtenBytes      uint64
	sizeDiffHint      uint64
	raftLogSizeHint   uint64
	deleteKeysHint    uint64
	cancelTaskIds     []uint64
	metrics           localMetrics
}

func createPeerReplicate(store *Store, cell *metapb.Cell) (*PeerReplicate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The peer can be created from another node with raft membership changes, and we only
// know the cell_id and peer_id when creating this replicated peer, the cell info
// will be retrieved later after applying snapshot.
func doReplicate(store *Store, msg *mraft.RaftMessage, peerID uint64) (*PeerReplicate, error) {
	_ = "STUB: not implemented"
	// We will remove tombstone key when apply snapshot
	return nil, nil
}

func newPeerReplicate(store *Store, cell *metapb.Cell, peerID uint64) (*PeerReplicate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If this region has only one peer and I am the one, campaign directly.

func (pr *PeerReplicate) maybeCampaign() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil

	// The peer campaigned when it was created, no need to do it again.
}

func (pr *PeerReplicate) onAdminRequest(adminReq *raftcmdpb.AdminRequest) {
	_ = "STUB: not implemented"
	return
}

func (pr *PeerReplicate) onReq(req *raftcmdpb.Request, cb func(*raftcmdpb.RaftCMDResponse)) {
	_ = "STUB: not implemented"
	return
}

func (pr *PeerReplicate) addAction(act action) { _ = "STUB: not implemented"; return }

func (pr *PeerReplicate) addRequest(req *reqCtx) { _ = "STUB: not implemented"; return }

func (pr *PeerReplicate) resetBatch() { _ = "STUB: not implemented"; return }

func (pr *PeerReplicate) handleHeartbeat() { _ = "STUB: not implemented"; return }

// cancel last if not complete

func (pr *PeerReplicate) setLastHBJob(job *task.Job) { _ = "STUB: not implemented"; return }

func (pr *PeerReplicate) doHeartbeat() error { _ = "STUB: not implemented"; return nil }

func (pr *PeerReplicate) checkPeers() { _ = "STUB: not implemented"; return }

// Insert heartbeats in case that some peers never response heartbeats.

func (pr *PeerReplicate) collectDownPeers(maxDuration time.Duration) []pdpb.PeerStats {
	_ = "STUB: not implemented"
	return nil
}

func (pr *PeerReplicate) collectPendingPeers() []metapb.Peer { _ = "STUB: not implemented"; return nil }

func (pr *PeerReplicate) stopEventLoop() { _ = "STUB: not implemented"; return }

func (pr *PeerReplicate) destroy() error { _ = "STUB: not implemented"; return nil }

func (pr *PeerReplicate) getStore() *peerStorage { _ = "STUB: not implemented"; return nil }

func (pr *PeerReplicate) getCell() metapb.Cell { _ = "STUB: not implemented"; return *new(metapb.Cell) }

func (pr *PeerReplicate) getPeer() metapb.Peer { _ = "STUB: not implemented"; return *new(metapb.Peer) }
