// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/blocks.proto

/*
Package safex is a generated protocol buffer package.

It is generated from these files:
	proto/blocks.proto
	proto/get_outs.proto
	proto/output_histogram.proto
	proto/safex_account.proto
	proto/transactions.proto

It has these top-level messages:
	BlockHeader
	Block
	Blocks
	OutEntry
	Outs
	Histogram
	Histograms
	CreateAccountData
	TxinGen
	TxinToKey
	TxinTokenToKey
	TxinTokenMigration
	TxinToScript
	TxinToScripthash
	TxinV
	TxoutToKey
	TxoutTokenToKey
	TxoutToScript
	TxoutToScriptHash
	TxoutTargetV
	Txout
	SigData
	Signature
	Transaction
	Transactions
*/
package safex

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BlockHeader struct {
	Depth        uint64 `protobuf:"varint,1,opt,name=depth" json:"depth,omitempty"`
	Hash         string `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	MajorVersion uint32 `protobuf:"varint,3,opt,name=major_version,json=majorVersion" json:"major_version,omitempty"`
	MinorVersion uint32 `protobuf:"varint,4,opt,name=minor_version,json=minorVersion" json:"minor_version,omitempty"`
	PrevHash     string `protobuf:"bytes,5,opt,name=prev_hash,json=prevHash" json:"prev_hash,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BlockHeader) GetDepth() uint64 {
	if m != nil {
		return m.Depth
	}
	return 0
}

func (m *BlockHeader) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *BlockHeader) GetMajorVersion() uint32 {
	if m != nil {
		return m.MajorVersion
	}
	return 0
}

func (m *BlockHeader) GetMinorVersion() uint32 {
	if m != nil {
		return m.MinorVersion
	}
	return 0
}

func (m *BlockHeader) GetPrevHash() string {
	if m != nil {
		return m.PrevHash
	}
	return ""
}

type Block struct {
	Header  *BlockHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Txs     []string     `protobuf:"bytes,2,rep,name=txs" json:"txs,omitempty"`
	MinerTx string       `protobuf:"bytes,3,opt,name=miner_tx,json=minerTx" json:"miner_tx,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetTxs() []string {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *Block) GetMinerTx() string {
	if m != nil {
		return m.MinerTx
	}
	return ""
}

type Blocks struct {
	Block     []*Block `protobuf:"bytes,1,rep,name=block" json:"block,omitempty"`
	Status    bool     `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	Untrusted bool     `protobuf:"varint,3,opt,name=untrusted" json:"untrusted,omitempty"`
	Error     string   `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
}

func (m *Blocks) Reset()                    { *m = Blocks{} }
func (m *Blocks) String() string            { return proto.CompactTextString(m) }
func (*Blocks) ProtoMessage()               {}
func (*Blocks) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Blocks) GetBlock() []*Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *Blocks) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *Blocks) GetUntrusted() bool {
	if m != nil {
		return m.Untrusted
	}
	return false
}

func (m *Blocks) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*BlockHeader)(nil), "safex.BlockHeader")
	proto.RegisterType((*Block)(nil), "safex.Block")
	proto.RegisterType((*Blocks)(nil), "safex.Blocks")
}

func init() { proto.RegisterFile("proto/blocks.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x4f, 0x4f, 0x84, 0x30,
	0x10, 0xc5, 0xd3, 0xe5, 0x8f, 0x74, 0xd8, 0x4d, 0xcc, 0xc4, 0x18, 0x8c, 0x1e, 0x08, 0x5e, 0x88,
	0x07, 0x4c, 0xd6, 0x6f, 0xe0, 0x69, 0xcf, 0x8d, 0xf1, 0x8a, 0x5d, 0xa9, 0x01, 0x75, 0x29, 0x69,
	0xcb, 0x86, 0x4f, 0xe3, 0x67, 0x35, 0x1d, 0x30, 0x72, 0x9b, 0xf7, 0xe3, 0x85, 0xf7, 0x66, 0x0a,
	0x38, 0x18, 0xed, 0xf4, 0xe3, 0xf1, 0x5b, 0xbf, 0x7f, 0xd9, 0x8a, 0x04, 0x46, 0x56, 0x7e, 0xa8,
	0xa9, 0xf8, 0x61, 0x90, 0x3e, 0x7b, 0x7e, 0x50, 0xb2, 0x51, 0x06, 0xaf, 0x20, 0x6a, 0xd4, 0xe0,
	0xda, 0x8c, 0xe5, 0xac, 0x0c, 0xc5, 0x2c, 0x10, 0x21, 0x6c, 0xa5, 0x6d, 0xb3, 0x4d, 0xce, 0x4a,
	0x2e, 0x68, 0xc6, 0x7b, 0xd8, 0x9d, 0xe4, 0xa7, 0x36, 0xf5, 0x59, 0x19, 0xdb, 0xe9, 0x3e, 0x0b,
	0x72, 0x56, 0xee, 0xc4, 0x96, 0xe0, 0xeb, 0xcc, 0xc8, 0xd4, 0xf5, 0x2b, 0x53, 0xb8, 0x98, 0x3c,
	0xfc, 0x33, 0xdd, 0x02, 0x1f, 0x8c, 0x3a, 0xd7, 0x14, 0x11, 0x51, 0x44, 0xe2, 0xc1, 0x41, 0xda,
	0xb6, 0x78, 0x83, 0x88, 0xfa, 0xe1, 0x03, 0xc4, 0x2d, 0x75, 0xa4, 0x6a, 0xe9, 0x1e, 0x2b, 0xda,
	0xa0, 0x5a, 0xb5, 0x17, 0x8b, 0x03, 0x2f, 0x21, 0x70, 0x93, 0xcd, 0x36, 0x79, 0x50, 0x72, 0xe1,
	0x47, 0xbc, 0x81, 0xe4, 0xd4, 0xf5, 0xca, 0xd4, 0x6e, 0xa2, 0xa2, 0x5c, 0x5c, 0x90, 0x7e, 0x99,
	0x8a, 0x09, 0x62, 0xfa, 0x87, 0xc5, 0x02, 0x22, 0xba, 0x51, 0xc6, 0xf2, 0xa0, 0x4c, 0xf7, 0xdb,
	0x75, 0x82, 0x98, 0x3f, 0xe1, 0x35, 0xc4, 0xd6, 0x49, 0x37, 0x5a, 0x3a, 0x46, 0x22, 0x16, 0x85,
	0x77, 0xc0, 0xc7, 0xde, 0x99, 0xd1, 0x3a, 0xd5, 0x50, 0x42, 0x22, 0xfe, 0x81, 0x3f, 0xab, 0x32,
	0x46, 0x1b, 0xda, 0x9f, 0x8b, 0x59, 0x1c, 0x63, 0x7a, 0x8a, 0xa7, 0xdf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x9d, 0x05, 0xe9, 0xa9, 0xa0, 0x01, 0x00, 0x00,
}
