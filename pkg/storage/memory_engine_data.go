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

package storage

import (
	"github.com/deepfabric/elasticell/pkg/util"
)

type memoryDataEngine struct {
	kv *util.KVTree
}

func newMemoryDataEngine(kv *util.KVTree) DataEngine {
	_ = "STUB: not implemented"
	return *new(DataEngine)
}

func (e *memoryDataEngine) RangeDelete(start, end []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *memoryDataEngine) GetTargetSizeKey(startKey []byte, endKey []byte, size uint64) (uint64, []byte, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (e *memoryDataEngine) ScanIndexInfo(start []byte, end []byte, skipEmpty bool, handler func(key, idxInfo []byte) error) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *memoryDataEngine) SetIndexInfo(key, idxInfo []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *memoryDataEngine) GetIndexInfo(key []byte) (idxInfo []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *memoryDataEngine) CreateSnapshot(path string, start, end []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *memoryDataEngine) ApplySnapshot(path string) error { _ = "STUB: not implemented"; return nil }
