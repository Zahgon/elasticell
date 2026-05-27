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
	"math"

	"github.com/coreos/etcd/clientv3"
	"github.com/deepfabric/elasticell/pkg/pb/metapb"
	"github.com/deepfabric/elasticell/pkg/pb/pdpb"
)

var (
	endID = uint64(math.MaxUint64)
)

// GetCurrentClusterMembers returns members in current etcd cluster
func (s *pdStore) GetCurrentClusterMembers() (*clientv3.MemberListResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetInitParams returns init params
func (s *pdStore) GetInitParams(clusterID uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetClusterID returns current cluster id
// if cluster is not init, return 0
func (s *pdStore) GetClusterID() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// If the key is "pdClusterIDPath", parse the cluster ID from it.

// Parse the cluster ID from any other keys for compatibility.

// CreateFirstClusterID create the first cluster
// More than one pd instance do this operation at the first time,
// only one can succ,
// others will get the committed id.
func (s *pdStore) CreateFirstClusterID() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// Generate a random cluster ID.

// Txn commits ok, return the generated cluster ID.

// Otherwise, parse the committed cluster ID.

// SetInitParams store the init cluster params
func (s *pdStore) SetInitParams(clusterID uint64, params string) error {
	_ = "STUB: not implemented"
	return nil
}

// SetClusterBootstrapped set cluster bootstrapped flag, only one can succ.
func (s *pdStore) SetClusterBootstrapped(clusterID uint64, cluster metapb.Cluster, store metapb.Store, cells []metapb.Cell) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// build operations

// txn

// already bootstrapped

// LoadClusterMeta returns cluster meta info
func (s *pdStore) LoadClusterMeta(clusterID uint64) (*metapb.Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadStoreMeta returns load error,
// do funcation will call on each loaded store meta info
func (s *pdStore) LoadStoreMeta(clusterID uint64, limit int64, do func(metapb.Store)) error {
	_ = "STUB: not implemented"
	return nil
}

// read complete

// LoadCellMeta returns load error,
// do funcation will call on each loaded cell meta info
func (s *pdStore) LoadCellMeta(clusterID uint64, limit int64, do func(metapb.Cell)) error {
	_ = "STUB: not implemented"
	return nil
}

// read complete

func (s *pdStore) LoadWatchers(clusterID uint64, limit int64, do func(pdpb.Watcher)) error {
	_ = "STUB: not implemented"
	return nil
}

// read complete

// SetStoreMeta returns nil if store is add or update succ
func (s *pdStore) SetStoreMeta(clusterID uint64, store metapb.Store) error {
	_ = "STUB: not implemented"
	return nil
}

// SetCellMeta returns nil if cell is add or update succ
func (s *pdStore) SetCellMeta(clusterID uint64, cell metapb.Cell) error {
	_ = "STUB: not implemented"
	return nil
}

// SetWatchers returns nil if add or update succ
func (s *pdStore) SetWatchers(clusterID uint64, watcher pdpb.Watcher) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *pdStore) getClusterMetaKey(clusterID uint64) string { _ = "STUB: not implemented"; return "" }

func (s *pdStore) getStoreMetaKey(clusterID, storeID uint64) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *pdStore) getInitParamsKey(clusterID uint64) string { _ = "STUB: not implemented"; return "" }

func (s *pdStore) getCellMetaKey(clusterID, cellID uint64) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *pdStore) getWatcherMetaKey(clusterID uint64, addr string) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *pdStore) getMaxWatcherMetaKey(clusterID uint64) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *pdStore) getMinWatcher() string { _ = "STUB: not implemented"; return "" }
