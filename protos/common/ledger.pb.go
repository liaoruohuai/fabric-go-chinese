
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:34</date>
//</624456114670407680>

//由Protoc Gen Go生成的代码。不要编辑。
//来源：common/ledger.proto

package common //导入“github.com/hyperledger/fabric/protos/common”

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

//引用导入以禁止错误（如果未使用）。
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

//这是一个编译时断言，以确保生成的文件
//与正在编译的proto包兼容。
//此行的编译错误可能意味着您的
//需要更新proto包。
const _ = proto.ProtoPackageIsVersion2 //请升级proto包

//包含有关区块链分类账的信息，如高度、当前
//块哈希和上一个块哈希。
type BlockchainInfo struct {
	Height               uint64   `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	CurrentBlockHash     []byte   `protobuf:"bytes,2,opt,name=currentBlockHash,proto3" json:"currentBlockHash,omitempty"`
	PreviousBlockHash    []byte   `protobuf:"bytes,3,opt,name=previousBlockHash,proto3" json:"previousBlockHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockchainInfo) Reset()         { *m = BlockchainInfo{} }
func (m *BlockchainInfo) String() string { return proto.CompactTextString(m) }
func (*BlockchainInfo) ProtoMessage()    {}
func (*BlockchainInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ledger_109361caaf4d9d5c, []int{0}
}
func (m *BlockchainInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockchainInfo.Unmarshal(m, b)
}
func (m *BlockchainInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockchainInfo.Marshal(b, m, deterministic)
}
func (dst *BlockchainInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockchainInfo.Merge(dst, src)
}
func (m *BlockchainInfo) XXX_Size() int {
	return xxx_messageInfo_BlockchainInfo.Size(m)
}
func (m *BlockchainInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockchainInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BlockchainInfo proto.InternalMessageInfo

func (m *BlockchainInfo) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockchainInfo) GetCurrentBlockHash() []byte {
	if m != nil {
		return m.CurrentBlockHash
	}
	return nil
}

func (m *BlockchainInfo) GetPreviousBlockHash() []byte {
	if m != nil {
		return m.PreviousBlockHash
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockchainInfo)(nil), "common.BlockchainInfo")
}

func init() { proto.RegisterFile("common/ledger.proto", fileDescriptor_ledger_109361caaf4d9d5c) }

var fileDescriptor_ledger_109361caaf4d9d5c = []byte{
//gzip文件描述符或协议的186字节
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0xcf, 0x49, 0x4d, 0x49, 0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x83, 0x08, 0x2a, 0x35, 0x31, 0x72, 0xf1, 0x39, 0xe5, 0xe4, 0x27, 0x67, 0x27, 0x67, 0x24,
	0x66, 0xe6, 0x79, 0xe6, 0xa5, 0xe5, 0x0b, 0x89, 0x71, 0xb1, 0x65, 0xa4, 0x66, 0xa6, 0x67, 0x94,
	0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x04, 0x41, 0x79, 0x42, 0x5a, 0x5c, 0x02, 0xc9, 0xa5, 0x45,
	0x45, 0xa9, 0x79, 0x25, 0x60, 0x0d, 0x1e, 0x89, 0xc5, 0x19, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x3c,
	0x41, 0x18, 0xe2, 0x42, 0x3a, 0x5c, 0x82, 0x05, 0x45, 0xa9, 0x65, 0x99, 0xf9, 0xa5, 0xc5, 0x08,
	0xc5, 0xcc, 0x60, 0xc5, 0x98, 0x12, 0x4e, 0xc1, 0x5c, 0x2a, 0xf9, 0x45, 0xe9, 0x7a, 0x19, 0x95,
	0x05, 0xa9, 0x45, 0x50, 0x57, 0xa6, 0x25, 0x26, 0x15, 0x65, 0x26, 0x43, 0x1c, 0x5b, 0xac, 0x07,
	0x71, 0x6c, 0x94, 0x76, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x12, 0x88, 0xab, 0x8f, 0xa4, 0x58, 0x1f,
	0xa2, 0x58, 0x1f, 0xa2, 0x58, 0x1f, 0xa2, 0x38, 0x89, 0x0d, 0xcc, 0x35, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x9f, 0xcc, 0x05, 0xd1, 0xff, 0x00, 0x00, 0x00,
}
