package proxy

import (
	"flag"
)

var (
	cfgFile = flag.String("cfg", "./cfg.json", "Configuration file of cell kv proxy, base on json formart.")
)

// Cfg returns cfg for proxy
type Cfg struct {
	Addr                string   `json:"addr"`
	AddrNotify          string   `json:"addrNotify"`
	WatcherHeartbeatSec int      `json:"watcherHeartbeatSec"`
	PDAddrs             []string `json:"pdAddrs"`
	MaxRetries          int      `json:"maxRetries"`
	RetryDuration       int64    `json:"retryDuration"`
	WorkerCount         uint64   `json:"workerCount"`
	SupportCMDs         []string `json:"supportCMDs"`
}

// GetCfg get cfg from command
func GetCfg() *Cfg { _ = "STUB: not implemented"; return nil }

func unmarshal(data []byte) (*Cfg, error) { _ = "STUB: not implemented"; return nil, nil }
