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

package util

import (
	"sync"
)

// OffsetQueue is a queue for sync.
type OffsetQueue struct {
	sync.Mutex

	start, end uint64
	items      []interface{}
}

// NewOffsetQueue returns a offset queue
func NewOffsetQueue() *OffsetQueue { _ = "STUB: not implemented"; return nil }

// Add add a item to the queue
func (q *OffsetQueue) Add(item interface{}) uint64 { _ = "STUB: not implemented"; return 0 }

// Get returns all the items after the offset, and remove all items before this offset
func (q *OffsetQueue) Get(offset uint64) ([]interface{}, uint64) {
	_ = "STUB: not implemented"
	return nil, 0
}

// GetMaxOffset returns the max offset in the queue
func (q *OffsetQueue) GetMaxOffset() uint64 { _ = "STUB: not implemented"; return 0 }

func (q *OffsetQueue) getMaxOffset0() uint64 { _ = "STUB: not implemented"; return 0 }
