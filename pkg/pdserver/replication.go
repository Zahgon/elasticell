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
)

const replicaBaseScore = 100

type replicaChecker struct {
	cfg     *Cfg
	cache   *cache
	filters []Filter
}

func newReplicaChecker(cfg *Cfg, cache *cache, filters ...Filter) *replicaChecker {
	_ = "STUB: not implemented"
	return nil
}

// Check return the Operator
func (r *replicaChecker) Check(target *CellInfo) Operator {
	_ = "STUB: not implemented"
	return *new(Operator)
}

func (r *replicaChecker) checkDownPeer(cell *CellInfo) Operator {
	_ = "STUB: not implemented"
	return *new(Operator)
}

func (r *replicaChecker) checkOfflinePeer(cell *CellInfo) Operator {
	_ = "STUB: not implemented"
	return *new(Operator)
}

// selectWorstPeer returns the worst peer in the cell.
func (r *replicaChecker) selectWorstPeer(cell *CellInfo, filters ...Filter) (*meta.Peer, float64) {
	_ = "STUB: not implemented"
	return nil, 0
}

// Select the store with lowest distinct score.
// If the scores are the same, select the store with maximal cell score.

// selectBestPeer returns the best peer in other stores.
func (r *replicaChecker) selectBestPeer(target *CellInfo, allocPeerID bool, filters ...Filter) (*meta.Peer, float64) {
	_ = "STUB: not implemented"
	// Add some must have filters.
	return nil, 0
}

// Select the store with best distinct score.
// If the scores are the same, select the store with minimal cells score.

func (r *replicaChecker) checkBestReplacement(cell *CellInfo) Operator {
	_ = "STUB: not implemented"
	return *new(Operator)
}

// Make sure the new peer is better than the old peer.

// selectBestReplacement returns the best peer to replace the cell peer.
func (r *replicaChecker) selectBestReplacement(cell *CellInfo, peer *meta.Peer) (*meta.Peer, float64) {
	_ = "STUB: not implemented"
	// selectBestReplacement returns the best peer to replace the cell peer.
	// Get a new cell without the peer we are going to replace.
	return nil, 0
}

// getDistinctScore returns the score that the other is distinct from the stores.
// A higher score means the other store is more different from the existed stores.
func getDistinctScore(cfg *Cfg, stores []*StoreInfo, other *StoreInfo) float64 {
	_ = "STUB: not implemented"
	return 0
}

// compareStoreScore compares which store is better for replication.
// Returns 0 if store A is as good as store B.
// Returns 1 if store A is better than store B.
// Returns -1 if store B is better than store A.
func compareStoreScore(cfg *Cfg, storeA *StoreInfo, scoreA float64, storeB *StoreInfo, scoreB float64) int {
	_ = "STUB: not implemented"
	// The store with higher score is better.
	return 0
}

// The store with lower region score is better.
