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
)

func newStoreInfo(store metapb.Store) *StoreInfo { _ = "STUB: not implemented"; return nil }

func newStoreCache() *storeCache { _ = "STUB: not implemented"; return nil }

type storeCache struct {
	sync.RWMutex
	stores map[uint64]*StoreInfo
}

func (sc *storeCache) createStoreInfo(store metapb.Store) { _ = "STUB: not implemented"; return }

func (sc *storeCache) updateStoreInfo(store *StoreInfo) { _ = "STUB: not implemented"; return }

func (sc *storeCache) getStores() []*StoreInfo { _ = "STUB: not implemented"; return nil }

func (sc *storeCache) getStore(storeID uint64) *StoreInfo { _ = "STUB: not implemented"; return nil }

func (sc *storeCache) setStoreOffline(storeID uint64) (*StoreInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sc *storeCache) setStoreTombstone(storeID uint64, force bool) (*StoreInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sc *storeCache) getCellStores(cell *CellInfo) []*StoreInfo {
	_ = "STUB: not implemented"
	return nil
}

func (sc *storeCache) getFollowerStores(cell *CellInfo) []*StoreInfo {
	_ = "STUB: not implemented"
	return nil
}

func (sc *storeCache) foreach(fn func(*StoreInfo) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (sc *storeCache) getStoreWithoutLock(storeID uint64) *StoreInfo {
	_ = "STUB: not implemented"
	return nil
}

func newStoreStatus() *StoreStatus { _ = "STUB: not implemented"; return nil }

// StoreStatus contains information about a store's status.
type StoreStatus struct {
	Stats           *pdpb.StoreStats
	LeaderCount     uint32
	LastHeartbeatTS time.Time

	// Blocked means that the store is blocked from balance.
	blocked bool
}

// StoreInfo store info
type StoreInfo struct {
	Meta   metapb.Store
	Status *StoreStatus
}

func (s *StoreInfo) getID() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *StoreInfo) isUp() bool { _ = "STUB: not implemented"; return false }

func (s *StoreInfo) isTombstone() bool { _ = "STUB: not implemented"; return false }

func (s *StoreInfo) isOffline() bool { _ = "STUB: not implemented"; return false }

func (s *StoreInfo) isBlocked() bool { _ = "STUB: not implemented"; return false }

func (s *StoreInfo) downTime() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (s *StoreInfo) resourceCount(kind ResourceKind) uint64 { _ = "STUB: not implemented"; return 0 }

func (s *StoreInfo) resourceScore(kind ResourceKind) float64 { _ = "STUB: not implemented"; return 0 }

func (s *StoreInfo) leaderCount() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *StoreInfo) leaderScore() float64 { _ = "STUB: not implemented"; return 0 }

func (s *StoreInfo) cellCount() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *StoreInfo) cellScore() float64 { _ = "STUB: not implemented"; return 0 }

func (s *StoreInfo) storageRatio() int { _ = "STUB: not implemented"; return 0 }

func (s *StoreInfo) storageSize() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *StoreInfo) getLocationID(keys []string) string { _ = "STUB: not implemented"; return "" }

func (s *StoreInfo) getLabelValue(key string) string { _ = "STUB: not implemented"; return "" }
