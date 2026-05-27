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

type nemoZSetEngine struct {
	limiter *util.Limiter
	db      *gonemo.NEMO
}

func newNemoZSetEngine(db *gonemo.NEMO, cfg *NemoCfg) ZSetEngine {
	_ = "STUB: not implemented"
	return *new(ZSetEngine)
}

func (e *nemoZSetEngine) ZAdd(key []byte, score float64, member []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoZSetEngine) ZCard(key []byte) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (e *nemoZSetEngine) ZCount(key []byte, min []byte, max []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoZSetEngine) ZIncrBy(key []byte, member []byte, by float64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *nemoZSetEngine) ZLexCount(key []byte, min []byte, max []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoZSetEngine) ZRange(key []byte, start int64, stop int64) ([]*raftcmdpb.ScorePair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *nemoZSetEngine) ZRangeByLex(key []byte, min []byte, max []byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *nemoZSetEngine) ZRangeByScore(key []byte, min []byte, max []byte) ([]*raftcmdpb.ScorePair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *nemoZSetEngine) ZRank(key []byte, member []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoZSetEngine) ZRem(key []byte, members ...[]byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoZSetEngine) ZRemRangeByLex(key []byte, min []byte, max []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoZSetEngine) ZRemRangeByRank(key []byte, start int64, stop int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoZSetEngine) ZRemRangeByScore(key []byte, min []byte, max []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoZSetEngine) ZScore(key []byte, member []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isInclude(value []byte) ([]byte, bool) { _ = "STUB: not implemented"; return nil, false }

var (
	max = []byte("+inf")
	min = []byte("-inf")
)

func parseInclude(value []byte) (float64, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}
