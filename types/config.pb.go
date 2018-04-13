// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Config struct {
	Title      string      `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Log        *Log        `protobuf:"bytes,2,opt,name=log" json:"log,omitempty"`
	Store      *Store      `protobuf:"bytes,3,opt,name=store" json:"store,omitempty"`
	LocalStore *LocalStore `protobuf:"bytes,4,opt,name=localStore" json:"localStore,omitempty"`
	Consensus  *Consensus  `protobuf:"bytes,5,opt,name=consensus" json:"consensus,omitempty"`
	MemPool    *MemPool    `protobuf:"bytes,6,opt,name=memPool" json:"memPool,omitempty"`
	BlockChain *BlockChain `protobuf:"bytes,7,opt,name=blockChain" json:"blockChain,omitempty"`
	Wallet     *Wallet     `protobuf:"bytes,8,opt,name=wallet" json:"wallet,omitempty"`
	P2P        *P2P        `protobuf:"bytes,9,opt,name=p2p" json:"p2p,omitempty"`
	Rpc        *Rpc        `protobuf:"bytes,10,opt,name=rpc" json:"rpc,omitempty"`
	Exec       *Exec       `protobuf:"bytes,11,opt,name=exec" json:"exec,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Config) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Config) GetLog() *Log {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *Config) GetStore() *Store {
	if m != nil {
		return m.Store
	}
	return nil
}

func (m *Config) GetLocalStore() *LocalStore {
	if m != nil {
		return m.LocalStore
	}
	return nil
}

func (m *Config) GetConsensus() *Consensus {
	if m != nil {
		return m.Consensus
	}
	return nil
}

func (m *Config) GetMemPool() *MemPool {
	if m != nil {
		return m.MemPool
	}
	return nil
}

func (m *Config) GetBlockChain() *BlockChain {
	if m != nil {
		return m.BlockChain
	}
	return nil
}

func (m *Config) GetWallet() *Wallet {
	if m != nil {
		return m.Wallet
	}
	return nil
}

func (m *Config) GetP2P() *P2P {
	if m != nil {
		return m.P2P
	}
	return nil
}

func (m *Config) GetRpc() *Rpc {
	if m != nil {
		return m.Rpc
	}
	return nil
}

func (m *Config) GetExec() *Exec {
	if m != nil {
		return m.Exec
	}
	return nil
}

type Log struct {
	// 日志级别，支持debug(dbug)/info/warn/error(eror)/crit
	Loglevel        string `protobuf:"bytes,1,opt,name=loglevel" json:"loglevel,omitempty"`
	LogConsoleLevel string `protobuf:"bytes,2,opt,name=logConsoleLevel" json:"logConsoleLevel,omitempty"`
	// 日志文件名，可带目录，所有生成的日志文件都放到此目录下
	LogFile string `protobuf:"bytes,3,opt,name=logFile" json:"logFile,omitempty"`
	// 单个日志文件的最大值（单位：兆）
	MaxFileSize uint32 `protobuf:"varint,4,opt,name=maxFileSize" json:"maxFileSize,omitempty"`
	// 最多保存的历史日志文件个数
	MaxBackups uint32 `protobuf:"varint,5,opt,name=maxBackups" json:"maxBackups,omitempty"`
	// 最多保存的历史日志消息（单位：天）
	MaxAge uint32 `protobuf:"varint,6,opt,name=maxAge" json:"maxAge,omitempty"`
	// 日志文件名是否使用本地事件（否则使用UTC时间）
	LocalTime bool `protobuf:"varint,7,opt,name=localTime" json:"localTime,omitempty"`
	// 历史日志文件是否压缩（压缩格式为gz）
	Compress bool `protobuf:"varint,8,opt,name=compress" json:"compress,omitempty"`
	// 是否打印调用源文件和行号
	CallerFile bool `protobuf:"varint,9,opt,name=callerFile" json:"callerFile,omitempty"`
	// 是否打印调用方法
	CallerFunction bool `protobuf:"varint,10,opt,name=callerFunction" json:"callerFunction,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *Log) GetLoglevel() string {
	if m != nil {
		return m.Loglevel
	}
	return ""
}

func (m *Log) GetLogConsoleLevel() string {
	if m != nil {
		return m.LogConsoleLevel
	}
	return ""
}

func (m *Log) GetLogFile() string {
	if m != nil {
		return m.LogFile
	}
	return ""
}

func (m *Log) GetMaxFileSize() uint32 {
	if m != nil {
		return m.MaxFileSize
	}
	return 0
}

func (m *Log) GetMaxBackups() uint32 {
	if m != nil {
		return m.MaxBackups
	}
	return 0
}

func (m *Log) GetMaxAge() uint32 {
	if m != nil {
		return m.MaxAge
	}
	return 0
}

func (m *Log) GetLocalTime() bool {
	if m != nil {
		return m.LocalTime
	}
	return false
}

func (m *Log) GetCompress() bool {
	if m != nil {
		return m.Compress
	}
	return false
}

func (m *Log) GetCallerFile() bool {
	if m != nil {
		return m.CallerFile
	}
	return false
}

func (m *Log) GetCallerFunction() bool {
	if m != nil {
		return m.CallerFunction
	}
	return false
}

type MemPool struct {
	PoolCacheSize int64 `protobuf:"varint,1,opt,name=poolCacheSize" json:"poolCacheSize,omitempty"`
	MinTxFee      int64 `protobuf:"varint,2,opt,name=minTxFee" json:"minTxFee,omitempty"`
}

func (m *MemPool) Reset()                    { *m = MemPool{} }
func (m *MemPool) String() string            { return proto.CompactTextString(m) }
func (*MemPool) ProtoMessage()               {}
func (*MemPool) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *MemPool) GetPoolCacheSize() int64 {
	if m != nil {
		return m.PoolCacheSize
	}
	return 0
}

func (m *MemPool) GetMinTxFee() int64 {
	if m != nil {
		return m.MinTxFee
	}
	return 0
}

type Consensus struct {
	Name             string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Genesis          string `protobuf:"bytes,2,opt,name=genesis" json:"genesis,omitempty"`
	Minerstart       bool   `protobuf:"varint,3,opt,name=minerstart" json:"minerstart,omitempty"`
	GenesisBlockTime int64  `protobuf:"varint,4,opt,name=genesisBlockTime" json:"genesisBlockTime,omitempty"`
	HotkeyAddr       string `protobuf:"bytes,5,opt,name=hotkeyAddr" json:"hotkeyAddr,omitempty"`
	NodeId           int32  `protobuf:"varint,6,opt,name=NodeId" json:"NodeId,omitempty"`
	PeersURL         string `protobuf:"bytes,7,opt,name=PeersURL" json:"PeersURL,omitempty"`
	ClientAddr       string `protobuf:"bytes,8,opt,name=ClientAddr" json:"ClientAddr,omitempty"`
}

func (m *Consensus) Reset()                    { *m = Consensus{} }
func (m *Consensus) String() string            { return proto.CompactTextString(m) }
func (*Consensus) ProtoMessage()               {}
func (*Consensus) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *Consensus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Consensus) GetGenesis() string {
	if m != nil {
		return m.Genesis
	}
	return ""
}

func (m *Consensus) GetMinerstart() bool {
	if m != nil {
		return m.Minerstart
	}
	return false
}

func (m *Consensus) GetGenesisBlockTime() int64 {
	if m != nil {
		return m.GenesisBlockTime
	}
	return 0
}

func (m *Consensus) GetHotkeyAddr() string {
	if m != nil {
		return m.HotkeyAddr
	}
	return ""
}

func (m *Consensus) GetNodeId() int32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *Consensus) GetPeersURL() string {
	if m != nil {
		return m.PeersURL
	}
	return ""
}

func (m *Consensus) GetClientAddr() string {
	if m != nil {
		return m.ClientAddr
	}
	return ""
}

type Wallet struct {
	MinFee   int64  `protobuf:"varint,1,opt,name=minFee" json:"minFee,omitempty"`
	Driver   string `protobuf:"bytes,2,opt,name=driver" json:"driver,omitempty"`
	DbPath   string `protobuf:"bytes,3,opt,name=dbPath" json:"dbPath,omitempty"`
	SignType string `protobuf:"bytes,4,opt,name=signType" json:"signType,omitempty"`
}

func (m *Wallet) Reset()                    { *m = Wallet{} }
func (m *Wallet) String() string            { return proto.CompactTextString(m) }
func (*Wallet) ProtoMessage()               {}
func (*Wallet) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *Wallet) GetMinFee() int64 {
	if m != nil {
		return m.MinFee
	}
	return 0
}

func (m *Wallet) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *Wallet) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

func (m *Wallet) GetSignType() string {
	if m != nil {
		return m.SignType
	}
	return ""
}

type Store struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Driver string `protobuf:"bytes,2,opt,name=driver" json:"driver,omitempty"`
	DbPath string `protobuf:"bytes,3,opt,name=dbPath" json:"dbPath,omitempty"`
}

func (m *Store) Reset()                    { *m = Store{} }
func (m *Store) String() string            { return proto.CompactTextString(m) }
func (*Store) ProtoMessage()               {}
func (*Store) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *Store) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Store) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *Store) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

type LocalStore struct {
	Driver string `protobuf:"bytes,1,opt,name=driver" json:"driver,omitempty"`
	DbPath string `protobuf:"bytes,2,opt,name=dbPath" json:"dbPath,omitempty"`
}

func (m *LocalStore) Reset()                    { *m = LocalStore{} }
func (m *LocalStore) String() string            { return proto.CompactTextString(m) }
func (*LocalStore) ProtoMessage()               {}
func (*LocalStore) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *LocalStore) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *LocalStore) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

type BlockChain struct {
	DefCacheSize        int64  `protobuf:"varint,1,opt,name=defCacheSize" json:"defCacheSize,omitempty"`
	MaxFetchBlockNum    int64  `protobuf:"varint,2,opt,name=maxFetchBlockNum" json:"maxFetchBlockNum,omitempty"`
	TimeoutSeconds      int64  `protobuf:"varint,3,opt,name=timeoutSeconds" json:"timeoutSeconds,omitempty"`
	BatchBlockNum       int64  `protobuf:"varint,4,opt,name=batchBlockNum" json:"batchBlockNum,omitempty"`
	Driver              string `protobuf:"bytes,5,opt,name=driver" json:"driver,omitempty"`
	DbPath              string `protobuf:"bytes,6,opt,name=dbPath" json:"dbPath,omitempty"`
	IsStrongConsistency bool   `protobuf:"varint,7,opt,name=isStrongConsistency" json:"isStrongConsistency,omitempty"`
	SingleMode          bool   `protobuf:"varint,8,opt,name=singleMode" json:"singleMode,omitempty"`
	Batchsync           bool   `protobuf:"varint,9,opt,name=batchsync" json:"batchsync,omitempty"`
}

func (m *BlockChain) Reset()                    { *m = BlockChain{} }
func (m *BlockChain) String() string            { return proto.CompactTextString(m) }
func (*BlockChain) ProtoMessage()               {}
func (*BlockChain) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *BlockChain) GetDefCacheSize() int64 {
	if m != nil {
		return m.DefCacheSize
	}
	return 0
}

func (m *BlockChain) GetMaxFetchBlockNum() int64 {
	if m != nil {
		return m.MaxFetchBlockNum
	}
	return 0
}

func (m *BlockChain) GetTimeoutSeconds() int64 {
	if m != nil {
		return m.TimeoutSeconds
	}
	return 0
}

func (m *BlockChain) GetBatchBlockNum() int64 {
	if m != nil {
		return m.BatchBlockNum
	}
	return 0
}

func (m *BlockChain) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *BlockChain) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

func (m *BlockChain) GetIsStrongConsistency() bool {
	if m != nil {
		return m.IsStrongConsistency
	}
	return false
}

func (m *BlockChain) GetSingleMode() bool {
	if m != nil {
		return m.SingleMode
	}
	return false
}

func (m *BlockChain) GetBatchsync() bool {
	if m != nil {
		return m.Batchsync
	}
	return false
}

type P2P struct {
	SeedPort     int32    `protobuf:"varint,1,opt,name=seedPort" json:"seedPort,omitempty"`
	DbPath       string   `protobuf:"bytes,2,opt,name=dbPath" json:"dbPath,omitempty"`
	GrpcLogFile  string   `protobuf:"bytes,3,opt,name=grpcLogFile" json:"grpcLogFile,omitempty"`
	IsSeed       bool     `protobuf:"varint,4,opt,name=isSeed" json:"isSeed,omitempty"`
	ServerStart  bool     `protobuf:"varint,5,opt,name=serverStart" json:"serverStart,omitempty"`
	Seeds        []string `protobuf:"bytes,6,rep,name=seeds" json:"seeds,omitempty"`
	Enable       bool     `protobuf:"varint,7,opt,name=enable" json:"enable,omitempty"`
	MsgCacheSize int32    `protobuf:"varint,8,opt,name=msgCacheSize" json:"msgCacheSize,omitempty"`
	Version      int32    `protobuf:"varint,9,opt,name=version" json:"version,omitempty"`
	VerMix       int32    `protobuf:"varint,10,opt,name=verMix" json:"verMix,omitempty"`
	VerMax       int32    `protobuf:"varint,11,opt,name=verMax" json:"verMax,omitempty"`
}

func (m *P2P) Reset()                    { *m = P2P{} }
func (m *P2P) String() string            { return proto.CompactTextString(m) }
func (*P2P) ProtoMessage()               {}
func (*P2P) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *P2P) GetSeedPort() int32 {
	if m != nil {
		return m.SeedPort
	}
	return 0
}

func (m *P2P) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

func (m *P2P) GetGrpcLogFile() string {
	if m != nil {
		return m.GrpcLogFile
	}
	return ""
}

func (m *P2P) GetIsSeed() bool {
	if m != nil {
		return m.IsSeed
	}
	return false
}

func (m *P2P) GetServerStart() bool {
	if m != nil {
		return m.ServerStart
	}
	return false
}

func (m *P2P) GetSeeds() []string {
	if m != nil {
		return m.Seeds
	}
	return nil
}

func (m *P2P) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *P2P) GetMsgCacheSize() int32 {
	if m != nil {
		return m.MsgCacheSize
	}
	return 0
}

func (m *P2P) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *P2P) GetVerMix() int32 {
	if m != nil {
		return m.VerMix
	}
	return 0
}

func (m *P2P) GetVerMax() int32 {
	if m != nil {
		return m.VerMax
	}
	return 0
}

type Rpc struct {
	JrpcBindAddr string   `protobuf:"bytes,1,opt,name=jrpcBindAddr" json:"jrpcBindAddr,omitempty"`
	GrpcBindAddr string   `protobuf:"bytes,2,opt,name=grpcBindAddr" json:"grpcBindAddr,omitempty"`
	Whitlist     []string `protobuf:"bytes,3,rep,name=whitlist" json:"whitlist,omitempty"`
}

func (m *Rpc) Reset()                    { *m = Rpc{} }
func (m *Rpc) String() string            { return proto.CompactTextString(m) }
func (*Rpc) ProtoMessage()               {}
func (*Rpc) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *Rpc) GetJrpcBindAddr() string {
	if m != nil {
		return m.JrpcBindAddr
	}
	return ""
}

func (m *Rpc) GetGrpcBindAddr() string {
	if m != nil {
		return m.GrpcBindAddr
	}
	return ""
}

func (m *Rpc) GetWhitlist() []string {
	if m != nil {
		return m.Whitlist
	}
	return nil
}

type Exec struct {
	MinExecFee int64 `protobuf:"varint,1,opt,name=minExecFee" json:"minExecFee,omitempty"`
	IsFree     bool  `protobuf:"varint,2,opt,name=isFree" json:"isFree,omitempty"`
}

func (m *Exec) Reset()                    { *m = Exec{} }
func (m *Exec) String() string            { return proto.CompactTextString(m) }
func (*Exec) ProtoMessage()               {}
func (*Exec) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{10} }

func (m *Exec) GetMinExecFee() int64 {
	if m != nil {
		return m.MinExecFee
	}
	return 0
}

func (m *Exec) GetIsFree() bool {
	if m != nil {
		return m.IsFree
	}
	return false
}

func init() {
	proto.RegisterType((*Config)(nil), "types.Config")
	proto.RegisterType((*Log)(nil), "types.Log")
	proto.RegisterType((*MemPool)(nil), "types.MemPool")
	proto.RegisterType((*Consensus)(nil), "types.Consensus")
	proto.RegisterType((*Wallet)(nil), "types.Wallet")
	proto.RegisterType((*Store)(nil), "types.Store")
	proto.RegisterType((*LocalStore)(nil), "types.LocalStore")
	proto.RegisterType((*BlockChain)(nil), "types.BlockChain")
	proto.RegisterType((*P2P)(nil), "types.P2P")
	proto.RegisterType((*Rpc)(nil), "types.Rpc")
	proto.RegisterType((*Exec)(nil), "types.Exec")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 933 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0x5f, 0x6f, 0xe3, 0x44,
	0x10, 0x57, 0xe3, 0x3a, 0x8d, 0xb7, 0x97, 0xe3, 0x58, 0x10, 0xb2, 0xd0, 0x09, 0x2a, 0x0b, 0x50,
	0xc5, 0x43, 0x05, 0xe5, 0x15, 0x21, 0x5d, 0x23, 0x2a, 0xa1, 0x4b, 0x4f, 0xd1, 0xa6, 0x88, 0x67,
	0x67, 0x3d, 0xe7, 0x2c, 0x5d, 0xef, 0x5a, 0xbb, 0x4e, 0x2e, 0xe1, 0xf3, 0xf0, 0x05, 0xf8, 0x08,
	0xbc, 0xf2, 0x8d, 0x78, 0x43, 0x33, 0xde, 0x38, 0x76, 0xef, 0xfa, 0xc0, 0x5b, 0x7e, 0xbf, 0x99,
	0xd9, 0xdd, 0xf9, 0xcd, 0x1f, 0x87, 0x3d, 0x93, 0xd6, 0xbc, 0x55, 0xe5, 0x55, 0xed, 0x6c, 0x63,
	0x79, 0xdc, 0xec, 0x6b, 0xf0, 0xd9, 0x9f, 0x11, 0x1b, 0xcf, 0x88, 0xe7, 0x9f, 0xb2, 0xb8, 0x51,
	0x8d, 0x86, 0xf4, 0xe4, 0xe2, 0xe4, 0x32, 0x11, 0x2d, 0xe0, 0x2f, 0x59, 0xa4, 0x6d, 0x99, 0x8e,
	0x2e, 0x4e, 0x2e, 0xcf, 0xaf, 0xd9, 0x15, 0x45, 0x5d, 0xcd, 0x6d, 0x29, 0x90, 0xe6, 0x19, 0x8b,
	0x7d, 0x63, 0x1d, 0xa4, 0x11, 0xd9, 0x9f, 0x05, 0xfb, 0x12, 0x39, 0xd1, 0x9a, 0xf8, 0xf7, 0x8c,
	0x69, 0x2b, 0x73, 0x4d, 0x64, 0x7a, 0x4a, 0x8e, 0x1f, 0x77, 0x07, 0x1d, 0x0c, 0xa2, 0xe7, 0xc4,
	0xaf, 0x58, 0x22, 0xad, 0xf1, 0x60, 0xfc, 0xc6, 0xa7, 0x31, 0x45, 0xbc, 0x08, 0x11, 0xb3, 0x03,
	0x2f, 0x8e, 0x2e, 0xfc, 0x92, 0x9d, 0x55, 0x50, 0x2d, 0xac, 0xd5, 0xe9, 0x98, 0xbc, 0x9f, 0x07,
	0xef, 0xbb, 0x96, 0x15, 0x07, 0x33, 0x3e, 0x66, 0xa5, 0xad, 0x7c, 0x98, 0xad, 0x73, 0x65, 0xd2,
	0xb3, 0xc1, 0x63, 0x6e, 0x3a, 0x83, 0xe8, 0x39, 0xf1, 0xaf, 0xd9, 0xf8, 0x5d, 0xae, 0x35, 0x34,
	0xe9, 0x84, 0xdc, 0xa7, 0xc1, 0xfd, 0x37, 0x22, 0x45, 0x30, 0xa2, 0x50, 0xf5, 0x75, 0x9d, 0x26,
	0x03, 0xa1, 0x16, 0xd7, 0x0b, 0x81, 0x34, 0x5a, 0x5d, 0x2d, 0x53, 0x36, 0xb0, 0x8a, 0x5a, 0x0a,
	0xa4, 0xf9, 0x97, 0xec, 0x14, 0x76, 0x20, 0xd3, 0x73, 0x32, 0x9f, 0x07, 0xf3, 0xcf, 0x3b, 0x90,
	0x82, 0x0c, 0xd9, 0xdf, 0x23, 0x16, 0xcd, 0x6d, 0xc9, 0x3f, 0x67, 0x13, 0x6d, 0x4b, 0x0d, 0x5b,
	0xd0, 0xa1, 0x4c, 0x1d, 0xe6, 0x97, 0xec, 0x23, 0x6d, 0x4b, 0xd4, 0xc7, 0x6a, 0x98, 0x93, 0xcb,
	0x88, 0x5c, 0x1e, 0xd3, 0x3c, 0x65, 0x67, 0xda, 0x96, 0xb7, 0x4a, 0xb7, 0x75, 0x4b, 0xc4, 0x01,
	0xf2, 0x0b, 0x76, 0x5e, 0xe5, 0x3b, 0xfc, 0xb9, 0x54, 0x7f, 0xb4, 0xc5, 0x9a, 0x8a, 0x3e, 0xc5,
	0xbf, 0x60, 0xac, 0xca, 0x77, 0x37, 0xb9, 0x7c, 0xd8, 0xd4, 0x6d, 0x6d, 0xa6, 0xa2, 0xc7, 0xf0,
	0xcf, 0xd8, 0xb8, 0xca, 0x77, 0xaf, 0x4a, 0xa0, 0x4a, 0x4c, 0x45, 0x40, 0xfc, 0x25, 0x4b, 0xa8,
	0xc0, 0xf7, 0xaa, 0x02, 0xd2, 0x7d, 0x22, 0x8e, 0x04, 0xe6, 0x25, 0x6d, 0x55, 0x3b, 0xf0, 0x9e,
	0x54, 0x9e, 0x88, 0x0e, 0xe3, 0x8d, 0x12, 0x25, 0x76, 0xf4, 0xe0, 0x84, 0xac, 0x3d, 0x86, 0x7f,
	0xc3, 0x9e, 0x07, 0xb4, 0x31, 0xb2, 0x51, 0xd6, 0x90, 0xca, 0x13, 0xf1, 0x88, 0xcd, 0x5e, 0xb3,
	0xb3, 0xd0, 0x0e, 0xfc, 0x2b, 0x36, 0xad, 0xad, 0xd5, 0xb3, 0x5c, 0xae, 0xdb, 0x44, 0x51, 0xcb,
	0x48, 0x0c, 0x49, 0x7c, 0x54, 0xa5, 0xcc, 0xfd, 0xee, 0x16, 0x80, 0x94, 0x8c, 0x44, 0x87, 0xb3,
	0x7f, 0x4f, 0x58, 0xd2, 0xb5, 0x22, 0xe7, 0xec, 0xd4, 0xe4, 0xd5, 0x61, 0x72, 0xe8, 0x37, 0x8a,
	0x5c, 0x82, 0x01, 0xaf, 0x7c, 0x28, 0xc3, 0x01, 0x92, 0x84, 0xca, 0x80, 0xf3, 0x4d, 0xee, 0x1a,
	0xaa, 0xc0, 0x44, 0xf4, 0x18, 0xfe, 0x2d, 0x7b, 0x11, 0x5c, 0xa9, 0x23, 0x49, 0xb1, 0x53, 0xba,
	0xff, 0x3d, 0x1e, 0xcf, 0x5a, 0xdb, 0xe6, 0x01, 0xf6, 0xaf, 0x8a, 0xc2, 0x51, 0x39, 0x12, 0xd1,
	0x63, 0xb0, 0x1c, 0x6f, 0x6c, 0x01, 0xbf, 0x14, 0x54, 0x8e, 0x58, 0x04, 0x84, 0xb9, 0x2d, 0x00,
	0x9c, 0xff, 0x55, 0xcc, 0xa9, 0x1a, 0x89, 0xe8, 0x30, 0x9e, 0x39, 0xd3, 0x0a, 0x4c, 0x43, 0x67,
	0x4e, 0xda, 0x33, 0x8f, 0x4c, 0xa6, 0xd9, 0xb8, 0xed, 0x7d, 0x2a, 0xb6, 0x32, 0xa8, 0x4f, 0x2b,
	0x60, 0x40, 0xc8, 0x17, 0x4e, 0x6d, 0xc1, 0x85, 0xd4, 0x03, 0x22, 0x7e, 0xb5, 0xc8, 0x9b, 0x75,
	0xe8, 0xbb, 0x80, 0xf0, 0x35, 0x5e, 0x95, 0xe6, 0x7e, 0x5f, 0xb7, 0x99, 0x26, 0xa2, 0xc3, 0xd9,
	0x6b, 0x16, 0xb7, 0x4b, 0xe1, 0x43, 0x22, 0xff, 0xcf, 0x8b, 0xb2, 0x1f, 0x19, 0x3b, 0xae, 0x9c,
	0x5e, 0xf4, 0xc9, 0x13, 0xd1, 0xa3, 0x41, 0xf4, 0x3f, 0x23, 0xc6, 0x8e, 0x4b, 0x82, 0x67, 0xec,
	0x59, 0x01, 0x6f, 0x1f, 0x37, 0xd1, 0x80, 0xc3, 0x5a, 0xe2, 0xf4, 0x40, 0x23, 0xd7, 0x14, 0xf9,
	0x66, 0x53, 0x85, 0x5e, 0x7a, 0x8f, 0xc7, 0x46, 0x6e, 0x54, 0x05, 0x76, 0xd3, 0x2c, 0x41, 0x5a,
	0x53, 0x78, 0x7a, 0x7c, 0x24, 0x1e, 0xb1, 0xd8, 0xbd, 0xab, 0xbc, 0x7f, 0x60, 0xdb, 0x1c, 0x43,
	0xb2, 0x97, 0x5c, 0xfc, 0x44, 0x72, 0xe3, 0x41, 0x0d, 0xbe, 0x63, 0x9f, 0x28, 0xbf, 0x6c, 0x9c,
	0x35, 0xb4, 0x2c, 0x94, 0x6f, 0xc0, 0xc8, 0x7d, 0x18, 0xd5, 0x0f, 0x99, 0xb0, 0x4f, 0xbc, 0x32,
	0xa5, 0x86, 0x3b, 0x5b, 0x40, 0x18, 0xdb, 0x1e, 0x83, 0x23, 0x4f, 0x4f, 0xf2, 0x7b, 0x23, 0xc3,
	0xdc, 0x1e, 0x89, 0xec, 0xaf, 0x11, 0x8b, 0x16, 0xd7, 0x0b, 0xaa, 0x3d, 0x40, 0xb1, 0xb0, 0xae,
	0x21, 0x05, 0x63, 0xd1, 0xe1, 0xa7, 0x0a, 0x81, 0x6b, 0xaa, 0x74, 0xb5, 0x9c, 0x0f, 0x96, 0x58,
	0x9f, 0xc2, 0x48, 0xe5, 0x97, 0x00, 0x05, 0x89, 0x33, 0x11, 0x01, 0x61, 0xa4, 0x07, 0xb7, 0x05,
	0xb7, 0xa4, 0xe1, 0x8b, 0xc9, 0xd8, 0xa7, 0xf0, 0x33, 0x88, 0xf7, 0xfb, 0x74, 0x7c, 0x11, 0xe1,
	0x67, 0x90, 0x00, 0x9e, 0x07, 0x26, 0x5f, 0xe9, 0xc3, 0xee, 0x0a, 0x08, 0x7b, 0xa0, 0xf2, 0xe5,
	0xb1, 0x07, 0x26, 0x94, 0xc1, 0x80, 0xc3, 0x4d, 0xb0, 0x05, 0xe7, 0x71, 0x33, 0x25, 0x64, 0x3e,
	0x40, 0x3c, 0x75, 0x0b, 0xee, 0x4e, 0xed, 0x68, 0x65, 0xc5, 0x22, 0xa0, 0x03, 0x9f, 0xef, 0xe8,
	0x8b, 0x10, 0xf8, 0x7c, 0x97, 0x29, 0x16, 0x89, 0x5a, 0xe2, 0xa5, 0xbf, 0xbb, 0x5a, 0xde, 0x28,
	0x53, 0xd0, 0x88, 0xb6, 0xdd, 0x3b, 0xe0, 0xd0, 0xa7, 0xec, 0xfb, 0xb4, 0x02, 0x0e, 0x38, 0x94,
	0xfe, 0xdd, 0x5a, 0x35, 0x5a, 0x79, 0x5c, 0x43, 0x98, 0x6d, 0x87, 0xb3, 0x9f, 0xd8, 0x29, 0x7e,
	0x7f, 0xc2, 0xb2, 0xc2, 0x9f, 0xc7, 0x31, 0xef, 0x31, 0xad, 0xd0, 0xb7, 0x2e, 0xac, 0x48, 0x12,
	0x1a, 0xd1, 0x6a, 0x4c, 0x7f, 0x33, 0x7e, 0xf8, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xd5, 0x94,
	0x6f, 0x76, 0x08, 0x00, 0x00,
}
