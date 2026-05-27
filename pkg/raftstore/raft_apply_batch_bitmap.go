package raftstore

import (
	"github.com/pilosa/pilosa/roaring"
)

const (
	opAdd = iota
	opRemove
	opClear
	opDel
)

type bitmapBatch struct {
	bitmaps       [][]byte
	bitmapAdds    []*roaring.Bitmap
	bitmapRemoves []*roaring.Bitmap
	ops           [][]int
}

func (rb *bitmapBatch) add(bm []byte, values ...uint64) { _ = "STUB: not implemented"; return }

func (rb *bitmapBatch) remove(bm []byte, values ...uint64) { _ = "STUB: not implemented"; return }

func (rb *bitmapBatch) clear(bm []byte) { _ = "STUB: not implemented"; return }

func (rb *bitmapBatch) del(bm []byte) { _ = "STUB: not implemented"; return }

func (rb *bitmapBatch) clean(bm []byte, op int) { _ = "STUB: not implemented"; return }

func (rb *bitmapBatch) appendAdds(idx int, values ...uint64) { _ = "STUB: not implemented"; return }

func (rb *bitmapBatch) appendRemoves(idx int, values ...uint64) { _ = "STUB: not implemented"; return }

func (rb *bitmapBatch) hasBatch() bool { _ = "STUB: not implemented"; return false }

func (rb *bitmapBatch) do(id uint64, store *Store) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (rb *bitmapBatch) reset() { _ = "STUB: not implemented"; return }
