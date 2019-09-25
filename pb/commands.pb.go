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

type WalletCreateResponse_Error_Code int32

const (
	WalletCreateResponse_Error_NULL                        WalletCreateResponse_Error_Code = 0
	WalletCreateResponse_Error_UNKNOWN_ERROR               WalletCreateResponse_Error_Code = 1
	WalletCreateResponse_Error_BAD_INPUT                   WalletCreateResponse_Error_Code = 2
	WalletCreateResponse_Error_FAILED_TO_CREATE_LOCAL_REPO WalletCreateResponse_Error_Code = 101
)

var WalletCreateResponse_Error_Code_name = map[int32]string{
	0:   "NULL",
	1:   "UNKNOWN_ERROR",
	2:   "BAD_INPUT",
	101: "FAILED_TO_CREATE_LOCAL_REPO",
}

var WalletCreateResponse_Error_Code_value = map[string]int32{
	"NULL":                        0,
	"UNKNOWN_ERROR":               1,
	"BAD_INPUT":                   2,
	"FAILED_TO_CREATE_LOCAL_REPO": 101,
}

func (x WalletCreateResponse_Error_Code) String() string {
	return proto.EnumName(WalletCreateResponse_Error_Code_name, int32(x))
}

func (WalletCreateResponse_Error_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{1, 0, 0}
}

type WalletRecoverResponse_Error_Code int32

const (
	WalletRecoverResponse_Error_NULL          WalletRecoverResponse_Error_Code = 0
	WalletRecoverResponse_Error_UNKNOWN_ERROR WalletRecoverResponse_Error_Code = 1
	WalletRecoverResponse_Error_BAD_INPUT     WalletRecoverResponse_Error_Code = 2
)

var WalletRecoverResponse_Error_Code_name = map[int32]string{
	0: "NULL",
	1: "UNKNOWN_ERROR",
	2: "BAD_INPUT",
}

var WalletRecoverResponse_Error_Code_value = map[string]int32{
	"NULL":          0,
	"UNKNOWN_ERROR": 1,
	"BAD_INPUT":     2,
}

func (x WalletRecoverResponse_Error_Code) String() string {
	return proto.EnumName(WalletRecoverResponse_Error_Code_name, int32(x))
}

func (WalletRecoverResponse_Error_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{3, 0, 0}
}

type AccountCreateResponse_Error_Code int32

const (
	AccountCreateResponse_Error_NULL                 AccountCreateResponse_Error_Code = 0
	AccountCreateResponse_Error_UNKNOWN_ERROR        AccountCreateResponse_Error_Code = 1
	AccountCreateResponse_Error_BAD_INPUT            AccountCreateResponse_Error_Code = 2
	AccountCreateResponse_Error_FAILED_TO_START_NODE AccountCreateResponse_Error_Code = 101
	AccountCreateResponse_Error_FAILED_TO_SET_NAME   AccountCreateResponse_Error_Code = 102
	AccountCreateResponse_Error_FAILED_TO_SET_AVATAR AccountCreateResponse_Error_Code = 103
)

var AccountCreateResponse_Error_Code_name = map[int32]string{
	0:   "NULL",
	1:   "UNKNOWN_ERROR",
	2:   "BAD_INPUT",
	101: "FAILED_TO_START_NODE",
	102: "FAILED_TO_SET_NAME",
	103: "FAILED_TO_SET_AVATAR",
}

var AccountCreateResponse_Error_Code_value = map[string]int32{
	"NULL":                 0,
	"UNKNOWN_ERROR":        1,
	"BAD_INPUT":            2,
	"FAILED_TO_START_NODE": 101,
	"FAILED_TO_SET_NAME":   102,
	"FAILED_TO_SET_AVATAR": 103,
}

func (x AccountCreateResponse_Error_Code) String() string {
	return proto.EnumName(AccountCreateResponse_Error_Code_name, int32(x))
}

func (AccountCreateResponse_Error_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{5, 0, 0}
}

type AccountSelectResponse_Error_Code int32

const (
	AccountSelectResponse_Error_NULL                            AccountSelectResponse_Error_Code = 0
	AccountSelectResponse_Error_UNKNOWN_ERROR                   AccountSelectResponse_Error_Code = 1
	AccountSelectResponse_Error_BAD_INPUT                       AccountSelectResponse_Error_Code = 2
	AccountSelectResponse_Error_FAILED_TO_CREATE_LOCAL_REPO     AccountSelectResponse_Error_Code = 101
	AccountSelectResponse_Error_LOCAL_REPO_EXISTS_BUT_CORRUPTED AccountSelectResponse_Error_Code = 102
	AccountSelectResponse_Error_FAILED_TO_RUN_NODE              AccountSelectResponse_Error_Code = 103
	AccountSelectResponse_Error_FAILED_TO_FIND_ACCOUNT_INFO     AccountSelectResponse_Error_Code = 104
)

var AccountSelectResponse_Error_Code_name = map[int32]string{
	0:   "NULL",
	1:   "UNKNOWN_ERROR",
	2:   "BAD_INPUT",
	101: "FAILED_TO_CREATE_LOCAL_REPO",
	102: "LOCAL_REPO_EXISTS_BUT_CORRUPTED",
	103: "FAILED_TO_RUN_NODE",
	104: "FAILED_TO_FIND_ACCOUNT_INFO",
}

var AccountSelectResponse_Error_Code_value = map[string]int32{
	"NULL":                            0,
	"UNKNOWN_ERROR":                   1,
	"BAD_INPUT":                       2,
	"FAILED_TO_CREATE_LOCAL_REPO":     101,
	"LOCAL_REPO_EXISTS_BUT_CORRUPTED": 102,
	"FAILED_TO_RUN_NODE":              103,
	"FAILED_TO_FIND_ACCOUNT_INFO":     104,
}

func (x AccountSelectResponse_Error_Code) String() string {
	return proto.EnumName(AccountSelectResponse_Error_Code_name, int32(x))
}

func (AccountSelectResponse_Error_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{7, 0, 0}
}

type WalletCreateRequest struct {
	RootPath             string   `protobuf:"bytes,1,opt,name=rootPath,proto3" json:"rootPath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WalletCreateRequest) Reset()         { *m = WalletCreateRequest{} }
func (m *WalletCreateRequest) String() string { return proto.CompactTextString(m) }
func (*WalletCreateRequest) ProtoMessage()    {}
func (*WalletCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{0}
}

func (m *WalletCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletCreateRequest.Unmarshal(m, b)
}
func (m *WalletCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletCreateRequest.Marshal(b, m, deterministic)
}
func (m *WalletCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletCreateRequest.Merge(m, src)
}
func (m *WalletCreateRequest) XXX_Size() int {
	return xxx_messageInfo_WalletCreateRequest.Size(m)
}
func (m *WalletCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WalletCreateRequest proto.InternalMessageInfo

func (m *WalletCreateRequest) GetRootPath() string {
	if m != nil {
		return m.RootPath
	}
	return ""
}

type WalletCreateResponse struct {
	Error                *WalletCreateResponse_Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Mnemonic             string                      `protobuf:"bytes,2,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *WalletCreateResponse) Reset()         { *m = WalletCreateResponse{} }
func (m *WalletCreateResponse) String() string { return proto.CompactTextString(m) }
func (*WalletCreateResponse) ProtoMessage()    {}
func (*WalletCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{1}
}

func (m *WalletCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletCreateResponse.Unmarshal(m, b)
}
func (m *WalletCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletCreateResponse.Marshal(b, m, deterministic)
}
func (m *WalletCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletCreateResponse.Merge(m, src)
}
func (m *WalletCreateResponse) XXX_Size() int {
	return xxx_messageInfo_WalletCreateResponse.Size(m)
}
func (m *WalletCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WalletCreateResponse proto.InternalMessageInfo

func (m *WalletCreateResponse) GetError() *WalletCreateResponse_Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *WalletCreateResponse) GetMnemonic() string {
	if m != nil {
		return m.Mnemonic
	}
	return ""
}

type WalletCreateResponse_Error struct {
	Code                 WalletCreateResponse_Error_Code `protobuf:"varint,1,opt,name=code,proto3,enum=anytype.WalletCreateResponse_Error_Code" json:"code,omitempty"`
	Descritrion          string                          `protobuf:"bytes,2,opt,name=descritrion,proto3" json:"descritrion,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *WalletCreateResponse_Error) Reset()         { *m = WalletCreateResponse_Error{} }
func (m *WalletCreateResponse_Error) String() string { return proto.CompactTextString(m) }
func (*WalletCreateResponse_Error) ProtoMessage()    {}
func (*WalletCreateResponse_Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{1, 0}
}

func (m *WalletCreateResponse_Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletCreateResponse_Error.Unmarshal(m, b)
}
func (m *WalletCreateResponse_Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletCreateResponse_Error.Marshal(b, m, deterministic)
}
func (m *WalletCreateResponse_Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletCreateResponse_Error.Merge(m, src)
}
func (m *WalletCreateResponse_Error) XXX_Size() int {
	return xxx_messageInfo_WalletCreateResponse_Error.Size(m)
}
func (m *WalletCreateResponse_Error) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletCreateResponse_Error.DiscardUnknown(m)
}

var xxx_messageInfo_WalletCreateResponse_Error proto.InternalMessageInfo

func (m *WalletCreateResponse_Error) GetCode() WalletCreateResponse_Error_Code {
	if m != nil {
		return m.Code
	}
	return WalletCreateResponse_Error_NULL
}

func (m *WalletCreateResponse_Error) GetDescritrion() string {
	if m != nil {
		return m.Descritrion
	}
	return ""
}

type WalletRecoverRequest struct {
	RootPath             string   `protobuf:"bytes,1,opt,name=rootPath,proto3" json:"rootPath,omitempty"`
	Mnemonic             string   `protobuf:"bytes,2,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WalletRecoverRequest) Reset()         { *m = WalletRecoverRequest{} }
func (m *WalletRecoverRequest) String() string { return proto.CompactTextString(m) }
func (*WalletRecoverRequest) ProtoMessage()    {}
func (*WalletRecoverRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{2}
}

func (m *WalletRecoverRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletRecoverRequest.Unmarshal(m, b)
}
func (m *WalletRecoverRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletRecoverRequest.Marshal(b, m, deterministic)
}
func (m *WalletRecoverRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletRecoverRequest.Merge(m, src)
}
func (m *WalletRecoverRequest) XXX_Size() int {
	return xxx_messageInfo_WalletRecoverRequest.Size(m)
}
func (m *WalletRecoverRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletRecoverRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WalletRecoverRequest proto.InternalMessageInfo

func (m *WalletRecoverRequest) GetRootPath() string {
	if m != nil {
		return m.RootPath
	}
	return ""
}

func (m *WalletRecoverRequest) GetMnemonic() string {
	if m != nil {
		return m.Mnemonic
	}
	return ""
}

type WalletRecoverResponse struct {
	Error                *WalletRecoverResponse_Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *WalletRecoverResponse) Reset()         { *m = WalletRecoverResponse{} }
func (m *WalletRecoverResponse) String() string { return proto.CompactTextString(m) }
func (*WalletRecoverResponse) ProtoMessage()    {}
func (*WalletRecoverResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{3}
}

func (m *WalletRecoverResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletRecoverResponse.Unmarshal(m, b)
}
func (m *WalletRecoverResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletRecoverResponse.Marshal(b, m, deterministic)
}
func (m *WalletRecoverResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletRecoverResponse.Merge(m, src)
}
func (m *WalletRecoverResponse) XXX_Size() int {
	return xxx_messageInfo_WalletRecoverResponse.Size(m)
}
func (m *WalletRecoverResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletRecoverResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WalletRecoverResponse proto.InternalMessageInfo

func (m *WalletRecoverResponse) GetError() *WalletRecoverResponse_Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type WalletRecoverResponse_Error struct {
	Code                 WalletRecoverResponse_Error_Code `protobuf:"varint,1,opt,name=code,proto3,enum=anytype.WalletRecoverResponse_Error_Code" json:"code,omitempty"`
	Description          string                           `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *WalletRecoverResponse_Error) Reset()         { *m = WalletRecoverResponse_Error{} }
func (m *WalletRecoverResponse_Error) String() string { return proto.CompactTextString(m) }
func (*WalletRecoverResponse_Error) ProtoMessage()    {}
func (*WalletRecoverResponse_Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{3, 0}
}

func (m *WalletRecoverResponse_Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletRecoverResponse_Error.Unmarshal(m, b)
}
func (m *WalletRecoverResponse_Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletRecoverResponse_Error.Marshal(b, m, deterministic)
}
func (m *WalletRecoverResponse_Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletRecoverResponse_Error.Merge(m, src)
}
func (m *WalletRecoverResponse_Error) XXX_Size() int {
	return xxx_messageInfo_WalletRecoverResponse_Error.Size(m)
}
func (m *WalletRecoverResponse_Error) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletRecoverResponse_Error.DiscardUnknown(m)
}

var xxx_messageInfo_WalletRecoverResponse_Error proto.InternalMessageInfo

func (m *WalletRecoverResponse_Error) GetCode() WalletRecoverResponse_Error_Code {
	if m != nil {
		return m.Code
	}
	return WalletRecoverResponse_Error_NULL
}

func (m *WalletRecoverResponse_Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type AccountCreateRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	AvatarLocalPath      string   `protobuf:"bytes,2,opt,name=avatarLocalPath,proto3" json:"avatarLocalPath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountCreateRequest) Reset()         { *m = AccountCreateRequest{} }
func (m *AccountCreateRequest) String() string { return proto.CompactTextString(m) }
func (*AccountCreateRequest) ProtoMessage()    {}
func (*AccountCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{4}
}

func (m *AccountCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountCreateRequest.Unmarshal(m, b)
}
func (m *AccountCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountCreateRequest.Marshal(b, m, deterministic)
}
func (m *AccountCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountCreateRequest.Merge(m, src)
}
func (m *AccountCreateRequest) XXX_Size() int {
	return xxx_messageInfo_AccountCreateRequest.Size(m)
}
func (m *AccountCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountCreateRequest proto.InternalMessageInfo

func (m *AccountCreateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AccountCreateRequest) GetAvatarLocalPath() string {
	if m != nil {
		return m.AvatarLocalPath
	}
	return ""
}

type AccountCreateResponse struct {
	Error                *AccountCreateResponse_Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Account              *Account                     `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *AccountCreateResponse) Reset()         { *m = AccountCreateResponse{} }
func (m *AccountCreateResponse) String() string { return proto.CompactTextString(m) }
func (*AccountCreateResponse) ProtoMessage()    {}
func (*AccountCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{5}
}

func (m *AccountCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountCreateResponse.Unmarshal(m, b)
}
func (m *AccountCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountCreateResponse.Marshal(b, m, deterministic)
}
func (m *AccountCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountCreateResponse.Merge(m, src)
}
func (m *AccountCreateResponse) XXX_Size() int {
	return xxx_messageInfo_AccountCreateResponse.Size(m)
}
func (m *AccountCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountCreateResponse proto.InternalMessageInfo

func (m *AccountCreateResponse) GetError() *AccountCreateResponse_Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *AccountCreateResponse) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type AccountCreateResponse_Error struct {
	Code                 AccountCreateResponse_Error_Code `protobuf:"varint,1,opt,name=code,proto3,enum=anytype.AccountCreateResponse_Error_Code" json:"code,omitempty"`
	Description          string                           `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *AccountCreateResponse_Error) Reset()         { *m = AccountCreateResponse_Error{} }
func (m *AccountCreateResponse_Error) String() string { return proto.CompactTextString(m) }
func (*AccountCreateResponse_Error) ProtoMessage()    {}
func (*AccountCreateResponse_Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{5, 0}
}

func (m *AccountCreateResponse_Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountCreateResponse_Error.Unmarshal(m, b)
}
func (m *AccountCreateResponse_Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountCreateResponse_Error.Marshal(b, m, deterministic)
}
func (m *AccountCreateResponse_Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountCreateResponse_Error.Merge(m, src)
}
func (m *AccountCreateResponse_Error) XXX_Size() int {
	return xxx_messageInfo_AccountCreateResponse_Error.Size(m)
}
func (m *AccountCreateResponse_Error) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountCreateResponse_Error.DiscardUnknown(m)
}

var xxx_messageInfo_AccountCreateResponse_Error proto.InternalMessageInfo

func (m *AccountCreateResponse_Error) GetCode() AccountCreateResponse_Error_Code {
	if m != nil {
		return m.Code
	}
	return AccountCreateResponse_Error_NULL
}

func (m *AccountCreateResponse_Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type AccountSelectRequest struct {
	Index                int64    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountSelectRequest) Reset()         { *m = AccountSelectRequest{} }
func (m *AccountSelectRequest) String() string { return proto.CompactTextString(m) }
func (*AccountSelectRequest) ProtoMessage()    {}
func (*AccountSelectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{6}
}

func (m *AccountSelectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountSelectRequest.Unmarshal(m, b)
}
func (m *AccountSelectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountSelectRequest.Marshal(b, m, deterministic)
}
func (m *AccountSelectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountSelectRequest.Merge(m, src)
}
func (m *AccountSelectRequest) XXX_Size() int {
	return xxx_messageInfo_AccountSelectRequest.Size(m)
}
func (m *AccountSelectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountSelectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountSelectRequest proto.InternalMessageInfo

func (m *AccountSelectRequest) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type AccountSelectResponse struct {
	Error                *AccountSelectResponse_Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Account              *Account                     `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *AccountSelectResponse) Reset()         { *m = AccountSelectResponse{} }
func (m *AccountSelectResponse) String() string { return proto.CompactTextString(m) }
func (*AccountSelectResponse) ProtoMessage()    {}
func (*AccountSelectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{7}
}

func (m *AccountSelectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountSelectResponse.Unmarshal(m, b)
}
func (m *AccountSelectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountSelectResponse.Marshal(b, m, deterministic)
}
func (m *AccountSelectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountSelectResponse.Merge(m, src)
}
func (m *AccountSelectResponse) XXX_Size() int {
	return xxx_messageInfo_AccountSelectResponse.Size(m)
}
func (m *AccountSelectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountSelectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountSelectResponse proto.InternalMessageInfo

func (m *AccountSelectResponse) GetError() *AccountSelectResponse_Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *AccountSelectResponse) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type AccountSelectResponse_Error struct {
	Code                 AccountSelectResponse_Error_Code `protobuf:"varint,1,opt,name=code,proto3,enum=anytype.AccountSelectResponse_Error_Code" json:"code,omitempty"`
	Description          string                           `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *AccountSelectResponse_Error) Reset()         { *m = AccountSelectResponse_Error{} }
func (m *AccountSelectResponse_Error) String() string { return proto.CompactTextString(m) }
func (*AccountSelectResponse_Error) ProtoMessage()    {}
func (*AccountSelectResponse_Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{7, 0}
}

func (m *AccountSelectResponse_Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountSelectResponse_Error.Unmarshal(m, b)
}
func (m *AccountSelectResponse_Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountSelectResponse_Error.Marshal(b, m, deterministic)
}
func (m *AccountSelectResponse_Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountSelectResponse_Error.Merge(m, src)
}
func (m *AccountSelectResponse_Error) XXX_Size() int {
	return xxx_messageInfo_AccountSelectResponse_Error.Size(m)
}
func (m *AccountSelectResponse_Error) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountSelectResponse_Error.DiscardUnknown(m)
}

var xxx_messageInfo_AccountSelectResponse_Error proto.InternalMessageInfo

func (m *AccountSelectResponse_Error) GetCode() AccountSelectResponse_Error_Code {
	if m != nil {
		return m.Code
	}
	return AccountSelectResponse_Error_NULL
}

func (m *AccountSelectResponse_Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterEnum("anytype.WalletCreateResponse_Error_Code", WalletCreateResponse_Error_Code_name, WalletCreateResponse_Error_Code_value)
	proto.RegisterEnum("anytype.WalletRecoverResponse_Error_Code", WalletRecoverResponse_Error_Code_name, WalletRecoverResponse_Error_Code_value)
	proto.RegisterEnum("anytype.AccountCreateResponse_Error_Code", AccountCreateResponse_Error_Code_name, AccountCreateResponse_Error_Code_value)
	proto.RegisterEnum("anytype.AccountSelectResponse_Error_Code", AccountSelectResponse_Error_Code_name, AccountSelectResponse_Error_Code_value)
	proto.RegisterType((*WalletCreateRequest)(nil), "anytype.WalletCreateRequest")
	proto.RegisterType((*WalletCreateResponse)(nil), "anytype.WalletCreateResponse")
	proto.RegisterType((*WalletCreateResponse_Error)(nil), "anytype.WalletCreateResponse.Error")
	proto.RegisterType((*WalletRecoverRequest)(nil), "anytype.WalletRecoverRequest")
	proto.RegisterType((*WalletRecoverResponse)(nil), "anytype.WalletRecoverResponse")
	proto.RegisterType((*WalletRecoverResponse_Error)(nil), "anytype.WalletRecoverResponse.Error")
	proto.RegisterType((*AccountCreateRequest)(nil), "anytype.AccountCreateRequest")
	proto.RegisterType((*AccountCreateResponse)(nil), "anytype.AccountCreateResponse")
	proto.RegisterType((*AccountCreateResponse_Error)(nil), "anytype.AccountCreateResponse.Error")
	proto.RegisterType((*AccountSelectRequest)(nil), "anytype.AccountSelectRequest")
	proto.RegisterType((*AccountSelectResponse)(nil), "anytype.AccountSelectResponse")
	proto.RegisterType((*AccountSelectResponse_Error)(nil), "anytype.AccountSelectResponse.Error")
}

func init() { proto.RegisterFile("commands.proto", fileDescriptor_0dff099eb2e3dfdb) }

var fileDescriptor_0dff099eb2e3dfdb = []byte{
	// 663 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0xe2, 0xa6, 0x5f, 0xdb, 0xdb, 0x1f, 0xcc, 0x90, 0xa2, 0x2a, 0x50, 0x5a, 0xb9, 0x2c,
	0x0a, 0x42, 0x91, 0x08, 0x2b, 0x10, 0x2c, 0x5c, 0xc7, 0x95, 0xa2, 0x86, 0x71, 0x34, 0x19, 0x53,
	0x84, 0x90, 0x46, 0xae, 0x33, 0x6d, 0x23, 0x25, 0x9e, 0xe0, 0x4c, 0x2b, 0xba, 0xe1, 0x49, 0x2a,
	0x9e, 0x80, 0x0d, 0x7b, 0x5e, 0x81, 0xf7, 0x40, 0x2c, 0x79, 0x02, 0x14, 0x3b, 0x4e, 0x33, 0x8e,
	0x9b, 0x26, 0x95, 0x58, 0xde, 0x9b, 0x73, 0xcf, 0x4d, 0xce, 0xb9, 0x39, 0x03, 0x6b, 0xbe, 0xe8,
	0x74, 0xbc, 0xa0, 0xd9, 0x2b, 0x75, 0x43, 0x21, 0x05, 0x5a, 0xf0, 0x82, 0x0b, 0x79, 0xd1, 0xe5,
	0xc5, 0x95, 0x8e, 0x68, 0xf2, 0xf6, 0xa0, 0x6d, 0x3c, 0x87, 0x7b, 0x87, 0x5e, 0xbb, 0xcd, 0xa5,
	0x15, 0x72, 0x4f, 0x72, 0xc2, 0x3f, 0x9d, 0xf1, 0x9e, 0x44, 0x45, 0x58, 0x0c, 0x85, 0x90, 0x75,
	0x4f, 0x9e, 0x6e, 0xe4, 0xb6, 0x73, 0xbb, 0x4b, 0x64, 0x58, 0x1b, 0xdf, 0x34, 0x28, 0xa8, 0x33,
	0xbd, 0xae, 0x08, 0x7a, 0x1c, 0xbd, 0x84, 0x79, 0x1e, 0x86, 0x22, 0x8c, 0x26, 0x96, 0xcb, 0x3b,
	0xa5, 0xc1, 0xca, 0x52, 0x16, 0xba, 0x64, 0xf7, 0xa1, 0x24, 0x9e, 0xe8, 0xef, 0xeb, 0x04, 0xbc,
	0x23, 0x82, 0x96, 0xbf, 0xa1, 0xc5, 0xfb, 0x92, 0xba, 0xf8, 0x23, 0x07, 0xf3, 0x11, 0x18, 0xbd,
	0x86, 0xbc, 0x2f, 0x9a, 0x3c, 0xe2, 0x5f, 0x2b, 0xef, 0x4e, 0xc1, 0x5f, 0xb2, 0x44, 0x93, 0x93,
	0x68, 0x0a, 0x6d, 0xc3, 0x72, 0x93, 0xf7, 0xfc, 0xb0, 0x25, 0xc3, 0x96, 0x08, 0x06, 0x6b, 0x46,
	0x5b, 0x46, 0x03, 0xf2, 0x7d, 0x3c, 0x5a, 0x84, 0x3c, 0x76, 0x6b, 0x35, 0xfd, 0x3f, 0x74, 0x17,
	0x56, 0x5d, 0x7c, 0x80, 0x9d, 0x43, 0xcc, 0x6c, 0x42, 0x1c, 0xa2, 0xe7, 0xd0, 0x2a, 0x2c, 0xed,
	0x99, 0x15, 0x56, 0xc5, 0x75, 0x97, 0xea, 0x1a, 0xda, 0x82, 0x07, 0xfb, 0x66, 0xb5, 0x66, 0x57,
	0x18, 0x75, 0x98, 0x45, 0x6c, 0x93, 0xda, 0xac, 0xe6, 0x58, 0x66, 0x8d, 0x11, 0xbb, 0xee, 0xe8,
	0xdc, 0xc0, 0x89, 0x5a, 0x84, 0xfb, 0xe2, 0x9c, 0x87, 0x53, 0x48, 0x3c, 0x49, 0x0e, 0xe3, 0x4f,
	0x0e, 0xd6, 0x53, 0x84, 0x03, 0xfd, 0x5f, 0xa9, 0xfa, 0x3f, 0x4e, 0xe9, 0x93, 0x82, 0x2b, 0x06,
	0x14, 0x2f, 0x87, 0x22, 0xbf, 0x51, 0x44, 0x7e, 0x32, 0x0d, 0x49, 0xa6, 0xca, 0x5d, 0x39, 0xa6,
	0x72, 0xd4, 0x32, 0xca, 0xb3, 0xab, 0x6c, 0x7c, 0x84, 0x82, 0xe9, 0xfb, 0xe2, 0x2c, 0x18, 0xbf,
	0xd3, 0xb3, 0x1e, 0x0f, 0x03, 0xaf, 0xc3, 0x13, 0x11, 0x93, 0x1a, 0xed, 0xc2, 0x1d, 0xef, 0xdc,
	0x93, 0x5e, 0x58, 0x13, 0xbe, 0xd7, 0x8e, 0x74, 0x8e, 0xbf, 0x4d, 0xba, 0x6d, 0xfc, 0xd6, 0x60,
	0x3d, 0x45, 0x7f, 0x93, 0xa4, 0x99, 0x70, 0xf5, 0xa6, 0x9f, 0xc2, 0x82, 0x17, 0xa3, 0xa2, 0xbd,
	0xcb, 0x65, 0x3d, 0x3d, 0x4d, 0x12, 0x40, 0xf1, 0xd7, 0x8d, 0xf2, 0x4f, 0x58, 0x38, 0x9b, 0xfc,
	0x5f, 0x6e, 0x71, 0xe4, 0x1b, 0x50, 0xb8, 0x3a, 0xf2, 0x06, 0x35, 0x09, 0x65, 0xd8, 0xa9, 0xd8,
	0x3a, 0x47, 0xf7, 0x01, 0x8d, 0x7c, 0x62, 0x53, 0x86, 0xcd, 0xb7, 0xb6, 0x7e, 0x9c, 0x9a, 0xb0,
	0x29, 0x33, 0xdf, 0x99, 0xd4, 0x24, 0xfa, 0x89, 0xf1, 0x6c, 0x68, 0x65, 0x83, 0xb7, 0xb9, 0x2f,
	0x13, 0x2b, 0x0b, 0x30, 0xdf, 0x0a, 0x9a, 0xfc, 0x73, 0xf4, 0xcb, 0xe7, 0x48, 0x5c, 0x18, 0x5f,
	0xe7, 0x86, 0xd6, 0x24, 0xf0, 0x29, 0xad, 0x51, 0xe1, 0xb7, 0xb7, 0xe6, 0x52, 0x9b, 0xd2, 0x9a,
	0xac, 0x85, 0xb3, 0x59, 0xf3, 0x3d, 0xf7, 0x0f, 0x02, 0x08, 0xed, 0xc0, 0xd6, 0x55, 0xcd, 0xec,
	0xf7, 0xd5, 0x06, 0x6d, 0xb0, 0x3d, 0x97, 0x32, 0xcb, 0x21, 0xc4, 0xad, 0x53, 0xbb, 0xa2, 0x1f,
	0xab, 0x3e, 0x12, 0x17, 0xc7, 0xfe, 0x9e, 0xa8, 0xec, 0xfb, 0x55, 0x5c, 0x61, 0xa6, 0x65, 0x39,
	0x2e, 0xa6, 0xac, 0x8a, 0xf7, 0x1d, 0xfd, 0xb4, 0xfc, 0x53, 0x83, 0x35, 0xab, 0xdd, 0xe2, 0x81,
	0xb4, 0x06, 0x0f, 0x0e, 0x3a, 0x80, 0x95, 0xd1, 0x44, 0x46, 0x0f, 0xaf, 0x09, 0xea, 0xc8, 0xf7,
	0xe2, 0xe6, 0xc4, 0x18, 0x47, 0x18, 0x56, 0x95, 0xe4, 0x41, 0x9b, 0xd7, 0x25, 0x52, 0x4c, 0xf7,
	0x68, 0x72, 0x60, 0xf5, 0xf9, 0x94, 0xbf, 0xd2, 0x08, 0x5f, 0x56, 0xc2, 0x8c, 0xf0, 0x65, 0x27,
	0xc4, 0x15, 0x5f, 0xec, 0xff, 0x38, 0x9f, 0x72, 0xe6, 0xe3, 0x7c, 0xea, 0xd9, 0xec, 0xe5, 0x3f,
	0x68, 0xdd, 0xa3, 0xa3, 0xff, 0xa3, 0xd7, 0xf9, 0xc5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfb,
	0x87, 0xdd, 0x80, 0xc6, 0x07, 0x00, 0x00,
}
