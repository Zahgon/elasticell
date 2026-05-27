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

//go:build freebsd || openbsd || netbsd || dragonfly || linux
// +build freebsd openbsd netbsd dragonfly linux

package storage

import (
	"runtime"

	gonemo "github.com/deepfabric/go-nemo"
)

// NemoCfg nemo cfg
type NemoCfg struct {
	DataPath              string
	OptionPath            string
	LimitConcurrencyWrite uint64
}

type nemoDrvier struct {
	db         *gonemo.NEMO
	metaEngine Engine
	dataEngine DataEngine
	kvEngine   KVEngine
	hashEngine HashEngine
	listEngine ListEngine
	setEngine  SetEngine
	zsetEngine ZSetEngine
}

// NewNemoDriver return a driver implemention by nemo
func NewNemoDriver(cfg *NemoCfg) (Driver, error) {
	_ = "STUB: not implemented"
	return *new(Driver), nil
}

func (n *nemoDrvier) init(cfg *NemoCfg) {
	if cfg.LimitConcurrencyWrite == 0 {
		cfg.LimitConcurrencyWrite = uint64(runtime.NumCPU())
	}

	n.metaEngine = newNemoMetaEngine(n.db, cfg)
	n.dataEngine = newNemoDataEngine(n.db, cfg)
	n.kvEngine = newNemoKVEngine(n.db, cfg)
	n.hashEngine = newNemoHashEngine(n.db, cfg)
	n.listEngine = newNemoListEngine(n.db, cfg)
	n.setEngine = newNemoSetEngine(n.db, cfg)
	n.zsetEngine = newNemoZSetEngine(n.db, cfg)
}

func (n *nemoDrvier) GetEngine() Engine { _ = "STUB: not implemented"; return *new(Engine) }

func (n *nemoDrvier) GetDataEngine() DataEngine { _ = "STUB: not implemented"; return *new(DataEngine) }

func (n *nemoDrvier) GetKVEngine() KVEngine { _ = "STUB: not implemented"; return *new(KVEngine) }

func (n *nemoDrvier) GetHashEngine() HashEngine { _ = "STUB: not implemented"; return *new(HashEngine) }

func (n *nemoDrvier) GetListEngine() ListEngine { _ = "STUB: not implemented"; return *new(ListEngine) }

func (n *nemoDrvier) GetSetEngine() SetEngine { _ = "STUB: not implemented"; return *new(SetEngine) }

func (n *nemoDrvier) GetZSetEngine() ZSetEngine { _ = "STUB: not implemented"; return *new(ZSetEngine) }

func (n *nemoDrvier) NewWriteBatch() WriteBatch { _ = "STUB: not implemented"; return *new(WriteBatch) }

func (n *nemoDrvier) Write(wb WriteBatch, sync bool) error { _ = "STUB: not implemented"; return nil }
