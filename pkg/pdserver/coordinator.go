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
	"context"
	"sync"

	"github.com/deepfabric/elasticell/pkg/pb/pdpb"
	"github.com/fagongzi/util/task"
)

var (
	emptyRsp = &pdpb.CellHeartbeatRsp{}
)

type coordinator struct {
	sync.RWMutex
	storeLock sync.RWMutex

	cfg        *Cfg
	cache      *cache
	checker    *replicaChecker
	limiter    *scheduleLimiter
	opts       map[uint64]Operator
	storeOpts  map[uint64]StoreOperator
	schedulers map[string]*scheduleController
	runner     *task.Runner
}

func newCoordinator(cfg *Cfg, cache *cache) *coordinator { _ = "STUB: not implemented"; return nil }

func (c *coordinator) run() { _ = "STUB: not implemented"; return }

func (c *coordinator) stop() { _ = "STUB: not implemented"; return }

func (c *coordinator) dispatchStore(target *StoreInfo) *pdpb.StoreHeartbeatRsp {
	_ = "STUB: not implemented"
	return nil
}

// dispatch is used for coordinator cell,
// it will coordinator when the heartbeat arrives
func (c *coordinator) dispatch(target *CellInfo) *pdpb.CellHeartbeatRsp {
	_ = "STUB: not implemented"
	// Check existed operator.
	return nil
}

// Check replica operator.

func (c *coordinator) getOperators() []interface{} { _ = "STUB: not implemented"; return nil }

func (c *coordinator) getOperator(cellID uint64) Operator {
	_ = "STUB: not implemented"
	return *new(Operator)
}

func (c *coordinator) getStoreOperator(storeID uint64) StoreOperator {
	_ = "STUB: not implemented"
	return *new(StoreOperator)
}

func (c *coordinator) getOperatorCount() int { _ = "STUB: not implemented"; return 0 }

func (c *coordinator) addOperator(op Operator) bool { _ = "STUB: not implemented"; return false }

func (c *coordinator) addStoreOperator(op StoreOperator) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *coordinator) removeOperator(op Operator) { _ = "STUB: not implemented"; return }

func (c *coordinator) removeStoreOperator(op StoreOperator) { _ = "STUB: not implemented"; return }

func (c *coordinator) getScheduler(name string) *scheduleController {
	_ = "STUB: not implemented"
	return nil
}

func (c *coordinator) addScheduler(scheduler Scheduler) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *coordinator) removeScheduler(name string) error { _ = "STUB: not implemented"; return nil }

func (c *coordinator) runScheduler(ctx context.Context, s *scheduleController) {
	_ = "STUB: not implemented"
	return
}
