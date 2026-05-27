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
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/deepfabric/elasticell/pkg/pb/metapb"
)

func (pr *PeerReplicate) startApplyingSnapJob() { _ = "STUB: not implemented"; return }

func (ps *peerStorage) startDestroyDataJob(cellID uint64, start, end []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (pr *PeerReplicate) startRegistrationJob() { _ = "STUB: not implemented"; return }

func (pr *PeerReplicate) startApplyCommittedEntriesJob(cellID uint64, term uint64, commitedEntries []raftpb.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

func (pr *PeerReplicate) startRaftLogGCJob(cellID, startIndex, endIndex uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) startDestroyJob(cellID uint64, peer metapb.Peer) error {
	_ = "STUB: not implemented"
	return nil
}

func (pr *PeerReplicate) startProposeJob(c *cmd, isConfChange bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (pr *PeerReplicate) startSplitCheckJob() error { _ = "STUB: not implemented"; return nil }

func (pr *PeerReplicate) startAskSplitJob(cell metapb.Cell, peer metapb.Peer, splitKey []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) startReportSpltJob(left metapb.Cell, right metapb.Cell) error {
	_ = "STUB: not implemented"
	return nil
}

func (ps *peerStorage) cancelApplyingSnapJob() bool { _ = "STUB: not implemented"; return false }

func (ps *peerStorage) resetApplyingSnapJob() { _ = "STUB: not implemented"; return }

func (ps *peerStorage) resetGenSnapJob() { _ = "STUB: not implemented"; return }

func (ps *peerStorage) doDestroyDataJob(cellID uint64, startKey, endKey []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (pr *PeerReplicate) doApplyingSnapshotJob() error { _ = "STUB: not implemented"; return nil }

func (ps *peerStorage) doGenerateSnapshotJob() error { _ = "STUB: not implemented"; return nil }

func (pr *PeerReplicate) doRegistrationJob(delegate *applyDelegate) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) doDestroy(cellID uint64, peer metapb.Peer) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: think send notify, then liner process this and other apply result

func (pr *PeerReplicate) doApplyCommittedEntries(cellID uint64, term uint64, commitedEntries []raftpb.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

func (pr *PeerReplicate) doRaftLogGC(cellID, startIndex, endIndex uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (ps *peerStorage) isApplyingSnap() bool { _ = "STUB: not implemented"; return false }

func (ps *peerStorage) isGeneratingSnap() bool { _ = "STUB: not implemented"; return false }

func (ps *peerStorage) isGenSnapJobComplete() bool { _ = "STUB: not implemented"; return false }
