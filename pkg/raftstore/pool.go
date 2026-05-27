package raftstore

import (
	"sync"

	"github.com/coreos/etcd/raft/raftpb"
	"github.com/deepfabric/elasticell/pkg/pb/mraft"
	"github.com/fagongzi/goetty"
	"github.com/pilosa/pilosa/roaring"
)

var (
	reqCtxPool           sync.Pool
	cmdPool              sync.Pool
	readyContextPool     sync.Pool
	asyncApplyResultPool sync.Pool
	applyContextPool     sync.Pool
	entryPool            sync.Pool
	bufPool              sync.Pool
	bitmapPool           sync.Pool
)

var (
	emptyRaftState    = mraft.RaftLocalState{}
	emptyApplyState   = mraft.RaftApplyState{}
	emptyApplyMetrics = applyMetrics{}
)

func acquireBitmap() *roaring.Bitmap { _ = "STUB: not implemented"; return nil }

func releaseBitmap(value *roaring.Bitmap) { _ = "STUB: not implemented"; return }

func acquireBuf() *goetty.ByteBuf { _ = "STUB: not implemented"; return nil }

func releaseBuf(buf *goetty.ByteBuf) { _ = "STUB: not implemented"; return }

func acquireEntry() *raftpb.Entry { _ = "STUB: not implemented"; return nil }

func releaseEntry(ent *raftpb.Entry) { _ = "STUB: not implemented"; return }

func acquireReqCtx() *reqCtx { _ = "STUB: not implemented"; return nil }

func releaseReqCtx(req *reqCtx) { _ = "STUB: not implemented"; return }

func acquireCmd() *cmd { _ = "STUB: not implemented"; return nil }

func releaseCmd(c *cmd) { _ = "STUB: not implemented"; return }

func acquireReadyContext() *readyContext { _ = "STUB: not implemented"; return nil }

func releaseReadyContext(ctx *readyContext) { _ = "STUB: not implemented"; return }

func acquireAsyncApplyResult() *asyncApplyResult { _ = "STUB: not implemented"; return nil }

func releaseAsyncApplyResult(res *asyncApplyResult) { _ = "STUB: not implemented"; return }

func acquireApplyContext() *applyContext { _ = "STUB: not implemented"; return nil }

func releaseApplyContext(ctx *applyContext) { _ = "STUB: not implemented"; return }
