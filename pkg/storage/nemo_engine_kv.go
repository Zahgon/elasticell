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

type nemoKVEngine struct {
	limiter *util.Limiter
	db      *gonemo.NEMO
}

func newNemoKVEngine(db *gonemo.NEMO, cfg *NemoCfg) KVEngine {
	_ = "STUB: not implemented"
	return *new(KVEngine)
}

func (e *nemoKVEngine) RangeDelete(start, end []byte) error { _ = "STUB: not implemented"; return nil }

func (e *nemoKVEngine) Set(key, value []byte) error { _ = "STUB: not implemented"; return nil }

func (e *nemoKVEngine) MSet(keys [][]byte, values [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *nemoKVEngine) Get(key []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *nemoKVEngine) IncrBy(key []byte, incrment int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoKVEngine) DecrBy(key []byte, incrment int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoKVEngine) GetSet(key, value []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *nemoKVEngine) Append(key, value []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoKVEngine) SetNX(key, value []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoKVEngine) StrLen(key []byte) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (e *nemoKVEngine) NewWriteBatch() WriteBatch {
	_ = "STUB: not implemented"
	return *new(WriteBatch)
}

func (e *nemoKVEngine) Write(wb WriteBatch) error { _ = "STUB: not implemented"; return nil }
