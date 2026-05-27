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

package server

import (
	"sync"

	"github.com/deepfabric/elasticell/pkg/node"
	"github.com/deepfabric/elasticell/pkg/raftstore"
	"github.com/deepfabric/elasticell/pkg/storage"
	"github.com/fagongzi/util/task"
)

// Server a server provide kv cache based on redis protocol
type Server struct {
	redisServer *RedisServer
	nodeServer  *node.Node

	stopOnce sync.Once
	stopWG   sync.WaitGroup
	stopC    chan interface{}

	runner *task.Runner
}

// NewServer create a server use spec cfg
func NewServer(cfg *Cfg) *Server { _ = "STUB: not implemented"; return nil }

// Start start the server
func (s *Server) Start() { _ = "STUB: not implemented"; return }

// Stop stop the server
func (s *Server) Stop() { _ = "STUB: not implemented"; return }

func (s *Server) listenToStop() { _ = "STUB: not implemented"; return }

func (s *Server) doStop() { _ = "STUB: not implemented"; return }

func (s *Server) startRedis(store *raftstore.Store) { _ = "STUB: not implemented"; return }

func (s *Server) stopRedis() { _ = "STUB: not implemented"; return }

func (s *Server) startNode() *raftstore.Store { _ = "STUB: not implemented"; return nil }

func (s *Server) stopNode() { _ = "STUB: not implemented"; return }

func (s *Server) initRedis() { _ = "STUB: not implemented"; return }

func (s *Server) initNode() { _ = "STUB: not implemented"; return }

func (s *Server) initDriver() ([]storage.Driver, error) { _ = "STUB: not implemented"; return nil, nil }
