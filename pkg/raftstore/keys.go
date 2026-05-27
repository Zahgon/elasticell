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
	"github.com/deepfabric/elasticell/pkg/pb/metapb"
	"github.com/fagongzi/goetty"
)

// for cell meta
const (
	cellStateSuffix = 0x01
)

const (
	// Following are the suffix after the local prefix.
	// For cell id
	raftLogSuffix    = 0x01
	raftStateSuffix  = 0x02
	applyStateSuffix = 0x03
	nextDocIDSuffix  = 0x04
)

// local is in (0x01, 0x02);
var (
	localPrefix byte = 0x01
	localMinKey      = []byte{localPrefix}
	localMaxKey      = []byte{localPrefix + 1}

	maxKey = []byte{}
	minKey = []byte{0xff}
)

var storeIdentKey = []byte{localPrefix, 0x01}

// data is in (z, z+1)
var (
	dataPrefix    byte = 'z'
	dataPrefixKey      = []byte{dataPrefix}
	dataMinKey         = []byte{dataPrefix}
	dataMaxKey         = []byte{dataPrefix + 1}

	dataPrefixKeySize = len(dataPrefixKey)
)

var (
	// We save two types region data in DB, for raft and other meta data.
	// When the store starts, we should iterate all region meta data to
	// construct peer, no need to travel large raft data, so we separate them
	// with different prefixes.
	cellRaftPrefix    byte = 0x02
	cellRaftPrefixKey      = []byte{localPrefix, cellRaftPrefix}
	cellMetaPrefix    byte = 0x03
	cellMetaPrefixKey      = []byte{localPrefix, cellMetaPrefix}
	cellMetaMinKey         = []byte{localPrefix, cellMetaPrefix}
	cellMetaMaxKey         = []byte{localPrefix, cellMetaPrefix + 1}

	// docID -> userKey
	cellDocIDPrefix    byte = 0x04
	cellDocIDPrefixKey      = []byte{localPrefix, cellDocIDPrefix}

	// index request queue key
	idxReqQueueKey = []byte{localPrefix, 0x05}
)

// GetStoreIdentKey return key of StoreIdent
func GetStoreIdentKey() []byte { _ = "STUB: not implemented"; return nil }

// GetMaxKey return max key
func GetMaxKey() []byte {
	_ = "STUB: not implemented"

	// GetMinKey return min key
	return nil
}

func GetMinKey() []byte { _ = "STUB: not implemented"; return nil }

func decodeCellMetaKey(key []byte) (uint64, byte, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func getCellStateKey(cellID uint64) []byte { _ = "STUB: not implemented"; return nil }

func getCellMetaKey(cellID uint64, suffix byte) []byte { _ = "STUB: not implemented"; return nil }

func getCellMetaPrefix(cellID uint64) []byte { _ = "STUB: not implemented"; return nil }

func getDataKey(key []byte) []byte { _ = "STUB: not implemented"; return nil }

func getDataKey0(key []byte, buf *goetty.ByteBuf) []byte { _ = "STUB: not implemented"; return nil }

func getOriginKey(key []byte) []byte { _ = "STUB: not implemented"; return nil }

func getDataEndKey(endKey []byte) []byte { _ = "STUB: not implemented"; return nil }

// Get the `startKey` of current cell in encoded form.
func encStartKey(cell *metapb.Cell) []byte {
	_ = "STUB: not implemented"
	// only initialized cell's startKey can be encoded, otherwise there must be bugs
	// somewhere.
	// cell
	return nil
}

// / Get the `endKey` of current region in encoded form.
func encEndKey(cell *metapb.Cell) []byte {
	_ = "STUB: not implemented"
	// only initialized region's end_key can be encoded, otherwise there must be bugs
	// somewhere.
	return nil
}

func getRaftStateKey(cellID uint64) []byte { _ = "STUB: not implemented"; return nil }

func getApplyStateKey(cellID uint64) []byte { _ = "STUB: not implemented"; return nil }

func getCellRaftPrefix(cellID uint64) []byte { _ = "STUB: not implemented"; return nil }

func getRaftLogKey(cellID uint64, logIndex uint64) []byte { _ = "STUB: not implemented"; return nil }

func getDocIDKey(docID uint64) []byte { _ = "STUB: not implemented"; return nil }

func getIdxReqQueueKey() []byte { _ = "STUB: not implemented"; return nil }

func getRaftLogIndex(key []byte) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func getCellIDKey(cellID uint64, suffix byte, extraCap int, extra uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}
