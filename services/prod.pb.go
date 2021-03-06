// Code generated by protoc-gen-go. DO NOT EDIT.
// source: prod.proto

package HelloTest

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type HelloRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_760dc27a66a60a4a, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type HelloResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_760dc27a66a60a4a, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReqMy struct {
	N1                   int64    `protobuf:"varint,1,opt,name=n1,proto3" json:"n1,omitempty"`
	N2                   int32    `protobuf:"varint,2,opt,name=n2,proto3" json:"n2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqMy) Reset()         { *m = ReqMy{} }
func (m *ReqMy) String() string { return proto.CompactTextString(m) }
func (*ReqMy) ProtoMessage()    {}
func (*ReqMy) Descriptor() ([]byte, []int) {
	return fileDescriptor_760dc27a66a60a4a, []int{2}
}

func (m *ReqMy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqMy.Unmarshal(m, b)
}
func (m *ReqMy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqMy.Marshal(b, m, deterministic)
}
func (m *ReqMy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqMy.Merge(m, src)
}
func (m *ReqMy) XXX_Size() int {
	return xxx_messageInfo_ReqMy.Size(m)
}
func (m *ReqMy) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqMy.DiscardUnknown(m)
}

var xxx_messageInfo_ReqMy proto.InternalMessageInfo

func (m *ReqMy) GetN1() int64 {
	if m != nil {
		return m.N1
	}
	return 0
}

func (m *ReqMy) GetN2() int32 {
	if m != nil {
		return m.N2
	}
	return 0
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "pb.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "pb.HelloResponse")
	proto.RegisterType((*ReqMy)(nil), "pb.ReqMy")
}

func init() { proto.RegisterFile("prod.proto", fileDescriptor_760dc27a66a60a4a) }

var fileDescriptor_760dc27a66a60a4a = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x28, 0xca, 0x4f,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe2, 0xe2, 0xf1, 0x48,
	0xcd, 0xc9, 0xc9, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe2, 0xe2, 0x28, 0x2d,
	0x4e, 0x2d, 0xca, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x95,
	0x34, 0xb9, 0x78, 0xa1, 0x6a, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x24, 0xb8, 0xd8, 0x73,
	0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x61, 0x6a, 0x61, 0x5c, 0x25, 0x75, 0x2e, 0xd6, 0xa0, 0xd4, 0x42,
	0xdf, 0x4a, 0x21, 0x3e, 0x2e, 0xa6, 0x3c, 0x43, 0xb0, 0x2c, 0x73, 0x10, 0x53, 0x9e, 0x21, 0x98,
	0x6f, 0x24, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x1a, 0xc4, 0x94, 0x67, 0x64, 0xb4, 0x88, 0x11, 0xea,
	0x80, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x43, 0x2e, 0x8e, 0xe0, 0xc4, 0x4a, 0xb0,
	0x90, 0x90, 0x80, 0x5e, 0x41, 0x92, 0x1e, 0xb2, 0xf3, 0xa4, 0x04, 0x91, 0x44, 0x20, 0x8e, 0x50,
	0x62, 0x10, 0xd2, 0xe4, 0xe2, 0xf4, 0xad, 0xf4, 0xcc, 0x2b, 0x09, 0x01, 0x79, 0x80, 0x13, 0xa4,
	0x02, 0x6c, 0x37, 0x16, 0xc5, 0x42, 0x26, 0x5c, 0xdc, 0x21, 0x99, 0xb9, 0xa9, 0xfe, 0xa5, 0x10,
	0xc5, 0xc4, 0x59, 0xe0, 0x24, 0x1e, 0x25, 0xaa, 0xa7, 0xa7, 0x5f, 0x0c, 0x71, 0x61, 0xb1, 0x35,
	0x58, 0x1a, 0xa4, 0x3f, 0x89, 0x0d, 0x1c, 0x90, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59,
	0x1d, 0xc2, 0xe5, 0x56, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	MyIntTest(ctx context.Context, in *ReqMy, opts ...grpc.CallOption) (*HelloResponse, error)
	// ????????????
	TimeOutTest(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/pb.HelloService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) MyIntTest(ctx context.Context, in *ReqMy, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/pb.HelloService/MyIntTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) TimeOutTest(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/pb.HelloService/TimeOutTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	MyIntTest(context.Context, *ReqMy) (*HelloResponse, error)
	// ????????????
	TimeOutTest(context.Context, *HelloRequest) (*HelloResponse, error)
}

// UnimplementedHelloServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (*UnimplementedHelloServiceServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedHelloServiceServer) MyIntTest(ctx context.Context, req *ReqMy) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MyIntTest not implemented")
}
func (*UnimplementedHelloServiceServer) TimeOutTest(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TimeOutTest not implemented")
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HelloService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_MyIntTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqMy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).MyIntTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HelloService/MyIntTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).MyIntTest(ctx, req.(*ReqMy))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_TimeOutTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).TimeOutTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HelloService/TimeOutTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).TimeOutTest(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
		{
			MethodName: "MyIntTest",
			Handler:    _HelloService_MyIntTest_Handler,
		},
		{
			MethodName: "TimeOutTest",
			Handler:    _HelloService_TimeOutTest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prod.proto",
}
