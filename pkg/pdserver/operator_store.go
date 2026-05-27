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

type storeOperatorKind int

const (
	setLogLevelKind storeOperatorKind = iota
)

// StoreOperator is an interface to operate store
type StoreOperator interface {
	GetStoreID() uint64
	Do(store *StoreInfo) (*pdpb.StoreHeartbeatRsp, bool)
}

func newSetLogLevelOperator(id uint64, newLevel int32) StoreOperator {
	_ = "STUB: not implemented"
	return *new(StoreOperator)
}

type setLogLevelOperator struct {
	id       uint64
	newLevel int32
}

func (op *setLogLevelOperator) GetStoreID() uint64 { _ = "STUB: not implemented"; return 0 }

func (op *setLogLevelOperator) Do(store *StoreInfo) (*pdpb.StoreHeartbeatRsp, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
