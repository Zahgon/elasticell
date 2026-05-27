package pdserver

import (
	"sync"

	"github.com/deepfabric/elasticell/pkg/pb/metapb"
	"github.com/deepfabric/elasticell/pkg/pb/pdpb"
)

// GetCellCluster returns current cell cluster
// if not bootstrap, return nil
func (s *Server) GetCellCluster() *CellCluster { _ = "STUB: not implemented"; return nil }

func (s *Server) isClusterBootstrapped() bool { _ = "STUB: not implemented"; return false }

func (s *Server) bootstrapCluster(req *pdpb.BootstrapClusterReq) (*pdpb.BootstrapClusterRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) listStore(req *pdpb.ListStoreReq) (*pdpb.ListStoreRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) putStore(req *pdpb.PutStoreReq) (*pdpb.PutStoreRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) getStore(req *pdpb.GetStoreReq) (*pdpb.GetStoreRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) cellHeartbeat(req *pdpb.CellHeartbeatReq) (*pdpb.CellHeartbeatRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) storeHeartbeat(req *pdpb.StoreHeartbeatReq) (*pdpb.StoreHeartbeatRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) askSplit(req *pdpb.AskSplitReq) (*pdpb.AskSplitRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) reportSplit(req *pdpb.ReportSplitReq) (*pdpb.ReportSplitRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) getLastRanges(req *pdpb.GetLastRangesReq) (*pdpb.GetLastRangesRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) registerWatcher(req *pdpb.RegisterWatcherReq) (*pdpb.RegisterWatcherRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) watcherHeartbeat(req *pdpb.WatcherHeartbeatReq) (*pdpb.WatcherHeartbeatRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetClusterID returns cluster id
func (s *Server) GetClusterID() uint64 {
	_ = "STUB: not implemented"

	// GetInitParamsValue returns cluster init params bytes
	return 0
}

func (s *Server) GetInitParamsValue() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Server) checkForBootstrap(req *pdpb.BootstrapClusterReq) (metapb.Store, error) {
	_ = "STUB: not implemented"
	return *new(metapb.Store), nil
}

// checkStore returns an error response if the store exists and is in tombstone state.
// It returns nil if it can't get the store.
func (s *Server) checkStore(storeID uint64) error { _ = "STUB: not implemented"; return nil }

// CellCluster is used for cluster config management.
type CellCluster struct {
	mux         sync.RWMutex
	s           *Server
	coordinator *coordinator
	cache       *cache
	running     bool
}

func newCellCluster(s *Server) *CellCluster { _ = "STUB: not implemented"; return nil }

func (c *CellCluster) doBootstrap(store metapb.Store, cells []metapb.Cell) (*pdpb.BootstrapClusterRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CellCluster) doCellHeartbeat(cr *CellInfo) (*pdpb.CellHeartbeatRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CellCluster) doStoreHeartbeat(req *pdpb.StoreHeartbeatReq) (*pdpb.StoreHeartbeatRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CellCluster) doPutStore(store metapb.Store) error { _ = "STUB: not implemented"; return nil }

func (c *CellCluster) doAskSplit(req *pdpb.AskSplitReq) (*pdpb.AskSplitRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the request epoch is less than current cell epoch, then returns an error.

func (c *CellCluster) doReportSplit(req *pdpb.ReportSplitReq) (*pdpb.ReportSplitRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CellCluster) doGetLastRanges(req *pdpb.GetLastRangesReq) (*pdpb.GetLastRangesRsp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CellCluster) checkSplitCell(left *metapb.Cell, right *metapb.Cell) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CellCluster) isRunning() bool { _ = "STUB: not implemented"; return false }

func (c *CellCluster) start() error { _ = "STUB: not implemented"; return nil }

// Here, we will load meta info from store.
// If the cluster is not bootstrapped, the running flag is not set to true

// cluster is not bootstrapped, skipped
