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
	"time"

	"github.com/fagongzi/util/task"
	"github.com/prometheus/client_golang/prometheus"
)

const contentTypeHeader = "Content-Type"

// MetricCfg is the metric configuration.
type MetricCfg struct {
	Job          string
	Instance     string
	Address      string
	DurationSync time.Duration
}

// NewMetricCfg returns metric cfg
func NewMetricCfg(job, instance, address string, durationSync time.Duration) *MetricCfg {
	_ = "STUB: not implemented"
	return nil
}

// InitMetric init the metric
func InitMetric(runner *task.Runner, cfg *MetricCfg) { _ = "STUB: not implemented"; return }

// instanceGroupingKey returns a label map with the only entry
// {instance="<instance>"}. If instance is empty, use hostname instead.
func instanceGroupingKey(instance string) map[string]string { _ = "STUB: not implemented"; return nil }

func doPush(job string, grouping map[string]string, pushURL string, g prometheus.Gatherer, method string) error {
	_ = "STUB: not implemented"
	return nil
}

// Check for pre-existing grouping labels:

// Ignore any further error as this is for an error message only.
