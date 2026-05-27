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

type nemoSetEngine struct {
	limiter *util.Limiter
	db      *gonemo.NEMO
}

func newNemoSetEngine(db *gonemo.NEMO, cfg *NemoCfg) SetEngine {
	_ = "STUB: not implemented"
	return *new(SetEngine)
}

func (e *nemoSetEngine) SAdd(key []byte, members ...[]byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoSetEngine) SRem(key []byte, members ...[]byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoSetEngine) SCard(key []byte) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (e *nemoSetEngine) SMembers(key []byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *nemoSetEngine) SIsMember(key []byte, member []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *nemoSetEngine) SPop(key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
