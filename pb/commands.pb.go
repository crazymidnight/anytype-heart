// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commands.proto

package pb

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

type WalletCreateR_Error_Code int32

const (
	WalletCreateR_Error_NULL                        WalletCreateR_Error_Code = 0
	WalletCreateR_Error_UNKNOWN_ERROR               WalletCreateR_Error_Code = 1
	WalletCreateR_Error_BAD_INPUT                   WalletCreateR_Error_Code = 2
	WalletCreateR_Error_FAILED_TO_CREATE_LOCAL_REPO WalletCreateR_Error_Code = 101
)

var WalletCreateR_Error_Code_name = map[int32]string{
	0:   "NULL",
	1:   "UNKNOWN_ERROR",
	2:   "BAD_INPUT",
	101: "FAILED_TO_CREATE_LOCAL_REPO",
}

var WalletCreateR_Error_Code_value = map[string]int32{
	"NULL":                        0,
	"UNKNOWN_ERROR":               1,
	"BAD_INPUT":                   2,
	"FAILED_TO_CREATE_LOCAL_REPO": 101,
}

func (x WalletCreateR_Error_Code) String() string {
	return proto.EnumName(WalletCreateR_Error_Code_name, int32(x))
}

func (WalletCreateR_Error_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{1, 0, 0}
}

type WalletRecoverR_Error_Code int32

const (
	WalletRecoverR_Error_NULL          WalletRecoverR_Error_Code = 0
	WalletRecoverR_Error_UNKNOWN_ERROR WalletRecoverR_Error_Code = 1
	WalletRecoverR_Error_BAD_INPUT     WalletRecoverR_Error_Code = 2
)

var WalletRecoverR_Error_Code_name = map[int32]string{
	0: "NULL",
	1: "UNKNOWN_ERROR",
	2: "BAD_INPUT",
}

var WalletRecoverR_Error_Code_value = map[string]int32{
	"NULL":          0,
	"UNKNOWN_ERROR": 1,
	"BAD_INPUT":     2,
}

func (x WalletRecoverR_Error_Code) String() string {
	return proto.EnumName(WalletRecoverR_Error_Code_name, int32(x))
}

func (WalletRecoverR_Error_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{3, 0, 0}
}

type AccountCreateR_Error_Code int32

const (
	AccountCreateR_Error_NULL                 AccountCreateR_Error_Code = 0
	AccountCreateR_Error_UNKNOWN_ERROR        AccountCreateR_Error_Code = 1
	AccountCreateR_Error_BAD_INPUT            AccountCreateR_Error_Code = 2
	AccountCreateR_Error_FAILED_TO_START_NODE AccountCreateR_Error_Code = 101
	AccountCreateR_Error_FAILED_TO_SET_NAME   AccountCreateR_Error_Code = 102
	AccountCreateR_Error_FAILED_TO_SET_AVATAR AccountCreateR_Error_Code = 103
)

var AccountCreateR_Error_Code_name = map[int32]string{
	0:   "NULL",
	1:   "UNKNOWN_ERROR",
	2:   "BAD_INPUT",
	101: "FAILED_TO_START_NODE",
	102: "FAILED_TO_SET_NAME",
	103: "FAILED_TO_SET_AVATAR",
}

var AccountCreateR_Error_Code_value = map[string]int32{
	"NULL":                 0,
	"UNKNOWN_ERROR":        1,
	"BAD_INPUT":            2,
	"FAILED_TO_START_NODE": 101,
	"FAILED_TO_SET_NAME":   102,
	"FAILED_TO_SET_AVATAR": 103,
}

func (x AccountCreateR_Error_Code) String() string {
	return proto.EnumName(AccountCreateR_Error_Code_name, int32(x))
}

func (AccountCreateR_Error_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{5, 0, 0}
}

type AccountSelectR_Error_Code int32

const (
	AccountSelectR_Error_NULL                            AccountSelectR_Error_Code = 0
	AccountSelectR_Error_UNKNOWN_ERROR                   AccountSelectR_Error_Code = 1
	AccountSelectR_Error_BAD_INPUT                       AccountSelectR_Error_Code = 2
	AccountSelectR_Error_FAILED_TO_CREATE_LOCAL_REPO     AccountSelectR_Error_Code = 101
	AccountSelectR_Error_LOCAL_REPO_EXISTS_BUT_CORRUPTED AccountSelectR_Error_Code = 102
	AccountSelectR_Error_FAILED_TO_RUN_NODE              AccountSelectR_Error_Code = 103
	AccountSelectR_Error_FAILED_TO_FIND_ACCOUNT_INFO     AccountSelectR_Error_Code = 104
)

var AccountSelectR_Error_Code_name = map[int32]string{
	0:   "NULL",
	1:   "UNKNOWN_ERROR",
	2:   "BAD_INPUT",
	101: "FAILED_TO_CREATE_LOCAL_REPO",
	102: "LOCAL_REPO_EXISTS_BUT_CORRUPTED",
	103: "FAILED_TO_RUN_NODE",
	104: "FAILED_TO_FIND_ACCOUNT_INFO",
}

var AccountSelectR_Error_Code_value = map[string]int32{
	"NULL":                            0,
	"UNKNOWN_ERROR":                   1,
	"BAD_INPUT":                       2,
	"FAILED_TO_CREATE_LOCAL_REPO":     101,
	"LOCAL_REPO_EXISTS_BUT_CORRUPTED": 102,
	"FAILED_TO_RUN_NODE":              103,
	"FAILED_TO_FIND_ACCOUNT_INFO":     104,
}

func (x AccountSelectR_Error_Code) String() string {
	return proto.EnumName(AccountSelectR_Error_Code_name, int32(x))
}

func (AccountSelectR_Error_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{7, 0, 0}
}

type WalletCreateQ struct {
	RootPath             string   `protobuf:"bytes,1,opt,name=rootPath,proto3" json:"rootPath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WalletCreateQ) Reset()         { *m = WalletCreateQ{} }
func (m *WalletCreateQ) String() string { return proto.CompactTextString(m) }
func (*WalletCreateQ) ProtoMessage()    {}
func (*WalletCreateQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{0}
}

func (m *WalletCreateQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletCreateQ.Unmarshal(m, b)
}
func (m *WalletCreateQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletCreateQ.Marshal(b, m, deterministic)
}
func (m *WalletCreateQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletCreateQ.Merge(m, src)
}
func (m *WalletCreateQ) XXX_Size() int {
	return xxx_messageInfo_WalletCreateQ.Size(m)
}
func (m *WalletCreateQ) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletCreateQ.DiscardUnknown(m)
}

var xxx_messageInfo_WalletCreateQ proto.InternalMessageInfo

func (m *WalletCreateQ) GetRootPath() string {
	if m != nil {
		return m.RootPath
	}
	return ""
}

type WalletCreateR struct {
	Error                *WalletCreateR_Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Mnemonic             string               `protobuf:"bytes,2,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *WalletCreateR) Reset()         { *m = WalletCreateR{} }
func (m *WalletCreateR) String() string { return proto.CompactTextString(m) }
func (*WalletCreateR) ProtoMessage()    {}
func (*WalletCreateR) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{1}
}

func (m *WalletCreateR) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletCreateR.Unmarshal(m, b)
}
func (m *WalletCreateR) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletCreateR.Marshal(b, m, deterministic)
}
func (m *WalletCreateR) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletCreateR.Merge(m, src)
}
func (m *WalletCreateR) XXX_Size() int {
	return xxx_messageInfo_WalletCreateR.Size(m)
}
func (m *WalletCreateR) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletCreateR.DiscardUnknown(m)
}

var xxx_messageInfo_WalletCreateR proto.InternalMessageInfo

func (m *WalletCreateR) GetError() *WalletCreateR_Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *WalletCreateR) GetMnemonic() string {
	if m != nil {
		return m.Mnemonic
	}
	return ""
}

type WalletCreateR_Error struct {
	Code                 WalletCreateR_Error_Code `protobuf:"varint,1,opt,name=code,proto3,enum=anytype.WalletCreateR_Error_Code" json:"code,omitempty"`
	Desc                 string                   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *WalletCreateR_Error) Reset()         { *m = WalletCreateR_Error{} }
func (m *WalletCreateR_Error) String() string { return proto.CompactTextString(m) }
func (*WalletCreateR_Error) ProtoMessage()    {}
func (*WalletCreateR_Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{1, 0}
}

func (m *WalletCreateR_Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletCreateR_Error.Unmarshal(m, b)
}
func (m *WalletCreateR_Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletCreateR_Error.Marshal(b, m, deterministic)
}
func (m *WalletCreateR_Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletCreateR_Error.Merge(m, src)
}
func (m *WalletCreateR_Error) XXX_Size() int {
	return xxx_messageInfo_WalletCreateR_Error.Size(m)
}
func (m *WalletCreateR_Error) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletCreateR_Error.DiscardUnknown(m)
}

var xxx_messageInfo_WalletCreateR_Error proto.InternalMessageInfo

func (m *WalletCreateR_Error) GetCode() WalletCreateR_Error_Code {
	if m != nil {
		return m.Code
	}
	return WalletCreateR_Error_NULL
}

func (m *WalletCreateR_Error) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type WalletRecoverQ struct {
	RootPath             string   `protobuf:"bytes,1,opt,name=rootPath,proto3" json:"rootPath,omitempty"`
	Mnemonic             string   `protobuf:"bytes,2,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WalletRecoverQ) Reset()         { *m = WalletRecoverQ{} }
func (m *WalletRecoverQ) String() string { return proto.CompactTextString(m) }
func (*WalletRecoverQ) ProtoMessage()    {}
func (*WalletRecoverQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{2}
}

func (m *WalletRecoverQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletRecoverQ.Unmarshal(m, b)
}
func (m *WalletRecoverQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletRecoverQ.Marshal(b, m, deterministic)
}
func (m *WalletRecoverQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletRecoverQ.Merge(m, src)
}
func (m *WalletRecoverQ) XXX_Size() int {
	return xxx_messageInfo_WalletRecoverQ.Size(m)
}
func (m *WalletRecoverQ) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletRecoverQ.DiscardUnknown(m)
}

var xxx_messageInfo_WalletRecoverQ proto.InternalMessageInfo

func (m *WalletRecoverQ) GetRootPath() string {
	if m != nil {
		return m.RootPath
	}
	return ""
}

func (m *WalletRecoverQ) GetMnemonic() string {
	if m != nil {
		return m.Mnemonic
	}
	return ""
}

type WalletRecoverR struct {
	Error                *WalletRecoverR_Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *WalletRecoverR) Reset()         { *m = WalletRecoverR{} }
func (m *WalletRecoverR) String() string { return proto.CompactTextString(m) }
func (*WalletRecoverR) ProtoMessage()    {}
func (*WalletRecoverR) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{3}
}

func (m *WalletRecoverR) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletRecoverR.Unmarshal(m, b)
}
func (m *WalletRecoverR) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletRecoverR.Marshal(b, m, deterministic)
}
func (m *WalletRecoverR) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletRecoverR.Merge(m, src)
}
func (m *WalletRecoverR) XXX_Size() int {
	return xxx_messageInfo_WalletRecoverR.Size(m)
}
func (m *WalletRecoverR) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletRecoverR.DiscardUnknown(m)
}

var xxx_messageInfo_WalletRecoverR proto.InternalMessageInfo

func (m *WalletRecoverR) GetError() *WalletRecoverR_Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type WalletRecoverR_Error struct {
	Code                 WalletRecoverR_Error_Code `protobuf:"varint,1,opt,name=code,proto3,enum=anytype.WalletRecoverR_Error_Code" json:"code,omitempty"`
	Desc                 string                    `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *WalletRecoverR_Error) Reset()         { *m = WalletRecoverR_Error{} }
func (m *WalletRecoverR_Error) String() string { return proto.CompactTextString(m) }
func (*WalletRecoverR_Error) ProtoMessage()    {}
func (*WalletRecoverR_Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{3, 0}
}

func (m *WalletRecoverR_Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletRecoverR_Error.Unmarshal(m, b)
}
func (m *WalletRecoverR_Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletRecoverR_Error.Marshal(b, m, deterministic)
}
func (m *WalletRecoverR_Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletRecoverR_Error.Merge(m, src)
}
func (m *WalletRecoverR_Error) XXX_Size() int {
	return xxx_messageInfo_WalletRecoverR_Error.Size(m)
}
func (m *WalletRecoverR_Error) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletRecoverR_Error.DiscardUnknown(m)
}

var xxx_messageInfo_WalletRecoverR_Error proto.InternalMessageInfo

func (m *WalletRecoverR_Error) GetCode() WalletRecoverR_Error_Code {
	if m != nil {
		return m.Code
	}
	return WalletRecoverR_Error_NULL
}

func (m *WalletRecoverR_Error) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type AccountCreateQ struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	AvatarLocalPath      string   `protobuf:"bytes,2,opt,name=avatarLocalPath,proto3" json:"avatarLocalPath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountCreateQ) Reset()         { *m = AccountCreateQ{} }
func (m *AccountCreateQ) String() string { return proto.CompactTextString(m) }
func (*AccountCreateQ) ProtoMessage()    {}
func (*AccountCreateQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{4}
}

func (m *AccountCreateQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountCreateQ.Unmarshal(m, b)
}
func (m *AccountCreateQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountCreateQ.Marshal(b, m, deterministic)
}
func (m *AccountCreateQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountCreateQ.Merge(m, src)
}
func (m *AccountCreateQ) XXX_Size() int {
	return xxx_messageInfo_AccountCreateQ.Size(m)
}
func (m *AccountCreateQ) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountCreateQ.DiscardUnknown(m)
}

var xxx_messageInfo_AccountCreateQ proto.InternalMessageInfo

func (m *AccountCreateQ) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AccountCreateQ) GetAvatarLocalPath() string {
	if m != nil {
		return m.AvatarLocalPath
	}
	return ""
}

type AccountCreateR struct {
	Error                *AccountCreateR_Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Account              *Account              `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AccountCreateR) Reset()         { *m = AccountCreateR{} }
func (m *AccountCreateR) String() string { return proto.CompactTextString(m) }
func (*AccountCreateR) ProtoMessage()    {}
func (*AccountCreateR) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{5}
}

func (m *AccountCreateR) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountCreateR.Unmarshal(m, b)
}
func (m *AccountCreateR) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountCreateR.Marshal(b, m, deterministic)
}
func (m *AccountCreateR) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountCreateR.Merge(m, src)
}
func (m *AccountCreateR) XXX_Size() int {
	return xxx_messageInfo_AccountCreateR.Size(m)
}
func (m *AccountCreateR) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountCreateR.DiscardUnknown(m)
}

var xxx_messageInfo_AccountCreateR proto.InternalMessageInfo

func (m *AccountCreateR) GetError() *AccountCreateR_Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *AccountCreateR) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type AccountCreateR_Error struct {
	Code                 AccountCreateR_Error_Code `protobuf:"varint,1,opt,name=code,proto3,enum=anytype.AccountCreateR_Error_Code" json:"code,omitempty"`
	Desc                 string                    `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *AccountCreateR_Error) Reset()         { *m = AccountCreateR_Error{} }
func (m *AccountCreateR_Error) String() string { return proto.CompactTextString(m) }
func (*AccountCreateR_Error) ProtoMessage()    {}
func (*AccountCreateR_Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{5, 0}
}

func (m *AccountCreateR_Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountCreateR_Error.Unmarshal(m, b)
}
func (m *AccountCreateR_Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountCreateR_Error.Marshal(b, m, deterministic)
}
func (m *AccountCreateR_Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountCreateR_Error.Merge(m, src)
}
func (m *AccountCreateR_Error) XXX_Size() int {
	return xxx_messageInfo_AccountCreateR_Error.Size(m)
}
func (m *AccountCreateR_Error) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountCreateR_Error.DiscardUnknown(m)
}

var xxx_messageInfo_AccountCreateR_Error proto.InternalMessageInfo

func (m *AccountCreateR_Error) GetCode() AccountCreateR_Error_Code {
	if m != nil {
		return m.Code
	}
	return AccountCreateR_Error_NULL
}

func (m *AccountCreateR_Error) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type AccountSelectQ struct {
	Index                int64    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountSelectQ) Reset()         { *m = AccountSelectQ{} }
func (m *AccountSelectQ) String() string { return proto.CompactTextString(m) }
func (*AccountSelectQ) ProtoMessage()    {}
func (*AccountSelectQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{6}
}

func (m *AccountSelectQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountSelectQ.Unmarshal(m, b)
}
func (m *AccountSelectQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountSelectQ.Marshal(b, m, deterministic)
}
func (m *AccountSelectQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountSelectQ.Merge(m, src)
}
func (m *AccountSelectQ) XXX_Size() int {
	return xxx_messageInfo_AccountSelectQ.Size(m)
}
func (m *AccountSelectQ) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountSelectQ.DiscardUnknown(m)
}

var xxx_messageInfo_AccountSelectQ proto.InternalMessageInfo

func (m *AccountSelectQ) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type AccountSelectR struct {
	Error                *AccountSelectR_Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Account              *Account              `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AccountSelectR) Reset()         { *m = AccountSelectR{} }
func (m *AccountSelectR) String() string { return proto.CompactTextString(m) }
func (*AccountSelectR) ProtoMessage()    {}
func (*AccountSelectR) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{7}
}

func (m *AccountSelectR) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountSelectR.Unmarshal(m, b)
}
func (m *AccountSelectR) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountSelectR.Marshal(b, m, deterministic)
}
func (m *AccountSelectR) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountSelectR.Merge(m, src)
}
func (m *AccountSelectR) XXX_Size() int {
	return xxx_messageInfo_AccountSelectR.Size(m)
}
func (m *AccountSelectR) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountSelectR.DiscardUnknown(m)
}

var xxx_messageInfo_AccountSelectR proto.InternalMessageInfo

func (m *AccountSelectR) GetError() *AccountSelectR_Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *AccountSelectR) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type AccountSelectR_Error struct {
	Code                 AccountSelectR_Error_Code `protobuf:"varint,1,opt,name=code,proto3,enum=anytype.AccountSelectR_Error_Code" json:"code,omitempty"`
	Desc                 string                    `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *AccountSelectR_Error) Reset()         { *m = AccountSelectR_Error{} }
func (m *AccountSelectR_Error) String() string { return proto.CompactTextString(m) }
func (*AccountSelectR_Error) ProtoMessage()    {}
func (*AccountSelectR_Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{7, 0}
}

func (m *AccountSelectR_Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountSelectR_Error.Unmarshal(m, b)
}
func (m *AccountSelectR_Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountSelectR_Error.Marshal(b, m, deterministic)
}
func (m *AccountSelectR_Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountSelectR_Error.Merge(m, src)
}
func (m *AccountSelectR_Error) XXX_Size() int {
	return xxx_messageInfo_AccountSelectR_Error.Size(m)
}
func (m *AccountSelectR_Error) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountSelectR_Error.DiscardUnknown(m)
}

var xxx_messageInfo_AccountSelectR_Error proto.InternalMessageInfo

func (m *AccountSelectR_Error) GetCode() AccountSelectR_Error_Code {
	if m != nil {
		return m.Code
	}
	return AccountSelectR_Error_NULL
}

func (m *AccountSelectR_Error) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func init() {
	proto.RegisterEnum("anytype.WalletCreateR_Error_Code", WalletCreateR_Error_Code_name, WalletCreateR_Error_Code_value)
	proto.RegisterEnum("anytype.WalletRecoverR_Error_Code", WalletRecoverR_Error_Code_name, WalletRecoverR_Error_Code_value)
	proto.RegisterEnum("anytype.AccountCreateR_Error_Code", AccountCreateR_Error_Code_name, AccountCreateR_Error_Code_value)
	proto.RegisterEnum("anytype.AccountSelectR_Error_Code", AccountSelectR_Error_Code_name, AccountSelectR_Error_Code_value)
	proto.RegisterType((*WalletCreateQ)(nil), "anytype.WalletCreateQ")
	proto.RegisterType((*WalletCreateR)(nil), "anytype.WalletCreateR")
	proto.RegisterType((*WalletCreateR_Error)(nil), "anytype.WalletCreateR.Error")
	proto.RegisterType((*WalletRecoverQ)(nil), "anytype.WalletRecoverQ")
	proto.RegisterType((*WalletRecoverR)(nil), "anytype.WalletRecoverR")
	proto.RegisterType((*WalletRecoverR_Error)(nil), "anytype.WalletRecoverR.Error")
	proto.RegisterType((*AccountCreateQ)(nil), "anytype.AccountCreateQ")
	proto.RegisterType((*AccountCreateR)(nil), "anytype.AccountCreateR")
	proto.RegisterType((*AccountCreateR_Error)(nil), "anytype.AccountCreateR.Error")
	proto.RegisterType((*AccountSelectQ)(nil), "anytype.AccountSelectQ")
	proto.RegisterType((*AccountSelectR)(nil), "anytype.AccountSelectR")
	proto.RegisterType((*AccountSelectR_Error)(nil), "anytype.AccountSelectR.Error")
}

func init() { proto.RegisterFile("commands.proto", fileDescriptor_0dff099eb2e3dfdb) }

var fileDescriptor_0dff099eb2e3dfdb = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0x49, 0xe8, 0xd8, 0xe6, 0x6d, 0x25, 0x58, 0xd3, 0x98, 0x02, 0x68, 0x10, 0x24, 0x34,
	0x81, 0xd4, 0x43, 0x26, 0x38, 0x22, 0x79, 0xa9, 0x27, 0x2a, 0x42, 0xd2, 0xba, 0xce, 0x86, 0xb8,
	0x58, 0x59, 0xe2, 0xbd, 0x48, 0x49, 0x3c, 0xa5, 0xd9, 0xc4, 0x2e, 0x5c, 0x27, 0x3e, 0x02, 0x9f,
	0x80, 0x33, 0x9f, 0x82, 0x23, 0x07, 0x3e, 0x0e, 0x17, 0xd4, 0xbc, 0x74, 0x75, 0xe8, 0x0b, 0x4c,
	0xe2, 0xd6, 0xc7, 0xcf, 0xdf, 0xff, 0xc7, 0xfe, 0xf9, 0xe9, 0x13, 0xd0, 0x0c, 0x44, 0x1c, 0xfb,
	0x49, 0x38, 0x68, 0x9d, 0xa5, 0x22, 0x13, 0x70, 0xd1, 0x4f, 0x2e, 0xb3, 0xcb, 0x33, 0xae, 0xaf,
	0xc6, 0x22, 0xe4, 0x51, 0xb9, 0x6c, 0xbc, 0x00, 0x6b, 0x07, 0x7e, 0x14, 0xf1, 0xcc, 0x4a, 0xb9,
	0x9f, 0xf1, 0x1e, 0xd4, 0xc1, 0x52, 0x2a, 0x44, 0xd6, 0xf5, 0xb3, 0x93, 0x4d, 0xe5, 0xb1, 0xb2,
	0xbd, 0x4c, 0x46, 0xb1, 0xf1, 0x59, 0x95, 0xd5, 0x04, 0x9a, 0x60, 0x81, 0xa7, 0xa9, 0x48, 0x73,
	0xe9, 0x8a, 0xf9, 0xb0, 0x55, 0x56, 0x69, 0x49, 0xb2, 0x16, 0x1e, 0x6a, 0x48, 0x21, 0x1d, 0x56,
	0x88, 0x13, 0x1e, 0x8b, 0xe4, 0x34, 0xd8, 0x54, 0x8b, 0x0a, 0x55, 0xac, 0x7f, 0x55, 0xc0, 0x42,
	0x2e, 0x86, 0x2f, 0x41, 0x23, 0x10, 0x21, 0xcf, 0x8d, 0x9b, 0xe6, 0x93, 0x59, 0xc6, 0x2d, 0x4b,
	0x84, 0x9c, 0xe4, 0x72, 0x08, 0x41, 0x23, 0xe4, 0x83, 0xca, 0x38, 0xff, 0x6d, 0xf4, 0x41, 0x63,
	0xa8, 0x80, 0x4b, 0xa0, 0xe1, 0x78, 0xb6, 0xad, 0xdd, 0x82, 0xf7, 0xc0, 0x9a, 0xe7, 0xbc, 0x75,
	0xdc, 0x03, 0x87, 0x61, 0x42, 0x5c, 0xa2, 0x29, 0x70, 0x0d, 0x2c, 0xef, 0xa2, 0x36, 0xeb, 0x38,
	0x5d, 0x8f, 0x6a, 0x2a, 0xdc, 0x02, 0x0f, 0xf6, 0x50, 0xc7, 0xc6, 0x6d, 0x46, 0x5d, 0x66, 0x11,
	0x8c, 0x28, 0x66, 0xb6, 0x6b, 0x21, 0x9b, 0x11, 0xdc, 0x75, 0x35, 0x6e, 0xbc, 0x01, 0xcd, 0xe2,
	0x28, 0x84, 0x07, 0xe2, 0x82, 0xa7, 0x33, 0xc9, 0xcd, 0xba, 0xb3, 0xf1, 0x43, 0xa9, 0x59, 0x11,
	0xb8, 0x23, 0x63, 0x7d, 0x54, 0xbb, 0x7d, 0xa5, 0x93, 0xb8, 0xea, 0x57, 0x23, 0x76, 0xaf, 0x24,
	0x76, 0xc6, 0xcc, 0xdd, 0xf3, 0xe0, 0x99, 0xff, 0x0e, 0xcf, 0xd8, 0x07, 0x4d, 0x14, 0x04, 0xe2,
	0x3c, 0x19, 0xef, 0xaa, 0xf3, 0x01, 0x4f, 0x13, 0x3f, 0xe6, 0x15, 0x9b, 0x2a, 0x86, 0xdb, 0xe0,
	0xae, 0x7f, 0xe1, 0x67, 0x7e, 0x6a, 0x8b, 0xc0, 0x8f, 0x72, 0x7c, 0xc5, 0x01, 0xea, 0xcb, 0xc6,
	0x77, 0xb5, 0x66, 0x3c, 0x83, 0x94, 0xac, 0x93, 0x3b, 0xf0, 0x39, 0x58, 0xf4, 0x8b, 0x74, 0x5e,
	0x69, 0xc5, 0xd4, 0xea, 0xdb, 0x48, 0x25, 0xd0, 0x7f, 0xce, 0xa5, 0x3a, 0xa9, 0xd2, 0x3c, 0xaa,
	0x9f, 0x6e, 0xd0, 0x92, 0x9b, 0x60, 0xfd, 0xba, 0x25, 0xfb, 0x14, 0x11, 0xca, 0x1c, 0xb7, 0x8d,
	0x35, 0x0e, 0x37, 0x00, 0x1c, 0xcb, 0x60, 0xca, 0x1c, 0xf4, 0x0e, 0x6b, 0x47, 0xb5, 0x1d, 0x98,
	0x32, 0xb4, 0x8f, 0x28, 0x22, 0xda, 0xb1, 0xf1, 0x6c, 0x04, 0xb2, 0xcf, 0x23, 0x1e, 0x64, 0x3d,
	0xb8, 0x0e, 0x16, 0x4e, 0x93, 0x90, 0x7f, 0xcc, 0xaf, 0x77, 0x9b, 0x14, 0x81, 0xf1, 0x4b, 0xad,
	0x09, 0xe7, 0x13, 0x2f, 0x75, 0x37, 0x27, 0x7e, 0xa5, 0xfe, 0x25, 0x71, 0xa9, 0xd2, 0x3c, 0xe2,
	0xdf, 0x94, 0xff, 0x30, 0x05, 0xe0, 0x53, 0xb0, 0x75, 0x1d, 0x33, 0xfc, 0xbe, 0xd3, 0xa7, 0x7d,
	0xb6, 0xeb, 0x51, 0x66, 0xb9, 0x84, 0x78, 0x5d, 0x8a, 0xdb, 0xda, 0x91, 0xfc, 0x3c, 0xc4, 0x73,
	0x8a, 0x67, 0x3b, 0x96, 0xdd, 0xf7, 0x3a, 0x4e, 0x9b, 0x21, 0xcb, 0x72, 0x3d, 0x87, 0xb2, 0x8e,
	0xb3, 0xe7, 0x6a, 0x27, 0xe6, 0x17, 0x15, 0x34, 0xad, 0xe8, 0x94, 0x27, 0x99, 0x55, 0x0e, 0x73,
	0xf8, 0x1a, 0xac, 0x8e, 0x4f, 0x40, 0xb8, 0x31, 0x71, 0x30, 0xf6, 0xf4, 0xc9, 0xeb, 0x04, 0xa2,
	0x6a, 0x82, 0x97, 0x53, 0x00, 0xde, 0x9f, 0x3c, 0x1d, 0x7a, 0xfa, 0x94, 0x44, 0x6e, 0x21, 0xb5,
	0xfc, 0x98, 0x85, 0xfc, 0xaf, 0xd7, 0xa7, 0x24, 0xc6, 0x2d, 0x8a, 0x37, 0xfc, 0xd3, 0xa2, 0x6c,
	0x4b, 0x7d, 0x4a, 0x82, 0xec, 0x36, 0x3e, 0xa8, 0x67, 0x87, 0x87, 0x77, 0xf2, 0xaf, 0xd8, 0xce,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0xf1, 0x86, 0x7f, 0xee, 0x06, 0x00, 0x00,
}
