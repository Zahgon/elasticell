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

	"github.com/deepfabric/elasticell/pkg/pb/metapb"
	"github.com/google/btree"
)

const (
	defaultBTreeDegree = 64
)

var (
	emptyCell metapb.Cell
	itemPool  sync.Pool
)

func acquireItem() *CellItem { _ = "STUB: not implemented"; return nil }

func releaseItem(item *CellItem) {
	_ = "STUB: not implemented"

	// CellItem is the cell btree item
	return
}

type CellItem struct {
	cell metapb.Cell
}

// CellTree is the btree for cell
type CellTree struct {
	sync.RWMutex
	tree *btree.BTree
}

// NewCellTree returns a default cell btree
func NewCellTree() *CellTree { _ = "STUB: not implemented"; return nil }

// Less returns true if the cell start key is greater than the other.
// So we will sort the cell with start key reversely.
func (r *CellItem) Less(other btree.Item) bool { _ = "STUB: not implemented"; return false }

// Contains returns the item contains the key
func (r *CellItem) Contains(key []byte) bool { _ = "STUB: not implemented"; return false }

// len(end) == 0: max field is positive infinity

func (t *CellTree) length() int { _ = "STUB: not implemented"; return 0 }

// Update updates the tree with the cell.
// It finds and deletes all the overlapped cells first, and then
// insert the cell.
func (t *CellTree) Update(cell metapb.Cell) { _ = "STUB: not implemented"; return }

// between [cell, first], so is iterator all.min >= cell.min' cell
// until all.min > cell.max

// cell.max <= i.start, so cell and i has no overlaps,
// otherwise cell and i has overlaps

// Remove removes a cell if the cell is in the tree.
// It will do nothing if it cannot find the cell or the found cell
// is not the same with the cell.
func (t *CellTree) Remove(cell metapb.Cell) bool { _ = "STUB: not implemented"; return false }

// Ascend asc iterator the tree until fn returns false
func (t *CellTree) Ascend(fn func(cell *metapb.Cell) bool) { _ = "STUB: not implemented"; return }

// NextCell return the next bigger key range cell
func (t *CellTree) NextCell(start []byte) *metapb.Cell { _ = "STUB: not implemented"; return nil }

// AscendRange asc iterator the tree in the range [start, end) until fn returns false
func (t *CellTree) AscendRange(start, end []byte, fn func(cell *metapb.Cell) bool) {
	_ = "STUB: not implemented"
	return
}

// Search returns a cell that contains the key.
func (t *CellTree) Search(key []byte) metapb.Cell {
	_ = "STUB: not implemented"
	return *new(metapb.Cell)
}

func (t *CellTree) find(cell metapb.Cell) *CellItem { _ = "STUB: not implemented"; return nil }
