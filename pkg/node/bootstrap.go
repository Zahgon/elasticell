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

package node

import (
	"github.com/deepfabric/elasticell/pkg/pb/metapb"
)

func (n *Node) checkClusterBootstrapped() bool { _ = "STUB: not implemented"; return false }

func (n *Node) checkStore() uint64 { _ = "STUB: not implemented"; return 0 }

func (n *Node) bootstrapStore() uint64 { _ = "STUB: not implemented"; return 0 }

func (n *Node) bootstrapCells() []metapb.Cell { _ = "STUB: not implemented"; return nil }

func (n *Node) createCell(start, end []byte) metapb.Cell {
	_ = "STUB: not implemented"
	return *new(metapb.Cell)
}

func (n *Node) bootstrapCluster(cells []metapb.Cell) { _ = "STUB: not implemented"; return }

// If more than one node try to bootstrap the cluster at the same time,
// Only one can succeed, others will get the `AlreadyBootstrapped` flag.
// If we get any error, we will delete local cells

func (n *Node) startStore() { _ = "STUB: not implemented"; return }

func (n *Node) putStore() { _ = "STUB: not implemented"; return }
