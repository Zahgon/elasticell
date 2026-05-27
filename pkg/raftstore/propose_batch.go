package raftstore

import (
	"github.com/deepfabric/elasticell/pkg/pb/raftcmdpb"
	"github.com/fagongzi/goetty"
)

const (
	read = iota
	write
	admin
)

type reqCtx struct {
	admin *raftcmdpb.AdminRequest
	req   *raftcmdpb.Request
	cb    func(*raftcmdpb.RaftCMDResponse)
}

func (r *reqCtx) reset() { _ = "STUB: not implemented"; return }

type proposeBatch struct {
	pr *PeerReplicate

	buf      *goetty.ByteBuf
	lastType int
	cmds     []*cmd
}

func newBatch(pr *PeerReplicate) *proposeBatch { _ = "STUB: not implemented"; return nil }

func (b *proposeBatch) getType(c *reqCtx) int { _ = "STUB: not implemented"; return 0 }

func (b *proposeBatch) size() int { _ = "STUB: not implemented"; return 0 }

func (b *proposeBatch) isEmpty() bool { _ = "STUB: not implemented"; return false }

func (b *proposeBatch) isFull(lastSize uint64) bool { _ = "STUB: not implemented"; return false }

func (b *proposeBatch) pop() *cmd { _ = "STUB: not implemented"; return nil }

func (b *proposeBatch) push(c *reqCtx) { _ = "STUB: not implemented"; return }

// use data key to store

// admin request must in a single batch

func (b *proposeBatch) lastCmd() *cmd { _ = "STUB: not implemented"; return nil }
