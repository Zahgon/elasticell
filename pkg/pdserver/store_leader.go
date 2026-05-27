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
	"github.com/coreos/etcd/clientv3"
	"github.com/deepfabric/elasticell/pkg/pb/pdpb"
)

// GetCurrentLeader return current leader
func (s *pdStore) GetCurrentLeader() (*pdpb.Leader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ResignLeader delete leader itself and let others start a new election again.
func (s *pdStore) ResignLeader(leaderSignature string) error { _ = "STUB: not implemented"; return nil }

// WatchLeader watch leader,
// this funcation will return unitl the leader's lease is timeout
// or server closed
func (s *pdStore) WatchLeader() { _ = "STUB: not implemented"; return }

// server closed, return

// CampaignLeader is for leader election
// if we are win the leader election, the enableLeaderFun will call
func (s *pdStore) CampaignLeader(leaderSignature string, leaderLeaseTTL int64, enableLeaderFun, disableLeaderFun func()) error {
	_ = "STUB: not implemented"
	return nil
}

// The leader key must not exist, so the CreateRevision is 0.

// Make the leader keepalived.

// txn returns an etcd client transaction wrapper.
// The wrapper will set a request timeout to the context and log slow transactions.
func (s *pdStore) txn() clientv3.Txn { _ = "STUB: not implemented"; return *new(clientv3.Txn) }

func (s *pdStore) leaderTxn(leaderSignature string, cs ...clientv3.Cmp) clientv3.Txn {
	_ = "STUB: not implemented"
	return *new(clientv3.Txn)
}

func (s *pdStore) leaderCmp(leaderSignature string) clientv3.Cmp {
	_ = "STUB: not implemented"
	return *new(clientv3.Cmp)
}
