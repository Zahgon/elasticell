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
	"sync"

	"time"

	"github.com/deepfabric/elasticell/pkg/pb/metapb"
)

type cellPeersMap struct {
	sync.RWMutex
	m map[uint64]*PeerReplicate
}

func newCellPeersMap() *cellPeersMap { _ = "STUB: not implemented"; return nil }

func (m *cellPeersMap) size() uint32 { _ = "STUB: not implemented"; return 0 }

func (m *cellPeersMap) put(key uint64, peers *PeerReplicate) { _ = "STUB: not implemented"; return }

func (m *cellPeersMap) get(key uint64) *PeerReplicate { _ = "STUB: not implemented"; return nil }

func (m *cellPeersMap) delete(key uint64) *PeerReplicate { _ = "STUB: not implemented"; return nil }

func (m *cellPeersMap) foreach(fn func(*PeerReplicate) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *cellPeersMap) values() []*PeerReplicate { _ = "STUB: not implemented"; return nil }

type peerCacheMap struct {
	sync.RWMutex
	m map[uint64]metapb.Peer
}

func newPeerCacheMap() *peerCacheMap { _ = "STUB: not implemented"; return nil }

func (m *peerCacheMap) put(key uint64, peer metapb.Peer) { _ = "STUB: not implemented"; return }

func (m *peerCacheMap) get(key uint64) (metapb.Peer, bool) {
	_ = "STUB: not implemented"
	return *new(metapb.Peer), false
}

func (m *peerCacheMap) delete(key uint64) { _ = "STUB: not implemented"; return }

type applyDelegateMap struct {
	sync.RWMutex
	m map[uint64]*applyDelegate
}

func newApplyDelegateMap() *applyDelegateMap { _ = "STUB: not implemented"; return nil }

func (m *applyDelegateMap) put(key uint64, value *applyDelegate) *applyDelegate {
	_ = "STUB: not implemented"
	return nil
}

func (m *applyDelegateMap) get(key uint64) *applyDelegate { _ = "STUB: not implemented"; return nil }

func (m *applyDelegateMap) delete(key uint64) *applyDelegate { _ = "STUB: not implemented"; return nil }

type peerHeartbeatsMap struct {
	sync.RWMutex
	m map[uint64]time.Time
}

func newPeerHeartbeatsMap() *peerHeartbeatsMap { _ = "STUB: not implemented"; return nil }

func (m *peerHeartbeatsMap) clear() { _ = "STUB: not implemented"; return }

func (m *peerHeartbeatsMap) has(key uint64) bool { _ = "STUB: not implemented"; return false }

func (m *peerHeartbeatsMap) put(key uint64, value time.Time) { _ = "STUB: not implemented"; return }

func (m *peerHeartbeatsMap) putOnlyNotExist(key uint64, value time.Time) {
	_ = "STUB: not implemented"
	return
}

func (m *peerHeartbeatsMap) size() int { _ = "STUB: not implemented"; return 0 }

func (m *peerHeartbeatsMap) get(key uint64) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (m *peerHeartbeatsMap) delete(key uint64) { _ = "STUB: not implemented"; return }
