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
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/embed"
	"github.com/coreos/etcd/pkg/types"
)

// unixToHTTP replace unix scheme with http.
var unixToHTTP = strings.NewReplacer("unix://", "http://", "unixs://", "http://")

var (
	maxCheckEtcdRunningCount = 60 * 10
	checkEtcdRunningDelay    = 1 * time.Second
)

func (s *Server) startEmbedEtcd() { _ = "STUB: not implemented"; return }

func (s *Server) doAfterEmbedEtcdServerReady(cfg *embed.Config) { _ = "STUB: not implemented"; return }

// See https://github.com/coreos/etcd/issues/6067
// Here may return "not capable" error because we don't start
// all etcds in initial_cluster at same time, so here just log
// an error.
// Note that pd can not work correctly if we don't start all etcds.

func (s *Server) waitEtcdStart(cfg *embed.Config) error { _ = "STUB: not implemented"; return nil }

// etcd may not start ok, we should wait and check again

// endpointStatus checks whether current etcd is running.
func (s *Server) endpointStatus(cfg *embed.Config) (*clientv3.StatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) initStore(cfg *embed.Config) { _ = "STUB: not implemented"; return }

func (s *Server) updateAdvertisePeerUrls() { _ = "STUB: not implemented"; return }

func (s *Server) checkEtcdCluster() { _ = "STUB: not implemented"; return }

func (s *Server) closeEmbedEtcd() { _ = "STUB: not implemented"; return }

func checkClusterID(localClusterID types.ID, um types.URLsMap) error {
	_ = "STUB: not implemented"
	return nil
}

// For tests, change scheme to http.
// etcdserver/api/v3rpc does not recognize unix protocol.

// Do not return error, because other members may be not ready.

func newHTTPTransport(scheme string) *http.Transport { _ = "STUB: not implemented"; return nil }

func unixDial(_, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}
