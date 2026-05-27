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

var (
	testPort        = 10000
	baseAddrPattern = "127.0.0.1:%d"
	httpAddrPattern = "http://127.0.0.1:%d"
	testNamePattern = "test-pd-%d"
)

func getTestPort() int { _ = "STUB: not implemented"; return 0 }

func genBaseAddr() string { _ = "STUB: not implemented"; return "" }

func genHTTPAddr() string { _ = "STUB: not implemented"; return "" }

func getTestName(index int) string { _ = "STUB: not implemented"; return "" }

// NewTestSingleServer returns a single pd server
func NewTestSingleServer() *Server { _ = "STUB: not implemented"; return nil }

// NewTestMultiServers returns multi pd server
func NewTestMultiServers(count int) []*Server { _ = "STUB: not implemented"; return nil }

func newTestConfig(name, addrClient, addrPeer, addrRPC, initCluster string) *Cfg {
	_ = "STUB: not implemented"
	return nil
}
