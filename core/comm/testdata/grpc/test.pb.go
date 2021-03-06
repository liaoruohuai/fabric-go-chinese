
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:06</date>
//</624455997393473536>

//由Protoc Gen Go生成的代码。不要编辑。
//来源：test.proto

package grpc //导入“github.com/hyperledger/fabric/core/comm/testdata/grpc”

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

//引用导入以禁止错误（如果未使用）。
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

//这是一个编译时断言，以确保生成的文件
//与正在编译的proto包兼容。
//此行的编译错误可能意味着您的
//需要更新proto包。
const _ = proto.ProtoPackageIsVersion2 //请升级proto包

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_cce5a74a641072e6, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Echo struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Echo) Reset()         { *m = Echo{} }
func (m *Echo) String() string { return proto.CompactTextString(m) }
func (*Echo) ProtoMessage()    {}
func (*Echo) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_cce5a74a641072e6, []int{1}
}
func (m *Echo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Echo.Unmarshal(m, b)
}
func (m *Echo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Echo.Marshal(b, m, deterministic)
}
func (dst *Echo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Echo.Merge(dst, src)
}
func (m *Echo) XXX_Size() int {
	return xxx_messageInfo_Echo.Size(m)
}
func (m *Echo) XXX_DiscardUnknown() {
	xxx_messageInfo_Echo.DiscardUnknown(m)
}

var xxx_messageInfo_Echo proto.InternalMessageInfo

func (m *Echo) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*Echo)(nil), "Echo")
}

//引用导入以禁止错误（如果未使用）。
var _ context.Context
var _ grpc.ClientConn

//这是一个编译时断言，以确保生成的文件
//与正在编译的GRPC包兼容。
const _ = grpc.SupportPackageIsVersion4

//TestService Client是TestService服务的客户端API。
//
//有关CTX使用和关闭/结束流式RPC的语义，请参阅https://godoc.org/google.golang.org/grpc clientconn.newstream。
type TestServiceClient interface {
	EmptyCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) EmptyCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/TestService/EmptyCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

//testservice server是testservice服务的服务器API。
type TestServiceServer interface {
	EmptyCall(context.Context, *Empty) (*Empty, error)
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_EmptyCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).EmptyCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TestService/EmptyCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).EmptyCall(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmptyCall",
			Handler:    _TestService_EmptyCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}

//EmptyServiceClient是EmptyService服务的客户端API。
//
//有关CTX使用和关闭/结束流式RPC的语义，请参阅https://godoc.org/google.golang.org/grpc clientconn.newstream。
type EmptyServiceClient interface {
	EmptyCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	EmptyStream(ctx context.Context, opts ...grpc.CallOption) (EmptyService_EmptyStreamClient, error)
}

type emptyServiceClient struct {
	cc *grpc.ClientConn
}

func NewEmptyServiceClient(cc *grpc.ClientConn) EmptyServiceClient {
	return &emptyServiceClient{cc}
}

func (c *emptyServiceClient) EmptyCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/EmptyService/EmptyCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emptyServiceClient) EmptyStream(ctx context.Context, opts ...grpc.CallOption) (EmptyService_EmptyStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EmptyService_serviceDesc.Streams[0], "/EmptyService/EmptyStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &emptyServiceEmptyStreamClient{stream}
	return x, nil
}

type EmptyService_EmptyStreamClient interface {
	Send(*Empty) error
	Recv() (*Empty, error)
	grpc.ClientStream
}

type emptyServiceEmptyStreamClient struct {
	grpc.ClientStream
}

func (x *emptyServiceEmptyStreamClient) Send(m *Empty) error {
	return x.ClientStream.SendMsg(m)
}

func (x *emptyServiceEmptyStreamClient) Recv() (*Empty, error) {
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

//EmptyServiceServer是EmptyService服务的服务器API。
type EmptyServiceServer interface {
	EmptyCall(context.Context, *Empty) (*Empty, error)
	EmptyStream(EmptyService_EmptyStreamServer) error
}

func RegisterEmptyServiceServer(s *grpc.Server, srv EmptyServiceServer) {
	s.RegisterService(&_EmptyService_serviceDesc, srv)
}

func _EmptyService_EmptyCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmptyServiceServer).EmptyCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EmptyService/EmptyCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmptyServiceServer).EmptyCall(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmptyService_EmptyStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EmptyServiceServer).EmptyStream(&emptyServiceEmptyStreamServer{stream})
}

type EmptyService_EmptyStreamServer interface {
	Send(*Empty) error
	Recv() (*Empty, error)
	grpc.ServerStream
}

type emptyServiceEmptyStreamServer struct {
	grpc.ServerStream
}

func (x *emptyServiceEmptyStreamServer) Send(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *emptyServiceEmptyStreamServer) Recv() (*Empty, error) {
	m := new(Empty)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EmptyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "EmptyService",
	HandlerType: (*EmptyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmptyCall",
			Handler:    _EmptyService_EmptyCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EmptyStream",
			Handler:       _EmptyService_EmptyStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "test.proto",
}

//EchoServiceClient是EchoService服务的客户端API。
//
//有关CTX使用和关闭/结束流式RPC的语义，请参阅https://godoc.org/google.golang.org/grpc clientconn.newstream。
type EchoServiceClient interface {
	EchoCall(ctx context.Context, in *Echo, opts ...grpc.CallOption) (*Echo, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) EchoCall(ctx context.Context, in *Echo, opts ...grpc.CallOption) (*Echo, error) {
	out := new(Echo)
	err := c.cc.Invoke(ctx, "/EchoService/EchoCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

//EchoServiceServer是EchoService服务的服务器API。
type EchoServiceServer interface {
	EchoCall(context.Context, *Echo) (*Echo, error)
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_EchoCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Echo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).EchoCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EchoService/EchoCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).EchoCall(ctx, req.(*Echo))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EchoCall",
			Handler:    _EchoService_EchoCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_test_cce5a74a641072e6) }

var fileDescriptor_test_cce5a74a641072e6 = []byte{
//gzip文件描述符或协议的211字节
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x65, 0x89, 0xb6, 0x70, 0xe9, 0xe4, 0x29, 0x2a, 0x4b, 0xe9, 0x42, 0xc4, 0x60, 0xa3,
	0x20, 0xc4, 0x0e, 0xca, 0xca, 0x00, 0x4c, 0x6c, 0x8e, 0x73, 0x24, 0x91, 0x6c, 0xd9, 0xba, 0x1c,
	0x48, 0xf9, 0xf7, 0xc8, 0x0e, 0x59, 0x98, 0x3a, 0xbd, 0xef, 0xc9, 0xcf, 0x4f, 0xa7, 0x07, 0xc0,
	0x38, 0xb1, 0x8a, 0x14, 0x38, 0x9c, 0x76, 0xb0, 0x69, 0x7c, 0xe4, 0xf9, 0x74, 0x84, 0x8b, 0xc6,
	0x0e, 0x41, 0x96, 0xb0, 0x8b, 0x66, 0x76, 0xc1, 0x74, 0xa5, 0x38, 0x8a, 0x6a, 0xff, 0xb6, 0xda,
	0xfa, 0x0e, 0x8a, 0x0f, 0x9c, 0xf8, 0x1d, 0xe9, 0x67, 0xb4, 0x28, 0xaf, 0xe1, 0x2a, 0xff, 0x7c,
	0x31, 0xce, 0xc9, 0xad, 0xca, 0x7c, 0xf8, 0xd3, 0xfa, 0x15, 0xf6, 0x19, 0xce, 0x09, 0xcb, 0x1b,
	0x28, 0x96, 0x30, 0x13, 0x1a, 0xff, 0xff, 0xb9, 0x12, 0xf7, 0xa2, 0xbe, 0x85, 0x22, 0x5d, 0xb7,
	0xd6, 0x95, 0x70, 0x99, 0x6c, 0x6e, 0xdb, 0xa8, 0x84, 0x87, 0x45, 0x9e, 0x9f, 0x3e, 0x1f, 0xfb,
	0x91, 0x87, 0xef, 0x56, 0xd9, 0xe0, 0xf5, 0x30, 0x47, 0x24, 0x87, 0x5d, 0x8f, 0xa4, 0xbf, 0x4c,
	0x4b, 0xa3, 0xd5, 0x36, 0x10, 0x6a, 0x1b, 0xbc, 0xd7, 0x69, 0x85, 0xce, 0xb0, 0xd1, 0x3d, 0x45,
	0xdb, 0x6e, 0xf3, 0x1e, 0x0f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x73, 0x53, 0x6a, 0x38, 0x1d,
	0x01, 0x00, 0x00,
}

