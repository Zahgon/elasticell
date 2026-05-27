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
)

const (
	bucketSize = int64(128)
	bucketM    = int64(127)
)

type routingMap struct {
	sync.RWMutex
	m map[int64]*session
}

func (m *routingMap) put(key int64, value *session) { _ = "STUB: not implemented"; return }

func (m *routingMap) delete(key int64) *session { _ = "STUB: not implemented"; return nil }

func (m *routingMap) get(key int64) *session { _ = "STUB: not implemented"; return nil }

type routing struct {
	rms []*routingMap
}

func newRouting() *routing { _ = "STUB: not implemented"; return nil }

func (r *routing) put(id int64, value *session) { _ = "STUB: not implemented"; return }

func (r *routing) get(id int64) *session { _ = "STUB: not implemented"; return nil }

func (r *routing) delete(id int64) *session { _ = "STUB: not implemented"; return nil }

func getIndex(id int64) int64 { _ = "STUB: not implemented"; return 0 }
