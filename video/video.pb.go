package video

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//
//message HealthRequest {}
type HealthResponse struct {
	Health               []byte   `protobuf:"bytes,1,opt,name=health,proto3" json:"health,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthResponse) Reset()         { *m = HealthResponse{} }
func (m *HealthResponse) String() string { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()    {}
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2d6993c7f86ea78, []int{0}
}

func (m *HealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthResponse.Unmarshal(m, b)
}
func (m *HealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthResponse.Marshal(b, m, deterministic)
}
func (m *HealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse.Merge(m, src)
}
func (m *HealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthResponse.Size(m)
}
func (m *HealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse proto.InternalMessageInfo

func (m *HealthResponse) GetHealth() []byte {
	if m != nil {
		return m.Health
	}
	return nil
}

type ThumbRequest struct {
	R                    string   `protobuf:"bytes,1,opt,name=r,proto3" json:"r,omitempty"`
	T                    string   `protobuf:"bytes,2,opt,name=t,proto3" json:"t,omitempty"`
	W                    string   `protobuf:"bytes,3,opt,name=w,proto3" json:"w,omitempty"`
	H                    string   `protobuf:"bytes,4,opt,name=h,proto3" json:"h,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThumbRequest) Reset()         { *m = ThumbRequest{} }
func (m *ThumbRequest) String() string { return proto.CompactTextString(m) }
func (*ThumbRequest) ProtoMessage()    {}
func (*ThumbRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2d6993c7f86ea78, []int{1}
}

func (m *ThumbRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThumbRequest.Unmarshal(m, b)
}
func (m *ThumbRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThumbRequest.Marshal(b, m, deterministic)
}
func (m *ThumbRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThumbRequest.Merge(m, src)
}
func (m *ThumbRequest) XXX_Size() int {
	return xxx_messageInfo_ThumbRequest.Size(m)
}
func (m *ThumbRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ThumbRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ThumbRequest proto.InternalMessageInfo

func (m *ThumbRequest) GetR() string {
	if m != nil {
		return m.R
	}
	return ""
}

func (m *ThumbRequest) GetT() string {
	if m != nil {
		return m.T
	}
	return ""
}

func (m *ThumbRequest) GetW() string {
	if m != nil {
		return m.W
	}
	return ""
}

func (m *ThumbRequest) GetH() string {
	if m != nil {
		return m.H
	}
	return ""
}

type ThumbResponse struct {
	Thumb                []byte   `protobuf:"bytes,1,opt,name=thumb,proto3" json:"thumb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThumbResponse) Reset()         { *m = ThumbResponse{} }
func (m *ThumbResponse) String() string { return proto.CompactTextString(m) }
func (*ThumbResponse) ProtoMessage()    {}
func (*ThumbResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2d6993c7f86ea78, []int{2}
}

func (m *ThumbResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThumbResponse.Unmarshal(m, b)
}
func (m *ThumbResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThumbResponse.Marshal(b, m, deterministic)
}
func (m *ThumbResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThumbResponse.Merge(m, src)
}
func (m *ThumbResponse) XXX_Size() int {
	return xxx_messageInfo_ThumbResponse.Size(m)
}
func (m *ThumbResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ThumbResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ThumbResponse proto.InternalMessageInfo

func (m *ThumbResponse) GetThumb() []byte {
	if m != nil {
		return m.Thumb
	}
	return nil
}

type MetaRequest struct {
	R                    string   `protobuf:"bytes,1,opt,name=r,proto3" json:"r,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetaRequest) Reset()         { *m = MetaRequest{} }
func (m *MetaRequest) String() string { return proto.CompactTextString(m) }
func (*MetaRequest) ProtoMessage()    {}
func (*MetaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2d6993c7f86ea78, []int{3}
}

func (m *MetaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaRequest.Unmarshal(m, b)
}
func (m *MetaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaRequest.Marshal(b, m, deterministic)
}
func (m *MetaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaRequest.Merge(m, src)
}
func (m *MetaRequest) XXX_Size() int {
	return xxx_messageInfo_MetaRequest.Size(m)
}
func (m *MetaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MetaRequest proto.InternalMessageInfo

func (m *MetaRequest) GetR() string {
	if m != nil {
		return m.R
	}
	return ""
}

type MetaResponse struct {
	Meta                 []byte   `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetaResponse) Reset()         { *m = MetaResponse{} }
func (m *MetaResponse) String() string { return proto.CompactTextString(m) }
func (*MetaResponse) ProtoMessage()    {}
func (*MetaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2d6993c7f86ea78, []int{4}
}

func (m *MetaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaResponse.Unmarshal(m, b)
}
func (m *MetaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaResponse.Marshal(b, m, deterministic)
}
func (m *MetaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaResponse.Merge(m, src)
}
func (m *MetaResponse) XXX_Size() int {
	return xxx_messageInfo_MetaResponse.Size(m)
}
func (m *MetaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MetaResponse proto.InternalMessageInfo

func (m *MetaResponse) GetMeta() []byte {
	if m != nil {
		return m.Meta
	}
	return nil
}

func init() {
	proto.RegisterType((*HealthResponse)(nil), "video.HealthResponse")
	proto.RegisterType((*ThumbRequest)(nil), "video.ThumbRequest")
	proto.RegisterType((*ThumbResponse)(nil), "video.ThumbResponse")
	proto.RegisterType((*MetaRequest)(nil), "video.MetaRequest")
	proto.RegisterType((*MetaResponse)(nil), "video.MetaResponse")
}

func init() { proto.RegisterFile("video/video.proto", fileDescriptor_c2d6993c7f86ea78) }

var fileDescriptor_c2d6993c7f86ea78 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x99, 0x7e, 0x4d, 0x3e, 0xbd, 0xa6, 0xa2, 0xb7, 0x7f, 0x08, 0xa9, 0x0b, 0x19, 0x10,
	0x4a, 0x17, 0x19, 0xd0, 0x5d, 0x77, 0x2e, 0x14, 0x41, 0x44, 0x88, 0xc5, 0x85, 0xbb, 0xd4, 0x8e,
	0x4d, 0xa0, 0xcd, 0xc4, 0x64, 0xd2, 0xe2, 0xd6, 0x57, 0xf0, 0xd1, 0x7c, 0x05, 0x1f, 0xc0, 0x47,
	0x90, 0x99, 0xb9, 0x42, 0x0b, 0x6e, 0x42, 0x7e, 0xf7, 0xcc, 0x3d, 0x97, 0x73, 0xe0, 0x78, 0x9d,
	0xcf, 0xa5, 0x12, 0xf6, 0x1b, 0x97, 0x95, 0xd2, 0x0a, 0x3d, 0x0b, 0xd1, 0xc9, 0x42, 0xa9, 0xc5,
	0x52, 0x8a, 0xb4, 0xcc, 0x45, 0x5a, 0x14, 0x4a, 0xa7, 0x3a, 0x57, 0x45, 0xed, 0x1e, 0x45, 0x43,
	0x52, 0x2d, 0xcd, 0x9a, 0x17, 0x21, 0x57, 0xa5, 0x7e, 0x73, 0x22, 0x1f, 0xc1, 0xe1, 0x8d, 0x4c,
	0x97, 0x3a, 0x4b, 0x64, 0x5d, 0xaa, 0xa2, 0x96, 0x38, 0x00, 0x3f, 0xb3, 0x93, 0x90, 0x9d, 0xb2,
	0x51, 0x90, 0x10, 0xf1, 0x6b, 0x08, 0xa6, 0x59, 0xb3, 0x9a, 0x25, 0xf2, 0xb5, 0x91, 0xb5, 0xc6,
	0x00, 0x58, 0x65, 0x9f, 0xec, 0x27, 0xac, 0x32, 0xa4, 0xc3, 0x96, 0x23, 0xab, 0x6d, 0xc2, 0x7f,
	0x8e, 0x36, 0x86, 0xb2, 0xb0, 0xed, 0x28, 0xe3, 0x67, 0xd0, 0x21, 0x1f, 0x3a, 0xd8, 0x03, 0x4f,
	0x9b, 0x01, 0xdd, 0x73, 0xc0, 0x87, 0x70, 0x70, 0x27, 0x75, 0xfa, 0xe7, 0x35, 0xce, 0x21, 0x70,
	0x22, 0x59, 0x20, 0xb4, 0x57, 0x52, 0xa7, 0xe4, 0x60, 0xff, 0xcf, 0xbf, 0x19, 0x74, 0x1e, 0xe7,
	0x6a, 0x9a, 0x35, 0x0f, 0xb2, 0x5a, 0xe7, 0xcf, 0x12, 0x6f, 0xc1, 0x77, 0x59, 0x71, 0x10, 0xbb,
	0x4e, 0xe2, 0xdf, 0x4e, 0xe2, 0x2b, 0xd3, 0x49, 0xd4, 0x8f, 0x5d, 0xbb, 0xbb, 0x95, 0x70, 0x7c,
	0xff, 0xfc, 0xfa, 0x68, 0x05, 0xfc, 0xbf, 0x70, 0x5d, 0x4c, 0xd8, 0x18, 0xef, 0xc1, 0xb3, 0x31,
	0xb0, 0x4b, 0x3b, 0xdb, 0xe5, 0x44, 0xbd, 0xdd, 0x21, 0xf9, 0x0c, 0xad, 0x4f, 0x9f, 0xfb, 0xc2,
	0x66, 0x9c, 0xb0, 0xf1, 0xd3, 0x1e, 0x12, 0xe0, 0x25, 0xb4, 0x4d, 0x26, 0x44, 0x5a, 0xdd, 0x4a,
	0x1f, 0x75, 0x77, 0x66, 0xe4, 0x76, 0x64, 0xdd, 0x80, 0x7b, 0xc2, 0xe4, 0x9d, 0xb0, 0xf1, 0xcc,
	0xb7, 0x71, 0x2e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x1f, 0xa0, 0xc3, 0x2a, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VdoThuServiceClient is the client API for VdoThuService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VdoThuServiceClient interface {
	Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthResponse, error)
	Thumb(ctx context.Context, in *ThumbRequest, opts ...grpc.CallOption) (*ThumbResponse, error)
	Meta(ctx context.Context, in *MetaRequest, opts ...grpc.CallOption) (*MetaResponse, error)
}

type vdoThuServiceClient struct {
	cc *grpc.ClientConn
}

func NewVdoThuServiceClient(cc *grpc.ClientConn) VdoThuServiceClient {
	return &vdoThuServiceClient{cc}
}

func (c *vdoThuServiceClient) Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/video.VdoThuService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vdoThuServiceClient) Thumb(ctx context.Context, in *ThumbRequest, opts ...grpc.CallOption) (*ThumbResponse, error) {
	out := new(ThumbResponse)
	err := c.cc.Invoke(ctx, "/video.VdoThuService/Thumb", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vdoThuServiceClient) Meta(ctx context.Context, in *MetaRequest, opts ...grpc.CallOption) (*MetaResponse, error) {
	out := new(MetaResponse)
	err := c.cc.Invoke(ctx, "/video.VdoThuService/Meta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VdoThuServiceServer is the server API for VdoThuService service.
type VdoThuServiceServer interface {
	Health(context.Context, *empty.Empty) (*HealthResponse, error)
	Thumb(context.Context, *ThumbRequest) (*ThumbResponse, error)
	Meta(context.Context, *MetaRequest) (*MetaResponse, error)
}

func RegisterVdoThuServiceServer(s *grpc.Server, srv VdoThuServiceServer) {
	s.RegisterService(&_VdoThuService_serviceDesc, srv)
}

func _VdoThuService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VdoThuServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.VdoThuService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VdoThuServiceServer).Health(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VdoThuService_Thumb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThumbRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VdoThuServiceServer).Thumb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.VdoThuService/Thumb",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VdoThuServiceServer).Thumb(ctx, req.(*ThumbRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VdoThuService_Meta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VdoThuServiceServer).Meta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.VdoThuService/Meta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VdoThuServiceServer).Meta(ctx, req.(*MetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VdoThuService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "video.VdoThuService",
	HandlerType: (*VdoThuServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _VdoThuService_Health_Handler,
		},
		{
			MethodName: "Thumb",
			Handler:    _VdoThuService_Thumb_Handler,
		},
		{
			MethodName: "Meta",
			Handler:    _VdoThuService_Meta_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "video/video.proto",
}
