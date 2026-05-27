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

package pdserver

import (
	"github.com/deepfabric/elasticell/pkg/pb/pdpb"
	"github.com/deepfabric/elasticell/pkg/pdapi"
)

func (s *Server) getInitParams() (*pdapi.InitParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetSystem returns the summary of elasticell cluster
func (s *Server) GetSystem() (*pdapi.System, error) { _ = "STUB: not implemented"; return nil, nil }

// InitCluster init cluster
func (s *Server) InitCluster(params *pdapi.InitParams) error { _ = "STUB: not implemented"; return nil }

// ListCellInStore returns all cells info in the store
func (s *Server) ListCellInStore(storeID uint64) ([]*pdapi.CellInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListCell returns all cells info
func (s *Server) ListCell() ([]*pdapi.CellInfo, error) { _ = "STUB: not implemented"; return nil, nil }

// GetCell return the cell with the id
func (s *Server) GetCell(id uint64) (*pdapi.CellInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListStore returns all store info
func (s *Server) ListStore() ([]*pdapi.StoreInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetStore return the store with the id
func (s *Server) GetStore(id uint64) (*pdapi.StoreInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteStore remove the store from cluster
// all cells on this store will move to another stores
func (s *Server) DeleteStore(id uint64, force bool) error { _ = "STUB: not implemented"; return nil }

// SetStoreLogLevel set store log level
func (s *Server) SetStoreLogLevel(set *pdapi.SetLogLevel) error {
	_ = "STUB: not implemented"
	return nil
}

// TransferLeader transfer cell leader to the spec peer
func (s *Server) TransferLeader(transfer *pdapi.TransferLeader) error {
	_ = "STUB: not implemented"
	return nil
}

// GetOperator get current operator with id
func (s *Server) GetOperator(id uint64) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetOperators returns the current schedule operators
func (s *Server) GetOperators() ([]interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Server) ListIndex() (idxDefs []*pdpb.IndexDef, err error) {
	_ = "STUB: not implemented"
	return nil,

		// GetIndex returns the info of given index
		nil
}

func (s *Server) GetIndex(id string) (idxDef *pdpb.IndexDef, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) CreateIndex(idxDef *pdpb.IndexDef) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) DeleteIndex(id string) (err error) { _ = "STUB: not implemented"; return nil }

func toAPIStore(store *StoreInfo) *pdapi.StoreInfo { _ = "STUB: not implemented"; return nil }

func toAPICell(cell *CellInfo) *pdapi.CellInfo { _ = "STUB: not implemented"; return nil }

func toAPIStoreSlice(stores []*StoreInfo) []*pdapi.StoreInfo { _ = "STUB: not implemented"; return nil }

func toAPICellSlice(cells []*CellInfo) []*pdapi.CellInfo { _ = "STUB: not implemented"; return nil }
