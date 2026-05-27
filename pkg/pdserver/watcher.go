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
	"errors"
	"sync"
	"time"

	"github.com/deepfabric/elasticell/pkg/pb/pdpb"
	"github.com/deepfabric/elasticell/pkg/util"
	"github.com/fagongzi/goetty"
	"github.com/fagongzi/util/task"
)

var (
	errConnect = errors.New("not connected")
)

const (
	ready = iota
	paused
)

func cmp(a, b interface{}) bool { _ = "STUB: not implemented"; return false }

type notify struct {
	watcher string
	offset  uint64
}

func newWatcherState(watcher pdpb.Watcher) *watcherState { _ = "STUB: not implemented"; return nil }

type watcherState struct {
	watcher          pdpb.Watcher
	state            int
	hbTimeout        *goetty.Timeout
	q                *util.OffsetQueue
	lastNotifyOffset uint64
}

func (state *watcherState) reset() { _ = "STUB: not implemented"; return }

func (state *watcherState) pause() { _ = "STUB: not implemented"; return }

func (state *watcherState) isReady() bool { _ = "STUB: not implemented"; return false }

func (state *watcherState) isPause() bool { _ = "STUB: not implemented"; return false }

func (state *watcherState) addNotify(event *pdpb.WatchEvent) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (state *watcherState) cancelTimeout() { _ = "STUB: not implemented"; return }

func (state *watcherState) resetTimeout(timeout time.Duration, fn func(interface{})) {
	_ = "STUB: not implemented"
	return
}

// watcherNotifier is used for notify the newest cell info to all watchers
type watcherNotifier struct {
	sync.RWMutex

	notifies *task.Queue
	pool     *goetty.AddressBasedPool
	watchers map[string]*watcherState
	timeout  time.Duration
}

func newWatcherNotifier(timeout time.Duration) *watcherNotifier {
	_ = "STUB: not implemented"
	return nil
}

func (wn *watcherNotifier) start() { _ = "STUB: not implemented"; return }

func (wn *watcherNotifier) stop() { _ = "STUB: not implemented"; return }

func (wn *watcherNotifier) removedAllWatcher() { _ = "STUB: not implemented"; return }

// addWatcher add a new watcher for notify the newest cells info.
func (wn *watcherNotifier) addWatcher(watcher pdpb.Watcher) { _ = "STUB: not implemented"; return }

func (wn *watcherNotifier) resetTimeout(addr string) { _ = "STUB: not implemented"; return }

// removeWatcher remove a watcher
func (wn *watcherNotifier) removeWatcher(addr string) { _ = "STUB: not implemented"; return }

// watcherHeartbeat return true if the watcher resume from pause.
func (wn *watcherNotifier) watcherHeartbeat(addr string, offset uint64) bool {
	_ = "STUB: not implemented"
	return false
}

func (wn *watcherNotifier) watcherTimeout(arg interface{}) { _ = "STUB: not implemented"; return }

func (wn *watcherNotifier) allowNotify(addr string) bool { _ = "STUB: not implemented"; return false }

func (wn *watcherNotifier) allowSend(nt *notify) bool { _ = "STUB: not implemented"; return false }

func (wn *watcherNotifier) pause(addr string, remove bool) { _ = "STUB: not implemented"; return }

func (wn *watcherNotifier) resume(addr string) bool { _ = "STUB: not implemented"; return false }

func (wn *watcherNotifier) notify(event *pdpb.WatchEvent) { _ = "STUB: not implemented"; return }

func (wn *watcherNotifier) notifyWithoutLock(state *watcherState, event *pdpb.WatchEvent) {
	_ = "STUB: not implemented"
	return
}

func (wn *watcherNotifier) sync(addr string, offset uint64) *pdpb.WatcherNotifyRsp {
	_ = "STUB: not implemented"
	return nil
}

// ConnectFailed pool status handler
func (wn *watcherNotifier) ConnectFailed(addr string, err error) {
	_ = "STUB: not implemented"

	// Connected pool status handler
	return
}

func (wn *watcherNotifier) Connected(addr string, conn goetty.IOSession) {
	_ = "STUB: not implemented"
	return
}

func createConn(addr string) goetty.IOSession {
	_ = "STUB: not implemented"
	return *new(goetty.IOSession)
}
