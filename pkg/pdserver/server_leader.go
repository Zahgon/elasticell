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
	"time"

	"github.com/deepfabric/elasticell/pkg/pb/pdpb"
)

var (
	loopInterval = 200 * time.Millisecond
)

func (s *Server) startLeaderLoop() { _ = "STUB: not implemented"; return }

// oh, we are already leader, we may meet something wrong
// in previous campaignLeader. we can resign and campaign again.

func (s *Server) enableLeader() {
	_ = "STUB: not implemented"
	// now, we are leader
	return
}

// if we start cell cluster failure, exit pd.
// Than other pd will become leader and try to start cell cluster

// load watchers

func (s *Server) disableLeader() {
	_ = "STUB: not implemented"
	// now we are not leader
	return
}

func (s *Server) isMatchLeader(leader *pdpb.Leader) bool { _ = "STUB: not implemented"; return false }

// IsLeader returns whether server is leader or not.
func (s *Server) IsLeader() bool { _ = "STUB: not implemented"; return false }

// GetLeader returns current leader pd for API
func (s *Server) GetLeader() (*pdpb.Leader, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Server) marshalLeader() string { _ = "STUB: not implemented"; return "" }

func marshal(leader *pdpb.Leader) string { _ = "STUB: not implemented"; return "" }

// can't fail, so panic here.
