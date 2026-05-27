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

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

type storeHandler struct {
	service Service
	rd      *render.Render
}

func initAPIForStore(router *mux.Router, service Service, rd *render.Render) {
	_ = "STUB: not implemented"
	return
}

func newStoreHandler(service Service, rd *render.Render) *storeHandler {
	_ = "STUB: not implemented"
	return nil
}

func (h *storeHandler) get(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *storeHandler) cells(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *storeHandler) delete(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *storeHandler) list(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *storeHandler) setLogLevel(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
