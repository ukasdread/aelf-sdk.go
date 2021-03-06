// Code generated by protoc-gen-go. DO NOT EDIT.
// source: token_converter_contract.proto

package client

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

//token converter
type TokenSymbol struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol" form:"symbol"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenSymbol) Reset()         { *m = TokenSymbol{} }
func (m *TokenSymbol) String() string { return proto.CompactTextString(m) }
func (*TokenSymbol) ProtoMessage()    {}
func (*TokenSymbol) Descriptor() ([]byte, []int) {
	return fileDescriptor_adbbaf6824dfdbb6, []int{0}
}

func (m *TokenSymbol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenSymbol.Unmarshal(m, b)
}
func (m *TokenSymbol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenSymbol.Marshal(b, m, deterministic)
}
func (m *TokenSymbol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenSymbol.Merge(m, src)
}
func (m *TokenSymbol) XXX_Size() int {
	return xxx_messageInfo_TokenSymbol.Size(m)
}
func (m *TokenSymbol) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenSymbol.DiscardUnknown(m)
}

var xxx_messageInfo_TokenSymbol proto.InternalMessageInfo

func (m *TokenSymbol) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

type Connector struct {
	Symbol                  string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol" form:"symbol"`
	VirtualBalance          int64    `protobuf:"zigzag64,2,opt,name=virtual_balance,json=virtualBalance,proto3" json:"virtual_balance" form:"virtual_balance"`
	Weight                  string   `protobuf:"bytes,3,opt,name=weight,proto3" json:"weight" form:"weight"`
	IsVirtualBalanceEnabled bool     `protobuf:"varint,4,opt,name=is_virtual_balance_enabled,json=isVirtualBalanceEnabled,proto3" json:"is_virtual_balance_enabled" form:"is_virtual_balance_enabled"`
	IsPurchaseEnabled       bool     `protobuf:"varint,5,opt,name=is_purchase_enabled,json=isPurchaseEnabled,proto3" json:"is_purchase_enabled" form:"is_purchase_enabled"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *Connector) Reset()         { *m = Connector{} }
func (m *Connector) String() string { return proto.CompactTextString(m) }
func (*Connector) ProtoMessage()    {}
func (*Connector) Descriptor() ([]byte, []int) {
	return fileDescriptor_adbbaf6824dfdbb6, []int{1}
}

func (m *Connector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connector.Unmarshal(m, b)
}
func (m *Connector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connector.Marshal(b, m, deterministic)
}
func (m *Connector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connector.Merge(m, src)
}
func (m *Connector) XXX_Size() int {
	return xxx_messageInfo_Connector.Size(m)
}
func (m *Connector) XXX_DiscardUnknown() {
	xxx_messageInfo_Connector.DiscardUnknown(m)
}

var xxx_messageInfo_Connector proto.InternalMessageInfo

func (m *Connector) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Connector) GetVirtualBalance() int64 {
	if m != nil {
		return m.VirtualBalance
	}
	return 0
}

func (m *Connector) GetWeight() string {
	if m != nil {
		return m.Weight
	}
	return ""
}

func (m *Connector) GetIsVirtualBalanceEnabled() bool {
	if m != nil {
		return m.IsVirtualBalanceEnabled
	}
	return false
}

func (m *Connector) GetIsPurchaseEnabled() bool {
	if m != nil {
		return m.IsPurchaseEnabled
	}
	return false
}

func init() {
	proto.RegisterType((*TokenSymbol)(nil), "client.TokenSymbol")
	proto.RegisterType((*Connector)(nil), "client.Connector")
}

func init() { proto.RegisterFile("token_converter_contract.proto", fileDescriptor_adbbaf6824dfdbb6) }

var fileDescriptor_adbbaf6824dfdbb6 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0xc7, 0x89, 0x1f, 0xc5, 0x46, 0x50, 0x8c, 0xa0, 0xc1, 0x83, 0x94, 0x82, 0xd8, 0x53, 0x2f,
	0x1e, 0xbd, 0x29, 0xde, 0xa5, 0x8a, 0xd7, 0x90, 0xc4, 0xc1, 0x06, 0x63, 0x52, 0x92, 0x69, 0x65,
	0x1f, 0x73, 0xdf, 0x68, 0x69, 0x13, 0x58, 0x76, 0x61, 0x6f, 0xf3, 0xff, 0xfa, 0x1d, 0x86, 0xde,
	0xa3, 0xff, 0x05, 0x27, 0xb4, 0x77, 0x13, 0x04, 0x84, 0x30, 0x5f, 0x18, 0xa4, 0xc6, 0x76, 0x08,
	0x1e, 0x3d, 0x2b, 0xb4, 0x35, 0xe0, 0xb0, 0x7e, 0xa0, 0xe7, 0x9f, 0x73, 0xf3, 0x63, 0xf5, 0xa7,
	0xbc, 0x65, 0x37, 0xb4, 0x88, 0xcb, 0xc5, 0x49, 0x45, 0x9a, 0xb2, 0xcb, 0xaa, 0x5e, 0x13, 0x5a,
	0xbe, 0x7a, 0xe7, 0x40, 0xa3, 0x0f, 0x87, 0x5a, 0xec, 0x91, 0x5e, 0x4e, 0x26, 0xe0, 0x28, 0xad,
	0x50, 0xd2, 0x4a, 0xa7, 0x81, 0x1f, 0x55, 0xa4, 0x61, 0xdd, 0x45, 0xb6, 0x5f, 0x92, 0x3b, 0x03,
	0xfe, 0xc1, 0xfc, 0xf4, 0xc8, 0x8f, 0x13, 0x20, 0x29, 0xf6, 0x4c, 0xef, 0x4c, 0x14, 0x7b, 0x0c,
	0x01, 0x4e, 0x2a, 0x0b, 0xdf, 0xfc, 0xa4, 0x22, 0xcd, 0x59, 0x77, 0x6b, 0xe2, 0xd7, 0x0e, 0xed,
	0x2d, 0xc5, 0xac, 0xa5, 0xd7, 0x26, 0x8a, 0x61, 0x0c, 0xba, 0x97, 0x71, 0xbb, 0x3a, 0x5d, 0x56,
	0x57, 0x26, 0xbe, 0xe7, 0x24, 0xf7, 0x55, 0xb1, 0x7c, 0xe2, 0x69, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x79, 0x7f, 0x35, 0x54, 0x2b, 0x01, 0x00, 0x00,
}
