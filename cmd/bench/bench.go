package main

import (
	"flag"
	"fmt"
	_ "net/http/pprof"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	con            = flag.Int64("c", 0, "The clients.")
	cn             = flag.Int64("cn", 64, "The concurrency per client.")
	readWeight     = flag.Int("r", 1, "read weight")
	writeWeight    = flag.Int("w", 1, "write weight")
	num            = flag.Int64("n", 0, "The total number.")
	size           = flag.Int("v", 256, "The value size.")
	readTimeout    = flag.Int("rt", 30, "The timeout for read in seconds")
	writeTimeout   = flag.Int("wt", 30, "The timeout for read in seconds")
	connectTimeout = flag.Int("ct", 10, "The timeout for connect to server")
	addrs          = flag.String("addrs", "127.0.0.1:6379", "The target address.")
)

var (
	idxWeightsSum []uint64
)

func main() {
	flag.Parse()

	if *readWeight == 0 && *writeWeight == 0 {
		fmt.Printf("read and write cann't be both zero")
		os.Exit(1)
	}

	gCount := *con
	total := *num
	if total < 0 {
		total = 0
	}

	ready := make(chan struct{}, gCount)
	complate := &sync.WaitGroup{}
	wg := &sync.WaitGroup{}

	countPerG := total / gCount

	ans := newAnalysis()

	var index int64
	proxies := strings.Split(*addrs, ",")
	for index = 0; index < gCount; index++ {
		start := index * countPerG
		end := (index + 1) * countPerG
		if index == gCount-1 {
			end = total
		}

		wg.Add(1)
		complate.Add(1)
		proxy := proxies[index%int64(len(proxies))]
		go startG(end-start, wg, complate, ready, ans, proxy)
	}

	wg.Wait()

	ans.start()

	for index = 0; index < gCount; index++ {
		ready <- struct{}{}
	}

	go func() {
		for {
			ans.print()
			time.Sleep(time.Second * 1)
		}
	}()

	complate.Wait()
	ans.print()
}

func startG(total int64, wg, complate *sync.WaitGroup, ready chan struct{}, ans *analysis, proxy string) {
	_ = "STUB: not implemented"
	return
}

type analysis struct {
	sync.RWMutex
	startAt                            time.Time
	recv, sent, prevRecv               int64
	avgLatency, maxLatency, minLatency int64
	totalCost, prevCost                int64
}

func newAnalysis() *analysis { _ = "STUB: not implemented"; return nil }

func (a *analysis) setLatency(latency int64) { _ = "STUB: not implemented"; return }

func (a *analysis) reset() { _ = "STUB: not implemented"; return }

func (a *analysis) start() { _ = "STUB: not implemented"; return }

func (a *analysis) incrRecv(latency int64) { _ = "STUB: not implemented"; return }

func (a *analysis) incrSent(n int64) { _ = "STUB: not implemented"; return }

func (a *analysis) print() { _ = "STUB: not implemented"; return }
