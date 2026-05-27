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

type balanceCellScheduler struct {
	cfg      *Cfg
	cache    *idCache
	limit    uint64
	selector Selector
}

func newBalanceCellScheduler(cfg *Cfg) *balanceCellScheduler { _ = "STUB: not implemented"; return nil }

func (s *balanceCellScheduler) GetName() string { _ = "STUB: not implemented"; return "" }

func (s *balanceCellScheduler) GetResourceKind() ResourceKind {
	_ = "STUB: not implemented"
	return *new(ResourceKind)
}

func (s *balanceCellScheduler) GetResourceLimit() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *balanceCellScheduler) Prepare(cache *cache) error { _ = "STUB: not implemented"; return nil }

func (s *balanceCellScheduler) Cleanup(cache *cache) { _ = "STUB: not implemented"; return }

func (s *balanceCellScheduler) Schedule(cache *cache) Operator {
	_ = "STUB: not implemented"
	// Select a peer from the store with most cells.
	return *new(Operator)
}

// We don't schedule cell with abnormal number of replicas.

// We can't transfer peer from this store now, so we add it to the cache
// and skip it for a while.

func (s *balanceCellScheduler) transferPeer(cache *cache, cell *CellInfo, oldPeer *metapb.Peer) Operator {
	_ = "STUB: not implemented"
	// scoreGuard guarantees that the distinct score will not decrease.
	return *new(Operator)
}
