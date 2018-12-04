// h 20181204
//
// Video Thumb Service

package main

import (
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gw "../video"
)

// run starts the VdoThu proxy service
func run(ctx context.Context, end, prx string) (err error) {
	for {
		// Hold Context & CancelFunc
		c, cancel := context.WithCancel(ctx)
		defer cancel()
		// New ServeMux
		m := runtime.NewServeMux()
		// In order to prompt endpoint
		if end == "" || end == ":0" {
			end = ":" + strconv.Itoa(randPort())
		}
		// Register handler
		err = gw.RegisterVdoThuServiceHandlerFromEndpoint(c, m, end, []grpc.DialOption{grpc.WithInsecure()})
		if err != nil {
			break
		}
		// Prompt
		if prx == "" || prx == ":0" {
			prx = ":" + strconv.Itoa(randPort())
		}
		fmt.Printf("Proxing %v & Listening %v, press ^C to quit\n", end, prx)
		// Serve
		err = http.ListenAndServe(prx, m)
		//
		// Finally
		if true {
			break
		}
	}
	return err
}

func randPort() int {
	l, _ := net.Listen("tcp", ":0")
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port
}
