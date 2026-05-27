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

import (
	"github.com/fagongzi/goetty"
)

var (
	decoder = goetty.NewIntLengthFieldBasedDecoder(newRaftDecoder())
	encoder = newRaftEncoder()
)

const (
	typeRaft = 1
	typeSnap = 2
	typeAck  = 3
)

type raftDecoder struct {
}

type raftEncoder struct {
}

func newRaftDecoder() *raftDecoder { _ = "STUB: not implemented"; return nil }

func newRaftEncoder() *raftEncoder { _ = "STUB: not implemented"; return nil }

func (decoder raftDecoder) Decode(in *goetty.ByteBuf) (bool, interface{}, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func (e raftEncoder) Encode(data interface{}, out *goetty.ByteBuf) error {
	_ = "STUB: not implemented"
	return nil
}
