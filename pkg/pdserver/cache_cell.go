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
	"time"

	"github.com/deepfabric/elasticell/pkg/pb/metapb"
	"github.com/deepfabric/elasticell/pkg/pb/pdpb"
	"github.com/deepfabric/elasticell/pkg/util"
)

// CellInfo The cell info
type CellInfo struct {
	Meta         metapb.Cell
	LeaderPeer   *metapb.Peer
	DownPeers    []pdpb.PeerStats
	PendingPeers []metapb.Peer
}

func newCellInfo(cell metapb.Cell, leader *metapb.Peer) *CellInfo {
	_ = "STUB: not implemented"
	return nil
}

type cellCache struct {
	sync.RWMutex

	tree      *util.CellTree
	cells     map[uint64]*CellInfo            // cellID -> cellRuntimeInfo
	leaders   map[uint64]map[uint64]*CellInfo // storeID -> cellID -> cellRuntimeInfo
	followers map[uint64]map[uint64]*CellInfo // storeID -> cellID -> cellRuntimeInfo
}

func newCellCache() *cellCache { _ = "STUB: not implemented"; return nil }

func (cc *cellCache) createAndAdd(cell metapb.Cell) { _ = "STUB: not implemented"; return }

func (cc *cellCache) addOrUpdate(cellInfo *CellInfo) { _ = "STUB: not implemented"; return }

func (cc *cellCache) addOrUpdateWithoutLock(origin *CellInfo) { _ = "STUB: not implemented"; return }

// Add to tree and regions.

// Add to leaders and followers.

// Add leader peer to leaders.

// Add follower peer to followers.

func (cc *cellCache) removeCell(origin *CellInfo) {
	_ = "STUB: not implemented"
	// Remove from tree and cells.
	return
}

// Remove from leaders and followers.

func (cc *cellCache) searchCell(startKey []byte) *CellInfo { _ = "STUB: not implemented"; return nil }

func (cc *cellCache) foreach(fn func(*CellInfo) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc *cellCache) getCells() []*CellInfo { _ = "STUB: not implemented"; return nil }

func (cc *cellCache) getCell(id uint64) *CellInfo { _ = "STUB: not implemented"; return nil }

func (cc *cellCache) randFollowerCell(storeID uint64) *CellInfo {
	_ = "STUB: not implemented"
	return nil
}

func (cc *cellCache) randLeaderCell(storeID uint64) *CellInfo {
	_ = "STUB: not implemented"
	return nil
}

func (cc *cellCache) getStoreLeaderCount(storeID uint64) int { _ = "STUB: not implemented"; return 0 }

func (cc *CellInfo) getFollowers() map[uint64]*metapb.Peer { _ = "STUB: not implemented"; return nil }

func (cc *CellInfo) getStorePeer(storeID uint64) *metapb.Peer {
	_ = "STUB: not implemented"
	return nil
}

func (cc *CellInfo) removeStorePeer(storeID uint64) { _ = "STUB: not implemented"; return }

func (cc *CellInfo) getPendingPeer(peerID uint64) *metapb.Peer {
	_ = "STUB: not implemented"
	return nil
}

func (cc *CellInfo) getPeer(peerID uint64) *metapb.Peer { _ = "STUB: not implemented"; return nil }

func (cc *CellInfo) getID() uint64 { _ = "STUB: not implemented"; return 0 }

func (cc *CellInfo) getPeers() []*metapb.Peer { _ = "STUB: not implemented"; return nil }

func (cc *CellInfo) getStoreIDs() map[uint64]struct{} { _ = "STUB: not implemented"; return nil }

type idCache struct {
	*expireCellCache
}

func newIDCache(interval, ttl time.Duration) *idCache { _ = "STUB: not implemented"; return nil }

func (c *idCache) set(id uint64) { _ = "STUB: not implemented"; return }

func (c *idCache) get(id uint64) bool { _ = "STUB: not implemented"; return false }

type cacheItem struct {
	key    uint64
	value  interface{}
	expire time.Time
}

// expireCellCache is an expired region cache.
type expireCellCache struct {
	sync.RWMutex

	items      map[uint64]cacheItem
	ttl        time.Duration
	gcInterval time.Duration
}

// newExpireCellCache returns a new expired region cache.
func neweExpireCellCache(gcInterval time.Duration, ttl time.Duration) *expireCellCache {
	_ = "STUB: not implemented"
	return nil
}

func (c *expireCellCache) get(key uint64) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *expireCellCache) set(key uint64, value interface{}) { _ = "STUB: not implemented"; return }

func (c *expireCellCache) setWithTTL(key uint64, value interface{}, ttl time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (c *expireCellCache) delete(key uint64) { _ = "STUB: not implemented"; return }

func (c *expireCellCache) count() int { _ = "STUB: not implemented"; return 0 }

func (c *expireCellCache) doGC() { _ = "STUB: not implemented"; return }
