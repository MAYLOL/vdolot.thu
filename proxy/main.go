// h 20181204
//
// Video Thumb Service

package main

import (
	"context"
	"flag"

	"github.com/golang/glog"
)

func main() {
	// Defer flushing all pending log I/O
	defer glog.Flush()
	// Run proxy service
	if err := run(context.Background(), *endpoint, *port); err != nil {
		glog.Fatal(err)
	}
}

func init() {
	// Parse CLI flags from os.Args[1:]
	flag.Parse()
}

var (
	endpoint = flag.String("e", ":0", "endpoint of the gRPC service to proxy")
	port     = flag.String("p", ":0", "port of the proxy service")
)
