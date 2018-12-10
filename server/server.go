// h 20181204
//
// Video Thumb Service

package server

import (
	"bytes"
	"context"
	//"fmt"

	"../video"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *vdoThuServiceServer) Health(context.Context, *empty.Empty) (*video.HealthResponse, error) {
	return &video.HealthResponse{Health: []byte("ok")}, nil
}

func (s *vdoThuServiceServer) Thumb(c context.Context, r *video.ThumbRequest) (*video.ThumbResponse, error) {
	//fmt.Printf("%v\n", r)
	buf := &bytes.Buffer{}
	video.Thumb(buf, r)
	return &video.ThumbResponse{Thumb: buf.Bytes()}, nil
}

func (s *vdoThuServiceServer) Meta(c context.Context, r *video.MetaRequest) (*video.MetaResponse, error) {
	//fmt.Printf("%v\n", r)
	buf := &bytes.Buffer{}
	video.Meta(buf, r)
	return &video.MetaResponse{Meta: buf.Bytes()}, nil
}

func newVdoThuServiceServer() video.VdoThuServiceServer {
	return &vdoThuServiceServer{}
}

type vdoThuServiceServer struct{}
