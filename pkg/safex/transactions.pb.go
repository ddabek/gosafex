// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/transactions.proto

package safex

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TxinToScriptCommandType int32

const (
	TxinToScript_nop                    TxinToScriptCommandType = 0
	TxinToScript_token_stake            TxinToScriptCommandType = 1
	TxinToScript_token_unstake          TxinToScriptCommandType = 2
	TxinToScript_token_collect          TxinToScriptCommandType = 3
	TxinToScript_donate_network_fee     TxinToScriptCommandType = 4
	TxinToScript_distribute_network_fee TxinToScriptCommandType = 5
	TxinToScript_simple_purchase        TxinToScriptCommandType = 6
	TxinToScript_create_account         TxinToScriptCommandType = 7
	TxinToScript_edit_account           TxinToScriptCommandType = 8
	TxinToScript_invalid_command        TxinToScriptCommandType = 9
)

var TxinToScriptCommandType_name = map[int32]string{
	0: "nop",
	1: "token_stake",
	2: "token_unstake",
	3: "token_collect",
	4: "donate_network_fee",
	5: "distribute_network_fee",
	6: "simple_purchase",
	7: "create_account",
	8: "edit_account",
	9: "invalid_command",
}
var TxinToScriptCommandType_value = map[string]int32{
	"nop":                    0,
	"token_stake":            1,
	"token_unstake":          2,
	"token_collect":          3,
	"donate_network_fee":     4,
	"distribute_network_fee": 5,
	"simple_purchase":        6,
	"create_account":         7,
	"edit_account":           8,
	"invalid_command":        9,
}

func (x TxinToScriptCommandType) String() string {
	return proto.EnumName(TxinToScriptCommandType_name, int32(x))
}
func (TxinToScriptCommandType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{4, 0} }

type TxinGen struct {
	Height uint64 `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
}

func (m *TxinGen) Reset()                    { *m = TxinGen{} }
func (m *TxinGen) String() string            { return proto.CompactTextString(m) }
func (*TxinGen) ProtoMessage()               {}
func (*TxinGen) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *TxinGen) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type TxinToKey struct {
	Amount     uint64   `protobuf:"varint,1,opt,name=amount" json:"amount,omitempty"`
	KImage     []byte   `protobuf:"bytes,2,opt,name=k_image,json=kImage,proto3" json:"k_image,omitempty"`
	KeyOffsets []uint64 `protobuf:"varint,3,rep,packed,name=key_offsets,json=keyOffsets" json:"key_offsets,omitempty"`
}

func (m *TxinToKey) Reset()                    { *m = TxinToKey{} }
func (m *TxinToKey) String() string            { return proto.CompactTextString(m) }
func (*TxinToKey) ProtoMessage()               {}
func (*TxinToKey) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *TxinToKey) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TxinToKey) GetKImage() []byte {
	if m != nil {
		return m.KImage
	}
	return nil
}

func (m *TxinToKey) GetKeyOffsets() []uint64 {
	if m != nil {
		return m.KeyOffsets
	}
	return nil
}

type TxinTokenToKey struct {
	TokenAmount uint64   `protobuf:"varint,1,opt,name=token_amount,json=tokenAmount" json:"token_amount,omitempty"`
	KImage      []byte   `protobuf:"bytes,2,opt,name=k_image,json=kImage,proto3" json:"k_image,omitempty"`
	KeyOffsets  []uint64 `protobuf:"varint,3,rep,packed,name=key_offsets,json=keyOffsets" json:"key_offsets,omitempty"`
}

func (m *TxinTokenToKey) Reset()                    { *m = TxinTokenToKey{} }
func (m *TxinTokenToKey) String() string            { return proto.CompactTextString(m) }
func (*TxinTokenToKey) ProtoMessage()               {}
func (*TxinTokenToKey) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *TxinTokenToKey) GetTokenAmount() uint64 {
	if m != nil {
		return m.TokenAmount
	}
	return 0
}

func (m *TxinTokenToKey) GetKImage() []byte {
	if m != nil {
		return m.KImage
	}
	return nil
}

func (m *TxinTokenToKey) GetKeyOffsets() []uint64 {
	if m != nil {
		return m.KeyOffsets
	}
	return nil
}

type TxinTokenMigration struct {
	TokenAmount            uint64 `protobuf:"varint,1,opt,name=token_amount,json=tokenAmount" json:"token_amount,omitempty"`
	BitcoinBurnTransaction string `protobuf:"bytes,2,opt,name=bitcoin_burn_transaction,json=bitcoinBurnTransaction" json:"bitcoin_burn_transaction,omitempty"`
	KImage                 []byte `protobuf:"bytes,3,opt,name=k_image,json=kImage,proto3" json:"k_image,omitempty"`
}

func (m *TxinTokenMigration) Reset()                    { *m = TxinTokenMigration{} }
func (m *TxinTokenMigration) String() string            { return proto.CompactTextString(m) }
func (*TxinTokenMigration) ProtoMessage()               {}
func (*TxinTokenMigration) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *TxinTokenMigration) GetTokenAmount() uint64 {
	if m != nil {
		return m.TokenAmount
	}
	return 0
}

func (m *TxinTokenMigration) GetBitcoinBurnTransaction() string {
	if m != nil {
		return m.BitcoinBurnTransaction
	}
	return ""
}

func (m *TxinTokenMigration) GetKImage() []byte {
	if m != nil {
		return m.KImage
	}
	return nil
}

type TxinToScript struct {
	Amount      uint64   `protobuf:"varint,1,opt,name=amount" json:"amount,omitempty"`
	TokenAmount uint64   `protobuf:"varint,2,opt,name=token_amount,json=tokenAmount" json:"token_amount,omitempty"`
	KImage      []byte   `protobuf:"bytes,3,opt,name=k_image,json=kImage,proto3" json:"k_image,omitempty"`
	KeyOffsets  []uint64 `protobuf:"varint,4,rep,packed,name=key_offsets,json=keyOffsets" json:"key_offsets,omitempty"`
	Script      string   `protobuf:"bytes,5,opt,name=script" json:"script,omitempty"`
}

func (m *TxinToScript) Reset()                    { *m = TxinToScript{} }
func (m *TxinToScript) String() string            { return proto.CompactTextString(m) }
func (*TxinToScript) ProtoMessage()               {}
func (*TxinToScript) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *TxinToScript) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TxinToScript) GetTokenAmount() uint64 {
	if m != nil {
		return m.TokenAmount
	}
	return 0
}

func (m *TxinToScript) GetKImage() []byte {
	if m != nil {
		return m.KImage
	}
	return nil
}

func (m *TxinToScript) GetKeyOffsets() []uint64 {
	if m != nil {
		return m.KeyOffsets
	}
	return nil
}

func (m *TxinToScript) GetScript() string {
	if m != nil {
		return m.Script
	}
	return ""
}

type TxinToScripthash struct {
	Prev    []byte         `protobuf:"bytes,1,opt,name=prev,proto3" json:"prev,omitempty"`
	Prevout uint64         `protobuf:"varint,2,opt,name=prevout" json:"prevout,omitempty"`
	Script  *TxoutToScript `protobuf:"bytes,3,opt,name=script" json:"script,omitempty"`
	Sigset  string         `protobuf:"bytes,4,opt,name=sigset" json:"sigset,omitempty"`
}

func (m *TxinToScripthash) Reset()                    { *m = TxinToScripthash{} }
func (m *TxinToScripthash) String() string            { return proto.CompactTextString(m) }
func (*TxinToScripthash) ProtoMessage()               {}
func (*TxinToScripthash) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *TxinToScripthash) GetPrev() []byte {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *TxinToScripthash) GetPrevout() uint64 {
	if m != nil {
		return m.Prevout
	}
	return 0
}

func (m *TxinToScripthash) GetScript() *TxoutToScript {
	if m != nil {
		return m.Script
	}
	return nil
}

func (m *TxinToScripthash) GetSigset() string {
	if m != nil {
		return m.Sigset
	}
	return ""
}

// Variants are implemented with optional fields
type TxinV struct {
	TxinGen            *TxinGen            `protobuf:"bytes,1,opt,name=txin_gen,json=txinGen" json:"txin_gen,omitempty"`
	TxinToKey          *TxinToKey          `protobuf:"bytes,2,opt,name=txin_to_key,json=txinToKey" json:"txin_to_key,omitempty"`
	TxinTokenToKey     *TxinTokenToKey     `protobuf:"bytes,3,opt,name=txin_token_to_key,json=txinTokenToKey" json:"txin_token_to_key,omitempty"`
	TxinTokenMigration *TxinTokenMigration `protobuf:"bytes,4,opt,name=txin_token_migration,json=txinTokenMigration" json:"txin_token_migration,omitempty"`
}

func (m *TxinV) Reset()                    { *m = TxinV{} }
func (m *TxinV) String() string            { return proto.CompactTextString(m) }
func (*TxinV) ProtoMessage()               {}
func (*TxinV) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *TxinV) GetTxinGen() *TxinGen {
	if m != nil {
		return m.TxinGen
	}
	return nil
}

func (m *TxinV) GetTxinToKey() *TxinToKey {
	if m != nil {
		return m.TxinToKey
	}
	return nil
}

func (m *TxinV) GetTxinTokenToKey() *TxinTokenToKey {
	if m != nil {
		return m.TxinTokenToKey
	}
	return nil
}

func (m *TxinV) GetTxinTokenMigration() *TxinTokenMigration {
	if m != nil {
		return m.TxinTokenMigration
	}
	return nil
}

type TxoutToKey struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *TxoutToKey) Reset()                    { *m = TxoutToKey{} }
func (m *TxoutToKey) String() string            { return proto.CompactTextString(m) }
func (*TxoutToKey) ProtoMessage()               {}
func (*TxoutToKey) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *TxoutToKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type TxoutTokenToKey struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *TxoutTokenToKey) Reset()                    { *m = TxoutTokenToKey{} }
func (m *TxoutTokenToKey) String() string            { return proto.CompactTextString(m) }
func (*TxoutTokenToKey) ProtoMessage()               {}
func (*TxoutTokenToKey) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *TxoutTokenToKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type TxoutToScript struct {
	Keys       [][]byte `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	Data       string   `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	OutputType uint64   `protobuf:"varint,3,opt,name=output_type,json=outputType" json:"output_type,omitempty"`
}

func (m *TxoutToScript) Reset()                    { *m = TxoutToScript{} }
func (m *TxoutToScript) String() string            { return proto.CompactTextString(m) }
func (*TxoutToScript) ProtoMessage()               {}
func (*TxoutToScript) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *TxoutToScript) GetKeys() [][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *TxoutToScript) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *TxoutToScript) GetOutputType() uint64 {
	if m != nil {
		return m.OutputType
	}
	return 0
}

type TxoutToScriptHash struct {
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *TxoutToScriptHash) Reset()                    { *m = TxoutToScriptHash{} }
func (m *TxoutToScriptHash) String() string            { return proto.CompactTextString(m) }
func (*TxoutToScriptHash) ProtoMessage()               {}
func (*TxoutToScriptHash) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func (m *TxoutToScriptHash) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type TxoutTargetV struct {
	TxoutToKey      *TxoutToKey      `protobuf:"bytes,1,opt,name=txout_to_key,json=txoutToKey" json:"txout_to_key,omitempty"`
	TxoutTokenToKey *TxoutTokenToKey `protobuf:"bytes,2,opt,name=txout_token_to_key,json=txoutTokenToKey" json:"txout_token_to_key,omitempty"`
}

func (m *TxoutTargetV) Reset()                    { *m = TxoutTargetV{} }
func (m *TxoutTargetV) String() string            { return proto.CompactTextString(m) }
func (*TxoutTargetV) ProtoMessage()               {}
func (*TxoutTargetV) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{11} }

func (m *TxoutTargetV) GetTxoutToKey() *TxoutToKey {
	if m != nil {
		return m.TxoutToKey
	}
	return nil
}

func (m *TxoutTargetV) GetTxoutTokenToKey() *TxoutTokenToKey {
	if m != nil {
		return m.TxoutTokenToKey
	}
	return nil
}

type Txout struct {
	Amount      uint64        `protobuf:"varint,1,opt,name=amount" json:"amount,omitempty"`
	TokenAmount uint64        `protobuf:"varint,2,opt,name=token_amount,json=tokenAmount" json:"token_amount,omitempty"`
	Target      *TxoutTargetV `protobuf:"bytes,3,opt,name=target" json:"target,omitempty"`
}

func (m *Txout) Reset()                    { *m = Txout{} }
func (m *Txout) String() string            { return proto.CompactTextString(m) }
func (*Txout) ProtoMessage()               {}
func (*Txout) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{12} }

func (m *Txout) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Txout) GetTokenAmount() uint64 {
	if m != nil {
		return m.TokenAmount
	}
	return 0
}

func (m *Txout) GetTarget() *TxoutTargetV {
	if m != nil {
		return m.Target
	}
	return nil
}

type SigData struct {
	R []byte `protobuf:"bytes,1,opt,name=r,proto3" json:"r,omitempty"`
	C []byte `protobuf:"bytes,2,opt,name=c,proto3" json:"c,omitempty"`
}

func (m *SigData) Reset()                    { *m = SigData{} }
func (m *SigData) String() string            { return proto.CompactTextString(m) }
func (*SigData) ProtoMessage()               {}
func (*SigData) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{13} }

func (m *SigData) GetR() []byte {
	if m != nil {
		return m.R
	}
	return nil
}

func (m *SigData) GetC() []byte {
	if m != nil {
		return m.C
	}
	return nil
}

type Signature struct {
	Signature []*SigData `protobuf:"bytes,1,rep,name=signature" json:"signature,omitempty"`
}

func (m *Signature) Reset()                    { *m = Signature{} }
func (m *Signature) String() string            { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()               {}
func (*Signature) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{14} }

func (m *Signature) GetSignature() []*SigData {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Transaction struct {
	Version         uint64       `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	UnlockTime      uint64       `protobuf:"varint,2,opt,name=unlock_time,json=unlockTime" json:"unlock_time,omitempty"`
	Extra           []byte       `protobuf:"bytes,3,opt,name=extra,proto3" json:"extra,omitempty"`
	Vin             []*TxinV     `protobuf:"bytes,4,rep,name=vin" json:"vin,omitempty"`
	Vout            []*Txout     `protobuf:"bytes,5,rep,name=vout" json:"vout,omitempty"`
	Signatures      []*Signature `protobuf:"bytes,6,rep,name=signatures" json:"signatures,omitempty"`
	BlockHeight     uint64       `protobuf:"varint,7,opt,name=block_height,json=blockHeight" json:"block_height,omitempty"`
	BlockTimestamp  uint64       `protobuf:"varint,8,opt,name=block_timestamp,json=blockTimestamp" json:"block_timestamp,omitempty"`
	DoubleSpendSeen bool         `protobuf:"varint,9,opt,name=double_spend_seen,json=doubleSpendSeen" json:"double_spend_seen,omitempty"`
	InPool          bool         `protobuf:"varint,10,opt,name=in_pool,json=inPool" json:"in_pool,omitempty"`
	OutputIndices   []uint64     `protobuf:"varint,11,rep,packed,name=output_indices,json=outputIndices" json:"output_indices,omitempty"`
	TxHash          string       `protobuf:"bytes,12,opt,name=tx_hash,json=txHash" json:"tx_hash,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{15} }

func (m *Transaction) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Transaction) GetUnlockTime() uint64 {
	if m != nil {
		return m.UnlockTime
	}
	return 0
}

func (m *Transaction) GetExtra() []byte {
	if m != nil {
		return m.Extra
	}
	return nil
}

func (m *Transaction) GetVin() []*TxinV {
	if m != nil {
		return m.Vin
	}
	return nil
}

func (m *Transaction) GetVout() []*Txout {
	if m != nil {
		return m.Vout
	}
	return nil
}

func (m *Transaction) GetSignatures() []*Signature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *Transaction) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *Transaction) GetBlockTimestamp() uint64 {
	if m != nil {
		return m.BlockTimestamp
	}
	return 0
}

func (m *Transaction) GetDoubleSpendSeen() bool {
	if m != nil {
		return m.DoubleSpendSeen
	}
	return false
}

func (m *Transaction) GetInPool() bool {
	if m != nil {
		return m.InPool
	}
	return false
}

func (m *Transaction) GetOutputIndices() []uint64 {
	if m != nil {
		return m.OutputIndices
	}
	return nil
}

func (m *Transaction) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

type Transactions struct {
	Tx        []*Transaction `protobuf:"bytes,1,rep,name=tx" json:"tx,omitempty"`
	MissedTxs []string       `protobuf:"bytes,2,rep,name=missed_txs,json=missedTxs" json:"missed_txs,omitempty"`
}

func (m *Transactions) Reset()                    { *m = Transactions{} }
func (m *Transactions) String() string            { return proto.CompactTextString(m) }
func (*Transactions) ProtoMessage()               {}
func (*Transactions) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{16} }

func (m *Transactions) GetTx() []*Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *Transactions) GetMissedTxs() []string {
	if m != nil {
		return m.MissedTxs
	}
	return nil
}

func init() {
	proto.RegisterType((*TxinGen)(nil), "safex.txin_gen")
	proto.RegisterType((*TxinToKey)(nil), "safex.txin_to_key")
	proto.RegisterType((*TxinTokenToKey)(nil), "safex.txin_token_to_key")
	proto.RegisterType((*TxinTokenMigration)(nil), "safex.txin_token_migration")
	proto.RegisterType((*TxinToScript)(nil), "safex.txin_to_script")
	proto.RegisterType((*TxinToScripthash)(nil), "safex.txin_to_scripthash")
	proto.RegisterType((*TxinV)(nil), "safex.txin_v")
	proto.RegisterType((*TxoutToKey)(nil), "safex.txout_to_key")
	proto.RegisterType((*TxoutTokenToKey)(nil), "safex.txout_token_to_key")
	proto.RegisterType((*TxoutToScript)(nil), "safex.txout_to_script")
	proto.RegisterType((*TxoutToScriptHash)(nil), "safex.txout_to_script_hash")
	proto.RegisterType((*TxoutTargetV)(nil), "safex.txout_target_v")
	proto.RegisterType((*Txout)(nil), "safex.txout")
	proto.RegisterType((*SigData)(nil), "safex.SigData")
	proto.RegisterType((*Signature)(nil), "safex.Signature")
	proto.RegisterType((*Transaction)(nil), "safex.Transaction")
	proto.RegisterType((*Transactions)(nil), "safex.Transactions")
	proto.RegisterEnum("safex.TxinToScriptCommandType", TxinToScriptCommandType_name, TxinToScriptCommandType_value)
}

func init() { proto.RegisterFile("proto/transactions.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 994 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdb, 0x8e, 0xe3, 0x44,
	0x13, 0xfe, 0x1d, 0xe7, 0x30, 0x29, 0x67, 0x92, 0x4c, 0xef, 0xfe, 0x83, 0x01, 0xa1, 0x0d, 0x96,
	0x16, 0x46, 0x23, 0x18, 0x50, 0x10, 0x12, 0x5c, 0x72, 0x10, 0xec, 0x0a, 0xad, 0x00, 0xcf, 0x5c,
	0x71, 0xd3, 0x72, 0x9c, 0x9a, 0xa4, 0xe5, 0xa4, 0xdb, 0xb8, 0xdb, 0xc1, 0x79, 0x05, 0x84, 0xc4,
	0x4b, 0xf0, 0x44, 0x3c, 0x0a, 0xd7, 0x5c, 0xa0, 0x3e, 0xd8, 0x71, 0xb2, 0xc3, 0x0a, 0x69, 0xaf,
	0xd2, 0xf5, 0x55, 0x75, 0x1d, 0xbe, 0xea, 0x2a, 0x07, 0xc2, 0xbc, 0x10, 0x4a, 0x7c, 0xa4, 0x8a,
	0x84, 0xcb, 0x24, 0x55, 0x4c, 0x70, 0x79, 0x63, 0x20, 0xd2, 0x93, 0xc9, 0x3d, 0x56, 0x51, 0x04,
	0x67, 0xaa, 0x62, 0x9c, 0xae, 0x90, 0x93, 0x4b, 0xe8, 0xaf, 0x91, 0xad, 0xd6, 0x2a, 0xf4, 0x66,
	0xde, 0x55, 0x37, 0x76, 0x52, 0x44, 0x21, 0x30, 0x36, 0x4a, 0xd0, 0x0c, 0xf7, 0xda, 0x2c, 0xd9,
	0x8a, 0x92, 0x37, 0x66, 0x56, 0x22, 0x6f, 0xc0, 0x20, 0xa3, 0x6c, 0x9b, 0xac, 0x30, 0xec, 0xcc,
	0xbc, 0xab, 0x51, 0xdc, 0xcf, 0x9e, 0x6b, 0x89, 0x3c, 0x81, 0x20, 0xc3, 0x3d, 0x15, 0xf7, 0xf7,
	0x12, 0x95, 0x0c, 0xfd, 0x99, 0x7f, 0xd5, 0x8d, 0x21, 0xc3, 0xfd, 0xf7, 0x16, 0x89, 0x72, 0xb8,
	0x70, 0x01, 0x32, 0x6c, 0xc2, 0xbc, 0x0b, 0x23, 0x2b, 0x1f, 0x05, 0x0b, 0x0c, 0xf6, 0xc5, 0xeb,
	0x46, 0xfc, 0xcd, 0x83, 0xc7, 0xad, 0x90, 0x5b, 0xb6, 0x2a, 0x12, 0xcd, 0xce, 0x7f, 0x89, 0xfa,
	0x19, 0x84, 0x0b, 0xa6, 0x52, 0xc1, 0x38, 0x5d, 0x94, 0x05, 0xa7, 0x2d, 0x72, 0x4d, 0x1a, 0xc3,
	0xf8, 0xd2, 0xe9, 0xbf, 0x2c, 0x0b, 0x7e, 0x77, 0xd0, 0xb6, 0xf3, 0xf5, 0xdb, 0xf9, 0x46, 0x7f,
	0x75, 0x60, 0x5c, 0x53, 0x2c, 0xd3, 0x82, 0xe5, 0xea, 0x5f, 0x59, 0x3e, 0x4d, 0xb0, 0xf3, 0x4a,
	0x5a, 0xfc, 0x57, 0xd1, 0xd2, 0x3d, 0xa5, 0x45, 0x07, 0xb5, 0xe1, 0xc3, 0x9e, 0x29, 0xc4, 0x49,
	0xd1, 0x9f, 0x1e, 0x8c, 0x52, 0xb1, 0xdd, 0x26, 0x7c, 0x49, 0xd5, 0x3e, 0x47, 0x32, 0x00, 0x9f,
	0x8b, 0x7c, 0xfa, 0x3f, 0x32, 0x01, 0x1b, 0x9a, 0x4a, 0x95, 0x64, 0x38, 0xf5, 0xc8, 0x05, 0x9c,
	0x5b, 0xa0, 0xe4, 0x16, 0xea, 0x1c, 0xa0, 0x54, 0x6c, 0x36, 0x98, 0xaa, 0xa9, 0x4f, 0x2e, 0x81,
	0x2c, 0x05, 0x4f, 0x14, 0x52, 0x8e, 0xea, 0x17, 0x51, 0x64, 0xf4, 0x1e, 0x71, 0xda, 0x25, 0x6f,
	0xc1, 0xe5, 0x92, 0x49, 0x55, 0xb0, 0x45, 0x79, 0xa2, 0xeb, 0x91, 0x47, 0x30, 0x91, 0x6c, 0x9b,
	0x6f, 0x90, 0xe6, 0x65, 0x91, 0xae, 0x13, 0x89, 0xd3, 0x3e, 0x21, 0x30, 0x4e, 0x0b, 0xd4, 0x8e,
	0x92, 0x34, 0xd5, 0xd5, 0x4f, 0x07, 0x64, 0x0a, 0x23, 0x5c, 0x32, 0xd5, 0x20, 0x67, 0xfa, 0x2a,
	0xe3, 0xbb, 0x64, 0xc3, 0x96, 0xd4, 0x95, 0x31, 0x1d, 0x46, 0xbf, 0x7a, 0x40, 0x8e, 0x49, 0x5f,
	0x27, 0x72, 0x4d, 0x08, 0x74, 0xf3, 0x02, 0x77, 0x86, 0xf6, 0x51, 0x6c, 0xce, 0x24, 0x84, 0x81,
	0xfe, 0x15, 0x65, 0xcd, 0x77, 0x2d, 0x92, 0x9b, 0x86, 0x31, 0x4d, 0x75, 0x30, 0xbf, 0xbc, 0x31,
	0x73, 0x75, 0xa3, 0x2a, 0x51, 0xaa, 0x83, 0xe7, 0x9a, 0x49, 0xc3, 0x30, 0x5b, 0x49, 0x54, 0x61,
	0xd7, 0x31, 0x6c, 0xa4, 0xe8, 0x6f, 0x0f, 0xfa, 0x26, 0x99, 0x1d, 0xb9, 0x3e, 0x8c, 0xa4, 0x49,
	0x22, 0x98, 0x4f, 0x1a, 0xa7, 0x16, 0x8e, 0x07, 0xfa, 0xf4, 0x2d, 0x72, 0x32, 0x3f, 0x1a, 0x4d,
	0x93, 0x5c, 0x30, 0x27, 0x6d, 0x73, 0xab, 0x89, 0x87, 0x5a, 0xb8, 0x13, 0xdf, 0xe1, 0x9e, 0x7c,
	0xf5, 0xc0, 0xb4, 0xb9, 0xec, 0xc3, 0xe3, 0x9b, 0x07, 0x7d, 0x3c, 0xb6, 0xf7, 0x33, 0x74, 0x4e,
	0x5e, 0x3c, 0x3c, 0x3f, 0xa6, 0xaa, 0x60, 0xfe, 0xf6, 0xcb, 0x7e, 0x1a, 0x93, 0x98, 0x34, 0xae,
	0x5e, 0xd4, 0x58, 0x34, 0x83, 0x51, 0xc3, 0x98, 0x1e, 0xfe, 0x29, 0xf8, 0x3a, 0x2b, 0xdb, 0x03,
	0x7d, 0x8c, 0xde, 0xd3, 0xcd, 0xb2, 0x16, 0xad, 0x25, 0xf1, 0xb2, 0xdd, 0x4f, 0x30, 0x39, 0xe1,
	0x5e, 0x77, 0x34, 0xc3, 0xbd, 0x0c, 0xbd, 0x99, 0xaf, 0x3b, 0xaa, 0xcf, 0x1a, 0x5b, 0x26, 0x2a,
	0x71, 0x03, 0x6b, 0xce, 0x7a, 0x3c, 0x44, 0xa9, 0x72, 0x7d, 0x77, 0x9f, 0xdb, 0xd9, 0xe9, 0xc6,
	0x60, 0xa1, 0xbb, 0x7d, 0x8e, 0xd1, 0xb5, 0x2e, 0xfa, 0xc8, 0x37, 0xad, 0x9f, 0x8c, 0xfe, 0xad,
	0x9f, 0x8c, 0x3e, 0x47, 0xbf, 0x7b, 0x7a, 0xa4, 0x8d, 0x71, 0x52, 0xac, 0x50, 0xd1, 0x1d, 0xf9,
	0xf4, 0xb8, 0x48, 0xd7, 0xdc, 0x47, 0xa7, 0x2f, 0x46, 0xd3, 0x0d, 0x46, 0xb2, 0x54, 0x7f, 0xf3,
	0x50, 0xe5, 0xae, 0xd5, 0x6f, 0x9e, 0x5c, 0x6e, 0x75, 0x6c, 0xe2, 0x5c, 0xd4, 0x2d, 0x8b, 0x7e,
	0x86, 0x9e, 0x81, 0x5e, 0x67, 0xb5, 0x7c, 0x08, 0x7d, 0x5b, 0x8e, 0x7b, 0x30, 0xff, 0x3f, 0x8e,
	0xef, 0x2a, 0x8d, 0x9d, 0x51, 0xf4, 0x14, 0x06, 0xb7, 0x6c, 0xf5, 0xb5, 0x26, 0x77, 0x04, 0x5e,
	0xe1, 0x08, 0xf2, 0x0a, 0x2d, 0xa5, 0x6e, 0x67, 0x7b, 0x69, 0xf4, 0x39, 0x0c, 0x6f, 0xd9, 0x8a,
	0x27, 0xaa, 0x2c, 0x90, 0x7c, 0x00, 0x43, 0x59, 0x0b, 0xa6, 0x65, 0xc1, 0x7c, 0xec, 0xa2, 0x38,
	0x5f, 0xf1, 0xc1, 0x20, 0xfa, 0xc3, 0x87, 0xa0, 0xbd, 0x62, 0x43, 0x18, 0xec, 0xb0, 0x90, 0xfa,
	0x29, 0xda, 0xe2, 0x6a, 0x51, 0x77, 0xb7, 0xe4, 0x1b, 0x91, 0x66, 0x54, 0xb1, 0x2d, 0xba, 0xe2,
	0xc0, 0x42, 0x77, 0x6c, 0x8b, 0xe4, 0x31, 0xf4, 0xb0, 0x52, 0x45, 0xe2, 0x96, 0xa6, 0x15, 0xc8,
	0x13, 0xf0, 0x77, 0x8c, 0x9b, 0x5d, 0x19, 0xcc, 0xcf, 0xdb, 0xef, 0x7a, 0x17, 0x6b, 0x0d, 0x99,
	0x41, 0xd7, 0x2c, 0x86, 0x9e, 0xb1, 0x18, 0xb5, 0x09, 0x89, 0x8d, 0x86, 0x7c, 0x0c, 0xd0, 0x24,
	0x2c, 0xc3, 0xbe, 0xb1, 0x9b, 0x1e, 0x4a, 0xb2, 0x8a, 0xb8, 0x65, 0xa3, 0x3b, 0xb1, 0x30, 0xa9,
	0xba, 0xef, 0xf1, 0xc0, 0x76, 0xc2, 0x60, 0xcf, 0x0c, 0x44, 0xde, 0x87, 0xc9, 0xa2, 0xa9, 0x46,
	0xaa, 0x64, 0x9b, 0x87, 0x67, 0xc6, 0x6a, 0xbc, 0xa8, 0x2b, 0x32, 0x28, 0xb9, 0x86, 0x8b, 0xa5,
	0x28, 0x17, 0x1b, 0xa4, 0x32, 0x47, 0xbe, 0xa4, 0x12, 0x91, 0x87, 0xc3, 0x99, 0x77, 0x75, 0x16,
	0x4f, 0xac, 0xe2, 0x56, 0xe3, 0xb7, 0x88, 0xe6, 0x03, 0xc5, 0x38, 0xcd, 0x85, 0xd8, 0x84, 0x60,
	0x2c, 0xfa, 0x8c, 0xff, 0x20, 0xc4, 0x86, 0x3c, 0x85, 0xb1, 0x1b, 0x0d, 0xc6, 0x97, 0x2c, 0x45,
	0x19, 0x06, 0xe6, 0xe3, 0x71, 0x6e, 0xd1, 0xe7, 0x16, 0xd4, 0xf7, 0x55, 0x65, 0x66, 0x22, 0x1c,
	0xd9, 0xf5, 0xa6, 0xaa, 0x67, 0x7a, 0x1a, 0x7e, 0x84, 0x51, 0xab, 0x4b, 0x92, 0x44, 0xd0, 0x51,
	0x95, 0xeb, 0x6e, 0xbd, 0xae, 0x5a, 0x06, 0x71, 0x47, 0x55, 0xe4, 0x1d, 0x80, 0x2d, 0x93, 0x12,
	0x97, 0x54, 0x55, 0x32, 0xec, 0xcc, 0xfc, 0xab, 0x61, 0x3c, 0xb4, 0xc8, 0x5d, 0x25, 0x17, 0x7d,
	0xf3, 0x3f, 0xe6, 0x93, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x57, 0xe6, 0x5c, 0xe3, 0x08,
	0x00, 0x00,
}
