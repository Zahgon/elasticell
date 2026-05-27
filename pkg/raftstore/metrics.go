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

func init() {
	initMetricsForRaft()
	initMetricsForCommand()
	initMetricsForStore()
	initMetricsForSnapshot()
	initMetricsForRequest()
}

type localMetrics struct {
	ready   raftReadyMetrics
	message raftMessageMetrics
	propose raftProposeMetrics
	admin   raftAdminMetrics
}

func (m *localMetrics) flush() { _ = "STUB: not implemented"; return }

type raftReadyMetrics struct {
	message   uint64
	commit    uint64
	append    uint64
	snapshort uint64
}

func (m *raftReadyMetrics) flush() { _ = "STUB: not implemented"; return }

type raftMessageMetrics struct {
	append        uint64
	appendResp    uint64
	vote          uint64
	voteResp      uint64
	snapshot      uint64
	heartbeat     uint64
	heartbeatResp uint64
	transfeLeader uint64
}

func (m *raftMessageMetrics) flush() { _ = "STUB: not implemented"; return }

type raftProposeMetrics struct {
	readLocal      uint64
	readIndex      uint64
	normal         uint64
	transferLeader uint64
	confChange     uint64
}

func (m *raftProposeMetrics) flush() { _ = "STUB: not implemented"; return }

type raftAdminMetrics struct {
	confChange uint64
	addPeer    uint64
	removePeer uint64
	split      uint64
	compact    uint64

	confChangeReject uint64

	confChangeSucceed uint64
	addPeerSucceed    uint64
	removePeerSucceed uint64
	splitSucceed      uint64
	compactSucceed    uint64
}

func (m *raftAdminMetrics) incBy(by raftAdminMetrics) { _ = "STUB: not implemented"; return }

func (m *raftAdminMetrics) flush() { _ = "STUB: not implemented"; return }
