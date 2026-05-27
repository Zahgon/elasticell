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

	"github.com/deepfabric/elasticell/pkg/pb/mraft"
	"github.com/deepfabric/elasticell/pkg/storage"
	"github.com/fagongzi/goetty"
	"golang.org/x/time/rate"
)

var (
	creating = 1
	sending  = 2
)

// SnapshotManager manager snapshot
type SnapshotManager interface {
	Register(msg *mraft.SnapshotMessage, step int) bool
	Deregister(msg *mraft.SnapshotMessage, step int)
	Create(msg *mraft.SnapshotMessage) error
	Exists(msg *mraft.SnapshotMessage) bool
	WriteTo(msg *mraft.SnapshotMessage, conn goetty.IOSession) (uint64, error)
	CleanSnap(msg *mraft.SnapshotMessage) error
	ReceiveSnapData(msg *mraft.SnapshotMessage) error
	Apply(msg *mraft.SnapshotMessage) error
}

type defaultSnapshotManager struct {
	sync.RWMutex

	limiter *rate.Limiter

	cfg *Cfg
	db  func(uint64) storage.DataEngine
	s   *Store
	dir string

	registry map[string]struct{}
}

func newDefaultSnapshotManager(cfg *Cfg, db func(uint64) storage.DataEngine, s *Store) SnapshotManager {
	_ = "STUB: not implemented"
	return *new(SnapshotManager)
}

func formatKey(msg *mraft.SnapshotMessage) string { _ = "STUB: not implemented"; return "" }

func formatKeyStep(msg *mraft.SnapshotMessage, step int) string {
	_ = "STUB: not implemented"
	return ""
}

func (m *defaultSnapshotManager) getPathOfSnapKey(msg *mraft.SnapshotMessage) string {
	_ = "STUB: not implemented"
	return ""
}

func (m *defaultSnapshotManager) getPathOfSnapKeyGZ(msg *mraft.SnapshotMessage) string {
	_ = "STUB: not implemented"
	return ""
}

func (m *defaultSnapshotManager) getTmpPathOfSnapKeyGZ(msg *mraft.SnapshotMessage) string {
	_ = "STUB: not implemented"
	return ""
}

func (m *defaultSnapshotManager) Register(msg *mraft.SnapshotMessage, step int) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *defaultSnapshotManager) Deregister(msg *mraft.SnapshotMessage, step int) {
	_ = "STUB: not implemented"
	return
}

func (m *defaultSnapshotManager) inRegistry(msg *mraft.SnapshotMessage, step int) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *defaultSnapshotManager) Create(msg *mraft.SnapshotMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *defaultSnapshotManager) Exists(msg *mraft.SnapshotMessage) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *defaultSnapshotManager) WriteTo(msg *mraft.SnapshotMessage, conn goetty.IOSession) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *defaultSnapshotManager) CleanSnap(msg *mraft.SnapshotMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *defaultSnapshotManager) ReceiveSnapData(msg *mraft.SnapshotMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *defaultSnapshotManager) Apply(msg *mraft.SnapshotMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// apply snapshot of data

func (m *defaultSnapshotManager) cleanTmp(msg *mraft.SnapshotMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *defaultSnapshotManager) check(msg *mraft.SnapshotMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func exist(name string) bool { _ = "STUB: not implemented"; return false }
