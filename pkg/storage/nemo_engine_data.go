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
	"github.com/deepfabric/elasticell/pkg/util"
	gonemo "github.com/deepfabric/go-nemo"
)

type nemoDataEngine struct {
	db                *gonemo.NEMO
	limiter           *util.Limiter
	limiterTargetScan *util.Limiter
}

func newNemoDataEngine(db *gonemo.NEMO, cfg *NemoCfg) DataEngine {
	_ = "STUB: not implemented"
	return *new(DataEngine)
}

func (e *nemoDataEngine) RangeDelete(start, end []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *nemoDataEngine) GetTargetSizeKey(startKey []byte, endKey []byte, size uint64) (uint64, []byte, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// CreateSnapshot create a snapshot file under the giving path
func (e *nemoDataEngine) CreateSnapshot(path string, start, end []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// ApplySnapshot apply a snapshort file from giving path
func (e *nemoDataEngine) ApplySnapshot(path string) error { _ = "STUB: not implemented"; return nil }

func (e *nemoDataEngine) ScanIndexInfo(startKey []byte, endKey []byte, skipEmpty bool, handler func(key, idxInfo []byte) error) (cnt int, firstErr error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoDataEngine) SetIndexInfo(key, idxInfo []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *nemoDataEngine) GetIndexInfo(key []byte) (idxInfo []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
