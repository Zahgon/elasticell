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

package raftstore

import (
	"sync"
	"time"

	"github.com/deepfabric/elasticell/pkg/pb/mraft"
	"github.com/deepfabric/elasticell/pkg/pd"
	"github.com/fagongzi/goetty"
	"github.com/fagongzi/util/task"
	"github.com/pkg/errors"
)

var (
	errConnect = errors.New("not connected")
)

const (
	defaultConnectTimeout = time.Second * 10
)

type transport struct {
	sync.RWMutex

	store *Store

	seq uint64

	server  *goetty.Server
	handler func(interface{})
	client  *pd.Client

	getStoreAddrFun func(storeID uint64) (string, error)

	conns     map[uint64]goetty.IOSessionPool
	msgs      []*task.Queue
	snapshots []*task.Queue
	mask      uint64
	snapMask  uint64

	waitACKs         map[uint64]*mraft.SnapshotMessage
	sendingSnapshots map[uint64]*mraft.RaftMessage
	pendingLock      sync.RWMutex

	addrs       map[uint64]string
	addrsRevert map[string]uint64
}

func newTransport(store *Store, client *pd.Client, handler func(interface{})) *transport {
	_ = "STUB: not implemented"
	return nil
}

func (t *transport) start() error { _ = "STUB: not implemented"; return nil }

func (t *transport) stop() { _ = "STUB: not implemented"; return }

func (t *transport) doConnection(session goetty.IOSession) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *transport) sendSnapshotMessage(msg *mraft.SnapshotMessage) {
	_ = "STUB: not implemented"
	return
}

func (t *transport) sendACK(msg *mraft.SnapshotMessage) { _ = "STUB: not implemented"; return }

func (t *transport) sendRaftMessage(msg *mraft.RaftMessage) { _ = "STUB: not implemented"; return }

func (t *transport) readyToSendRaft(q *task.Queue) { _ = "STUB: not implemented"; return }

func (t *transport) readyToSendSnapshots(q *task.Queue) { _ = "STUB: not implemented"; return }

// If we write succ, wait ack messages,
// Otherwise, resent if timeout

func (t *transport) waitACK(msg *mraft.SnapshotMessage) { _ = "STUB: not implemented"; return }

func (t *transport) removeACK(id uint64) { _ = "STUB: not implemented"; return }

func (t *transport) isWaitting(id uint64) bool { _ = "STUB: not implemented"; return false }

func (t *transport) retrySnapshot(msg *mraft.SnapshotMessage) { _ = "STUB: not implemented"; return }

func (t *transport) addPendingSnapshot(msg *mraft.RaftMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *transport) getSendingSnapshot(peerID uint64) *mraft.RaftMessage {
	_ = "STUB: not implemented"
	return nil
}

// remove sending snapshots
// it must called only if peer will removed
func (t *transport) forceRemoveSendingSnapshot(peerID uint64) { _ = "STUB: not implemented"; return }

// remove sending snapshots, it will skip to retry if not received ack by peer
func (t *transport) removeSendingSnapshot(peerID uint64, head mraft.SnapshotMessageHeader) {
	_ = "STUB: not implemented"
	return
}

func (t *transport) doRetrySnapshot(arg interface{}) { _ = "STUB: not implemented"; return }

func (t *transport) doWaitTimeoutSnapshot(arg interface{}) { _ = "STUB: not implemented"; return }

func (t *transport) doSend(msg interface{}, to uint64) error { _ = "STUB: not implemented"; return nil }

func (t *transport) doWrite(msg interface{}, conn goetty.IOSession) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *transport) doSendSnapshotMessage(msg *mraft.SnapshotMessage, conn goetty.IOSession) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *transport) postSend(msg *mraft.RaftMessage, err error) { _ = "STUB: not implemented"; return }

func (t *transport) getStoreAddr(storeID uint64) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (t *transport) putConn(id uint64, conn goetty.IOSession) { _ = "STUB: not implemented"; return }

func (t *transport) getConn(storeID uint64) (goetty.IOSession, error) {
	_ = "STUB: not implemented"
	return *new(goetty.IOSession), nil
}

func (t *transport) getConnLocked(id uint64) (goetty.IOSession, error) {
	_ = "STUB: not implemented"
	return *new(goetty.IOSession), nil
}

func (t *transport) checkConnect(id uint64, conn goetty.IOSession) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *transport) createConn(id uint64) (goetty.IOSession, error) {
	_ = "STUB: not implemented"
	return *new(goetty.IOSession), nil
}

func (t *transport) nexSeq() uint64 { _ = "STUB: not implemented"; return 0 }

func isOnlyHeader(msg *mraft.SnapshotMessage) bool { _ = "STUB: not implemented"; return false }
