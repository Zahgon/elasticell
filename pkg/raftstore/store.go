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

	"github.com/coreos/etcd/raft/raftpb"
	"github.com/deepfabric/elasticell/pkg/pb/metapb"
	"github.com/deepfabric/elasticell/pkg/pb/mraft"
	"github.com/deepfabric/elasticell/pkg/pb/raftcmdpb"
	"github.com/deepfabric/elasticell/pkg/pd"
	"github.com/deepfabric/elasticell/pkg/redis"
	"github.com/deepfabric/elasticell/pkg/storage"
	"github.com/deepfabric/elasticell/pkg/util"
	"github.com/fagongzi/util/task"
)

const (
	applyWorker   = "apply-worker-%d"
	snapWorker    = "snap-worerk"
	pdWorker      = "pd-worker"
	splitWorker   = "split-worker"
	raftGCWorker  = "raft-gc-worker"
	queryChanSize = 10240
)

var (
	globalCfg *Cfg
)

// Store is the store for raft
type Store struct {
	id                 uint64
	clusterID          uint64
	startAt            uint32
	meta               metapb.Store
	snapshotManager    SnapshotManager
	pdClient           *pd.Client
	keyConvertFun      func([]byte, func([]byte) metapb.Cell) metapb.Cell
	replicatesMap      *cellPeersMap // cellid -> peer replicate
	keyRanges          *util.CellTree
	peerCache          *peerCacheMap
	delegates          *applyDelegateMap
	pendingLock        sync.RWMutex
	pendingSnapshots   map[uint64]mraft.SnapshotMessageHeader
	trans              *transport
	engines            []storage.Driver
	enginesMask        uint64
	runner             *task.Runner
	redisReadHandles   map[raftcmdpb.CMDType]func(uint64, *raftcmdpb.Request) *raftcmdpb.Response
	redisWriteHandles  map[raftcmdpb.CMDType]func(*applyContext, *raftcmdpb.Request) *raftcmdpb.Response
	sendingSnapCount   uint32
	reveivingSnapCount uint32

	droppedLock     sync.Mutex
	droppedVoteMsgs map[uint64]raftpb.Message
}

// NewStore returns store
func NewStore(clusterID uint64, pdClient *pd.Client, meta metapb.Store, engines []storage.Driver, cfg *Cfg) *Store {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) startCells() { _ = "STUB: not implemented"; return }

// Start returns the error when start store
func (s *Store) Start() { _ = "STUB: not implemented"; return }

func (s *Store) startTransfer() { _ = "STUB: not implemented"; return }

func (s *Store) startStoreHeartbeatTask() { _ = "STUB: not implemented"; return }

// cancel last if not complete

func (s *Store) startCellHeartbeatTask() { _ = "STUB: not implemented"; return }

func (s *Store) startGCTask() { _ = "STUB: not implemented"; return }

func (s *Store) startCellReportTask() { _ = "STUB: not implemented"; return }

func (s *Store) startCellSplitCheckTask() { _ = "STUB: not implemented"; return }

// Stop returns the error when stop store
func (s *Store) Stop() error { _ = "STUB: not implemented"; return nil }

// GetID returns store id
func (s *Store) GetID() uint64 {
	_ = "STUB: not implemented"

	// GetMeta returns store meta
	return 0
}

func (s *Store) GetMeta() metapb.Store { _ = "STUB: not implemented"; return *new(metapb.Store) }

func (s *Store) getTargetCell(key []byte) (*PeerReplicate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) searchCell(value []byte) metapb.Cell {
	_ = "STUB: not implemented"
	return *new(metapb.Cell)
}

// OnProxyReq process proxy req
func (s *Store) OnProxyReq(req *raftcmdpb.Request, cb func(*raftcmdpb.RaftCMDResponse)) error {
	_ = "STUB: not implemented"
	return nil
}

// OnRedisCommand process redis command
func (s *Store) OnRedisCommand(sessionID int64, cmdType raftcmdpb.CMDType, cmd redis.Command, cb func(*raftcmdpb.RaftCMDResponse)) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) notify(n interface{}) { _ = "STUB: not implemented"; return }

func (s *Store) onRaftMessage(msg *mraft.RaftMessage) { _ = "STUB: not implemented"; return }

// we receive a message tells us to remove ourself.

// In some case, the vote raft msg maybe dropped, so follwer node can't response the vote msg
// Cell a has 3 peers p1, p2, p3. The p1 split to new cell b
// case 1: in most sence, p1 apply split raft log is before p2 and p3.
//
//	At this time, if p2, p3 received the cell b's vote msg,
//	and this vote will dropped by p2 and p3 node,
//	because cell a and cell b has overlapped range at p2 and p3 node
//
// case 2: p2 or p3 apply split log is before p1, we can't mock cell b's vote msg
func (s *Store) cacheDroppedVoteMsg(cellID uint64, msg raftpb.Message) {
	_ = "STUB: not implemented"
	return
}

func (s *Store) removeDroppedVoteMsg(cellID uint64) (raftpb.Message, bool) {
	_ = "STUB: not implemented"
	return *new(raftpb.Message), false
}

func (s *Store) onSplitCheckResult(result *splitCheckResult) { _ = "STUB: not implemented"; return }

func (s *Store) onSnapshotMessage(msg *mraft.SnapshotMessage) { _ = "STUB: not implemented"; return }

func (s *Store) onSnapshotAsk(msg *mraft.SnapshotMessage) { _ = "STUB: not implemented"; return }

// reject a stale snap with current pr

// reject a stale snap with pending

// reject a stale snap

func (s *Store) onSnapshotAck(msg *mraft.SnapshotMessage) { _ = "STUB: not implemented"; return }

// check if this accept is stale

func (s *Store) onSnapshotChunk(msg *mraft.SnapshotMessage) { _ = "STUB: not implemented"; return }

func (s *Store) handleGCPeerMsg(msg *mraft.RaftMessage) { _ = "STUB: not implemented"; return }

func (s *Store) handleStaleMsg(msg *mraft.RaftMessage, currEpoch metapb.CellEpoch, needGC bool) {
	_ = "STUB: not implemented"
	return
}

// If target peer doesn't exist, create it.
//
// return false to indicate that target peer is in invalid state or
// doesn't exist and can't be created.
func (s *Store) tryToCreatePeerReplicate(cellID uint64, msg *mraft.RaftMessage) bool {
	_ = "STUB: not implemented"
	return false
}

// we may encounter a message with larger peer id, which means
// current peer is stale, then we should remove current peer

// cancel snapshotting op

// If we found stale peer, we will destory it

// arrive here means target peer not found, we will try to create it

// check range overlapped

// Maybe split, but not registered yet.

// now we can create a replicate

// following snapshot may overlap, should insert into keyRanges after
// snapshot is applied.

func (s *Store) destroyPeer(cellID uint64, target metapb.Peer, async bool) {
	_ = "STUB: not implemented"
	return
}

func (s *Store) cleanup() {
	_ = "STUB: not implemented"
	// clean up all possible garbage data
	return
}

func (s *Store) clearMeta(cellID uint64, wb storage.WriteBatch) error {
	_ = "STUB: not implemented"
	return nil
}

// meta must in the range [cellID, cellID + 1)

func (s *Store) getPeerReplicate(cellID uint64) *PeerReplicate {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) addPeerToCache(peer metapb.Peer) { _ = "STUB: not implemented"; return }

func (s *Store) getPeer(id uint64) metapb.Peer { _ = "STUB: not implemented"; return *new(metapb.Peer) }

func (s *Store) addPDJob(task func() error, cb func(*task.Job)) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) addRaftLogGCJob(task func() error) error { _ = "STUB: not implemented"; return nil }

func (s *Store) addSnapJob(task func() error, cb func(*task.Job)) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) addApplyJob(cellID uint64, desc string, task func() error, cb func(*task.Job)) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) addSplitJob(task func() error) error { _ = "STUB: not implemented"; return nil }

func (s *Store) addNamedJob(desc, worker string, task func() error) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) addNamedJobWithCB(desc, worker string, task func() error, cb func(*task.Job)) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) handleStoreHeartbeat() error { _ = "STUB: not implemented"; return nil }

func (s *Store) getApplySnapshotCount() (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Store) handleCellHeartbeat() { _ = "STUB: not implemented"; return }

func (s *Store) handleCellSplitCheck() { _ = "STUB: not implemented"; return }

func (s *Store) handleCellReport() { _ = "STUB: not implemented"; return }

func (s *Store) handleRaftGCLog() { _ = "STUB: not implemented"; return }

// SetKeyConvertFun set key convert function
func (s *Store) SetKeyConvertFun(fn func([]byte, func([]byte) metapb.Cell) metapb.Cell) {
	_ = "STUB: not implemented"
	return
}

func (s *Store) getDriver(id uint64) storage.Driver {
	_ = "STUB: not implemented"
	return *new(storage.Driver)
}

func (s *Store) getEngine(id uint64) storage.Engine {
	_ = "STUB: not implemented"
	return *new(storage.Engine)
}

func (s *Store) getDataEngine(id uint64) storage.DataEngine {
	_ = "STUB: not implemented"
	return *new(storage.DataEngine)
}

func (s *Store) getKVEngine(id uint64) storage.KVEngine {
	_ = "STUB: not implemented"
	return *new(storage.KVEngine)
}

func (s *Store) getHashEngine(id uint64) storage.HashEngine {
	_ = "STUB: not implemented"
	return *new(storage.HashEngine)
}

func (s *Store) getListEngine(id uint64) storage.ListEngine {
	_ = "STUB: not implemented"
	return *new(storage.ListEngine)
}

func (s *Store) getSetEngine(id uint64) storage.SetEngine {
	_ = "STUB: not implemented"
	return *new(storage.SetEngine)
}

func (s *Store) getZSetEngine(id uint64) storage.ZSetEngine {
	_ = "STUB: not implemented"
	return *new(storage.ZSetEngine)
}
