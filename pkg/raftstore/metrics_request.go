package raftstore

import (
	"time"

	"github.com/deepfabric/elasticell/pkg/pb/raftcmdpb"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	labelRequestWaitting = "waitting"
	labelRequestProposal = "proposal"
	labelRequestRaft     = "raft"
	labelRequestStored   = "stored"
	labelRequestResponse = "response"
	labelRequestInQueue  = "in-queue"

	labelQueueReport      = "report"
	labelQueueReq         = "req"
	labelQueueBatchSize   = "batch-size"
	labelQueueBatch       = "batch"
	labelQueueTick        = "tick"
	labelQueueApplyResult = "apply-result"
	labelQueueStep        = "step"
	labelQueueMsgs        = "msgs"
	labelQueueSnaps       = "snaps"
)

var (
	requestDurationHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "elasticell",
			Subsystem: "cell",
			Name:      "request_duration_seconds",
			Help:      "Bucketed histogram of request cycle time duration",
			Buckets:   prometheus.ExponentialBuckets(0.0005, 2.0, 20),
		}, []string{"stage"})

	queueGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "elasticell",
			Subsystem: "cell",
			Name:      "queue_size",
			Help:      "Total size of queue size.",
		}, []string{"type"})
)

func initMetricsForRequest() { _ = "STUB: not implemented"; return }

func observeRequestInQueue(start time.Time) { _ = "STUB: not implemented"; return }

func observeRequestWaitting(c *cmd) { _ = "STUB: not implemented"; return }

func observeRequestProposal(c *cmd) { _ = "STUB: not implemented"; return }

func observeRequestRaft(c *cmd) { _ = "STUB: not implemented"; return }

func observeRequestStored(c *cmd) { _ = "STUB: not implemented"; return }

func observeRequestResponse(c *cmd) { _ = "STUB: not implemented"; return }

func observeRequestWithlabel(c *cmd, label string) { _ = "STUB: not implemented"; return }

func stage(req *raftcmdpb.Request, now int64) { _ = "STUB: not implemented"; return }
