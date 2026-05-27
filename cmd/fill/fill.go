package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	readTimeout    = flag.Int("rt", 10, "The timeout for read in seconds")
	connectTimeout = flag.Int("ct", 10, "The timeout for connect to server")
	addr           = flag.String("addr", "127.0.0.1:6379", "The target address.")
	numKeys        = 10000 //number of keys
)

func main() {
	flag.Parse()
	if err := fill(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}

func fill() (err error) { _ = "STUB: not implemented"; return nil }

//"title", fmt.Sprintf("Exploring AI, volume %d", i),
