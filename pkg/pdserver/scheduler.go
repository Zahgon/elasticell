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
	"time"

	"github.com/deepfabric/elasticell/pkg/pb/metapb"
)

const (
	maxScheduleRetries  = 10
	maxScheduleInterval = time.Minute
	minScheduleInterval = time.Millisecond * 10
)

// ResourceKind distinguishes different kinds of resources.
type ResourceKind int

const (
	adminKind ResourceKind = iota
	leaderKind
	cellKind
)

// Scheduler is an interface to schedule resources.
type Scheduler interface {
	GetName() string
	GetResourceKind() ResourceKind
	GetResourceLimit() uint64
	Prepare(cache *cache) error
	Cleanup(cache *cache)
	Schedule(cache *cache) Operator
}

type scheduleController struct {
	sync.Mutex

	Scheduler
	cfg      *Cfg
	limiter  *scheduleLimiter
	interval time.Duration
}

func newScheduleController(c *coordinator, s Scheduler) *scheduleController {
	_ = "STUB: not implemented"
	return nil
}

func (s *scheduleController) Schedule(cache *cache) Operator {
	_ = "STUB: not implemented"
	// If we have schedule, reset interval to the minimal interval.
	return *new(Operator)
}

// If we have no schedule, increase the interval exponentially.

func (s *scheduleController) GetInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (s *scheduleController) AllowSchedule() bool { _ = "STUB: not implemented"; return false }

type scheduleLimiter struct {
	sync.RWMutex
	counts map[ResourceKind]uint64
}

func newScheduleLimiter() *scheduleLimiter { _ = "STUB: not implemented"; return nil }

func (l *scheduleLimiter) addOperator(op Operator) { _ = "STUB: not implemented"; return }

func (l *scheduleLimiter) removeOperator(op Operator) { _ = "STUB: not implemented"; return }

func (l *scheduleLimiter) operatorCount(kind ResourceKind) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// scheduleRemovePeer schedules a cell to remove the peer.
func scheduleRemovePeer(cache *cache, s Selector, filters ...Filter) (*CellInfo, *metapb.Peer) {
	_ = "STUB: not implemented"
	return nil, nil
}
