// h 20181204
//
// Video Thumb Service

package server

import (
	"context"
	"fmt"
	"net"

	"../video"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

// Run starts the VdoThu gRPC service
func Run(ctx context.Context, network, address string) (ret error) {
	for {
		// Listen
		l, err := net.Listen(network, address)
		if err != nil {
			break
		}
		// Defer close
		defer func() {
			if err := l.Close(); err != nil {
				glog.Errorf("Failed to close %s %s: %v", network, address, err)
			}
		}()
		// Register server
		s := grpc.NewServer()
		video.RegisterVdoThuServiceServer(s, newVdoThuServiceServer())
		// Graceful stop
		go func() {
			defer s.GracefulStop()
			<-ctx.Done()
		}()
		// Prompt
		fmt.Printf("Listening %v:%v, press ^C to quit\n", network, l.Addr().(*net.TCPAddr).Port)
		// Serve
		ret = s.Serve(l)
		//
		// Finally
		if true {
			break
		}
	}
	return ret
}
