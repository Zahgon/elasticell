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
)

const (
	batch = uint64(1000)
)

type idAllocator struct {
	sync.Mutex
	store             Store
	leaderSignatureFn func() string
	base              uint64
	end               uint64
}

func newIDAllocator(store Store, leaderSignatureFn func() string) *idAllocator {
	_ = "STUB: not implemented"
	return nil
}

func (alloc *idAllocator) newID() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (alloc *idAllocator) generate() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// create id
