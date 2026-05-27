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

package pdapi

import (
	"net/http"
)

const (
	headerAccess        = "Access-Control-Allow-Origin"
	headerAccessMethods = "Access-Control-Allow-Methods"
	headerAccessHeaders = "Access-Control-Allow-Headers"
	headerAccessValue   = "*"
)

type cross struct {
}

func newCross() *cross { _ = "STUB: not implemented"; return nil }

func (c *cross) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_ = "STUB: not implemented"
	return
}

func options(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }
