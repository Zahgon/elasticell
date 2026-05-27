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

	"github.com/coreos/etcd/embed"
	"github.com/deepfabric/elasticell/pkg/pd"
	"google.golang.org/grpc"
)

// Server the pd server
type Server struct {
	cfg *Cfg

	// etcd fields
	id   uint64
	etcd *embed.Etcd

	// rpc fields
	rpcServer *grpc.Server

	store Store

	// cluster fields
	isLeaderValue   int64
	leaderSignature string
	clusterID       uint64
	cluster         *CellCluster
	leaderProxy     *pd.Client
	leaderMux       sync.RWMutex
	idAlloc         *idAllocator

	notifier *watcherNotifier

	// status
	callStop bool
	closed   int64

	// stop fields
	stopOnce sync.Once
	stopWG   sync.WaitGroup
	stopC    chan interface{}

	complete chan struct{}
}

// NewServer create a pd server
func NewServer(cfg *Cfg) *Server { _ = "STUB: not implemented"; return nil }

// Name returns name of current pd server
func (s *Server) Name() string {
	_ = "STUB: not implemented"

	// Start start the pd server
	return ""
}

func (s *Server) Start() { _ = "STUB: not implemented"; return }

// Stop the server
func (s *Server) Stop() { _ = "STUB: not implemented"; return }

func (s *Server) listenToStop() { _ = "STUB: not implemented"; return }

func (s *Server) doStop() { _ = "STUB: not implemented"; return }

func (s *Server) notifyElectionComplete() { _ = "STUB: not implemented"; return }

// GetCfg returns cfg, just for test
func (s *Server) GetCfg() *Cfg { _ = "STUB: not implemented"; return nil }

func (s *Server) initCluster() { _ = "STUB: not implemented"; return }

func (s *Server) isClosed() bool { _ = "STUB: not implemented"; return false }

func (s *Server) setServerIsStopped() { _ = "STUB: not implemented"; return }

func (s *Server) setServerIsStarted() { _ = "STUB: not implemented"; return }
