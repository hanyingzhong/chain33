// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateTx struct {
	To          string `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
	Amount      int64  `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	Fee         int64  `protobuf:"varint,3,opt,name=fee" json:"fee,omitempty"`
	Note        string `protobuf:"bytes,4,opt,name=note" json:"note,omitempty"`
	IsToken     bool   `protobuf:"varint,5,opt,name=isToken" json:"isToken,omitempty"`
	TokenSymbol string `protobuf:"bytes,6,opt,name=tokenSymbol" json:"tokenSymbol,omitempty"`
}

func (m *CreateTx) Reset()                    { *m = CreateTx{} }
func (m *CreateTx) String() string            { return proto.CompactTextString(m) }
func (*CreateTx) ProtoMessage()               {}
func (*CreateTx) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *CreateTx) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *CreateTx) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *CreateTx) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *CreateTx) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *CreateTx) GetIsToken() bool {
	if m != nil {
		return m.IsToken
	}
	return false
}

func (m *CreateTx) GetTokenSymbol() string {
	if m != nil {
		return m.TokenSymbol
	}
	return ""
}

type UnsignTx struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *UnsignTx) Reset()                    { *m = UnsignTx{} }
func (m *UnsignTx) String() string            { return proto.CompactTextString(m) }
func (*UnsignTx) ProtoMessage()               {}
func (*UnsignTx) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *UnsignTx) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type SignedTx struct {
	Unsign []byte `protobuf:"bytes,1,opt,name=unsign,proto3" json:"unsign,omitempty"`
	Sign   []byte `protobuf:"bytes,2,opt,name=sign,proto3" json:"sign,omitempty"`
	Pubkey []byte `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Ty     int32  `protobuf:"varint,4,opt,name=ty" json:"ty,omitempty"`
}

func (m *SignedTx) Reset()                    { *m = SignedTx{} }
func (m *SignedTx) String() string            { return proto.CompactTextString(m) }
func (*SignedTx) ProtoMessage()               {}
func (*SignedTx) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *SignedTx) GetUnsign() []byte {
	if m != nil {
		return m.Unsign
	}
	return nil
}

func (m *SignedTx) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *SignedTx) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *SignedTx) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

type Transaction struct {
	Execer    []byte     `protobuf:"bytes,1,opt,name=execer,proto3" json:"execer,omitempty"`
	Payload   []byte     `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature *Signature `protobuf:"bytes,3,opt,name=signature" json:"signature,omitempty"`
	Fee       int64      `protobuf:"varint,4,opt,name=fee" json:"fee,omitempty"`
	Expire    int64      `protobuf:"varint,5,opt,name=expire" json:"expire,omitempty"`
	// 随机ID，可以防止payload 相同的时候，交易重复
	Nonce int64 `protobuf:"varint,6,opt,name=nonce" json:"nonce,omitempty"`
	// 对方地址，如果没有对方地址，可以为空
	To string `protobuf:"bytes,7,opt,name=to" json:"to,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func (m *Transaction) GetExecer() []byte {
	if m != nil {
		return m.Execer
	}
	return nil
}

func (m *Transaction) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Transaction) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Transaction) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *Transaction) GetExpire() int64 {
	if m != nil {
		return m.Expire
	}
	return 0
}

func (m *Transaction) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Transaction) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

// 对于一个交易组中的交易，要么全部成功，要么全部失败
// 这个要好好设计一下
// 最好交易构成一个链条[prevhash].独立的交易构成链条
// 只要这个组中有一个执行是出错的，那么就执行不成功
// 三种签名支持
// ty = 1 -> secp256k1
// ty = 2 -> ed25519
// ty = 3 -> sm2
type Signature struct {
	Ty        int32  `protobuf:"varint,1,opt,name=ty" json:"ty,omitempty"`
	Pubkey    []byte `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Signature) Reset()                    { *m = Signature{} }
func (m *Signature) String() string            { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()               {}
func (*Signature) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

func (m *Signature) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

func (m *Signature) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *Signature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type AddrOverview struct {
	Reciver int64 `protobuf:"varint,1,opt,name=reciver" json:"reciver,omitempty"`
	Balance int64 `protobuf:"varint,2,opt,name=balance" json:"balance,omitempty"`
	TxCount int64 `protobuf:"varint,3,opt,name=txCount" json:"txCount,omitempty"`
}

func (m *AddrOverview) Reset()                    { *m = AddrOverview{} }
func (m *AddrOverview) String() string            { return proto.CompactTextString(m) }
func (*AddrOverview) ProtoMessage()               {}
func (*AddrOverview) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{5} }

func (m *AddrOverview) GetReciver() int64 {
	if m != nil {
		return m.Reciver
	}
	return 0
}

func (m *AddrOverview) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *AddrOverview) GetTxCount() int64 {
	if m != nil {
		return m.TxCount
	}
	return 0
}

type ReqAddr struct {
	Addr string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	// 表示取所有/from/to/其他的hash列表
	Flag      int32 `protobuf:"varint,2,opt,name=flag" json:"flag,omitempty"`
	Count     int32 `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	Direction int32 `protobuf:"varint,4,opt,name=direction" json:"direction,omitempty"`
	Height    int64 `protobuf:"varint,5,opt,name=height" json:"height,omitempty"`
	Index     int64 `protobuf:"varint,6,opt,name=index" json:"index,omitempty"`
}

func (m *ReqAddr) Reset()                    { *m = ReqAddr{} }
func (m *ReqAddr) String() string            { return proto.CompactTextString(m) }
func (*ReqAddr) ProtoMessage()               {}
func (*ReqAddr) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{6} }

func (m *ReqAddr) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ReqAddr) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *ReqAddr) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqAddr) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *ReqAddr) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ReqAddr) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type HexTx struct {
	Tx string `protobuf:"bytes,1,opt,name=tx" json:"tx,omitempty"`
}

func (m *HexTx) Reset()                    { *m = HexTx{} }
func (m *HexTx) String() string            { return proto.CompactTextString(m) }
func (*HexTx) ProtoMessage()               {}
func (*HexTx) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{7} }

func (m *HexTx) GetTx() string {
	if m != nil {
		return m.Tx
	}
	return ""
}

type ReplyTxInfo struct {
	Hash   []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Height int64  `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	Index  int64  `protobuf:"varint,3,opt,name=index" json:"index,omitempty"`
}

func (m *ReplyTxInfo) Reset()                    { *m = ReplyTxInfo{} }
func (m *ReplyTxInfo) String() string            { return proto.CompactTextString(m) }
func (*ReplyTxInfo) ProtoMessage()               {}
func (*ReplyTxInfo) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{8} }

func (m *ReplyTxInfo) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *ReplyTxInfo) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ReplyTxInfo) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type ReqTxList struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *ReqTxList) Reset()                    { *m = ReqTxList{} }
func (m *ReqTxList) String() string            { return proto.CompactTextString(m) }
func (*ReqTxList) ProtoMessage()               {}
func (*ReqTxList) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{9} }

func (m *ReqTxList) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ReplyTxList struct {
	Txs []*Transaction `protobuf:"bytes,1,rep,name=txs" json:"txs,omitempty"`
}

func (m *ReplyTxList) Reset()                    { *m = ReplyTxList{} }
func (m *ReplyTxList) String() string            { return proto.CompactTextString(m) }
func (*ReplyTxList) ProtoMessage()               {}
func (*ReplyTxList) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{10} }

func (m *ReplyTxList) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

type TxHashList struct {
	Hashes [][]byte `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty"`
	Count  int64    `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *TxHashList) Reset()                    { *m = TxHashList{} }
func (m *TxHashList) String() string            { return proto.CompactTextString(m) }
func (*TxHashList) ProtoMessage()               {}
func (*TxHashList) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{11} }

func (m *TxHashList) GetHashes() [][]byte {
	if m != nil {
		return m.Hashes
	}
	return nil
}

func (m *TxHashList) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ReplyTxInfos struct {
	TxInfos []*ReplyTxInfo `protobuf:"bytes,1,rep,name=txInfos" json:"txInfos,omitempty"`
}

func (m *ReplyTxInfos) Reset()                    { *m = ReplyTxInfos{} }
func (m *ReplyTxInfos) String() string            { return proto.CompactTextString(m) }
func (*ReplyTxInfos) ProtoMessage()               {}
func (*ReplyTxInfos) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{12} }

func (m *ReplyTxInfos) GetTxInfos() []*ReplyTxInfo {
	if m != nil {
		return m.TxInfos
	}
	return nil
}

type ReceiptLog struct {
	Ty  int32  `protobuf:"varint,1,opt,name=ty" json:"ty,omitempty"`
	Log []byte `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
}

func (m *ReceiptLog) Reset()                    { *m = ReceiptLog{} }
func (m *ReceiptLog) String() string            { return proto.CompactTextString(m) }
func (*ReceiptLog) ProtoMessage()               {}
func (*ReceiptLog) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{13} }

func (m *ReceiptLog) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

func (m *ReceiptLog) GetLog() []byte {
	if m != nil {
		return m.Log
	}
	return nil
}

// ty = 0 -> error Receipt
// ty = 1 -> CutFee //cut fee ,bug exec not ok
// ty = 2 -> exec ok
type Receipt struct {
	Ty   int32         `protobuf:"varint,1,opt,name=ty" json:"ty,omitempty"`
	KV   []*KeyValue   `protobuf:"bytes,2,rep,name=KV" json:"KV,omitempty"`
	Logs []*ReceiptLog `protobuf:"bytes,3,rep,name=logs" json:"logs,omitempty"`
}

func (m *Receipt) Reset()                    { *m = Receipt{} }
func (m *Receipt) String() string            { return proto.CompactTextString(m) }
func (*Receipt) ProtoMessage()               {}
func (*Receipt) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{14} }

func (m *Receipt) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

func (m *Receipt) GetKV() []*KeyValue {
	if m != nil {
		return m.KV
	}
	return nil
}

func (m *Receipt) GetLogs() []*ReceiptLog {
	if m != nil {
		return m.Logs
	}
	return nil
}

type ReceiptData struct {
	Ty   int32         `protobuf:"varint,1,opt,name=ty" json:"ty,omitempty"`
	Logs []*ReceiptLog `protobuf:"bytes,3,rep,name=logs" json:"logs,omitempty"`
}

func (m *ReceiptData) Reset()                    { *m = ReceiptData{} }
func (m *ReceiptData) String() string            { return proto.CompactTextString(m) }
func (*ReceiptData) ProtoMessage()               {}
func (*ReceiptData) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{15} }

func (m *ReceiptData) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

func (m *ReceiptData) GetLogs() []*ReceiptLog {
	if m != nil {
		return m.Logs
	}
	return nil
}

type TxResult struct {
	Height      int64        `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
	Index       int32        `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
	Tx          *Transaction `protobuf:"bytes,3,opt,name=tx" json:"tx,omitempty"`
	Receiptdate *ReceiptData `protobuf:"bytes,4,opt,name=receiptdate" json:"receiptdate,omitempty"`
	Blocktime   int64        `protobuf:"varint,5,opt,name=blocktime" json:"blocktime,omitempty"`
	ActionName  string       `protobuf:"bytes,6,opt,name=actionName" json:"actionName,omitempty"`
}

func (m *TxResult) Reset()                    { *m = TxResult{} }
func (m *TxResult) String() string            { return proto.CompactTextString(m) }
func (*TxResult) ProtoMessage()               {}
func (*TxResult) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{16} }

func (m *TxResult) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *TxResult) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *TxResult) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *TxResult) GetReceiptdate() *ReceiptData {
	if m != nil {
		return m.Receiptdate
	}
	return nil
}

func (m *TxResult) GetBlocktime() int64 {
	if m != nil {
		return m.Blocktime
	}
	return 0
}

func (m *TxResult) GetActionName() string {
	if m != nil {
		return m.ActionName
	}
	return ""
}

type TransactionDetail struct {
	Tx         *Transaction `protobuf:"bytes,1,opt,name=tx" json:"tx,omitempty"`
	Receipt    *ReceiptData `protobuf:"bytes,2,opt,name=receipt" json:"receipt,omitempty"`
	Proofs     [][]byte     `protobuf:"bytes,3,rep,name=proofs,proto3" json:"proofs,omitempty"`
	Height     int64        `protobuf:"varint,4,opt,name=height" json:"height,omitempty"`
	Index      int64        `protobuf:"varint,5,opt,name=index" json:"index,omitempty"`
	Blocktime  int64        `protobuf:"varint,6,opt,name=blocktime" json:"blocktime,omitempty"`
	Amount     int64        `protobuf:"varint,7,opt,name=amount" json:"amount,omitempty"`
	Fromaddr   string       `protobuf:"bytes,8,opt,name=fromaddr" json:"fromaddr,omitempty"`
	ActionName string       `protobuf:"bytes,9,opt,name=actionName" json:"actionName,omitempty"`
}

func (m *TransactionDetail) Reset()                    { *m = TransactionDetail{} }
func (m *TransactionDetail) String() string            { return proto.CompactTextString(m) }
func (*TransactionDetail) ProtoMessage()               {}
func (*TransactionDetail) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{17} }

func (m *TransactionDetail) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *TransactionDetail) GetReceipt() *ReceiptData {
	if m != nil {
		return m.Receipt
	}
	return nil
}

func (m *TransactionDetail) GetProofs() [][]byte {
	if m != nil {
		return m.Proofs
	}
	return nil
}

func (m *TransactionDetail) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *TransactionDetail) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *TransactionDetail) GetBlocktime() int64 {
	if m != nil {
		return m.Blocktime
	}
	return 0
}

func (m *TransactionDetail) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TransactionDetail) GetFromaddr() string {
	if m != nil {
		return m.Fromaddr
	}
	return ""
}

func (m *TransactionDetail) GetActionName() string {
	if m != nil {
		return m.ActionName
	}
	return ""
}

type TransactionDetails struct {
	Txs []*TransactionDetail `protobuf:"bytes,1,rep,name=txs" json:"txs,omitempty"`
}

func (m *TransactionDetails) Reset()                    { *m = TransactionDetails{} }
func (m *TransactionDetails) String() string            { return proto.CompactTextString(m) }
func (*TransactionDetails) ProtoMessage()               {}
func (*TransactionDetails) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{18} }

func (m *TransactionDetails) GetTxs() []*TransactionDetail {
	if m != nil {
		return m.Txs
	}
	return nil
}

type ReqAddrs struct {
	Addrs []string `protobuf:"bytes,1,rep,name=addrs" json:"addrs,omitempty"`
}

func (m *ReqAddrs) Reset()                    { *m = ReqAddrs{} }
func (m *ReqAddrs) String() string            { return proto.CompactTextString(m) }
func (*ReqAddrs) ProtoMessage()               {}
func (*ReqAddrs) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{19} }

func (m *ReqAddrs) GetAddrs() []string {
	if m != nil {
		return m.Addrs
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTx)(nil), "types.CreateTx")
	proto.RegisterType((*UnsignTx)(nil), "types.UnsignTx")
	proto.RegisterType((*SignedTx)(nil), "types.SignedTx")
	proto.RegisterType((*Transaction)(nil), "types.Transaction")
	proto.RegisterType((*Signature)(nil), "types.Signature")
	proto.RegisterType((*AddrOverview)(nil), "types.AddrOverview")
	proto.RegisterType((*ReqAddr)(nil), "types.ReqAddr")
	proto.RegisterType((*HexTx)(nil), "types.HexTx")
	proto.RegisterType((*ReplyTxInfo)(nil), "types.ReplyTxInfo")
	proto.RegisterType((*ReqTxList)(nil), "types.ReqTxList")
	proto.RegisterType((*ReplyTxList)(nil), "types.ReplyTxList")
	proto.RegisterType((*TxHashList)(nil), "types.TxHashList")
	proto.RegisterType((*ReplyTxInfos)(nil), "types.ReplyTxInfos")
	proto.RegisterType((*ReceiptLog)(nil), "types.ReceiptLog")
	proto.RegisterType((*Receipt)(nil), "types.Receipt")
	proto.RegisterType((*ReceiptData)(nil), "types.ReceiptData")
	proto.RegisterType((*TxResult)(nil), "types.TxResult")
	proto.RegisterType((*TransactionDetail)(nil), "types.TransactionDetail")
	proto.RegisterType((*TransactionDetails)(nil), "types.TransactionDetails")
	proto.RegisterType((*ReqAddrs)(nil), "types.ReqAddrs")
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 842 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x51, 0x8f, 0xe3, 0x34,
	0x10, 0x56, 0x92, 0xb6, 0xdb, 0x4e, 0x2b, 0xd8, 0xb5, 0xd0, 0x11, 0x9d, 0xd0, 0x51, 0x2c, 0x90,
	0x56, 0xe8, 0xd4, 0x87, 0x3b, 0x9e, 0x10, 0x0f, 0xa0, 0xdb, 0x87, 0x43, 0x7b, 0xe2, 0x84, 0x37,
	0xec, 0x13, 0x42, 0x72, 0x9b, 0x69, 0x1b, 0x6d, 0x1a, 0xf7, 0x12, 0x77, 0x49, 0x7f, 0x07, 0xbc,
	0xf2, 0x5b, 0xf8, 0x0b, 0xfc, 0x24, 0xe4, 0xb1, 0xdd, 0x78, 0xdb, 0x0a, 0xf1, 0x36, 0xdf, 0x64,
	0xec, 0x6f, 0xe6, 0x9b, 0xf1, 0x04, 0xae, 0x74, 0x2d, 0xab, 0x46, 0x2e, 0x74, 0xa1, 0xaa, 0xd9,
	0xb6, 0x56, 0x5a, 0xb1, 0xbe, 0xde, 0x6f, 0xb1, 0x79, 0x3e, 0x59, 0xa8, 0xcd, 0xc6, 0x3b, 0xf9,
	0x9f, 0x11, 0x0c, 0xdf, 0xd4, 0x28, 0x35, 0x66, 0x2d, 0xfb, 0x08, 0x62, 0xad, 0xd2, 0x68, 0x1a,
	0x5d, 0x8f, 0x44, 0xac, 0x15, 0x7b, 0x06, 0x03, 0xb9, 0x51, 0xbb, 0x4a, 0xa7, 0xf1, 0x34, 0xba,
	0x4e, 0x84, 0x43, 0xec, 0x12, 0x92, 0x25, 0x62, 0x9a, 0x90, 0xd3, 0x98, 0x8c, 0x41, 0xaf, 0x52,
	0x1a, 0xd3, 0x1e, 0x9d, 0x25, 0x9b, 0xa5, 0x70, 0x51, 0x34, 0x99, 0x7a, 0xc0, 0x2a, 0xed, 0x4f,
	0xa3, 0xeb, 0xa1, 0xf0, 0x90, 0x4d, 0x61, 0xac, 0x8d, 0x71, 0xb7, 0xdf, 0xcc, 0x55, 0x99, 0x0e,
	0xe8, 0x50, 0xe8, 0xe2, 0x2f, 0x60, 0xf8, 0x4b, 0xd5, 0x14, 0xab, 0x2a, 0x6b, 0xcd, 0xdd, 0xb9,
	0xd4, 0x92, 0xf2, 0x9a, 0x08, 0xb2, 0xf9, 0x6f, 0x30, 0xbc, 0x2b, 0x56, 0x15, 0xe6, 0x59, 0x6b,
	0xb2, 0xdc, 0x51, 0xac, 0x8b, 0x70, 0xc8, 0x9c, 0x23, 0x6f, 0x6c, 0xcf, 0x91, 0xef, 0x19, 0x0c,
	0xb6, 0xbb, 0xf9, 0x03, 0xee, 0x29, 0xf9, 0x89, 0x70, 0x88, 0x2a, 0xdf, 0x53, 0xf6, 0x7d, 0x11,
	0xeb, 0x3d, 0xff, 0x3b, 0x82, 0x71, 0xd6, 0x29, 0x68, 0xce, 0x61, 0x8b, 0x0b, 0xac, 0x3d, 0x87,
	0x45, 0xa6, 0xc6, 0xad, 0xdc, 0x97, 0x4a, 0xe6, 0x8e, 0xc6, 0x43, 0x36, 0x83, 0x91, 0x61, 0x94,
	0x7a, 0x57, 0x5b, 0xa5, 0xc6, 0xaf, 0x2e, 0x67, 0xd4, 0x81, 0xd9, 0x9d, 0xf7, 0x8b, 0x2e, 0xc4,
	0x6b, 0xda, 0xeb, 0x34, 0x25, 0xce, 0x6d, 0x51, 0x23, 0xc9, 0x97, 0x08, 0x87, 0xd8, 0x27, 0xd0,
	0xaf, 0x54, 0xb5, 0x40, 0xd2, 0x2d, 0x11, 0x16, 0xb8, 0xde, 0x5d, 0xf8, 0xde, 0xf1, 0x9f, 0x61,
	0x74, 0xe0, 0x71, 0xe5, 0x45, 0xbe, 0xbc, 0x40, 0x86, 0xf8, 0x89, 0x0c, 0x9f, 0x1d, 0x27, 0x3d,
	0x09, 0x52, 0xe4, 0xbf, 0xc2, 0xe4, 0x87, 0x3c, 0xaf, 0xdf, 0x3f, 0x62, 0xfd, 0x58, 0xe0, 0xef,
	0xa6, 0xf8, 0x1a, 0x17, 0xc5, 0xa3, 0x53, 0x25, 0x11, 0x1e, 0x9a, 0x2f, 0x73, 0x59, 0x4a, 0x93,
	0xa4, 0x9d, 0x1c, 0x0f, 0xcd, 0x17, 0xdd, 0xbe, 0xa1, 0x99, 0xb2, 0xe3, 0xe3, 0x21, 0xff, 0x23,
	0x82, 0x0b, 0x81, 0x1f, 0x0c, 0x83, 0x69, 0x9d, 0xcc, 0xf3, 0xda, 0x8d, 0x22, 0xd9, 0xc6, 0xb7,
	0x2c, 0xe5, 0x8a, 0x2e, 0xec, 0x0b, 0xb2, 0x8d, 0x14, 0x8b, 0xc3, 0x5d, 0x7d, 0x61, 0x81, 0xa9,
	0x22, 0x2f, 0x6a, 0xa4, 0xce, 0xb9, 0x9e, 0x76, 0x0e, 0x53, 0xfb, 0x1a, 0x8b, 0xd5, 0x5a, 0x7b,
	0x59, 0x2d, 0x32, 0x77, 0x15, 0x55, 0x8e, 0xad, 0x97, 0x95, 0x00, 0xff, 0x14, 0xfa, 0x6f, 0xb1,
	0x75, 0x6f, 0xa3, 0x3d, 0xbc, 0x8d, 0x96, 0xbf, 0x87, 0xb1, 0xc0, 0x6d, 0xb9, 0xcf, 0xda, 0x1f,
	0xab, 0xa5, 0x32, 0xd9, 0xad, 0x65, 0xb3, 0xf6, 0x43, 0x6a, 0xec, 0x80, 0x29, 0x3e, 0xcf, 0x94,
	0x84, 0x4c, 0x5f, 0xc0, 0x48, 0xe0, 0x87, 0xac, 0x7d, 0x57, 0x34, 0xba, 0x2b, 0xcc, 0x0a, 0x6b,
	0x01, 0x7f, 0x7d, 0xe0, 0xa4, 0xa0, 0x2f, 0x21, 0xd1, 0x6d, 0x93, 0x46, 0xd3, 0xe4, 0x7a, 0xfc,
	0x8a, 0xb9, 0xe1, 0x0a, 0xa6, 0x56, 0x98, 0xcf, 0xfc, 0x5b, 0x80, 0xac, 0x7d, 0x2b, 0x9b, 0x35,
	0x9d, 0x31, 0x39, 0xc9, 0x66, 0x8d, 0xf6, 0xd8, 0x44, 0x38, 0xd4, 0x11, 0xc6, 0x21, 0xe1, 0x77,
	0x30, 0x09, 0x8a, 0x6c, 0xd8, 0x4b, 0xd3, 0x3d, 0x32, 0x8f, 0x58, 0x83, 0x28, 0xe1, 0x43, 0xf8,
	0x0c, 0x40, 0xe0, 0x02, 0x8b, 0xad, 0x7e, 0xa7, 0x56, 0x27, 0x33, 0x78, 0x09, 0x49, 0xa9, 0x56,
	0x6e, 0x00, 0x8d, 0xc9, 0xa5, 0x19, 0x00, 0x8a, 0x3f, 0x09, 0xfe, 0x1c, 0xe2, 0xdb, 0xfb, 0x34,
	0x26, 0xce, 0x8f, 0x1d, 0xe7, 0x2d, 0xee, 0xef, 0x65, 0xb9, 0x43, 0x11, 0xdf, 0xde, 0xb3, 0xaf,
	0xa0, 0x57, 0xaa, 0x55, 0x93, 0x26, 0x14, 0x72, 0x75, 0x48, 0xcb, 0xd3, 0x0b, 0xfa, 0xcc, 0x6f,
	0x8c, 0x82, 0xe4, 0xbb, 0x91, 0x5a, 0x9e, 0xd0, 0xfc, 0xcf, 0x5b, 0xfe, 0x89, 0x60, 0x98, 0xb5,
	0x02, 0x9b, 0x5d, 0xa9, 0x83, 0x2e, 0x47, 0xe7, 0xbb, 0x6c, 0x07, 0xd6, 0x02, 0xc6, 0x69, 0x8c,
	0xec, 0x3e, 0x38, 0xd7, 0xb2, 0x58, 0xb7, 0xec, 0x1b, 0x18, 0xd7, 0x96, 0x32, 0x97, 0x6e, 0xa7,
	0x86, 0x4a, 0x1f, 0xd2, 0x17, 0x61, 0x98, 0x99, 0xfa, 0x79, 0xa9, 0x16, 0x0f, 0xba, 0xd8, 0xf8,
	0x8d, 0xd1, 0x39, 0xd8, 0x0b, 0x00, 0xcb, 0xf0, 0x93, 0xdc, 0xa0, 0xdb, 0xb8, 0x81, 0x87, 0xff,
	0x15, 0xc3, 0x55, 0x90, 0xc7, 0x0d, 0x6a, 0x59, 0x94, 0x2e, 0xdb, 0xe8, 0x3f, 0xb3, 0x7d, 0x49,
	0x5b, 0xc0, 0xa4, 0x41, 0x95, 0x9e, 0xcf, 0xd4, 0x87, 0xd0, 0xe6, 0xa9, 0x95, 0x5a, 0x5a, 0x8d,
	0xcd, 0xe6, 0x21, 0x14, 0xa8, 0xd8, 0x3b, 0xaf, 0x62, 0x3f, 0x78, 0x2b, 0x4f, 0x6b, 0x1d, 0x1c,
	0xd7, 0xda, 0xfd, 0xb6, 0x2e, 0x9e, 0xfc, 0xb6, 0x9e, 0xc3, 0x70, 0x59, 0xab, 0x0d, 0x6d, 0x96,
	0x21, 0x29, 0x70, 0xc0, 0x47, 0xfa, 0x8c, 0x4e, 0xf4, 0xf9, 0x1e, 0xd8, 0x89, 0x3c, 0x0d, 0xfb,
	0x3a, 0x7c, 0x81, 0xe9, 0xa9, 0x40, 0x36, 0xce, 0xbe, 0xc3, 0x29, 0x0c, 0xdd, 0x7a, 0xa3, 0xd7,
	0x66, 0x58, 0xed, 0xc9, 0x91, 0xb0, 0x60, 0x3e, 0xa0, 0x5f, 0xf2, 0xeb, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x3c, 0xcc, 0x47, 0xfb, 0xbc, 0x07, 0x00, 0x00,
}
