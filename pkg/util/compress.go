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
	"archive/tar"
	"os"
)

// GZIP compress a path to a gzip file
func GZIP(path string) error { _ = "STUB: not implemented"; return nil }

// UnGZIP ungip file
func UnGZIP(file string, dest string) error { _ = "STUB: not implemented"; return nil }

func compress(file *os.File, prefix string, tw *tar.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func deCompress(tarFile, dest string) error { _ = "STUB: not implemented"; return nil }

func createFile(name string) (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }
