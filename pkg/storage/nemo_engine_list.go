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
	"github.com/deepfabric/elasticell/pkg/util"
	gonemo "github.com/deepfabric/go-nemo"
)

type nemoListEngine struct {
	limiter *util.Limiter
	db      *gonemo.NEMO
}

func newNemoListEngine(db *gonemo.NEMO, cfg *NemoCfg) ListEngine {
	_ = "STUB: not implemented"
	return *new(ListEngine)
}

func (e *nemoListEngine) LIndex(key []byte, index int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *nemoListEngine) LInsert(key []byte, pos int, pivot []byte, value []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoListEngine) LLen(key []byte) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (e *nemoListEngine) LPop(key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *nemoListEngine) LPush(key []byte, values ...[]byte) (int64, error) {
	_ = "STUB: not implemented"
	// TODO: nemo must support more value push
	return 0, nil
}

func (e *nemoListEngine) LPushX(key []byte, value []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoListEngine) LRange(key []byte, begin int64, end int64) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *nemoListEngine) LRem(key []byte, count int64, value []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoListEngine) LSet(key []byte, index int64, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *nemoListEngine) LTrim(key []byte, begin int64, end int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *nemoListEngine) RPop(key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *nemoListEngine) RPush(key []byte, values ...[]byte) (int64, error) {
	_ = "STUB: not implemented"
	// TODO: nemo must support more value push
	return 0, nil
}

func (e *nemoListEngine) RPushX(key []byte, value []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
