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
)

// ListIndex list indices definion
func (s *pdStore) ListIndex() (idxDefs []*pdpb.IndexDef, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetIndex returns index definion
func (s *pdStore) GetIndex(id string) (idxDef *pdpb.IndexDef, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateIndex creates index
func (s *pdStore) CreateIndex(idxDef *pdpb.IndexDef) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// DeleteIndex deletes index
func (s *pdStore) DeleteIndex(id string) (err error) { _ = "STUB: not implemented"; return nil }

func (s *pdStore) getIndicesKey() string { _ = "STUB: not implemented"; return "" }

func (s *pdStore) getIndexKey(id string) string { _ = "STUB: not implemented"; return "" }
