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

// Filter is used for filter store
type Filter interface {
	FilterSource(store *StoreInfo) bool
	FilterTarget(store *StoreInfo) bool
}

func filterSource(store *StoreInfo, filters []Filter) bool { _ = "STUB: not implemented"; return false }

func filterTarget(store *StoreInfo, filters []Filter) bool { _ = "STUB: not implemented"; return false }

type stateFilter struct {
	cfg *Cfg
}

// storageThresholdFilter ensures that we will not use an almost full store as a target.
type storageThresholdFilter struct {
	cfg *Cfg
}

type excludedFilter struct {
	sources map[uint64]struct{}
	targets map[uint64]struct{}
}

func newStorageThresholdFilter(cfg *Cfg) *storageThresholdFilter {
	_ = "STUB: not implemented"
	return nil
}

func newStateFilter(cfg *Cfg) Filter { _ = "STUB: not implemented"; return *new(Filter) }

func newExcludedFilter(sources, targets map[uint64]struct{}) *excludedFilter {
	_ = "STUB: not implemented"
	return nil
}

func (f *stateFilter) filter(store *StoreInfo) bool { _ = "STUB: not implemented"; return false }

func (f *stateFilter) FilterSource(store *StoreInfo) bool { _ = "STUB: not implemented"; return false }

func (f *stateFilter) FilterTarget(store *StoreInfo) bool { _ = "STUB: not implemented"; return false }

func (f *storageThresholdFilter) FilterSource(store *StoreInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *storageThresholdFilter) FilterTarget(store *StoreInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *excludedFilter) FilterSource(store *StoreInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *excludedFilter) FilterTarget(store *StoreInfo) bool {
	_ = "STUB: not implemented"
	return false
}

type blockFilter struct{}

func newBlockFilter() *blockFilter { _ = "STUB: not implemented"; return nil }

func (f *blockFilter) FilterSource(store *StoreInfo) bool { _ = "STUB: not implemented"; return false }

func (f *blockFilter) FilterTarget(store *StoreInfo) bool { _ = "STUB: not implemented"; return false }

type healthFilter struct {
	cfg *Cfg
}

func newHealthFilter(cfg *Cfg) *healthFilter { _ = "STUB: not implemented"; return nil }

func (f *healthFilter) filter(store *StoreInfo) bool { _ = "STUB: not implemented"; return false }

func (f *healthFilter) FilterSource(store *StoreInfo) bool { _ = "STUB: not implemented"; return false }

func (f *healthFilter) FilterTarget(store *StoreInfo) bool { _ = "STUB: not implemented"; return false }

type cacheFilter struct {
	cache *idCache
}

func newCacheFilter(cache *idCache) *cacheFilter { _ = "STUB: not implemented"; return nil }

func (f *cacheFilter) FilterSource(store *StoreInfo) bool { _ = "STUB: not implemented"; return false }

func (f *cacheFilter) FilterTarget(store *StoreInfo) bool { _ = "STUB: not implemented"; return false }

type snapshotCountFilter struct {
	cfg *Cfg
}

func newSnapshotCountFilter(cfg *Cfg) *snapshotCountFilter { _ = "STUB: not implemented"; return nil }

func (f *snapshotCountFilter) filter(store *StoreInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *snapshotCountFilter) FilterSource(store *StoreInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *snapshotCountFilter) FilterTarget(store *StoreInfo) bool {
	_ = "STUB: not implemented"
	return false

	// distinctScoreFilter ensures that distinct score will not decrease.
}

type distinctScoreFilter struct {
	cfg       *Cfg
	stores    []*StoreInfo
	safeScore float64
}

func newDistinctScoreFilter(cfg *Cfg, stores []*StoreInfo, source *StoreInfo) *distinctScoreFilter {
	_ = "STUB: not implemented"
	return nil
}

func (f *distinctScoreFilter) FilterSource(store *StoreInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *distinctScoreFilter) FilterTarget(store *StoreInfo) bool {
	_ = "STUB: not implemented"
	return false
}
