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
	"context"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/embed"
	"github.com/deepfabric/elasticell/pkg/pb/metapb"
	"github.com/deepfabric/elasticell/pkg/pb/pdpb"
)

const (
	// DefaultTimeout default timeout
	DefaultTimeout = time.Second * 3
	// DefaultRequestTimeout default request timeout
	DefaultRequestTimeout = 10 * time.Second
	// DefaultSlowRequestTime default slow request time
	DefaultSlowRequestTime = time.Second * 1

	pdRootPath         = "/pd"
	pdIDPath           = "/pd/meta/id"
	pdLeaderPath       = "/pd/meta/leader"
	pdClusterIDPath    = "/pd/meta/cluster_id"
	pdBootstrappedPath = "/pd/meta/boot"
	pdClusterRootPath  = "/pd/cluster"
	pdIndicesRootPath  = "/pd/indices"
)

// ClusterStore is the store interface for cluster info
type ClusterStore interface {
	SetInitParams(clusterID uint64, params string) error
	GetInitParams(clusterID uint64) ([]byte, error)
	GetCurrentClusterMembers() (*clientv3.MemberListResponse, error)
	GetClusterID() (uint64, error)
	CreateFirstClusterID() (uint64, error)
	SetClusterBootstrapped(clusterID uint64, cluster metapb.Cluster, store metapb.Store, cells []metapb.Cell) (bool, error)
	LoadClusterMeta(clusterID uint64) (*metapb.Cluster, error)
	LoadStoreMeta(clusterID uint64, limit int64, do func(metapb.Store)) error
	LoadCellMeta(clusterID uint64, limit int64, do func(metapb.Cell)) error
	LoadWatchers(clusterID uint64, limit int64, do func(pdpb.Watcher)) error
	SetStoreMeta(clusterID uint64, store metapb.Store) error
	SetCellMeta(clusterID uint64, cell metapb.Cell) error
	SetWatchers(clusterID uint64, watcher pdpb.Watcher) error
}

// LeaderStore is the store interface for leader info
type LeaderStore interface {
	// CampaignLeader is for leader election
	// if we are win the leader election, the enableLeaderFun will call, otherwrise disableLeaderFun will call
	CampaignLeader(leaderSignature string, leaderLeaseTTL int64, enableLeaderFun, disableLeaderFun func()) error
	// WatchLeader watch leader,
	// this funcation will return unitl the leader's lease is timeout
	// or server closed
	WatchLeader()
	// ResignLeader delete leader itself and let others start a new election again.
	ResignLeader(leaderSignature string) error
	// GetCurrentLeader return current leader
	GetCurrentLeader() (*pdpb.Leader, error)
}

// IDStore is the store interface for id info
type IDStore interface {
	GetID() (uint64, error)
	CreateID(leaderSignature string, value uint64) error
	UpdateID(leaderSignature string, old, value uint64) error
}

// IndexStore is the store interface for index info
type IndexStore interface {
	ListIndex() (idxDefs []*pdpb.IndexDef, err error)
	GetIndex(id string) (idxDef *pdpb.IndexDef, err error)
	CreateIndex(idxDef *pdpb.IndexDef) (err error)
	DeleteIndex(id string) (err error)
}

// Store is the store interface for all pd store info
type Store interface {
	ClusterStore
	IDStore
	LeaderStore
	IndexStore

	Close() error
	RawClient() *clientv3.Client
}

// Store used for  metedata
type pdStore struct {
	client *clientv3.Client
}

// NewStore create a store
func NewStore(cfg *embed.Config) (Store, error) { _ = "STUB: not implemented"; return *new(Store), nil }

func initEtcdClient(cfg *embed.Config) (*clientv3.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close close etcd client
func (s *pdStore) Close() error { _ = "STUB: not implemented"; return nil }

// RawClient return raw etcd client
func (s *pdStore) RawClient() *clientv3.Client {
	_ = "STUB: not implemented"

	// slowLogTxn wraps etcd transaction and log slow one.
	return nil
}

type slowLogTxn struct {
	clientv3.Txn
	cancel context.CancelFunc
}

func newSlowLogTxn(client *clientv3.Client) clientv3.Txn {
	_ = "STUB: not implemented"
	return *new(clientv3.Txn)
}

func (t *slowLogTxn) If(cs ...clientv3.Cmp) clientv3.Txn {
	_ = "STUB: not implemented"
	return *new(clientv3.Txn)
}

func (t *slowLogTxn) Then(ops ...clientv3.Op) clientv3.Txn {
	_ = "STUB: not implemented"
	return *new(clientv3.Txn)
}

// Commit implements Txn Commit interface.
func (t *slowLogTxn) Commit() (*clientv3.TxnResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *pdStore) getValue(key string, opts ...clientv3.OpOption) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *pdStore) get(key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *pdStore) save(key, value string) error { _ = "STUB: not implemented"; return nil }

func (s *pdStore) create(key, value string) error { _ = "STUB: not implemented"; return nil }

func (s *pdStore) delete(key string, opts ...clientv3.OpOption) error {
	_ = "STUB: not implemented"
	return nil
}
