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

//go:build freebsd || openbsd || netbsd || dragonfly || linux
// +build freebsd openbsd netbsd dragonfly linux

package storage

import (
	"github.com/deepfabric/elasticell/pkg/pb/raftcmdpb"
	"github.com/deepfabric/elasticell/pkg/util"
	gonemo "github.com/deepfabric/go-nemo"
)

var (
	endScan = []byte("")
)

type nemoHashEngine struct {
	limiter *util.Limiter
	db      *gonemo.NEMO
}

func newNemoHashEngine(db *gonemo.NEMO, cfg *NemoCfg) HashEngine {
	_ = "STUB: not implemented"
	return *new(HashEngine)
}

func (e *nemoHashEngine) HSet(key, field, value []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoHashEngine) HGet(key, field []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *nemoHashEngine) HDel(key []byte, fields ...[]byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoHashEngine) HExists(key, field []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *nemoHashEngine) HKeys(key []byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *nemoHashEngine) HVals(key []byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *nemoHashEngine) HScanGet(key, start []byte, count int) ([]*raftcmdpb.FVPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *nemoHashEngine) HGetAll(key []byte) ([]*raftcmdpb.FVPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *nemoHashEngine) HLen(key []byte) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (e *nemoHashEngine) HMGet(key []byte, fields ...[]byte) ([][]byte, []error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *nemoHashEngine) HMSet(key []byte, fields, values [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *nemoHashEngine) HSetNX(key, field, value []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoHashEngine) HStrLen(key, field []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoHashEngine) HIncrBy(key, field []byte, incrment int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
