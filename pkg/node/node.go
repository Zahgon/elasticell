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

package node

import (
	"sync"

	"github.com/deepfabric/elasticell/pkg/pb/metapb"
	"github.com/deepfabric/elasticell/pkg/pd"
	"github.com/deepfabric/elasticell/pkg/pdapi"
	"github.com/deepfabric/elasticell/pkg/raftstore"
	"github.com/deepfabric/elasticell/pkg/storage"
	"github.com/fagongzi/util/task"
)

// Node node
type Node struct {
	sync.RWMutex

	cfg         *Cfg
	clusterID   uint64
	pdClient    *pd.Client
	drivers     []storage.Driver
	driversMask uint64
	storeMeta   metapb.Store
	store       *raftstore.Store

	runner *task.Runner
}

// NewNode create a node instance, then init store, pd connection and init the cluster ID
func NewNode(clientAddr string, cfg *Cfg, drivers []storage.Driver) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start start the node.
// if cluster is not bootstrapped, bootstrap cluster and create the first cell.
func (n *Node) Start() *raftstore.Store { _ = "STUB: not implemented"; return nil }

// Stop the node
func (n *Node) Stop() error { _ = "STUB: not implemented"; return nil }

func (n *Node) closePDClient() { _ = "STUB: not implemented"; return }

func (n *Node) initPDClient() error { _ = "STUB: not implemented"; return nil }

func (n *Node) getAllocID() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (n *Node) getInitParam() (*pdapi.InitParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newStore(clientAddr string, cfg *Cfg) metapb.Store {
	_ = "STUB: not implemented"
	return *new(metapb.Store)
}
