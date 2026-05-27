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

	"github.com/deepfabric/elasticell/pkg/pb/raftcmdpb"
	"github.com/fagongzi/goetty"
	"github.com/fagongzi/util/task"
)

type session struct {
	sync.RWMutex

	id int64

	closed bool
	resps  *task.Queue

	conn goetty.IOSession
	addr string

	fromProxy bool
}

func newSession(conn goetty.IOSession) *session { _ = "STUB: not implemented"; return nil }

func (s *session) close() { _ = "STUB: not implemented"; return }

func (s *session) setFromProxy() { _ = "STUB: not implemented"; return }

func (s *session) onResp(resp *raftcmdpb.Response) { _ = "STUB: not implemented"; return }

func (s *session) writeLoop() { _ = "STUB: not implemented"; return }

// If in the read goroutine, the connection is closed, so we need a lock

func (s *session) doResp(resp *raftcmdpb.Response, buf *goetty.ByteBuf) {
	_ = "STUB: not implemented"
	return
}
