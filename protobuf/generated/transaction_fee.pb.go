// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction_fee.proto

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

type TransactionSizeFeeSymbols struct {
	TransactionSizeFeeSymbolList []*TransactionSizeFeeSymbol `protobuf:"bytes,1,rep,name=transaction_size_fee_symbol_list,json=transactionSizeFeeSymbolList,proto3" json:"transaction_size_fee_symbol_list" form:"transaction_size_fee_symbol_list"`
	XXX_NoUnkeyedLiteral         struct{}                    `json:"-"`
	XXX_unrecognized             []byte                      `json:"-"`
	XXX_sizecache                int32                       `json:"-"`
}

func (m *TransactionSizeFeeSymbols) Reset()         { *m = TransactionSizeFeeSymbols{} }
func (m *TransactionSizeFeeSymbols) String() string { return proto.CompactTextString(m) }
func (*TransactionSizeFeeSymbols) ProtoMessage()    {}
func (*TransactionSizeFeeSymbols) Descriptor() ([]byte, []int) {
	return fileDescriptor_a56c8c1c5f610389, []int{0}
}

func (m *TransactionSizeFeeSymbols) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionSizeFeeSymbols.Unmarshal(m, b)
}
func (m *TransactionSizeFeeSymbols) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionSizeFeeSymbols.Marshal(b, m, deterministic)
}
func (m *TransactionSizeFeeSymbols) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionSizeFeeSymbols.Merge(m, src)
}
func (m *TransactionSizeFeeSymbols) XXX_Size() int {
	return xxx_messageInfo_TransactionSizeFeeSymbols.Size(m)
}
func (m *TransactionSizeFeeSymbols) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionSizeFeeSymbols.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionSizeFeeSymbols proto.InternalMessageInfo

func (m *TransactionSizeFeeSymbols) GetTransactionSizeFeeSymbolList() []*TransactionSizeFeeSymbol {
	if m != nil {
		return m.TransactionSizeFeeSymbolList
	}
	return nil
}

type TransactionSizeFeeSymbol struct {
	TokenSymbol          string   `protobuf:"bytes,1,opt,name=token_symbol,json=tokenSymbol,proto3" json:"token_symbol" form:"token_symbol"`
	BaseTokenWeight      int32    `protobuf:"zigzag32,2,opt,name=base_token_weight,json=baseTokenWeight,proto3" json:"base_token_weight" form:"base_token_weight"`
	AddedTokenWeight     int32    `protobuf:"zigzag32,3,opt,name=added_token_weight,json=addedTokenWeight,proto3" json:"added_token_weight" form:"added_token_weight"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionSizeFeeSymbol) Reset()         { *m = TransactionSizeFeeSymbol{} }
func (m *TransactionSizeFeeSymbol) String() string { return proto.CompactTextString(m) }
func (*TransactionSizeFeeSymbol) ProtoMessage()    {}
func (*TransactionSizeFeeSymbol) Descriptor() ([]byte, []int) {
	return fileDescriptor_a56c8c1c5f610389, []int{1}
}

func (m *TransactionSizeFeeSymbol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionSizeFeeSymbol.Unmarshal(m, b)
}
func (m *TransactionSizeFeeSymbol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionSizeFeeSymbol.Marshal(b, m, deterministic)
}
func (m *TransactionSizeFeeSymbol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionSizeFeeSymbol.Merge(m, src)
}
func (m *TransactionSizeFeeSymbol) XXX_Size() int {
	return xxx_messageInfo_TransactionSizeFeeSymbol.Size(m)
}
func (m *TransactionSizeFeeSymbol) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionSizeFeeSymbol.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionSizeFeeSymbol proto.InternalMessageInfo

func (m *TransactionSizeFeeSymbol) GetTokenSymbol() string {
	if m != nil {
		return m.TokenSymbol
	}
	return ""
}

func (m *TransactionSizeFeeSymbol) GetBaseTokenWeight() int32 {
	if m != nil {
		return m.BaseTokenWeight
	}
	return 0
}

func (m *TransactionSizeFeeSymbol) GetAddedTokenWeight() int32 {
	if m != nil {
		return m.AddedTokenWeight
	}
	return 0
}

type TransactionFeeCharged struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol" form:"symbol"`
	Amount               int64    `protobuf:"zigzag64,2,opt,name=amount,proto3" json:"amount" form:"amount"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionFeeCharged) Reset()         { *m = TransactionFeeCharged{} }
func (m *TransactionFeeCharged) String() string { return proto.CompactTextString(m) }
func (*TransactionFeeCharged) ProtoMessage()    {}
func (*TransactionFeeCharged) Descriptor() ([]byte, []int) {
	return fileDescriptor_a56c8c1c5f610389, []int{2}
}

func (m *TransactionFeeCharged) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionFeeCharged.Unmarshal(m, b)
}
func (m *TransactionFeeCharged) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionFeeCharged.Marshal(b, m, deterministic)
}
func (m *TransactionFeeCharged) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionFeeCharged.Merge(m, src)
}
func (m *TransactionFeeCharged) XXX_Size() int {
	return xxx_messageInfo_TransactionFeeCharged.Size(m)
}
func (m *TransactionFeeCharged) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionFeeCharged.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionFeeCharged proto.InternalMessageInfo

func (m *TransactionFeeCharged) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *TransactionFeeCharged) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type ResourceTokenCharged struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol" form:"symbol"`
	Amount               int64    `protobuf:"zigzag64,2,opt,name=amount,proto3" json:"amount" form:"amount"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceTokenCharged) Reset()         { *m = ResourceTokenCharged{} }
func (m *ResourceTokenCharged) String() string { return proto.CompactTextString(m) }
func (*ResourceTokenCharged) ProtoMessage()    {}
func (*ResourceTokenCharged) Descriptor() ([]byte, []int) {
	return fileDescriptor_a56c8c1c5f610389, []int{3}
}

func (m *ResourceTokenCharged) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceTokenCharged.Unmarshal(m, b)
}
func (m *ResourceTokenCharged) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceTokenCharged.Marshal(b, m, deterministic)
}
func (m *ResourceTokenCharged) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceTokenCharged.Merge(m, src)
}
func (m *ResourceTokenCharged) XXX_Size() int {
	return xxx_messageInfo_ResourceTokenCharged.Size(m)
}
func (m *ResourceTokenCharged) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceTokenCharged.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceTokenCharged proto.InternalMessageInfo

func (m *ResourceTokenCharged) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ResourceTokenCharged) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type ResourceTokenOwned struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol" form:"symbol"`
	Amount               int64    `protobuf:"zigzag64,2,opt,name=amount,proto3" json:"amount" form:"amount"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceTokenOwned) Reset()         { *m = ResourceTokenOwned{} }
func (m *ResourceTokenOwned) String() string { return proto.CompactTextString(m) }
func (*ResourceTokenOwned) ProtoMessage()    {}
func (*ResourceTokenOwned) Descriptor() ([]byte, []int) {
	return fileDescriptor_a56c8c1c5f610389, []int{4}
}

func (m *ResourceTokenOwned) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceTokenOwned.Unmarshal(m, b)
}
func (m *ResourceTokenOwned) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceTokenOwned.Marshal(b, m, deterministic)
}
func (m *ResourceTokenOwned) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceTokenOwned.Merge(m, src)
}
func (m *ResourceTokenOwned) XXX_Size() int {
	return xxx_messageInfo_ResourceTokenOwned.Size(m)
}
func (m *ResourceTokenOwned) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceTokenOwned.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceTokenOwned proto.InternalMessageInfo

func (m *ResourceTokenOwned) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ResourceTokenOwned) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*TransactionSizeFeeSymbols)(nil), "client.TransactionSizeFeeSymbols")
	proto.RegisterType((*TransactionSizeFeeSymbol)(nil), "client.TransactionSizeFeeSymbol")
	proto.RegisterType((*TransactionFeeCharged)(nil), "client.TransactionFeeCharged")
	proto.RegisterType((*ResourceTokenCharged)(nil), "client.ResourceTokenCharged")
	proto.RegisterType((*ResourceTokenOwned)(nil), "client.ResourceTokenOwned")
}

func init() { proto.RegisterFile("transaction_fee.proto", fileDescriptor_a56c8c1c5f610389) }

var fileDescriptor_a56c8c1c5f610389 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x59, 0x2b, 0x05, 0xa7, 0x8a, 0x76, 0xb1, 0x12, 0xc5, 0x43, 0xcc, 0x29, 0x88, 0xe4,
	0xa0, 0x37, 0xaf, 0x42, 0x0f, 0x52, 0x11, 0xb6, 0x05, 0x8f, 0x61, 0x93, 0x8c, 0xcd, 0x62, 0xba,
	0x5b, 0xb2, 0x5b, 0x8a, 0xbd, 0xfb, 0x1b, 0xbc, 0xfa, 0x3f, 0xfc, 0x73, 0xb2, 0xbb, 0x01, 0x93,
	0x43, 0x2e, 0x7a, 0x9c, 0xf7, 0xde, 0x7c, 0xf3, 0x0e, 0x03, 0x13, 0x53, 0x73, 0xa9, 0x79, 0x6e,
	0x84, 0x92, 0xe9, 0x2b, 0x62, 0xb2, 0xae, 0x95, 0x51, 0x74, 0x98, 0x57, 0x02, 0xa5, 0xb9, 0x38,
	0x52, 0x6b, 0xeb, 0x68, 0x2f, 0x47, 0x1f, 0x04, 0xce, 0x17, 0xbf, 0x0b, 0x73, 0xb1, 0xc3, 0x29,
	0xe2, 0xfc, 0x7d, 0x95, 0xa9, 0x4a, 0xd3, 0x12, 0xc2, 0x36, 0x4d, 0x8b, 0x1d, 0x5a, 0x64, 0xaa,
	0x9d, 0x9f, 0x56, 0x42, 0x9b, 0x80, 0x84, 0x83, 0x78, 0x74, 0x1b, 0x26, 0x9e, 0x9f, 0xf4, 0xc1,
	0xd8, 0xa5, 0xe9, 0x71, 0x66, 0x42, 0x9b, 0xe8, 0x93, 0x40, 0xd0, 0xb7, 0x4a, 0xaf, 0xe0, 0xd0,
	0xa8, 0x37, 0x94, 0xcd, 0xdd, 0x80, 0x84, 0x24, 0x3e, 0x60, 0x23, 0xa7, 0x35, 0x91, 0x6b, 0x18,
	0x67, 0x5c, 0x63, 0xea, 0x73, 0x5b, 0x14, 0xcb, 0xd2, 0x04, 0x7b, 0x21, 0x89, 0xc7, 0xec, 0xd8,
	0x1a, 0x0b, 0xab, 0xbf, 0x38, 0x99, 0xde, 0x00, 0xe5, 0x45, 0x81, 0x45, 0x37, 0x3c, 0x70, 0xe1,
	0x13, 0xe7, 0xb4, 0xd2, 0xd1, 0x13, 0x4c, 0x5a, 0xc5, 0xa6, 0x88, 0x0f, 0x25, 0xaf, 0x97, 0x58,
	0xd0, 0x33, 0x18, 0x76, 0xfa, 0x34, 0x93, 0xd5, 0xf9, 0x4a, 0x6d, 0xa4, 0xbf, 0x4f, 0x59, 0x33,
	0xdd, 0xef, 0x7f, 0x7d, 0x07, 0x24, 0x9a, 0xc1, 0x29, 0x43, 0xad, 0x36, 0x75, 0xee, 0x3b, 0xfd,
	0x8f, 0xf6, 0x08, 0xb4, 0x43, 0x7b, 0xde, 0xca, 0xbf, 0xb2, 0xb2, 0xa1, 0xfb, 0x88, 0xbb, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x91, 0x92, 0x18, 0x41, 0x02, 0x00, 0x00,
}