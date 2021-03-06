
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:35</date>
//</624456119233810432>

//由Protoc Gen Go生成的代码。不要编辑。
//来源：order/cluster.proto

package orderer //导入“github.com/hyperledger/fabric/protos/order”

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/hyperledger/fabric/protos/common"

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

//stepRequest包含共识实现
//发送到群集成员的特定消息
type StepRequest struct {
	Channel              string   `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StepRequest) Reset()         { *m = StepRequest{} }
func (m *StepRequest) String() string { return proto.CompactTextString(m) }
func (*StepRequest) ProtoMessage()    {}
func (*StepRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_fa8fc28779c58bfb, []int{0}
}
func (m *StepRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StepRequest.Unmarshal(m, b)
}
func (m *StepRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StepRequest.Marshal(b, m, deterministic)
}
func (dst *StepRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StepRequest.Merge(dst, src)
}
func (m *StepRequest) XXX_Size() int {
	return xxx_messageInfo_StepRequest.Size(m)
}
func (m *StepRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StepRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StepRequest proto.InternalMessageInfo

func (m *StepRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *StepRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

//StepResponse包含共识实施
//从接收到的特定消息
//作为对stepRequest的响应的集群成员
type StepResponse struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StepResponse) Reset()         { *m = StepResponse{} }
func (m *StepResponse) String() string { return proto.CompactTextString(m) }
func (*StepResponse) ProtoMessage()    {}
func (*StepResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_fa8fc28779c58bfb, []int{1}
}
func (m *StepResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StepResponse.Unmarshal(m, b)
}
func (m *StepResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StepResponse.Marshal(b, m, deterministic)
}
func (dst *StepResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StepResponse.Merge(dst, src)
}
func (m *StepResponse) XXX_Size() int {
	return xxx_messageInfo_StepResponse.Size(m)
}
func (m *StepResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StepResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StepResponse proto.InternalMessageInfo

func (m *StepResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

//SubmitRequest包装要发送用于订购的事务
type SubmitRequest struct {
	Channel string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
//Last_validation_seq表示最后一个
//配置顺序
//发件人已验证此邮件
	LastValidationSeq uint64 `protobuf:"varint,2,opt,name=last_validation_seq,json=lastValidationSeq,proto3" json:"last_validation_seq,omitempty"`
//内容是结构事务
//转发给群集成员的
	Content              *common.Envelope `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SubmitRequest) Reset()         { *m = SubmitRequest{} }
func (m *SubmitRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitRequest) ProtoMessage()    {}
func (*SubmitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_fa8fc28779c58bfb, []int{2}
}
func (m *SubmitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitRequest.Unmarshal(m, b)
}
func (m *SubmitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitRequest.Marshal(b, m, deterministic)
}
func (dst *SubmitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitRequest.Merge(dst, src)
}
func (m *SubmitRequest) XXX_Size() int {
	return xxx_messageInfo_SubmitRequest.Size(m)
}
func (m *SubmitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitRequest proto.InternalMessageInfo

func (m *SubmitRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *SubmitRequest) GetLastValidationSeq() uint64 {
	if m != nil {
		return m.LastValidationSeq
	}
	return 0
}

func (m *SubmitRequest) GetContent() *common.Envelope {
	if m != nil {
		return m.Content
	}
	return nil
}

//SubmitResponse返回成功
//或发送程序的失败状态
type SubmitResponse struct {
//状态代码，可用于对成功/失败进行编程响应
	Status common.Status `protobuf:"varint,1,opt,name=status,proto3,enum=common.Status" json:"status,omitempty"`
//可能包含有关返回状态的其他信息的信息字符串
	Info                 string   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitResponse) Reset()         { *m = SubmitResponse{} }
func (m *SubmitResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitResponse) ProtoMessage()    {}
func (*SubmitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_fa8fc28779c58bfb, []int{3}
}
func (m *SubmitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitResponse.Unmarshal(m, b)
}
func (m *SubmitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitResponse.Marshal(b, m, deterministic)
}
func (dst *SubmitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitResponse.Merge(dst, src)
}
func (m *SubmitResponse) XXX_Size() int {
	return xxx_messageInfo_SubmitResponse.Size(m)
}
func (m *SubmitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitResponse proto.InternalMessageInfo

func (m *SubmitResponse) GetStatus() common.Status {
	if m != nil {
		return m.Status
	}
	return common.Status_UNKNOWN
}

func (m *SubmitResponse) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func init() {
	proto.RegisterType((*StepRequest)(nil), "orderer.StepRequest")
	proto.RegisterType((*StepResponse)(nil), "orderer.StepResponse")
	proto.RegisterType((*SubmitRequest)(nil), "orderer.SubmitRequest")
	proto.RegisterType((*SubmitResponse)(nil), "orderer.SubmitResponse")
}

//引用导入以禁止错误（如果未使用）。
var _ context.Context
var _ grpc.ClientConn

//这是一个编译时断言，以确保生成的文件
//与正在编译的GRPC包兼容。
const _ = grpc.SupportPackageIsVersion4

//clusterClient是群集服务的客户端API。
//
//有关CTX使用和关闭/结束流式RPC的语义，请参阅https://godoc.org/google.golang.org/grpc clientconn.newstream。
type ClusterClient interface {
//提交向群集成员提交事务
	Submit(ctx context.Context, opts ...grpc.CallOption) (Cluster_SubmitClient, error)
//步骤将特定于实现的消息传递给另一个集群成员。
	Step(ctx context.Context, in *StepRequest, opts ...grpc.CallOption) (*StepResponse, error)
}

type clusterClient struct {
	cc *grpc.ClientConn
}

func NewClusterClient(cc *grpc.ClientConn) ClusterClient {
	return &clusterClient{cc}
}

func (c *clusterClient) Submit(ctx context.Context, opts ...grpc.CallOption) (Cluster_SubmitClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Cluster_serviceDesc.Streams[0], "/orderer.Cluster/Submit", opts...)
	if err != nil {
		return nil, err
	}
	x := &clusterSubmitClient{stream}
	return x, nil
}

type Cluster_SubmitClient interface {
	Send(*SubmitRequest) error
	Recv() (*SubmitResponse, error)
	grpc.ClientStream
}

type clusterSubmitClient struct {
	grpc.ClientStream
}

func (x *clusterSubmitClient) Send(m *SubmitRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clusterSubmitClient) Recv() (*SubmitResponse, error) {
	m := new(SubmitResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clusterClient) Step(ctx context.Context, in *StepRequest, opts ...grpc.CallOption) (*StepResponse, error) {
	out := new(StepResponse)
	err := c.cc.Invoke(ctx, "/orderer.Cluster/Step", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

//cluster server是群集服务的服务器API。
type ClusterServer interface {
//提交向群集成员提交事务
	Submit(Cluster_SubmitServer) error
//步骤将特定于实现的消息传递给另一个集群成员。
	Step(context.Context, *StepRequest) (*StepResponse, error)
}

func RegisterClusterServer(s *grpc.Server, srv ClusterServer) {
	s.RegisterService(&_Cluster_serviceDesc, srv)
}

func _Cluster_Submit_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClusterServer).Submit(&clusterSubmitServer{stream})
}

type Cluster_SubmitServer interface {
	Send(*SubmitResponse) error
	Recv() (*SubmitRequest, error)
	grpc.ServerStream
}

type clusterSubmitServer struct {
	grpc.ServerStream
}

func (x *clusterSubmitServer) Send(m *SubmitResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clusterSubmitServer) Recv() (*SubmitRequest, error) {
	m := new(SubmitRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Cluster_Step_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).Step(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderer.Cluster/Step",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).Step(ctx, req.(*StepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cluster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "orderer.Cluster",
	HandlerType: (*ClusterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Step",
			Handler:    _Cluster_Step_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Submit",
			Handler:       _Cluster_Submit_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "orderer/cluster.proto",
}

func init() { proto.RegisterFile("orderer/cluster.proto", fileDescriptor_cluster_fa8fc28779c58bfb) }

var fileDescriptor_cluster_fa8fc28779c58bfb = []byte{
//gzip文件描述符或协议的347字节
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4f, 0xf2, 0x40,
	0x10, 0xc5, 0xb3, 0xdf, 0x47, 0x68, 0x18, 0x90, 0xe8, 0x22, 0xda, 0x70, 0x22, 0x24, 0x9a, 0xc6,
	0x98, 0xd6, 0xc0, 0xd9, 0x83, 0x1a, 0x6f, 0x9e, 0xda, 0xe8, 0xc1, 0x0b, 0xd9, 0xb6, 0x03, 0x34,
	0x59, 0x76, 0xcb, 0xee, 0x96, 0x84, 0x83, 0x47, 0xff, 0x6f, 0xd3, 0x6e, 0x2b, 0xa0, 0x07, 0x4f,
	0xed, 0xbc, 0xf7, 0x9b, 0xdd, 0x37, 0xb3, 0x30, 0x94, 0x2a, 0x45, 0x85, 0x2a, 0x48, 0x78, 0xa1,
	0x0d, 0x2a, 0x3f, 0x57, 0xd2, 0x48, 0xea, 0xd4, 0xf2, 0x68, 0x90, 0xc8, 0xf5, 0x5a, 0x8a, 0xc0,
	0x7e, 0xac, 0x3b, 0x79, 0x80, 0x6e, 0x64, 0x30, 0x0f, 0x71, 0x53, 0xa0, 0x36, 0xd4, 0x05, 0x27,
	0x59, 0x31, 0x21, 0x90, 0xbb, 0x64, 0x4c, 0xbc, 0x4e, 0xd8, 0x94, 0xa5, 0x93, 0xb3, 0x1d, 0x97,
	0x2c, 0x75, 0xff, 0x8d, 0x89, 0xd7, 0x0b, 0x9b, 0x72, 0xe2, 0x41, 0xcf, 0x1e, 0xa1, 0x73, 0x29,
	0x34, 0x1e, 0x92, 0xe4, 0x98, 0xfc, 0x24, 0x70, 0x12, 0x15, 0xf1, 0x3a, 0x33, 0x7f, 0xdf, 0xe7,
	0xc3, 0x80, 0x33, 0x6d, 0xe6, 0x5b, 0xc6, 0xb3, 0x94, 0x99, 0x4c, 0x8a, 0xb9, 0xc6, 0x4d, 0x75,
	0x77, 0x2b, 0x3c, 0x2b, 0xad, 0xb7, 0x6f, 0x27, 0xc2, 0x0d, 0xbd, 0x01, 0x27, 0x91, 0xc2, 0xa0,
	0x30, 0xee, 0xff, 0x31, 0xf1, 0xba, 0xd3, 0x53, 0xbf, 0x1e, 0xf4, 0x59, 0x6c, 0x91, 0xcb, 0x1c,
	0xc3, 0x06, 0x98, 0xbc, 0x40, 0xbf, 0x89, 0x51, 0x67, 0xbe, 0x86, 0xb6, 0x36, 0xcc, 0x14, 0xba,
	0x8a, 0xd1, 0x9f, 0xf6, 0x9b, 0xe6, 0xa8, 0x52, 0xc3, 0xda, 0xa5, 0x14, 0x5a, 0x99, 0x58, 0xc8,
	0x2a, 0x46, 0x27, 0xac, 0xfe, 0xa7, 0x1f, 0xe0, 0x3c, 0xd9, 0x8d, 0xd3, 0x7b, 0x68, 0xdb, 0x83,
	0xe9, 0x85, 0x5f, 0xaf, 0xdd, 0x3f, 0x1a, 0x78, 0x74, 0xf9, 0x4b, 0xb7, 0x09, 0x3c, 0x72, 0x47,
	0xe8, 0x0c, 0x5a, 0xe5, 0x26, 0xe9, 0xf9, 0x1e, 0xda, 0xbf, 0xcd, 0x68, 0xf8, 0x43, 0xb5, 0x8d,
	0x8f, 0xaf, 0x70, 0x25, 0xd5, 0xd2, 0x5f, 0xed, 0x72, 0x54, 0x1c, 0xd3, 0x25, 0x2a, 0x7f, 0xc1,
	0x62, 0x95, 0x25, 0xf6, 0x85, 0x75, 0xd3, 0xf5, 0x7e, 0xbb, 0xcc, 0xcc, 0xaa, 0x88, 0xcb, 0xc9,
	0x82, 0x03, 0x3a, 0xb0, 0x74, 0x60, 0xe9, 0xa0, 0xa6, 0xe3, 0x76, 0x55, 0xcf, 0xbe, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xc7, 0xdf, 0xa0, 0x73, 0x56, 0x02, 0x00, 0x00,
}

