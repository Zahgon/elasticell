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

package util

import (
	"sync"

	"github.com/google/btree"
)

type treeItem struct {
	key   []byte
	value []byte
}

// Less returns true if the item key is less than the other.
func (item *treeItem) Less(other btree.Item) bool { _ = "STUB: not implemented"; return false }

// Equals returns true if the item key is equals the other.
func (item *treeItem) Equals(other btree.Item) bool { _ = "STUB: not implemented"; return false }

// KVTree kv btree
type KVTree struct {
	sync.RWMutex
	tree *btree.BTree
}

// NewKVTree return a kv btree
func NewKVTree() *KVTree { _ = "STUB: not implemented"; return nil }

// Put puts a key, value to the tree
func (kv *KVTree) Put(key, value []byte) { _ = "STUB: not implemented"; return }

// Delete deletes a key, return false if not the key is not exists
func (kv *KVTree) Delete(key []byte) bool { _ = "STUB: not implemented"; return false }

// RangeDelete deletes key in [start, end)
func (kv *KVTree) RangeDelete(start, end []byte) { _ = "STUB: not implemented"; return }

// Get get value, return nil if not the key is not exists
func (kv *KVTree) Get(key []byte) []byte { _ = "STUB: not implemented"; return nil }

// Seek returns the next key and value which key >= spec key
func (kv *KVTree) Seek(key []byte) ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }

// Scan scans in [start, end]
func (kv *KVTree) Scan(start, end []byte, handler func(key, value []byte) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}
