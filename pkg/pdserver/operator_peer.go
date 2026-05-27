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
	"github.com/deepfabric/elasticell/pkg/pb/pdpb"
)

// changePeerOperator is sub operator of cellOperator
type changePeerOperator struct {
	Name       string          `json:"name"`
	CellID     uint64          `json:"cellID"`
	ChangePeer pdpb.ChangePeer `json:"changePeer"`
}

func (op *changePeerOperator) String() string { _ = "STUB: not implemented"; return "" }

func (op *changePeerOperator) GetCellID() uint64 { _ = "STUB: not implemented"; return 0 }

func (op *changePeerOperator) GetResourceKind() ResourceKind {
	_ = "STUB: not implemented"
	return *new(ResourceKind)
}

func (op *changePeerOperator) Do(cell *CellInfo) (*pdpb.CellHeartbeatRsp, bool) {
	_ = "STUB: not implemented"
	// Check if operator is finished.
	return nil, false
}

// Peer is added but not finished.

// Peer is added and finished.

// Peer is removed.
