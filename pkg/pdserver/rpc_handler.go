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
	"github.com/deepfabric/elasticell/pkg/pb"
	"github.com/deepfabric/elasticell/pkg/pb/pdpb"
	"github.com/deepfabric/elasticell/pkg/pd"
	"golang.org/x/net/context"
)

// RPCHandler it's a grpc interface implemention
type RPCHandler struct {
	server *Server
}

// NewRPCHandler create a new instance
func NewRPCHandler(server *Server) pdpb.PDServiceServer {
	_ = "STUB: not implemented"
	return *new(pdpb.PDServiceServer)
}

// RegisterWatcher regsiter a watcher for newest cell info notify
func (h *RPCHandler) RegisterWatcher(c context.Context, req *pdpb.RegisterWatcherReq) (*pdpb.RegisterWatcherRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatcherHeartbeat update the watcher lastest alive time
func (h *RPCHandler) WatcherHeartbeat(c context.Context, req *pdpb.WatcherHeartbeatReq) (*pdpb.WatcherHeartbeatRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetClusterID returns cluster id
func (h *RPCHandler) GetClusterID(c context.Context, req *pdpb.GetClusterIDReq) (*pdpb.GetClusterIDRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetInitParams returns cluster init params
func (h *RPCHandler) GetInitParams(c context.Context, req *pdpb.GetInitParamsReq) (*pdpb.GetInitParamsRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AllocID returns alloc id for kv node
func (h *RPCHandler) AllocID(c context.Context, req *pdpb.AllocIDReq) (*pdpb.AllocIDRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetLeader returns current leader
func (h *RPCHandler) GetLeader(c context.Context, req *pdpb.LeaderReq) (*pdpb.LeaderRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IsClusterBootstrap returns cluster is bootstrap already
func (h *RPCHandler) IsClusterBootstrap(c context.Context, req *pdpb.IsClusterBootstrapReq) (*pdpb.IsClusterBootstrapRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BootstrapCluster returns bootstrap cluster response
func (h *RPCHandler) BootstrapCluster(c context.Context, req *pdpb.BootstrapClusterReq) (*pdpb.BootstrapClusterRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListStore puts store
func (h *RPCHandler) ListStore(c context.Context, req *pdpb.ListStoreReq) (*pdpb.ListStoreRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PutStore puts store
func (h *RPCHandler) PutStore(c context.Context, req *pdpb.PutStoreReq) (*pdpb.PutStoreRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetStore get store info
func (h *RPCHandler) GetStore(c context.Context, req *pdpb.GetStoreReq) (*pdpb.GetStoreRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CellHeartbeat returns cell heartbeat response
func (h *RPCHandler) CellHeartbeat(c context.Context, req *pdpb.CellHeartbeatReq) (*pdpb.CellHeartbeatRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StoreHeartbeat returns store heartbeat response
func (h *RPCHandler) StoreHeartbeat(c context.Context, req *pdpb.StoreHeartbeatReq) (*pdpb.StoreHeartbeatRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AskSplit returns ask split response
func (h *RPCHandler) AskSplit(c context.Context, req *pdpb.AskSplitReq) (*pdpb.AskSplitRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReportSplit returns report split response
func (h *RPCHandler) ReportSplit(c context.Context, req *pdpb.ReportSplitReq) (*pdpb.ReportSplitRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetLastRanges returns lastest key ranges
func (h *RPCHandler) GetLastRanges(c context.Context, req *pdpb.GetLastRangesReq) (*pdpb.GetLastRangesRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *RPCHandler) doHandle(name string, req pb.BaseReq, forwardFun func(*pd.Client) (interface{}, error), doFun func() (interface{}, error)) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// forward to leader
