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
	"context"

	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/deepfabric/elasticell/pkg/pb/metapb"
	"github.com/deepfabric/elasticell/pkg/pb/raftcmdpb"
)

var (
	emptyStruct = struct{}{}
)

type requestPolicy int

const (
	readLocal             = requestPolicy(0)
	readIndex             = requestPolicy(1)
	proposeNormal         = requestPolicy(2)
	proposeTransferLeader = requestPolicy(3)
	proposeChange         = requestPolicy(4)

	transferLeaderAllowLogLag = 10

	batch = 1024
)

func (pr *PeerReplicate) onRaftTick(arg interface{}) { _ = "STUB: not implemented"; return }

func (pr *PeerReplicate) addEvent() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (pr *PeerReplicate) readyToServeRaft(ctx context.Context) { _ = "STUB: not implemented"; return }

// resp all stale requests in batch and queue

func (pr *PeerReplicate) handleAction(items []interface{}) { _ = "STUB: not implemented"; return }

func (pr *PeerReplicate) handleCheckCompact() {
	_ = "STUB: not implemented"
	// Leader will replicate the compact log command to followers,
	// If we use current replicated_index (like 10) as the compact index,
	// when we replicate this log, the newest replicated_index will be 11,
	// but we only compact the log to 10, not 11, at that time,
	// the first index is 10, and replicated_index is 11, with an extra log,
	// and we will do compact again with compact index 11, in cycles...
	// So we introduce a threshold, if replicated index - first index > threshold,
	// we will try to compact log.
	// raft log entries[..............................................]
	//
	//	    ^                                       ^
	//	    |-----------------threshold------------ |
	//	first_index                         replicated_index
	return
}

// When an election happened or a new peer is added, replicated_idx can be 0.

// Have no idea why subtract 1 here, but original code did this by magic.

// avoid leader send snapshot to the a little lag peer.

// In case compactIdx == firstIdx before subtraction.

func (pr *PeerReplicate) handleCheckSplit() { _ = "STUB: not implemented"; return }

// If a peer is apply snapshot, skip split, avoid sent snapshot again in future.

func (pr *PeerReplicate) handleTick(items []interface{}) { _ = "STUB: not implemented"; return }

func (pr *PeerReplicate) handleStep(items []interface{}) { _ = "STUB: not implemented"; return }

func (pr *PeerReplicate) handleReport(items []interface{}) { _ = "STUB: not implemented"; return }

func (pr *PeerReplicate) handleApplyResult(items []interface{}) { _ = "STUB: not implemented"; return }

func (pr *PeerReplicate) handleRequest(items []interface{}) { _ = "STUB: not implemented"; return }

func (pr *PeerReplicate) handleReady() {
	_ = "STUB: not implemented"
	// If we continue to handle all the messages, it may cause too many messages because
	// leader will send all the remaining messages to this follower, which can lead
	// to full message queue under high load.
	return
}

// wait apply committed entries complete

// If snapshot is received, further handling

// When we apply snapshot, stop raft tick and resume until the snapshot applied

func (pr *PeerReplicate) addApplyResult(result *asyncApplyResult) {
	_ = "STUB: not implemented"
	return
}

func (pr *PeerReplicate) doPollApply(result *asyncApplyResult) { _ = "STUB: not implemented"; return }

func (pr *PeerReplicate) handleRaftReadyAppend(ctx *readyContext, rd *raft.Ready) {
	_ = "STUB: not implemented"
	return

	// If we become leader, send heartbeat to pd
}

// The leader can write to disk and replicate to the followers concurrently
// For more details, check raft thesis 10.2.1.

func (pr *PeerReplicate) handleRaftReadyApply(ctx *readyContext, rd *raft.Ready) {
	_ = "STUB: not implemented"
	return

	// When apply snapshot, there is no log applied and not compacted yet.
}

// Because we only handle raft ready when not applying snapshot, so following
// line won't be called twice for the same snapshot.

func (pr *PeerReplicate) handleAppendSnapshot(ctx *readyContext, rd *raft.Ready) {
	_ = "STUB: not implemented"
	return
}

func (pr *PeerReplicate) handleAppendEntries(ctx *readyContext, rd *raft.Ready) {
	_ = "STUB: not implemented"
	return
}

func (pr *PeerReplicate) handleSaveRaftState(ctx *readyContext) { _ = "STUB: not implemented"; return }

func (pr *PeerReplicate) handleSaveApplyState(ctx *readyContext) { _ = "STUB: not implemented"; return }

func (pr *PeerReplicate) checkProposal(c *cmd) bool {
	_ = "STUB: not implemented"
	// we handle all read, write and admin cmd here
	return false
}

// Note:
// The peer that is being checked is a leader. It might step down to be a follower later. It
// doesn't matter whether the peer is a leader or not. If it's not a leader, the proposing
// command log entry can't be committed.

func (pr *PeerReplicate) propose(c *cmd) { _ = "STUB: not implemented"; return }

func (pr *PeerReplicate) proposeNormal(c *cmd) bool { _ = "STUB: not implemented"; return false }

func (pr *PeerReplicate) proposeConfChange(c *cmd) bool { _ = "STUB: not implemented"; return false }

func (pr *PeerReplicate) proposeTransferLeader(c *cmd) bool {
	_ = "STUB: not implemented"
	return false
}

// transfer leader command doesn't need to replicate log and apply, so we
// return immediately. Note that this command may fail, we can view it just as an advice

func (pr *PeerReplicate) doTransferLeader(peer *metapb.Peer) { _ = "STUB: not implemented"; return }

func (pr *PeerReplicate) isTransferLeaderAllowed(newLeaderPeer *metapb.Peer) bool {
	_ = "STUB: not implemented"
	return false
}

// / Check whether it's safe to propose the specified conf change request.
// / It's safe iff at least the quorum of the Raft group is still healthy
// / right after that conf change is applied.
// / Define the total number of nodes in current Raft cluster to be `total`.
// / To ensure the above safety, if the cmd is
// / 1. A `AddNode` request
// /    Then at least '(total + 1)/2 + 1' nodes need to be up to date for now.
// / 2. A `RemoveNode` request
// /    Then at least '(total - 1)/2 + 1' other nodes (the node about to be removed is excluded)
// /    need to be up to date for now.
func (pr *PeerReplicate) checkConfChange(c *cmd) error { _ = "STUB: not implemented"; return nil }

// It's always safe if there is only one node in the cluster.

// / Count the number of the healthy nodes.
// / A node is healthy when
// / 1. it's the leader of the Raft group, which has the latest logs
// / 2. it's a follower, and it does not lag behind the leader a lot.
// /    If a snapshot is involved between it and the Raft leader, it's not healthy since
// /    it cannot works as a node in the quorum to receive replicating logs from leader.
func (pr *PeerReplicate) countHealthyNode() int { _ = "STUB: not implemented"; return 0 }

func (pr *PeerReplicate) nextProposalIndex() uint64 { _ = "STUB: not implemented"; return 0 }

func (pr *PeerReplicate) pendingReadCount() int { _ = "STUB: not implemented"; return 0 }

func (pr *PeerReplicate) readyReadCount() int { _ = "STUB: not implemented"; return 0 }

func (pr *PeerReplicate) isLeader() bool { _ = "STUB: not implemented"; return false }

func (pr *PeerReplicate) getLeaderPeerID() uint64 { _ = "STUB: not implemented"; return 0 }

func (pr *PeerReplicate) send(msgs []raftpb.Message) { _ = "STUB: not implemented"; return }

// We don't care that the message is sent failed, so here just log this error

func (pr *PeerReplicate) sendRaftMsg(msg raftpb.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// There could be two cases:
// 1. Target peer already exists but has not established communication with leader yet
// 2. Target peer is added newly due to member change or region split, but it's not
//    created yet
// For both cases the region start key and end key are attached in RequestVote and
// Heartbeat message for the store of that peer to check whether to create a new peer
// when receiving these messages, or just to wait for a pending region split to perform
// later.

// the peer has not been known to this leader, it may exist or not.

func (pr *PeerReplicate) isRead(req *raftcmdpb.Request) bool {
	_ = "STUB: not implemented"
	return false
}

func (pr *PeerReplicate) isWrite(req *raftcmdpb.Request) bool {
	_ = "STUB: not implemented"
	return false
}

func (pr *PeerReplicate) getHandlePolicy(req *raftcmdpb.RaftCMDRequest) (requestPolicy, error) {
	_ = "STUB: not implemented"
	return *new(requestPolicy), nil
}

func (pr *PeerReplicate) getCurrentTerm() uint64 { _ = "STUB: not implemented"; return 0 }

func (pr *PeerReplicate) step(msg raftpb.Message) { _ = "STUB: not implemented"; return }

func (pr *PeerReplicate) report(report interface{}) { _ = "STUB: not implemented"; return }

func getRaftConfig(id, appliedIndex uint64, store raft.Storage) *raft.Config {
	_ = "STUB: not implemented"
	return nil
}
