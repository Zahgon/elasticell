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
	"sync"

	"github.com/deepfabric/elasticell/pkg/pb/metapb"
)

const (
	batchLimit = 10000
)

func newCache(clusterID uint64, store Store, allocator *idAllocator, notify *watcherNotifier) *cache {
	_ = "STUB: not implemented"
	return nil
}

func newClusterRuntime(cluster metapb.Cluster) *ClusterInfo { _ = "STUB: not implemented"; return nil }

// ClusterInfo The cluster info
type ClusterInfo struct {
	Meta metapb.Cluster `json:"meta"`
}

type cache struct {
	sync.RWMutex

	clusterID uint64
	cluster   *ClusterInfo
	sc        *storeCache
	cc        *cellCache

	allocator *idAllocator
	notify    *watcherNotifier

	store Store
}

func (c *cache) getStoreCache() *storeCache { _ = "STUB: not implemented"; return nil }

func (c *cache) getCellCache() *cellCache { _ = "STUB: not implemented"; return nil }

func (c *cache) allocPeer(storeID uint64, allocPeerID bool) (metapb.Peer, error) {
	_ = "STUB: not implemented"
	return *new(metapb.Peer), nil
}

func (c *cache) handleCellHeartbeat(source *CellInfo) error { _ = "STUB: not implemented"; return nil }

// add new cell

// update cell

// cell meta is stale, return an error.

// cell meta is updated, update kv and cache.

// cell meta is the same, update cache only.

func (c *cache) doSaveCellInfo(source *CellInfo) error { _ = "STUB: not implemented"; return nil }

func (c *cache) notifyStoreRange(id uint64) { _ = "STUB: not implemented"; return }

func (c *cache) notifyChangedRange(id uint64, event uint32) { _ = "STUB: not implemented"; return }

func randCell(cells map[uint64]*CellInfo) *CellInfo { _ = "STUB: not implemented"; return nil }
