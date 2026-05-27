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
	"sync"
)

type opt struct {
	key      []byte
	value    []byte
	isDelete bool
}

type memoryWriteBatch struct {
	sync.Mutex

	opts []*opt
}

func newMemoryWriteBatch() WriteBatch { _ = "STUB: not implemented"; return *new(WriteBatch) }

func (wb *memoryWriteBatch) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

func (wb *memoryWriteBatch) Set(key []byte, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

type memoryDriver struct {
	metaEngine Engine
	dataEngine DataEngine
	kvEngine   KVEngine
	hashEngine HashEngine
	listEngine ListEngine
	setEngine  SetEngine
	zsetEngine ZSetEngine
}

// NewMemoryDriver returns Driver with memory implemention
func NewMemoryDriver() Driver { _ = "STUB: not implemented"; return *new(Driver) }

func (d *memoryDriver) GetEngine() Engine { _ = "STUB: not implemented"; return *new(Engine) }

func (d *memoryDriver) GetDataEngine() DataEngine {
	_ = "STUB: not implemented"
	return *new(DataEngine)
}

func (d *memoryDriver) GetKVEngine() KVEngine { _ = "STUB: not implemented"; return *new(KVEngine) }

func (d *memoryDriver) GetHashEngine() HashEngine {
	_ = "STUB: not implemented"
	return *new(HashEngine)
}

func (d *memoryDriver) GetListEngine() ListEngine {
	_ = "STUB: not implemented"
	return *new(ListEngine)
}

func (d *memoryDriver) GetSetEngine() SetEngine { _ = "STUB: not implemented"; return *new(SetEngine) }

func (d *memoryDriver) GetZSetEngine() ZSetEngine {
	_ = "STUB: not implemented"
	return *new(ZSetEngine)
}

func (d *memoryDriver) NewWriteBatch() WriteBatch {
	_ = "STUB: not implemented"
	return *new(WriteBatch)
}

func (d *memoryDriver) Write(wb WriteBatch, sync bool) error { _ = "STUB: not implemented"; return nil }
