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

package pd

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/deepfabric/elasticell/pkg/pb/pdpb"
	"github.com/fagongzi/goetty"
)

var (
	// ErrWatcherStopped watcher is stopped
	ErrWatcherStopped = errors.New("watcher is stopped")
)

// Watcher is watch pd event
type Watcher struct {
	sync.RWMutex
	addr      string
	eventFlag uint32
	client    *Client
	heartbeat time.Duration
	listen    *goetty.Server
	readyC    chan *pdpb.WatchEvent
	ctx       context.Context
	cancel    context.CancelFunc
	// about sync protocol
	localOffset, serverOffset uint64
	syncing                   bool
}

// NewWatcher returns a watcher.
// The watcher will listen form pd at the addr parameter
func NewWatcher(client *Client, addr string, eventFlag uint32, heartbeat time.Duration) *Watcher {
	_ = "STUB: not implemented"
	return nil
}

// Ready returns the event
// return a error if watcher was stopped
func (w *Watcher) Ready() (*pdpb.WatchEvent, error) { _ = "STUB: not implemented"; return nil, nil }

// Start start the watch
// If watcher was started, use Ready method in a loop to receive the newest notify
func (w *Watcher) Start() error { _ = "STUB: not implemented"; return nil }

// Stop stop the watcher
func (w *Watcher) Stop() { _ = "STUB: not implemented"; return }

func (w *Watcher) reset() { _ = "STUB: not implemented"; return }

func (w *Watcher) startHeartbeat() { _ = "STUB: not implemented"; return }

func (w *Watcher) doHeartbeat() { _ = "STUB: not implemented"; return }

// If we are paused from pd, we need refresh all ranges

func (w *Watcher) getServerOffset() uint64 { _ = "STUB: not implemented"; return 0 }

func (w *Watcher) resetServerOffset(offset uint64) { _ = "STUB: not implemented"; return }

func (w *Watcher) resetLocalOffset(offset uint64) { _ = "STUB: not implemented"; return }

func (w *Watcher) getLocalOffset() uint64 { _ = "STUB: not implemented"; return 0 }

func (w *Watcher) initNotify() { _ = "STUB: not implemented"; return }

func (w *Watcher) syncNotify(events ...*pdpb.WatchEvent) { _ = "STUB: not implemented"; return }

func (w *Watcher) doNotify(events ...*pdpb.WatchEvent) { _ = "STUB: not implemented"; return }

func (w *Watcher) doConnection(conn goetty.IOSession) error { _ = "STUB: not implemented"; return nil }

func (w *Watcher) sync(conn goetty.IOSession) error { _ = "STUB: not implemented"; return nil }
