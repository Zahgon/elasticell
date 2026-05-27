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

package redis

import (
	"github.com/deepfabric/elasticell/pkg/pb/raftcmdpb"
	"github.com/fagongzi/goetty"
)

const (
	pong = "PONG"
)

var (
	ErrNotSupportCommand  = []byte("command is not support")
	ErrInvalidCommandResp = []byte("invalid command")
	PongResp              = []byte("PONG")
	OKStatusResp          = []byte("OK")
)

// WriteFVPairArray write field value pair array resp
func WriteFVPairArray(lst []*raftcmdpb.FVPair, buf *goetty.ByteBuf) {
	_ = "STUB: not implemented"
	return
}

// WriteScorePairArray write score member pair array resp
func WriteScorePairArray(lst []*raftcmdpb.ScorePair, withScores bool, buf *goetty.ByteBuf) {
	_ = "STUB: not implemented"
	return
}
