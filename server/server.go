// h 20181204
//
// Video Thumb Service

package server

import (
	"context"

	"../video"
)

func (s *vdoThuServiceServer) Health(context.Context, *video.HealthRequest) (*video.HealthResponse, error) {
	return &video.HealthResponse{}, nil
}

func (s *vdoThuServiceServer) Thumb(context.Context, *video.ThumbRequest) (*video.ThumbResponse, error) {
	return &video.ThumbResponse{}, nil
}

func (s *vdoThuServiceServer) Meta(context.Context, *video.MetaRequest) (*video.MetaResponse, error) {
	return &video.MetaResponse{}, nil
}

func newVdoThuServiceServer() video.VdoThuServiceServer {
	return &vdoThuServiceServer{}
}

type vdoThuServiceServer struct{}
