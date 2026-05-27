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

package pool

import (
	"sync"

	"github.com/deepfabric/elasticell/pkg/pb/mraft"
	"github.com/deepfabric/elasticell/pkg/pb/raftcmdpb"
)

var (
	requestsPool           sync.Pool
	responsePool           sync.Pool
	raftMessagePool        sync.Pool
	raftCMDRequestPool     sync.Pool
	raftCMDResponsePool    sync.Pool
	raftRequestHeaderPool  sync.Pool
	raftResponseHeaderPool sync.Pool
)

// AcquireRaftMessage returns a raft message from pool
func AcquireRaftMessage() *mraft.RaftMessage { _ = "STUB: not implemented"; return nil }

// ReleaseRaftMessage returns a raft message to pool
func ReleaseRaftMessage(msg *mraft.RaftMessage) { _ = "STUB: not implemented"; return }

// AcquireRaftCMDRequest returns a raft cmd request from pool
func AcquireRaftCMDRequest() *raftcmdpb.RaftCMDRequest { _ = "STUB: not implemented"; return nil }

// ReleaseRaftCMDRequest returns a raft cmd request to pool
func ReleaseRaftCMDRequest(req *raftcmdpb.RaftCMDRequest) { _ = "STUB: not implemented"; return }

// AcquireRaftRequestHeader returns a raft request header from pool
func AcquireRaftRequestHeader() *raftcmdpb.RaftRequestHeader { _ = "STUB: not implemented"; return nil }

// ReleaseRaftRequestHeader returns a raft request header to pool
func ReleaseRaftRequestHeader(header *raftcmdpb.RaftRequestHeader) {
	_ = "STUB: not implemented"
	return
}

// AcquireRequest returns a raft request from pool
func AcquireRequest() *raftcmdpb.Request { _ = "STUB: not implemented"; return nil }

// ReleaseRequest returns a request to pool
func ReleaseRequest(req *raftcmdpb.Request) { _ = "STUB: not implemented"; return }

// AcquireResponse returns a response from pool
func AcquireResponse() *raftcmdpb.Response { _ = "STUB: not implemented"; return nil }

// ReleaseResponse returns a response to pool
func ReleaseResponse(resp *raftcmdpb.Response) { _ = "STUB: not implemented"; return }

// AcquireRaftCMDResponse returns a raft cmd response from pool
func AcquireRaftCMDResponse() *raftcmdpb.RaftCMDResponse { _ = "STUB: not implemented"; return nil }

// ReleaseRaftCMDResponse returns a raft cmd response to pool
func ReleaseRaftCMDResponse(resp *raftcmdpb.RaftCMDResponse) { _ = "STUB: not implemented"; return }

// AcquireRaftResponseHeader returns a raft response header from pool
func AcquireRaftResponseHeader() *raftcmdpb.RaftResponseHeader {
	_ = "STUB: not implemented"
	return nil
}

// ReleaseRaftResponseHeader returns a raft response header to pool
func ReleaseRaftResponseHeader(header *raftcmdpb.RaftResponseHeader) {
	_ = "STUB: not implemented"
	return
}

// ReleaseRaftRequestAll release requests, header and self to pool
func ReleaseRaftRequestAll(req *raftcmdpb.RaftCMDRequest) { _ = "STUB: not implemented"; return }

// ReleaseRaftResponseAll release responses, header and self to pool
func ReleaseRaftResponseAll(resp *raftcmdpb.RaftCMDResponse) { _ = "STUB: not implemented"; return }
