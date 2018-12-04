// h 20181204
//
// Video Thumb Service

package main

import (
	"context"
	"flag"

	"./server"
	"github.com/golang/glog"
)

func main() {
	// Defer flushing all pending log I/O
	defer glog.Flush()
	// Run gRPC service
	if err := server.Run(context.Background(), *network, *addr); err != nil {
		glog.Fatal(err)
	}
}

func init() {
	flag.Parse()
}

var (
	network = flag.String("n", "tcp", "a valid network type which is consistent to -a")
	addr    = flag.String("a", ":0", "endpoint of the gRPC service")
)
