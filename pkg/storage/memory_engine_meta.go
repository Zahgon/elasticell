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

type memoryMetaEngine struct {
	kv *util.KVTree
}

func newMemoryMetaEngine() Engine { _ = "STUB: not implemented"; return *new(Engine) }

func (e *memoryMetaEngine) Set(key []byte, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *memoryMetaEngine) Get(key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *memoryMetaEngine) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

func (e *memoryMetaEngine) RangeDelete(start, end []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Scan scans the range and execute the handler fun.
// returns false means end the scan.
func (e *memoryMetaEngine) Scan(start, end []byte, handler func(key, value []byte) (bool, error), pooledKey bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Free free unsafe the key or value
func (e *memoryMetaEngine) Free(unsafe []byte) {
	_ = "STUB: not implemented"

	// Seek the first key >= given key, if no found, return None.
	return
}

func (e *memoryMetaEngine) Seek(key []byte) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
