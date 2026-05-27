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
	"github.com/deepfabric/elasticell/pkg/pb/metapb"
)

const (
	balanceLeaderSchedulerName = "balance-leader-scheduler"
)

type balanceLeaderScheduler struct {
	cfg      *Cfg
	limit    uint64
	selector Selector
}

func newBalanceLeaderScheduler(cfg *Cfg) *balanceLeaderScheduler {
	_ = "STUB: not implemented"
	return nil
}

func (l *balanceLeaderScheduler) GetName() string { _ = "STUB: not implemented"; return "" }

func (l *balanceLeaderScheduler) GetResourceKind() ResourceKind {
	_ = "STUB: not implemented"
	return *new(ResourceKind)
}

func (l *balanceLeaderScheduler) GetResourceLimit() uint64 { _ = "STUB: not implemented"; return 0 }

func (l *balanceLeaderScheduler) Prepare(cache *cache) error { _ = "STUB: not implemented"; return nil }

func (l *balanceLeaderScheduler) Cleanup(cache *cache) { _ = "STUB: not implemented"; return }

func (l *balanceLeaderScheduler) Schedule(cache *cache) Operator {
	_ = "STUB: not implemented"
	return *new(Operator)
}

// scheduleTransferLeader schedules a cell to transfer leader to the peer.
func scheduleTransferLeader(cache *cache, s Selector, filters ...Filter) (*CellInfo, *metapb.Peer) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transfer a leader out of mostLeaderStore.

// Transfer a leader into leastLeaderStore.
