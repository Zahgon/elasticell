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
	"github.com/deepfabric/elasticell/pkg/pb/pdpb"
	"github.com/deepfabric/elasticell/pkg/pb/raftcmdpb"
	"github.com/deepfabric/elasticell/pkg/storage"
)

const (
	invalidIndex = 0
)

// check whether epoch is staler than checkEpoch.
func isEpochStale(epoch metapb.CellEpoch, checkEpoch metapb.CellEpoch) bool {
	_ = "STUB: not implemented"
	return false
}

func findPeer(cell *metapb.Cell, storeID uint64) *metapb.Peer {
	_ = "STUB: not implemented"
	return nil
}

func removePeer(cell *metapb.Cell, storeID uint64) { _ = "STUB: not implemented"; return }

func newPeer(peerID, storeID uint64) metapb.Peer {
	_ = "STUB: not implemented"
	return *new(metapb.Peer)
}

func removedPeers(new, old metapb.Cell) []uint64 { _ = "STUB: not implemented"; return nil }

// Check if key in cell range [`startKey`, `endKey`).
func checkKeyInCell(key []byte, cell *metapb.Cell) *errorpb.Error {
	_ = "STUB: not implemented"
	return nil
}

func newChangePeerRequest(changeType pdpb.ConfChangeType, peer metapb.Peer) *raftcmdpb.AdminRequest {
	_ = "STUB: not implemented"
	return nil
}

func newTransferLeaderRequest(rsp *pdpb.TransferLeader) *raftcmdpb.AdminRequest {
	_ = "STUB: not implemented"
	return nil
}

func newCompactLogRequest(index, term uint64) *raftcmdpb.AdminRequest {
	_ = "STUB: not implemented"
	return nil
}

// SaveCell save  cell with state, raft state and apply state.
func SaveCell(driver storage.Driver, cell metapb.Cell) error { _ = "STUB: not implemented"; return nil }

// save state

// DeleteCell delete cell with state, raft state and apply state.
func DeleteCell(id uint64, wb storage.WriteBatch) error {
	_ = "STUB: not implemented"
	// save state
	return nil
}

// HasOverlap check cells has overlap
func HasOverlap(c1, c2 *metapb.Cell) bool { _ = "STUB: not implemented"; return false }

// StalEpoch returns true if the target epoch is stale
func StalEpoch(target, check metapb.CellEpoch) bool { _ = "STUB: not implemented"; return false }
