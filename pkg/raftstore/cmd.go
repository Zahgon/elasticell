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
	"github.com/deepfabric/elasticell/pkg/pb/metapb"
	"github.com/deepfabric/elasticell/pkg/pb/raftcmdpb"
)

type cmd struct {
	req  *raftcmdpb.RaftCMDRequest
	cb   func(*raftcmdpb.RaftCMDResponse)
	term uint64
}

func (c *cmd) reset() { _ = "STUB: not implemented"; return }

func newCMD(req *raftcmdpb.RaftCMDRequest, cb func(*raftcmdpb.RaftCMDResponse)) *cmd {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) respStoreNotMatch(err error, req *raftcmdpb.Request, cb func(*raftcmdpb.RaftCMDResponse)) {
	_ = "STUB: not implemented"
	return
}

func (c *cmd) resp(resp *raftcmdpb.RaftCMDResponse) { _ = "STUB: not implemented"; return }

func (c *cmd) release() { _ = "STUB: not implemented"; return }

func (c *cmd) respCellNotFound(cellID uint64) { _ = "STUB: not implemented"; return }

func (c *cmd) respLargeRaftEntrySize(cellID uint64, size uint64) { _ = "STUB: not implemented"; return }

func (c *cmd) respOtherError(err error) { _ = "STUB: not implemented"; return }

func (c *cmd) respNotLeader(cellID uint64, leader metapb.Peer) { _ = "STUB: not implemented"; return }

func (c *cmd) getUUID() []byte { _ = "STUB: not implemented"; return nil }

func (pr *PeerReplicate) execReadLocal(c *cmd) { _ = "STUB: not implemented"; return }

func (pr *PeerReplicate) execReadIndex(c *cmd) { _ = "STUB: not implemented"; return }
