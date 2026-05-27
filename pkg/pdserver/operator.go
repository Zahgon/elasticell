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

package pdserver

import (
	meta "github.com/deepfabric/elasticell/pkg/pb/metapb"
	"github.com/deepfabric/elasticell/pkg/pb/pdpb"
)

// Operator is an interface to scheduler cell
type Operator interface {
	GetCellID() uint64
	GetResourceKind() ResourceKind
	Do(cell *CellInfo) (*pdpb.CellHeartbeatRsp, bool)
}

func newAddPeerAggregationOp(cell *CellInfo, peer *meta.Peer) Operator {
	_ = "STUB: not implemented"
	return *new(Operator)
}

func newTransferLeaderAggregationOp(cell *CellInfo, newLeader *meta.Peer) Operator {
	_ = "STUB: not implemented"
	return *new(Operator)
}

func newTransferPeerAggregationOp(cell *CellInfo, oldPeer, newPeer *meta.Peer) Operator {
	_ = "STUB: not implemented"
	return *new(Operator)
}

func newAddPeerOp(cellID uint64, peer *meta.Peer) *changePeerOperator {
	_ = "STUB: not implemented"
	return nil
}

func newRemovePeerOp(cellID uint64, peer *meta.Peer) *changePeerOperator {
	_ = "STUB: not implemented"
	return nil
}

func newAggregationOp(cell *CellInfo, ops ...Operator) *aggregationOperator {
	_ = "STUB: not implemented"
	return nil
}
