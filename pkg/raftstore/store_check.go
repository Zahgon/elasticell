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
	"github.com/deepfabric/elasticell/pkg/pb/errorpb"
	"github.com/deepfabric/elasticell/pkg/pb/metapb"
	"github.com/deepfabric/elasticell/pkg/pb/mraft"
	"github.com/deepfabric/elasticell/pkg/pb/raftcmdpb"
)

func (s *Store) isRaftMsgValid(msg *mraft.RaftMessage) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *Store) isMsgStale(msg *mraft.RaftMessage) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Let's consider following cases with three nodes [1, 2, 3] and 1 is leader:
// a. 1 removes 2, 2 may still send MsgAppendResponse to 1.
//  We should ignore this stale message and let 2 remove itself after
//  applying the ConfChange log.
// b. 2 is isolated, 1 removes 2. When 2 rejoins the cluster, 2 will
//  send stale MsgRequestVote to 1 and 3, at this time, we should tell 2 to gc itself.
// c. 2 is isolated but can communicate with 3. 1 removes 3.
//  2 will send stale MsgRequestVote to 3, 3 should ignore this message.
// d. 2 is isolated but can communicate with 3. 1 removes 2, then adds 4, remove 3.
//  2 will send stale MsgRequestVote to 3, 3 should tell 2 to gc itself.
// e. 2 is isolated. 1 adds 4, 5, 6, removes 3, 1. Now assume 4 is leader.
//  After 2 rejoins the cluster, 2 may send stale MsgRequestVote to 1 and 3,
//  1 and 3 will ignore this message. Later 4 will send messages to 2 and 2 will
//  rejoin the raft group again.
// f. 2 is isolated. 1 adds 4, 5, 6, removes 3, 1. Now assume 4 is leader, and 4 removes 2.
//  unlike case e, 2 will be stale forever.
// TODO: for case f, if 2 is stale for a long time, 2 will communicate with pd and pd will
// tell 2 is stale, so 2 can remove itself.

// no exist, check with tombstone key.

// Maybe split, but not registered yet.

// The cell in this peer is already destroyed

func (s *Store) checkSnapshot(msg *mraft.RaftMessage) (bool, error) {
	_ = "STUB: not implemented"
	// Check if we can accept the snapshot
	return false, nil
}

func (s *Store) addPendingSnapshot(msg mraft.SnapshotMessageHeader) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *Store) getPendingSnapshot(id uint64) mraft.SnapshotMessageHeader {
	_ = "STUB: not implemented"
	return *new(mraft.SnapshotMessageHeader)
}

func (s *Store) removePendingSnapshot(id uint64) { _ = "STUB: not implemented"; return }

func (s *Store) hasOverlapInPendingSnapshots(cell *metapb.Cell) bool {
	_ = "STUB: not implemented"
	return false
}

// Same cell can overlap, we will apply the latest version of snapshot.

func checkEpoch(cell metapb.Cell, req *raftcmdpb.RaftCMDRequest) bool {
	_ = "STUB: not implemented"
	return false
}

// for redis command, we don't care conf version.

func (s *Store) validateStoreID(req *raftcmdpb.RaftCMDRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) validateCell(req *raftcmdpb.RaftCMDRequest) *errorpb.Error {
	_ = "STUB: not implemented"
	return nil
}

// If header's term is 2 verions behind current term,
// leadership may have been changed away.

// Attach the next cell which might be split from the current cell. But it doesn't
// matter if the next cell is not split from the current cell. If the cell meta
// received by the KV driver is newer than the meta cached in the driver, the meta is
// updated.
