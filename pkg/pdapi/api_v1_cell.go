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

type cellHandler struct {
	service Service
	rd      *render.Render
}

func initAPIForCell(router *mux.Router, service Service, rd *render.Render) {
	_ = "STUB: not implemented"
	return
}

func newCellHandler(service Service, rd *render.Render) *cellHandler {
	_ = "STUB: not implemented"
	return nil
}

func (h *cellHandler) operator(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *cellHandler) leader(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *cellHandler) get(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *cellHandler) list(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
