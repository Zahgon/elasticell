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

// GetID returns current id
func (s *pdStore) GetID() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// CreateID create id alloc info.
func (s *pdStore) CreateID(leaderSignature string, value uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateID update id for alloc.
func (s *pdStore) UpdateID(leaderSignature string, old, value uint64) error {
	_ = "STUB: not implemented"
	return nil
}
