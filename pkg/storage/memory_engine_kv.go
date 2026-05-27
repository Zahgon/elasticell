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

type memoryKVEngine struct {
	kv *util.KVTree
}

func newMemoryKVEngine(kv *util.KVTree) KVEngine { _ = "STUB: not implemented"; return *new(KVEngine) }

func (e *memoryKVEngine) RangeDelete(start, end []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *memoryKVEngine) Set(key, value []byte) error { _ = "STUB: not implemented"; return nil }

func (e *memoryKVEngine) MSet(keys [][]byte, values [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *memoryKVEngine) Get(key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *memoryKVEngine) IncrBy(key []byte, incrment int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *memoryKVEngine) DecrBy(key []byte, incrment int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *memoryKVEngine) GetSet(key, value []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *memoryKVEngine) Append(key, value []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *memoryKVEngine) SetNX(key, value []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *memoryKVEngine) StrLen(key []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *memoryKVEngine) NewWriteBatch() WriteBatch {
	_ = "STUB: not implemented"
	return *new(WriteBatch)
}

func (e *memoryKVEngine) Write(wb WriteBatch) error { _ = "STUB: not implemented"; return nil }
