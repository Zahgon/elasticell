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
	"sync"
	"time"

	"github.com/deepfabric/elasticell/pkg/pb"
	"github.com/deepfabric/elasticell/pkg/pb/pdpb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

var (
	defaultConnectTimeout = 5 * time.Second
	defaultTimeout        = 20 * time.Second

	// ErrNotLeader pd is not leader
	ErrNotLeader = errors.New("PD Server Not Leader")
)

// Client pd client
type Client struct {
	sync.RWMutex

	name                   string
	addrs                  []string
	continuousFailureCount int64
	conn                   *grpc.ClientConn
	pd                     pdpb.PDServiceClient
	lastAddr               string
	seq                    uint64
}

// NewClient create a pd client use init pd pdAddrs
func NewClient(name string, initAddrs ...string) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetName set name of client
func (c *Client) SetName(name string) {
	_ = "STUB: not implemented"

	// Close close conn
	return
}

func (c *Client) Close() error { _ = "STUB: not implemented"; return nil }

// GetLastPD returns last pd server
func (c *Client) GetLastPD() string { _ = "STUB: not implemented"; return "" }

func (c *Client) resetConn() error { _ = "STUB: not implemented"; return nil }

func createConn(addr string) (*grpc.ClientConn, error) { _ = "STUB: not implemented"; return nil, nil }

// GetLeader returns current leader
func (c *Client) GetLeader(ctx context.Context, req *pdpb.LeaderReq) (*pdpb.LeaderRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AllocID returns a uniq id
func (c *Client) AllocID(ctx context.Context, req *pdpb.AllocIDReq) (*pdpb.AllocIDRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RegisterWatcher register a watcher for newest cell info notify
func (c *Client) RegisterWatcher(ctx context.Context, req *pdpb.RegisterWatcherReq) (*pdpb.RegisterWatcherRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatcherHeartbeat update the watcher lastest alive time
func (c *Client) WatcherHeartbeat(ctx context.Context, req *pdpb.WatcherHeartbeatReq) (*pdpb.WatcherHeartbeatRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetClusterID returns cluster id
func (c *Client) GetClusterID(ctx context.Context, req *pdpb.GetClusterIDReq) (*pdpb.GetClusterIDRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetInitParams returns cluster init params
func (c *Client) GetInitParams(ctx context.Context, req *pdpb.GetInitParamsReq) (*pdpb.GetInitParamsRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IsClusterBootstrapped returns cluster is bootstrapped response
func (c *Client) IsClusterBootstrapped(ctx context.Context, req *pdpb.IsClusterBootstrapReq) (*pdpb.IsClusterBootstrapRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BootstrapCluster returns bootstrap cluster response
func (c *Client) BootstrapCluster(ctx context.Context, req *pdpb.BootstrapClusterReq) (*pdpb.BootstrapClusterRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListStore returns list store response
func (c *Client) ListStore(ctx context.Context, req *pdpb.ListStoreReq) (*pdpb.ListStoreRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PutStore returns put store response
func (c *Client) PutStore(ctx context.Context, req *pdpb.PutStoreReq) (*pdpb.PutStoreRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetStore returns get store response
func (c *Client) GetStore(ctx context.Context, req *pdpb.GetStoreReq) (*pdpb.GetStoreRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CellHeartbeat returns cell heartbeat response
func (c *Client) CellHeartbeat(ctx context.Context, req *pdpb.CellHeartbeatReq) (*pdpb.CellHeartbeatRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StoreHeartbeat returns store heartbeat response
func (c *Client) StoreHeartbeat(ctx context.Context, req *pdpb.StoreHeartbeatReq) (*pdpb.StoreHeartbeatRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AskSplit returns ask split response
func (c *Client) AskSplit(ctx context.Context, req *pdpb.AskSplitReq) (*pdpb.AskSplitRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReportSplit returns report split response
func (c *Client) ReportSplit(ctx context.Context, req *pdpb.ReportSplitReq) (*pdpb.ReportSplitRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetLastRanges returns lastest key ranges
func (c *Client) GetLastRanges(ctx context.Context, req *pdpb.GetLastRangesReq) (*pdpb.GetLastRangesRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) proxyRPC(ctx context.Context, req pb.BaseReq, setFromFun func(), doRPC func(context.Context) (interface{}, error)) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func needRetry(err error) bool { _ = "STUB: not implemented"; return false }
